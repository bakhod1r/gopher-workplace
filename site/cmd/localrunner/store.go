package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

// store persists submissions to SQLite. Rows are kept forever unless a
// retention window is configured, in which case older rows are swept
// periodically.
type store struct {
	db *sql.DB
}

// sqlOpen is a seam: the sqlite driver opens lazily and never fails here, but
// the error path is still handled, and tests exercise it through this var.
var sqlOpen = sql.Open

func openStore(path string) (*store, error) {
	db, err := sqlOpen("sqlite", path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1) // sqlite: serialize writes
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS submissions (
			id           INTEGER PRIMARY KEY AUTOINCREMENT,
			challenge_id TEXT    NOT NULL,
			files        TEXT    NOT NULL,
			ok           INTEGER NOT NULL,
			total        INTEGER NOT NULL,
			passed       INTEGER NOT NULL,
			race         INTEGER NOT NULL,
			submitted    INTEGER NOT NULL DEFAULT 0,
			created_at   INTEGER NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_sub_ch  ON submissions(challenge_id);
		CREATE INDEX IF NOT EXISTS idx_sub_at  ON submissions(created_at);
	`); err != nil {
		db.Close()
		return nil, err
	}
	migrateSubmitted(db)
	migrateIDs(db)
	return &store{db: db}, nil
}

// migrateIDs rewrites challenge ids stored under the old level-first layout
// (<level>/<topic>/<subtopic>/<name>) to the current topic-first one
// (<topic>/<subtopic>/<level>/<name>). Without it a tree reshuffle silently
// orphans every past solve: the ids no longer match anything in the catalog, so
// /solved returns puzzles the site cannot recognise.
func migrateIDs(db *sql.DB) {
	rows, err := db.Query(`SELECT DISTINCT challenge_id FROM submissions`)
	if err != nil {
		log.Printf("migrate ids: %v", err)
		return
	}
	var olds []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			break
		}
		olds = append(olds, id)
	}
	rows.Close()

	n := 0
	for _, old := range olds {
		neu := modernID(old)
		if neu == old {
			continue
		}
		if _, err := db.Exec(
			`UPDATE submissions SET challenge_id = ? WHERE challenge_id = ?`, neu, old); err != nil {
			log.Printf("migrate id %s: %v", old, err)
			continue
		}
		n++
	}
	if n > 0 {
		log.Printf("migrated %d challenge id(s) to the topic-first layout", n)
	}
}

// levels are the level dir names, in learning-path order.
var levels = map[string]bool{"junior": true, "middle": true, "senior": true, "staff": true}

// modernID moves a leading level segment into third position; anything already
// in the current layout is returned unchanged.
func modernID(id string) string {
	p := strings.Split(id, "/")
	if len(p) != 4 || !levels[strings.ToLower(p[0])] {
		return id
	}
	return strings.Join([]string{p[1], p[2], p[0], p[3]}, "/")
}

// migrateSubmitted adds the submitted column to DBs that predate it. Already
// migrated is the normal case and stays quiet; anything else is worth a line in
// the log, but never fatal — the column only gates the solved set.
func migrateSubmitted(db *sql.DB) {
	if _, err := db.Exec(`ALTER TABLE submissions ADD COLUMN submitted INTEGER NOT NULL DEFAULT 0`); err != nil {
		if !strings.Contains(err.Error(), "duplicate column") {
			log.Printf("migrate submitted column: %v", err)
		}
	}
}

func (s *store) close() {
	if s != nil && s.db != nil {
		s.db.Close()
	}
}

// save records a run. submitted marks it as a Submit (not just a Run), which is
// what counts toward a challenge being "solved". Best-effort: logs and swallows.
func (s *store) save(challengeID string, files map[string]string, ok bool, total, passed int, race, submitted bool) {
	blob, _ := json.Marshal(files)
	_, err := s.db.Exec(
		`INSERT INTO submissions (challenge_id, files, ok, total, passed, race, submitted, created_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		challengeID, string(blob), b2i(ok), total, passed, b2i(race), b2i(submitted), time.Now().Unix(),
	)
	if err != nil {
		log.Printf("save submission: %v", err)
	}
}

// solvedIDs returns the challenge ids that have at least one passing Submit —
// the authoritative "solved" set, persisted in SQLite.
func (s *store) solvedIDs() ([]string, error) {
	rows, err := s.db.Query(
		`SELECT DISTINCT challenge_id FROM submissions WHERE ok = 1 AND submitted = 1`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []string{}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, rows.Err()
}

type historyRow struct {
	ID        int64             `json:"id"`
	OK        bool              `json:"ok"`
	Total     int               `json:"total"`
	Passed    int               `json:"passed"`
	Race      bool              `json:"race"`
	CreatedAt int64             `json:"createdAt"`
	Files     map[string]string `json:"files,omitempty"`
}

func (s *store) history(challengeID string, limit int) ([]historyRow, error) {
	rows, err := s.db.Query(
		`SELECT id, files, ok, total, passed, race, created_at
		 FROM submissions WHERE challenge_id = ?
		 ORDER BY created_at DESC, id DESC LIMIT ?`, challengeID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []historyRow{}
	for rows.Next() {
		var r historyRow
		var files string
		var ok, race int
		if err := rows.Scan(&r.ID, &files, &ok, &r.Total, &r.Passed, &race, &r.CreatedAt); err != nil {
			return nil, err
		}
		r.OK, r.Race = ok != 0, race != 0
		_ = json.Unmarshal([]byte(files), &r.Files)
		out = append(out, r)
	}
	return out, rows.Err()
}

// sweepInterval is how often retention runs after the startup sweep. A var so
// tests can drive the ticker without waiting an hour.
var sweepInterval = time.Hour

// startRetention sweeps rows older than window at startup and hourly
// thereafter. window <= 0 disables pruning entirely — rows are kept forever,
// which is the default, because the solved set is derived from this table.
func (s *store) startRetention(window time.Duration) {
	if window <= 0 {
		return
	}
	s.sweep(window)
	go func() {
		t := time.NewTicker(sweepInterval)
		defer t.Stop()
		for range t.C {
			s.sweep(window)
		}
	}()
}

func (s *store) sweep(window time.Duration) {
	cutoff := time.Now().Add(-window).Unix()
	res, err := s.db.Exec(`DELETE FROM submissions WHERE created_at < ?`, cutoff)
	if err != nil {
		log.Printf("retention sweep: %v", err)
		return
	}
	if n, _ := res.RowsAffected(); n > 0 {
		log.Printf("retention: pruned %d submission(s) older than %s", n, window)
	}
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
