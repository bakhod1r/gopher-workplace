/* Gopher Workplace — frontend logic. Data comes from problems.js; Go execution
   runs on the local Go toolchain backend (site/cmd/localrunner) via localrunner.js. */

const $ = s => document.querySelector(s), $$ = s => document.querySelectorAll(s);

/* ---- theme (persisted) ---- */
(function themeInit(){
  const sel = $('#theme');
  const saved = localStorage.getItem('gw-theme') || 'light';
  document.documentElement.setAttribute('data-theme', saved);
  sel.value = saved;
  sel.addEventListener('change', () => {
    document.documentElement.setAttribute('data-theme', sel.value);
    localStorage.setItem('gw-theme', sel.value);
  });
})();

/* ---- active problem ---- */
let CUR = window.PROBLEMS[Object.keys(window.PROBLEMS)[0]] || null;

const codeEl = $('#code'), linesEl = $('#lines'), hlEl = $('#hl'), runEl = $('#run'), subEl = $('#submit');
const pResult = $('#p-result'), tcOut = $('#tcOut'), tcInput = $('#tcInput'), descEl = $('#desc');

// Editor mode: 'learn' = implement-from-scratch stub (p.starter),
// 'debug' = fix a planted bug (p.debug, only when the puzzle ships one).
let MODE = 'learn';
const modeBtn = $('#btnMode');
const hasDebug = p => !!(p && p.debug);
const modeBase = () => (MODE === 'debug' && hasDebug(CUR)) ? CUR.debug : CUR.starter;

// Per-puzzle, per-mode editor draft, persisted in localStorage so in-progress
// typing survives a page refresh. This is the CURRENT draft, not past
// submissions (those live in the Submissions tab).
const draftKey = () => CUR ? 'gw-draft:' + CUR.id + ':' + MODE : '';
function saveDraft(){
  if(!CUR) return;
  try{ localStorage.setItem(draftKey(), codeEl.value); }catch(e){}
}
function clearDraft(){
  if(!CUR) return;
  try{ localStorage.removeItem(draftKey()); }catch(e){}
}
// The saved draft for the current puzzle/mode, or the pristine stub if none.
function baseOrDraft(){
  const base = modeBase();
  if(!CUR) return base;
  try{ const d = localStorage.getItem(draftKey()); if(d != null) return d; }catch(e){}
  return base;
}

function syncModeBtn(){
  if(!modeBtn) return;
  if(hasDebug(CUR)){
    modeBtn.hidden = false;
    modeBtn.textContent = MODE === 'debug' ? '◨ debug' : '◧ learn';
    modeBtn.title = MODE === 'debug'
      ? 'Debug mode: fix the planted bug. Click to build from scratch instead.'
      : 'Learn mode: build from scratch. Click to fix a planted bug instead.';
  } else {
    modeBtn.hidden = true;
  }
}

const solvedBadge = $('#solvedBadge');
const isSolved = id => !!(id && window.GWProgress && window.GWProgress.isSolved(id));
function syncSolvedBadge(){
  if(solvedBadge) solvedBadge.hidden = !isSolved(CUR && CUR.id);
}

// Put a ✓ on a sidebar row when its puzzle is solved.
function markDrawerSolved(id){
  if(!id) return;
  const it = document.querySelector('#drawer .pitem[data-id="'+CSS.escape(id)+'"]');
  if(!it) return;
  it.classList.add('done');
  const st = it.querySelector('.st'); if(st) st.textContent = '✓';
}
// Reflect all currently-solved puzzles in the sidebar (call after it is built).
function refreshDrawerSolved(){
  document.querySelectorAll('#drawer .pitem').forEach(it => {
    if(isSolved(it.dataset.id)) markDrawerSolved(it.dataset.id);
  });
}

function loadProblem(p){
  CUR = p;
  MODE = 'learn';                 // every problem opens in learn mode
  if(!p){
    descEl.innerHTML = '<div class="verdict mut">no problems loaded</div>';
    codeEl.value = ''; syncLines(); syncModeBtn(); syncSolvedBadge(); return;
  }
  descEl.innerHTML = p.description;
  codeEl.value = baseOrDraft();
  tcInput.value = p.customDefault || '';
  $('#edFile').textContent = p.file;
  pResult.innerHTML = isSolved(p.id)
    ? '<div class="verdict ok">✓ submitted — solved earlier. Run again anytime.</div>'
    : '<div class="verdict mut">run your code<span class="blink">_</span></div>';
  syncLines();
  syncModeBtn();
  syncSolvedBadge();
}
// Note: the editor always opens on the pristine stub (modeBase). Past
// submissions are not auto-restored — view/load them from the Submissions tab.

if(modeBtn){
  modeBtn.addEventListener('click', () => {
    if(!hasDebug(CUR)) return;
    MODE = MODE === 'debug' ? 'learn' : 'debug';
    codeEl.value = baseOrDraft();  // swap to the chosen mode's draft (or stub)
    syncLines(); syncModeBtn(); codeEl.focus();
  });
}

/* ---- Go syntax highlight (overlay) ---- */
const GO_KW = new Set("break case chan const continue default defer else fallthrough for func go goto if import interface map package range return select struct switch type var".split(" "));
// Every Go keyword gets its OWN distinct colour. Palette is spread evenly
// around the hue wheel at fixed saturation/lightness so all 25 stay legible.
const KW_LIST = "break case chan const continue default defer else fallthrough for func go goto if import interface map package range return select struct switch type var".split(" ");
const KW_COLOR = (() => {
  const m = {}, n = KW_LIST.length;
  KW_LIST.forEach((w, i) => { m[w] = 'hsl(' + Math.round(i * 360 / n) + ' 68% 62%)'; });
  return m;
})();
const kwColor = w => KW_COLOR[w] || '#c678dd';
const GO_TY = new Set("int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex64 complex128 byte rune string bool error any comparable true false nil iota make len cap append copy delete new clear min max".split(" "));
const escH = s => s.replace(/[&<>]/g, m => ({'&':'&amp;','<':'&lt;','>':'&gt;'}[m]));

function highlight(src){
  // tokenize by regex; order matters (comments/strings first)
  const re = /(\/\/[^\n]*|\/\*[\s\S]*?\*\/)|(`[^`]*`|"(?:\\.|[^"\\])*"|'(?:\\.|[^'\\])*')|(\b\d[\d_.eExXa-fA-F]*\b)|([A-Za-z_]\w*)/g;
  let out = '', last = 0, m;
  while((m = re.exec(src))){
    out += escH(src.slice(last, m.index));
    last = re.lastIndex;
    if(m[1]){ // comment
      const cls = /CHANGE CODE/.test(m[1]) ? 'tk-marker' : 'tk-com';
      out += '<span class="'+cls+'">'+escH(m[1])+'</span>';
    } else if(m[2]){ out += '<span class="tk-str">'+escH(m[2])+'</span>'; }
    else if(m[3]){ out += '<span class="tk-num">'+escH(m[3])+'</span>'; }
    else if(m[4]){
      const w = m[4];
      const after = src.slice(re.lastIndex).match(/^\s*\(/); // call?
      if(GO_KW.has(w)) out += '<span class="tk-kw" style="color:'+kwColor(w)+'">'+w+'</span>';
      else if(GO_TY.has(w)) out += '<span class="tk-ty">'+w+'</span>';
      else if(after) out += '<span class="tk-fn">'+escH(w)+'</span>';
      else out += escH(w);
    }
  }
  out += escH(src.slice(last));
  return out + '\n';
}

/* ---- editor: line numbers, highlight, tab, shortcut ---- */
function syncLines(){
  const n = codeEl.value.split('\n').length;
  let s=''; for(let i=1;i<=n;i++) s+=i+'\n';
  linesEl.textContent = s;
  hlEl.innerHTML = highlight(codeEl.value);
  saveDraft();   // persist current editor content so a refresh keeps it
}
function syncScroll(){
  hlEl.scrollTop = codeEl.scrollTop; hlEl.scrollLeft = codeEl.scrollLeft;
  linesEl.scrollTop = codeEl.scrollTop;
}
codeEl.addEventListener('input', syncLines);
codeEl.addEventListener('scroll', syncScroll);
const PAIR_OPEN = {'(':')','[':']','{':'}'};
const PAIR_QUOTE = new Set(['"',"'",'`']);
const PAIR_CLOSE = new Set([')',']','}']);
function setCaret(pos){ codeEl.selectionStart = codeEl.selectionEnd = pos; }
codeEl.addEventListener('keydown', e => {
  const s=codeEl.selectionStart, en=codeEl.selectionEnd, v=codeEl.value;

  if(e.key === 'Tab'){ e.preventDefault();
    codeEl.value = v.slice(0,s)+'\t'+v.slice(en);
    setCaret(s+1); syncLines(); return;
  }
  const mod = e.ctrlKey || e.metaKey;
  // Ctrl/Cmd+Enter runs; add Shift to submit. Ctrl/Cmd+S also submits.
  if(mod && e.key === 'Enter'){ e.preventDefault(); run(e.shiftKey); return; }
  if(mod && (e.key === 's' || e.key === 'S')){ e.preventDefault(); run(true); return; }

  // Ctrl/Cmd+/ toggles line comments over the selection (LeetCode-style).
  if(mod && e.key === '/'){
    e.preventDefault();
    const lineStart = v.lastIndexOf('\n', s-1)+1;
    let lineEnd = v.indexOf('\n', en); if(lineEnd === -1) lineEnd = v.length;
    const block = v.slice(lineStart, lineEnd);
    const lines = block.split('\n');
    const commented = lines.every(l => l.trim()==='' || /^\s*\/\/ ?/.test(l));
    const out = lines.map(l => {
      if(commented) return l.replace(/^(\s*)\/\/ ?/, '$1');
      if(l.trim()==='') return l;
      const m = l.match(/^\s*/)[0];
      return m + '// ' + l.slice(m.length);
    }).join('\n');
    codeEl.value = v.slice(0,lineStart) + out + v.slice(lineEnd);
    const delta = out.length - block.length;
    codeEl.selectionStart = lineStart;
    codeEl.selectionEnd = lineEnd + delta;
    syncLines(); return;
  }

  // Skip over an auto-inserted closing bracket/quote instead of typing a second.
  if(s===en && (PAIR_CLOSE.has(e.key) || PAIR_QUOTE.has(e.key)) && v[s]===e.key){
    e.preventDefault(); setCaret(s+1); return;
  }
  // Auto-close: opening bracket or quote wraps the selection (or inserts a pair).
  if(PAIR_OPEN[e.key] || PAIR_QUOTE.has(e.key)){
    e.preventDefault();
    const close = PAIR_OPEN[e.key] || e.key;
    const sel = v.slice(s,en);
    codeEl.value = v.slice(0,s)+e.key+sel+close+v.slice(en);
    if(sel){ codeEl.selectionStart=s+1; codeEl.selectionEnd=en+1; }  // keep selection wrapped
    else setCaret(s+1);
    syncLines(); return;
  }
  // Backspace inside an empty pair removes both halves.
  if(e.key==='Backspace' && s===en && s>0){
    const a=v[s-1], b=v[s];
    if((PAIR_OPEN[a] && b===PAIR_OPEN[a]) || (PAIR_QUOTE.has(a) && b===a)){
      e.preventDefault();
      codeEl.value = v.slice(0,s-1)+v.slice(s+1); setCaret(s-1); syncLines(); return;
    }
  }
  // Enter between {} expands to an indented block.
  if(e.key==='Enter' && s===en && v[s-1]==='{' && v[s]==='}'){
    e.preventDefault();
    const lineStart = v.lastIndexOf('\n', s-1)+1;
    const indent = (v.slice(lineStart, s).match(/^\t*/) || [''])[0];
    codeEl.value = v.slice(0,s)+'\n'+indent+'\t\n'+indent+v.slice(s);
    setCaret(s+1+indent.length+1); syncLines(); return;
  }
});

/* ---- editor toolbar ---- */
(function toolbar(){
  const box = $('.codebox');
  let fs = parseInt(localStorage.getItem('gw-font')||'13',10);
  const applyFont = () => { [codeEl,hlEl,linesEl].forEach(el=>el.style.fontSize=fs+'px'); localStorage.setItem('gw-font',fs); };
  applyFont();
  $('#btnReset').addEventListener('click', () => { clearDraft(); codeEl.value = modeBase(); syncLines(); codeEl.focus(); });
  $('#btnFormat').addEventListener('click', () => {
    const b=$('#btnFormat'), t=b.textContent;
    // Prefer the local runner's real gofmt when connected.
    if(window.GW_LOCAL && window.GW_LOCAL.connected && typeof gwLocalFmt === 'function'){
      gwLocalFmt(codeEl.value).then(r => {
        if(r && r.ok){ codeEl.value=r.source; syncLines(); b.textContent='✓ formatted'; setTimeout(()=>b.textContent=t,1000); }
        else{ b.textContent='⚠ syntax'; setTimeout(()=>b.textContent=t,1200); }
      });
      return;
    }
    // Backend down: fall back to a light client-side tidy.
    codeEl.value = codeEl.value.replace(/[ \t]+$/gm,'').replace(/\n{3,}/g,'\n\n'); syncLines();
  });
  $('#btnCopy').addEventListener('click', async () => {
    try{ await navigator.clipboard.writeText(codeEl.value); const b=$('#btnCopy'); const t=b.textContent; b.textContent='✓ copied'; setTimeout(()=>b.textContent=t,1200);}catch(e){}
  });
  $('#btnFontUp').addEventListener('click', () => { fs=Math.min(fs+1,22); applyFont(); });
  $('#btnFontDn').addEventListener('click', () => { fs=Math.max(fs-1,10); applyFont(); });
  $('#btnFull').addEventListener('click', () => { document.body.classList.toggle('ed-full'); });
})();

/* ---- console tabs ---- */
$$('.console .tab').forEach(t => t.addEventListener('click', () => {
  $$('.console .tab').forEach(x => x.classList.remove('on')); t.classList.add('on');
  const p = t.dataset.panel;
  $$('.console .cpanel').forEach(el => el.classList.remove('on'));
  const panel = $('#p-'+p); if(panel) panel.classList.add('on');
  if(p==='subs') loadSubmissions();
}));

/* ---- submissions history (LeetCode-style) ---- */
const subsEl = $('#p-subs');
function timeAgo(sec){
  const d = Math.max(0, Math.floor(Date.now()/1000 - sec));
  if(d < 60) return d+'s ago';
  if(d < 3600) return Math.floor(d/60)+'m ago';
  if(d < 86400) return Math.floor(d/3600)+'h ago';
  return Math.floor(d/86400)+'d ago';
}
function loadSubmissions(){
  if(!subsEl) return;
  if(!(window.GW_LOCAL && window.GW_LOCAL.connected)){
    subsEl.innerHTML = '<div class="verdict mut">connect the local runner to see submissions.</div>'; return;
  }
  if(!CUR){ subsEl.innerHTML = ''; return; }
  subsEl.innerHTML = '<div class="verdict mut">loading<span class="blink">_</span></div>';
  const forId = CUR.id;
  gwLocalHistory(forId).then(rows => {
    if(!CUR || CUR.id !== forId) return;         // navigated away
    if(!rows || !rows.length){
      subsEl.innerHTML = '<div class="verdict mut">no submissions yet — Run or Submit to record one.</div>'; return;
    }
    let html = '<table class="subs"><thead><tr><th>#</th><th>status</th><th>tests</th><th>lang</th><th>when</th></tr></thead><tbody>';
    rows.forEach((r, i) => {
      const ok = r.ok;
      const st = ok ? '<span class="sv ok">Accepted</span>' : '<span class="sv bad">Wrong Answer</span>';
      html += '<tr class="subrow" data-idx="'+i+'">'
        + '<td>'+(rows.length - i)+'</td>'
        + '<td>'+st+'</td>'
        + '<td class="gw">'+r.passed+'/'+r.total+'</td>'
        + '<td>Go'+(r.race?' <span class="gw">·race</span>':'')+'</td>'
        + '<td class="gw">'+timeAgo(r.createdAt)+'</td></tr>';
    });
    html += '</tbody></table><div class="subs-hint mut">click a row to load its code into the editor</div>';
    subsEl.innerHTML = html;
    subsEl.querySelectorAll('.subrow').forEach(tr => tr.addEventListener('click', () => {
      const row = rows[+tr.dataset.idx];
      const code = (typeof gwRowCode === 'function') ? gwRowCode(row) : '';
      if(code){ codeEl.value = code; syncLines(); codeEl.focus(); }
    }));
  });
}

/* ---- resizable split ---- */
(function splitInit(){
  const g = $('#gutter'), left = $('.left'); let drag=false;
  g.addEventListener('mousedown', () => { drag=true; document.body.style.userSelect='none'; });
  window.addEventListener('mouseup', () => { drag=false; document.body.style.userSelect=''; });
  window.addEventListener('mousemove', e => {
    if(!drag) return;
    left.style.width = Math.max(280, Math.min(e.clientX, window.innerWidth-360)) + 'px';
  });
})();

/* ---- vertical split: editor / console ---- */
(function vsplitInit(){
  const g = $('#vgutter'), con = $('#console'), right = $('.right'); let drag=false;
  g.addEventListener('mousedown', () => { drag=true; document.body.style.userSelect='none'; });
  window.addEventListener('mouseup', () => { drag=false; document.body.style.userSelect=''; });
  window.addEventListener('mousemove', e => {
    if(!drag) return;
    const r = right.getBoundingClientRect();
    const h = Math.max(90, Math.min(r.bottom - e.clientY, r.height - 120));
    con.style.height = h + 'px';
  });
})();

/* ---- runner boot ----
   Local-runner-only build: no wasm. Run/Submit stay enabled and route through
   the localhost Go backend; run() shows a hint if the backend is not up yet. */
function boot(){
  if (typeof gwRenderRunButtons === 'function') gwRenderRunButtons();
  else { runEl.disabled = false; subEl.disabled = false; }
}

const esc = s => String(s).replace(/[&<>]/g, m => ({'&':'&amp;','<':'&lt;','>':'&gt;'}[m]));

/* ---- run suite ---- */
function render(json, submit){
  let r; try{ r=JSON.parse(json); }catch(e){ pResult.innerHTML='<div class="verdict bad">runtime error</div>'; return; }
  if(!r.compileOk){ pResult.innerHTML='<div class="verdict bad">compile error</div><div class="cerr">'+esc(r.error||'')+'</div>'; return; }
  const pass = r.cases.filter(c=>c.pass).length, total = r.cases.length;
  const warns = (r.warnings || []);
  const warned = warns.length > 0;
  // A Run just checks correctness (warnings are advisory). A Submit is only
  // accepted when it passes AND is warning-free — hardcoded / unformatted
  // answers run green but are rejected on Submit.
  let head;
  if(submit){
    if(!r.ok) head = '<div class="verdict bad">WRONG_ANSWER <span class="gw">'+pass+'/'+total+'</span></div>';
    else if(warned) head = '<div class="verdict bad">NOT ACCEPTED <span class="gw">resolve warnings to submit</span></div>';
    else head = '<div class="verdict ok">ACCEPTED</div>';
  } else {
    head = '<div class="verdict '+(r.ok?'ok':'bad')+'">'+(r.ok?'PASSED':'FAILED')+' <span class="gw">'+pass+'/'+total+'</span></div>';
  }
  // Solved only on an accepted submit (passing + warning-free) — matches the
  // sqlite submitted flag the backend writes.
  if(submit && r.ok && !warned && CUR){
    try{ window.GWProgress && window.GWProgress.markSolved(CUR.id); }catch(e){}
    syncSolvedBadge();
    markDrawerSolved(CUR.id);
  }
  let html = head;
  if(warned){
    html += '<div class="warns">';
    for(const w of warns) html += '<div class="warn">⚠ '+esc(w)+'</div>';
    html += '</div>';
  }
  for(const c of r.cases)
    html += '<div class="case '+(c.pass?'pass':'fail')+'"><span class="t">'+(c.pass?'✓':'✗')+'</span>'
      + '<span class="n">'+esc(c.name)+'</span>'
      + (c.pass?'':'<span class="gw">got '+esc(c.got)+' · want '+esc(c.want)+'</span>')+'</div>';
  pResult.innerHTML = html;
  // A run just recorded a new submission — refresh the list if it's showing.
  if($('#p-subs') && $('#p-subs').classList.contains('on')) loadSubmissions();
}
function run(submit){
  if(!CUR) return;
  const local = window.GW_LOCAL && window.GW_LOCAL.connected;
  $$('.console .tab')[0].click();
  // Every problem runs against the real Go toolchain via the local backend.
  if(!local){
    pResult.innerHTML = gwOfflineHTML(submit ? 'Submitting' : 'Running');
    return;
  }
  pResult.innerHTML = '<div class="verdict mut">'+(submit?'submitting':'running')+'<span class="blink">_</span></div>';
  gwLocalRun(CUR.id, codeEl.value, submit)
    .then(txt => render(txt, submit))
    .catch(e => { pResult.innerHTML = '<div class="verdict bad">local runner error</div><div class="cerr">'+esc(e)+'</div>'; });
}
runEl.addEventListener('click', () => run(false));
subEl.addEventListener('click', () => run(true));

/* ---- custom testcase (input only) ---- */
$('#tcRun').addEventListener('click', () => {
  if(!CUR) return;
  // The local backend runs the puzzle's own test suite (real `go test`); it has
  // no single-input eval endpoint, so custom input is not supported here.
  tcOut.innerHTML = '<div class="cerr">There is no single-input eval: the runner executes the '
    + "puzzle's own test suite with the real <code>go test</code>. Use <b>Run</b> once the local "
    + 'runner is up.</div>';
});

/* ---- sidebar from catalog ---- */
(function buildDrawer(){
  const drawer = $('#drawer'), scrim = $('#scrim'), btn = $('#probsBtn');
  let html = '';
  let num = 0;
  for(const grp of window.CATALOG){
    const openAttr = grp.items.some(it => it.id===CUR.id) ? ' open' : '';
    html += '<details class="tgroup"'+openAttr+'>'
      + '<summary class="topic"><span class="tname">'+esc(grp.topic)+'</span>'
      + '<span class="tcount">'+grp.items.length+'</span></summary><ul class="plist">';
    for(const it of grp.items){
      const cls = 'pitem'+(it.locked?' locked':'')+(it.done?' done':'')+(it.id===CUR.id?' on':'');
      num++;
      html += '<li class="'+cls+'" data-id="'+esc(it.id||'')+'">'
        + '<span class="pnum">'+num+'.</span>'
        + '<span class="st">'+(it.done?'✓':'▫')+'</span>'
        + '<span class="nm">'+esc(it.name)+(it.sub?'<span class="sub">'+esc(it.sub)+'</span>':'')+'</span>'
        + (it.lv?'<span class="lv">'+esc(it.lv)+'</span>':'')+'</li>';
    }
    html += '</ul></details>';
  }
  drawer.innerHTML = html;

  const toggle = o => {
    drawer.classList.toggle('open',o); scrim.classList.toggle('open',o);
    if(o){ // bring the current problem into view
      const cur = drawer.querySelector('.pitem.on');
      if(cur) cur.scrollIntoView({block:'center'});
    }
  };
  btn.addEventListener('click', () => toggle(!drawer.classList.contains('open')));
  scrim.addEventListener('click', () => toggle(false));
  drawer.querySelectorAll('.pitem').forEach(it => it.addEventListener('click', () => {
    const id = it.dataset.id;
    const prob = window.PROBLEMS[id];
    const local = window.GW_LOCAL && window.GW_LOCAL.connected;
    const backend = prob && typeof gwNeedsBackend === 'function' && gwNeedsBackend(prob);
    // A locked (no-wasm) problem becomes playable when the local runner is
    // connected — it serves every level. Open it instead of showing the lock.
    const playable = prob && (!it.classList.contains('locked') || local);
    if(!playable){
      $$('.console .tab')[0].click();
      const tag = it.querySelector('.badge2');
      const isBackend = backend || (tag && tag.textContent==='backend');
      pResult.innerHTML = '<div class="verdict mut">🔒 '+esc(it.querySelector('.nm').textContent)+' — '
        + (isBackend
            ? 'needs the local runner (real <code>go test</code>).<br>start it with: <code>make dev</code>'
            : 'coming soon.')+'</div>';
      toggle(false); return;
    }
    drawer.querySelectorAll('.pitem').forEach(x => x.classList.remove('on'));
    it.classList.add('on');
    // Persist selection in the URL so a refresh reopens this problem, not the first.
    history.replaceState(null, '', '?p=' + encodeURIComponent(id));
    loadProblem(window.PROBLEMS[id]);
    toggle(false);
  }));
  refreshDrawerSolved();   // mark already-solved puzzles from stored progress
})();

/* ---- init ---- */
(function deepLink(){
  const id = new URLSearchParams(location.search).get('p');
  if(id && window.PROBLEMS[id]) CUR = window.PROBLEMS[id];
})();
loadProblem(CUR);
// When the local runner connects, enable Run/Submit and pull the authoritative
// solved set from SQLite so the "submitted" state survives reloads / browsers.
document.addEventListener('gw-runner', e => {
  if(!(e.detail && e.detail.connected)) return;
  if (typeof gwRenderRunButtons === 'function') gwRenderRunButtons();
  else { runEl.disabled = false; subEl.disabled = false; }
  if(typeof gwLocalSolved === 'function'){
    gwLocalSolved().then(ids => {
      if(!ids || !ids.length){ syncSolvedBadge(); return; }
      try{ ids.forEach(id => window.GWProgress && window.GWProgress.markSolved(id)); }catch(e){}
      // Reflect any newly-known solves in the current view + drawer marks.
      syncSolvedBadge();
      if(CUR && isSolved(CUR.id) && /run your code/.test(pResult.innerHTML)){
        pResult.innerHTML = '<div class="verdict ok">✓ submitted — solved earlier. Run again anytime.</div>';
      }
      refreshDrawerSolved();
    });
  }
});
boot();

// Block copy/cut/context-menu inside Solution blocks (belt-and-braces on top of
// CSS user-select:none) so the answer can be read but not lifted out.
['copy','cut','contextmenu'].forEach(function(ev){
  document.addEventListener(ev, function(e){
    if (e.target.closest && e.target.closest('.nocopy')) e.preventDefault();
  });
});
