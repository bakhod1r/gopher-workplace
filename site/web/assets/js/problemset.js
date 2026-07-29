/* Gopher Workplace — LeetCode-style problemset list. Data: window.CATALOG
   (from problems.js) + solved state (window.GWProgress). */
(function(){
  const esc = s => String(s==null?'':s).replace(/[&<>"]/g,m=>({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;'}[m]));
  const LV = { J:{k:'Junior',c:'junior'}, M:{k:'Middle',c:'middle'},
               S:{k:'Senior',c:'senior'}, ST:{k:'Staff',c:'staff'} };

  // Flatten CATALOG into a single row list, carrying the topic label.
  const ROWS = [];
  (window.CATALOG || []).forEach(g => g.items.forEach((it, i) => {
    ROWS.push({
      id: it.id, name: it.name, sub: it.sub || '', topic: g.topic,
      lv: it.lv || 'J', locked: !!it.locked, tag: it.tag || '',
      num: ROWS.length + 1,
    });
  }));

  const P = window.GWProgress;
  const statusOf = r => P.isSolved(r.id) ? 'solved' : (r.locked ? 'locked' : 'todo');

  const els = {
    body: document.getElementById('psBody'),
    empty: document.getElementById('psEmpty'),
    search: document.getElementById('psSearch'),
    reset: document.getElementById('psReset'),
    random: document.getElementById('psRandom'),
    stats: document.getElementById('psStats'),
  };

  let SHOWN = ROWS;  // rows currently passing the filters

  // populate the tag multi-select with the subtopics, sorted
  const tagMenu = document.querySelector('#psTag .ps-ms-menu');
  [...new Set(ROWS.map(r => r.sub).filter(Boolean))].sort().forEach(t => {
    const l = document.createElement('label');
    l.innerHTML = '<input type="checkbox"><span></span>';
    l.querySelector('input').value = t;
    l.querySelector('span').textContent = t;
    tagMenu.appendChild(l);
  });

  // Multi-select dropdown: each returns a live Set of checked values and calls
  // onChange whenever the selection changes. Menu closes on outside click.
  function multiSelect(id, onChange){
    const root = document.getElementById(id);
    const btn = root.querySelector('.ps-ms-btn');
    const menu = root.querySelector('.ps-ms-menu');
    const allLabel = root.dataset.all;
    const selected = new Set();
    function label(){
      const boxes = [...menu.querySelectorAll('input:checked')];
      if(!boxes.length){ btn.textContent = allLabel + ' ▾'; root.classList.remove('on'); return; }
      const txt = boxes.length === 1
        ? boxes[0].parentNode.textContent.trim()
        : boxes.length + ' selected';
      btn.textContent = txt + ' ▾'; root.classList.add('on');
    }
    function sync(){
      selected.clear();
      menu.querySelectorAll('input:checked').forEach(b => selected.add(b.value));
      label(); onChange();
    }
    btn.addEventListener('click', e => {
      e.stopPropagation();
      const open = menu.hidden;
      document.querySelectorAll('.ps-ms-menu').forEach(m => { m.hidden = true; });
      menu.hidden = !open; btn.setAttribute('aria-expanded', String(open));
    });
    menu.addEventListener('click', e => e.stopPropagation());
    menu.addEventListener('change', sync);
    root._clear = () => { menu.querySelectorAll('input:checked').forEach(b => b.checked = false); sync(); };
    return selected;
  }
  document.addEventListener('click', () => {
    document.querySelectorAll('.ps-ms-menu').forEach(m => { m.hidden = true; });
    document.querySelectorAll('.ps-ms-btn').forEach(b => b.setAttribute('aria-expanded','false'));
  });

  const selLevel = multiSelect('psLevel', apply);
  const selTag   = multiSelect('psTag', apply);
  const selStatus= multiSelect('psStatus', apply);

  function statusCell(s){
    if(s==='solved') return '<span class="st-ico ok" title="Solved">✓</span>';
    if(s==='locked') return '<span class="st-ico lock" title="Locked">🔒</span>';
    return '<span class="st-ico todo" title="Todo">▫</span>';
  }

  function rowHTML(r){
    const s = statusOf(r), lv = LV[r.lv] || LV.J;
    const clickable = s !== 'locked';
    const href = clickable ? 'playground.html?p='+encodeURIComponent(r.id) : '';
    const title = clickable
      ? '<a class="ti-link" href="'+esc(href)+'">'+r.num+'. '+esc(r.name)+'</a>'
      : '<span class="ti-link locked">'+r.num+'. '+esc(r.name)+'</span>';
    const tags = (r.sub ? '<span class="tg">'+esc(r.sub)+'</span>' : '');
    return '<tr class="'+s+'">'
      + '<td class="c-st">'+statusCell(s)+'</td>'
      + '<td class="c-ti">'+title+'</td>'
      + '<td class="c-tp">'+esc(r.topic)+'</td>'
      + '<td class="c-df"><span class="df '+lv.c+'">'+lv.k+'</span></td>'
      + '<td class="c-tg">'+tags+'</td>'
      + '</tr>';
  }

  function apply(){
    const q = els.search.value.trim().toLowerCase();
    const rows = ROWS.filter(r => {
      if(selLevel.size && !selLevel.has(r.lv)) return false;
      if(selTag.size && !selTag.has(r.sub)) return false;
      if(selStatus.size && !selStatus.has(statusOf(r))) return false;
      if(q && !((r.name+' '+r.sub+' '+r.topic).toLowerCase().includes(q))) return false;
      return true;
    });
    SHOWN = rows;
    els.body.innerHTML = rows.map(rowHTML).join('');
    els.empty.hidden = rows.length > 0;
    const anyTodo = rows.some(r => statusOf(r) === 'todo');
    els.random.disabled = !anyTodo;
  }

  function pickRandom(){
    const pool = SHOWN.filter(r => statusOf(r) === 'todo');
    if(!pool.length) return;
    const r = pool[Math.floor(Math.random() * pool.length)];
    location.href = 'playground.html?p=' + encodeURIComponent(r.id);
  }

  function renderStats(){
    const total = ROWS.length;
    const solved = ROWS.filter(r => statusOf(r)==='solved').length;
    const open = ROWS.filter(r => !r.locked).length;
    els.stats.innerHTML =
        '<span class="s ok">'+solved+' solved</span>'
      + '<span class="s">'+open+' playable</span>'
      + '<span class="s mut">'+total+' total</span>';
  }

  els.search.addEventListener('input', apply);
  els.random.addEventListener('click', pickRandom);
  els.reset.addEventListener('click', () => {
    els.search.value='';
    ['psLevel','psTag','psStatus'].forEach(id => document.getElementById(id)._clear());
    apply();
  });

  renderStats();
  apply();
  // Solved state is authoritative in the runner's SQLite and lands after load.
  document.addEventListener('gw-solved', () => { renderStats(); apply(); });
})();
