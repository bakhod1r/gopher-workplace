package main

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func testStore(t *testing.T) *store {
	t.Helper()
	s, err := openStore(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(s.close)
	return s
}

func TestOpenStoreIsIdempotent(t *testing.T) {
	path := filepath.Join(t.TempDir(), "reopen.db")
	s, err := openStore(path)
	if err != nil {
		t.Fatal(err)
	}
	s.save("a", map[string]string{"a.go": "x"}, true, 1, 1, false, true)
	s.close()

	// Reopening must find the existing schema (and the ALTER for the submitted
	// column must not blow up on an already-migrated db).
	s2, err := openStore(path)
	if err != nil {
		t.Fatalf("reopen: %v", err)
	}
	defer s2.close()
	ids, err := s2.solvedIDs()
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != 1 || ids[0] != "a" {
		t.Errorf("solvedIDs after reopen = %q, want [a]", ids)
	}
}

func TestOpenStoreBadPath(t *testing.T) {
	// A path whose parent does not exist cannot be opened.
	if s, err := openStore(filepath.Join(t.TempDir(), "no-such-dir", "x.db")); err == nil {
		s.close()
		t.Error("openStore on a missing directory: want an error")
	}
}

// solvedIDs is the authoritative solved set: only a passing *Submit* counts. A
// passing Run, or a failing Submit, must not appear.
func TestSolvedIDsOnlyCountsPassingSubmits(t *testing.T) {
	s := testStore(t)
	s.save("passing-submit", nil, true, 2, 2, false, true)
	s.save("passing-run", nil, true, 2, 2, false, false)
	s.save("failing-submit", nil, false, 2, 1, false, true)
	s.save("passing-submit", nil, true, 2, 2, false, true) // duplicate

	ids, err := s.solvedIDs()
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != 1 || ids[0] != "passing-submit" {
		t.Errorf("solvedIDs = %q, want [passing-submit]", ids)
	}
}

func TestSolvedIDsEmpty(t *testing.T) {
	ids, err := testStore(t).solvedIDs()
	if err != nil {
		t.Fatal(err)
	}
	if ids == nil || len(ids) != 0 {
		t.Errorf("solvedIDs = %#v, want an empty non-nil slice", ids)
	}
}

func TestHistoryOrderLimitAndRoundTrip(t *testing.T) {
	s := testStore(t)
	files := map[string]string{"swap.go": "package swap\n"}
	for i := 0; i < 5; i++ {
		s.save("ch", files, i%2 == 0, 3, i, i == 4, false)
	}
	s.save("other", nil, true, 1, 1, false, false)

	rows, err := s.history("ch", 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 3 {
		t.Fatalf("rows = %d, want 3 (limit)", len(rows))
	}
	// Most recent first: the last insert has the highest id.
	if rows[0].ID <= rows[1].ID || rows[1].ID <= rows[2].ID {
		t.Errorf("not ordered newest-first: %d, %d, %d", rows[0].ID, rows[1].ID, rows[2].ID)
	}
	if !rows[0].Race {
		t.Error("race flag lost on the newest row")
	}
	if rows[0].Passed != 4 || rows[0].Total != 3 {
		t.Errorf("passed/total = %d/%d, want 4/3", rows[0].Passed, rows[0].Total)
	}
	if rows[0].Files["swap.go"] != files["swap.go"] {
		t.Errorf("files round-trip = %#v", rows[0].Files)
	}

	empty, err := s.history("nope", 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(empty) != 0 {
		t.Errorf("history for an unknown challenge = %d rows", len(empty))
	}
}

// A row with unparseable files JSON must still be returned — history is a
// convenience, not a correctness boundary.
func TestHistorySurvivesBadFilesBlob(t *testing.T) {
	s := testStore(t)
	if _, err := s.db.Exec(
		`INSERT INTO submissions (challenge_id, files, ok, total, passed, race, submitted, created_at)
		 VALUES ('ch', 'not json', 1, 1, 1, 0, 0, ?)`, time.Now().Unix()); err != nil {
		t.Fatal(err)
	}
	rows, err := s.history("ch", 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("rows = %d, want 1", len(rows))
	}
	if len(rows[0].Files) != 0 {
		t.Errorf("files = %#v, want empty", rows[0].Files)
	}
}

func TestSweepPrunesOnlyOldRows(t *testing.T) {
	s := testStore(t)
	old := time.Now().Add(-48 * time.Hour).Unix()
	if _, err := s.db.Exec(
		`INSERT INTO submissions (challenge_id, files, ok, total, passed, race, submitted, created_at)
		 VALUES ('ch', '{}', 1, 1, 1, 0, 1, ?)`, old); err != nil {
		t.Fatal(err)
	}
	s.save("ch", nil, true, 1, 1, false, true) // fresh

	s.sweep(24 * time.Hour)

	rows, err := s.history("ch", 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("rows after sweep = %d, want 1 (the fresh one)", len(rows))
	}
	if rows[0].CreatedAt == old {
		t.Error("sweep kept the stale row and dropped the fresh one")
	}
}

func TestSweepOnEmptyTableIsQuiet(t *testing.T) {
	testStore(t).sweep(time.Hour) // must not panic or error out
}

func TestStartRetentionSweepsImmediately(t *testing.T) {
	s := testStore(t)
	if _, err := s.db.Exec(
		`INSERT INTO submissions (challenge_id, files, ok, total, passed, race, submitted, created_at)
		 VALUES ('ch', '{}', 1, 1, 1, 0, 1, ?)`, time.Now().Add(-48*time.Hour).Unix()); err != nil {
		t.Fatal(err)
	}
	s.startRetention(time.Hour)
	rows, err := s.history("ch", 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 0 {
		t.Errorf("rows = %d, want 0: startRetention must sweep at startup", len(rows))
	}
}

// save is best-effort: a write against a closed db logs and returns rather than
// failing the request that triggered it.
func TestSaveSwallowsErrors(t *testing.T) {
	s, err := openStore(filepath.Join(t.TempDir(), "closed.db"))
	if err != nil {
		t.Fatal(err)
	}
	s.close()
	s.save("ch", nil, true, 1, 1, false, true) // must not panic
}

func TestStoreCloseIsSafeOnNil(t *testing.T) {
	var s *store
	s.close()
	(&store{}).close()
}

func TestB2I(t *testing.T) {
	if b2i(true) != 1 || b2i(false) != 0 {
		t.Error("b2i")
	}
}

// permissiveStore rebuilds the table without NOT NULL constraints so scan
// failures — a row the schema should never allow — can be exercised.
func permissiveStore(t *testing.T) *store {
	t.Helper()
	s := testStore(t)
	if _, err := s.db.Exec(`DROP TABLE submissions;
		CREATE TABLE submissions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			challenge_id TEXT, files TEXT, ok INTEGER, total INTEGER,
			passed INTEGER, race INTEGER, submitted INTEGER, created_at INTEGER)`); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestSolvedIDsScanError(t *testing.T) {
	s := permissiveStore(t)
	if _, err := s.db.Exec(
		`INSERT INTO submissions (challenge_id, ok, submitted) VALUES (NULL, 1, 1)`); err != nil {
		t.Fatal(err)
	}
	if _, err := s.solvedIDs(); err == nil {
		t.Error("a NULL challenge_id must fail the scan, not be skipped")
	}
}

func TestHistoryScanError(t *testing.T) {
	s := permissiveStore(t)
	// created_at is read into an int64; sqlite's loose typing lets text land in
	// an INTEGER column, and that must surface as an error.
	if _, err := s.db.Exec(
		`INSERT INTO submissions (challenge_id, files, ok, total, passed, race, created_at)
		 VALUES ('ch', '{}', 1, 1, 1, 0, 'not-a-time')`); err != nil {
		t.Fatal(err)
	}
	if _, err := s.history("ch", 10); err == nil {
		t.Error("a non-numeric created_at must fail the scan")
	}
}

func TestOpenStoreSurfacesDriverError(t *testing.T) {
	old := sqlOpen
	sqlOpen = func(string, string) (*sql.DB, error) { return nil, errors.New("driver exploded") }
	t.Cleanup(func() { sqlOpen = old })

	if s, err := openStore(filepath.Join(t.TempDir(), "x.db")); err == nil {
		s.close()
		t.Error("want the driver error")
	}
}

// An unexpected migration failure (here: no table at all) is logged, not
// swallowed and not fatal — the column only gates the solved set.
func TestMigrateSubmittedLogsUnexpectedError(t *testing.T) {
	logs := captureLog(t)
	db, err := sql.Open("sqlite", filepath.Join(t.TempDir(), "bare.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	migrateSubmitted(db) // no submissions table exists
	if !strings.Contains(logs.String(), "migrate submitted column") {
		t.Errorf("logs = %q, want the migration error", logs.String())
	}
}

// The already-migrated case must stay silent.
func TestMigrateSubmittedQuietWhenAlreadyMigrated(t *testing.T) {
	logs := captureLog(t)
	s := testStore(t)
	migrateSubmitted(s.db)
	if strings.Contains(logs.String(), "migrate submitted column") {
		t.Errorf("logs = %q, want silence on a duplicate column", logs.String())
	}
}

func TestOpenStoreLogsUnexpectedMigrationError(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("running as root: a read-only file is still writable")
	}
	path := filepath.Join(t.TempDir(), "old.db")
	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatal(err)
	}
	// The pre-migration schema: no `submitted` column.
	if _, err := db.Exec(`CREATE TABLE submissions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		challenge_id TEXT NOT NULL, files TEXT NOT NULL, ok INTEGER NOT NULL,
		total INTEGER NOT NULL, passed INTEGER NOT NULL, race INTEGER NOT NULL,
		created_at INTEGER NOT NULL)`); err != nil {
		t.Fatal(err)
	}
	db.Close()
	if err := os.Chmod(path, 0o444); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(path, 0o644) })

	logs := captureLog(t)
	s, err := openStore(path)
	if err != nil {
		t.Skipf("read-only db rejected earlier than the migration: %v", err)
	}
	defer s.close()
	if !strings.Contains(logs.String(), "migrate submitted column") {
		t.Errorf("logs = %q, want the migration error", logs.String())
	}
}
