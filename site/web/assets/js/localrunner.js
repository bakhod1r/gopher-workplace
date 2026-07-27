/* Gopher Workplace — local runner bridge.

   The backend (site/cmd/localrunner) is a localhost HTTP server that runs the
   real Go toolchain: go test, -race, GC/finalizer timing, benchmarks. It binds
   127.0.0.1 and refuses non-loopback origins, because it executes arbitrary Go
   code as the person running it — so it is never hosted.

   On load we probe GET /health. Connected: every Run/Submit goes to POST /run.
   Absent — which is the normal state on the published site — the UI stays
   browse-only and says so, instead of offering buttons that cannot work. */

/* Backend address. Default :7070 matches the runner's default port. Override
   without editing this file: ?runner=http://localhost:9090 (sticky — remembered
   in localStorage under gw-runner). Must match RUNNER_PORT if you changed it. */
function gwRunnerBase() {
  const q = new URLSearchParams(location.search).get('runner');
  if (q) { try { localStorage.setItem('gw-runner', q); } catch (_) {} return q; }
  try { return localStorage.getItem('gw-runner') || 'http://localhost:7070'; }
  catch (_) { return 'http://localhost:7070'; }
}

window.GW_LOCAL = { connected: false, base: gwRunnerBase(), version: '' };

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

/* The setup a visitor needs to actually execute code. The published site is
   browse-only by design: the runner drives the real Go toolchain on your own
   machine and only accepts loopback connections, so it is never hosted. */
const GW_SETUP_STEPS = [
  'git clone https://github.com/bakhod1r/gopher-workplace.git',
  'cd gopher-workplace',
  'make setup && make dev',
];

/* Shared explanation, rendered wherever a run is refused. */
function gwOfflineHTML(action) {
  return '<div class="verdict mut">🔌 ' + (action || 'Running code') +
    ' happens on your machine, not here.</div>' +
    '<div class="cerr">This is the public, browse-only site: read the puzzle, ' +
    'try the editor, copy the code. To run the real <code>go test</code>, ' +
    'clone the repo and start the local runner:\n\n' +
    GW_SETUP_STEPS.join('\n') + '\n\nThen open http://localhost:7070 and the ' +
    'same puzzle becomes runnable.</div>';
}

/* A standing notice while no runner answers, so nobody has to press a dead
   button to find out why it is dead. Anchored to whatever the page leads with:
   the description pane on the playground, the dashboard elsewhere. */
function gwRenderOfflineNotice() {
  const host = document.querySelector('.left') || document.querySelector('.dash')
    || document.querySelector('.ps') || document.querySelector('.rm-wrap')
    || document.querySelector('.home') || document.querySelector('.wrap');
  if (!host) return;
  const anchor = document.getElementById('desc') || host.firstElementChild;
  let el = document.getElementById('gw-offline-note');
  if (window.GW_LOCAL.connected) {
    if (el) el.remove();
    return;
  }
  if (!el) {
    el = document.createElement('div');
    el.id = 'gw-offline-note';
    el.style.cssText = 'margin:0 0 1em;padding:.7em .9em;border-radius:8px;font-size:13px;' +
      'line-height:1.5;border:1px solid currentColor;opacity:.85;';
    el.innerHTML = '<strong>Browse-only.</strong> Run and Submit need the local ' +
      'runner — it executes the real Go toolchain on your own machine, so it ' +
      'cannot be hosted here. <code>make setup &amp;&amp; make dev</code>, then ' +
      'open <code>http://localhost:7070</code>. ' +
      '<a href="https://github.com/bakhod1r/gopher-workplace#install" ' +
      'target="_blank" rel="noopener">Setup guide →</a>';
    host.insertBefore(el, anchor || host.firstChild);
  }

  // The result pane's idle prompt invites a click that cannot work.
  const res = document.getElementById('p-result');
  if (res && /run your code/i.test(res.textContent)) {
    res.innerHTML = '<div class="verdict mut">browse-only — start the local runner to run code</div>';
  }
}

/* Reflect connection state on the Run/Submit buttons: disabled and explained
   rather than silently failing. */
function gwRenderRunButtons() {
  const local = window.GW_LOCAL.connected;
  for (const [id, label] of [['run', 'Run'], ['submit', 'Submit']]) {
    const btn = document.getElementById(id);
    if (!btn) continue;
    btn.disabled = !local;
    btn.style.opacity = local ? '' : '.5';
    btn.style.cursor = local ? '' : 'not-allowed';
    btn.title = local ? '' : label + ' needs the local runner — see the notice above';
  }
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
    el.title = 'Local runner not detected — this site is browse-only. Start: make dev';
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
  gwRenderOfflineNotice();
  gwRenderRunButtons();
  document.dispatchEvent(new CustomEvent('gw-runner', { detail: window.GW_LOCAL }));
}

document.addEventListener('DOMContentLoaded', gwProbeRunner);
