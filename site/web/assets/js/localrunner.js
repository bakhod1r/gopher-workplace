/* Gopher Workplace — local runner bridge.

   Optional full-fidelity backend: a localhost HTTP server (site/cmd/localrunner)
   that runs the real Go toolchain — go test, -race, GC/finalizer timing,
   benchmarks — things the in-browser yaegi-wasm interpreter cannot do.

   On load we probe GET /health. If it answers, window.GW_LOCAL.connected is set
   and app.js routes EVERY Run/Submit through POST /run (all levels), rendering
   the same Report shape. When absent, app.js falls back to the wasm path
   (junior/middle only). */

window.GW_LOCAL = { connected: false, base: 'http://localhost:7070', version: '' };

/* A problem needs the real backend when it can't run under wasm: senior/staff
   levels, or anything explicitly tagged backend. */
function gwNeedsBackend(p) {
  if (!p) return false;
  if (p.tag === 'backend') return true;
  const lv = (p.level || (p.id || '').split('/')[0] || '').toLowerCase();
  return lv === 'senior' || lv === 'staff';
}

/* POST /run with the candidate source; resolves to a Report JSON *string*
   (same shape render() already consumes). submit=true records the run as a
   Submit, which is what counts toward the persisted solved set. */
async function gwLocalRun(challengeId, src, submit) {
  const res = await fetch(window.GW_LOCAL.base + '/run', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ challengeId, src, submit: !!submit }),
  });
  return res.text();
}

/* GET /solved — the authoritative solved set from SQLite. Resolves to an array
   of challenge ids (empty on any failure). */
async function gwLocalSolved() {
  try {
    const res = await fetch(window.GW_LOCAL.base + '/solved');
    const j = await res.json();
    return (j && j.solved) || [];
  } catch (e) { return []; }
}

/* GET /history — recent saved submissions for a challenge (most recent first).
   Resolves to an array of rows (empty on any failure). */
async function gwLocalHistory(challengeId) {
  try {
    const res = await fetch(window.GW_LOCAL.base + '/history?challengeId=' + encodeURIComponent(challengeId));
    const j = await res.json();
    return Array.isArray(j) ? j : [];
  } catch (e) { return []; }
}

/* The candidate source of a history row, if it stored one. */
function gwRowCode(row) {
  const f = row && row.files;
  if (!f) return '';
  if (typeof f.src === 'string') return f.src;
  const keys = Object.keys(f);
  return keys.length ? f[keys[0]] : '';
}

/* gofmt via the backend; resolves to {ok, source, error} or null on failure. */
async function gwLocalFmt(src) {
  try {
    const res = await fetch(window.GW_LOCAL.base + '/fmt', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ src }),
    });
    return await res.json();
  } catch (e) { return null; }
}

/* Small nav badge reflecting connection state. */
function gwRenderBadge() {
  let el = document.getElementById('gw-runner-badge');
  if (!el) {
    const actions = document.querySelector('.nav .actions') || document.querySelector('.nav');
    if (!actions) return;
    el = document.createElement('span');
    el.id = 'gw-runner-badge';
    el.style.cssText = 'display:inline-flex;align-items:center;gap:.4em;font-size:12px;'
      + 'padding:.25em .6em;border-radius:999px;border:1px solid currentColor;opacity:.85;'
      + 'white-space:nowrap;margin-right:.4em;';
    actions.insertBefore(el, actions.firstChild);
  }
  if (window.GW_LOCAL.connected) {
    el.style.color = '#16a34a';
    el.textContent = '● local runner';
    el.title = 'Real Go toolchain connected (v' + window.GW_LOCAL.version + ') — all puzzles runnable';
  } else {
    el.style.color = '#9ca3af';
    el.textContent = '○ offline';
    el.title = 'Local runner not detected. Start: go run ./site/cmd/localrunner';
  }
}

/* Probe /health with a short timeout so a missing runner never blocks the UI. */
async function gwProbeRunner() {
  const ctl = new AbortController();
  const t = setTimeout(() => ctl.abort(), 800);
  try {
    const res = await fetch(window.GW_LOCAL.base + '/health', { signal: ctl.signal });
    const j = await res.json();
    window.GW_LOCAL.connected = !!(j && j.ok);
    window.GW_LOCAL.version = (j && j.version) || '';
  } catch (e) {
    window.GW_LOCAL.connected = false;
  } finally {
    clearTimeout(t);
  }
  gwRenderBadge();
  document.dispatchEvent(new CustomEvent('gw-runner', { detail: window.GW_LOCAL }));
}

document.addEventListener('DOMContentLoaded', gwProbeRunner);
