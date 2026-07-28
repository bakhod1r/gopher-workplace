/* Gopher Workplace — solved-state + progress derived from real data.
   Truth: window.CATALOG (topics + items) from problems.js.
   Solved challenge ids persist in localStorage under gw-solved. */
(function(){
  const KEY = 'gw-solved';   // solved challenge ids
  const DAYS = 'gw-days';    // set of YYYY-MM-DD days with at least one solve

  const dayStr = d => d.toISOString().slice(0,10);

  function solvedSet(){
    try{ return new Set(JSON.parse(localStorage.getItem(KEY) || '[]')); }
    catch(e){ return new Set(); }
  }
  function daysSet(){
    try{ return new Set(JSON.parse(localStorage.getItem(DAYS) || '[]')); }
    catch(e){ return new Set(); }
  }
  function markSolved(id){
    if(!id) return;
    const d = daysSet(); d.add(dayStr(new Date()));
    localStorage.setItem(DAYS, JSON.stringify([...d]));
    const s = solvedSet(); if(s.has(id)) return;
    s.add(id); localStorage.setItem(KEY, JSON.stringify([...s]));
  }
  function isSolved(id){ return solvedSet().has(id); }

  // Consecutive-day streak ending today (or yesterday, so it survives until
  // end of the next day). Counts back day-by-day through gw-days.
  function streak(){
    const d = daysSet(); if(!d.size) return 0;
    const cur = new Date(); cur.setHours(0,0,0,0);
    if(!d.has(dayStr(cur))){                 // no solve today — allow yesterday
      cur.setDate(cur.getDate()-1);
      if(!d.has(dayStr(cur))) return 0;
    }
    let n=0;
    while(d.has(dayStr(cur))){ n++; cur.setDate(cur.getDate()-1); }
    return n;
  }

  // Solved / total per difficulty level (J/M/S/ST) across the whole catalog.
  function byLevel(){
    const solved = solvedSet();
    const acc = { J:{done:0,total:0}, M:{done:0,total:0}, S:{done:0,total:0}, ST:{done:0,total:0} };
    (window.CATALOG || []).forEach(g => g.items.forEach(it => {
      const b = acc[it.lv] || acc.J;
      b.total++; if(solved.has(it.id)) b.done++;
    }));
    return acc;
  }

  // Flat list of all challenges with topic + status, catalog order.
  function all(){
    const solved = solvedSet(); const out=[];
    (window.CATALOG || []).forEach(g => g.items.forEach(it => out.push({
      id: it.id, name: it.name, sub: it.sub||'', topic: g.topic,
      lv: it.lv||'J', locked: !!it.locked, tag: it.tag||'',
      status: solved.has(it.id) ? 'solved' : (it.locked ? 'locked' : 'todo'),
      num: out.length+1,
    })));
    return out;
  }

  // Per-topic stats derived from CATALOG + solved state.
  function topics(){
    const solved = solvedSet();
    return (window.CATALOG || []).map(grp => {
      const total   = grp.items.length;
      const open    = grp.items.filter(it => !it.locked).length;
      const done    = grp.items.filter(it => solved.has(it.id)).length;
      const next     = grp.items.find(it => !it.locked && !solved.has(it.id))
                    || grp.items.find(it => !it.locked) || grp.items[0];
      return { topic: grp.topic, items: grp.items, total, open, done, next,
               locked: open === 0, pct: total ? Math.round(done/total*100) : 0 };
    });
  }
  function overall(){
    const t = topics();
    const total = t.reduce((a,x)=>a+x.total,0);
    const done  = t.reduce((a,x)=>a+x.done,0);
    const open  = t.reduce((a,x)=>a+x.open,0);
    return { total, done, open, pct: total ? Math.round(done/total*100) : 0 };
  }

  // SQLite is the source of truth for solved puzzles; localStorage is only a
  // cache (and is per-browser-profile, so it silently looks empty in a new one).
  // Every page — dashboard, problemset, roadmap, playground — merges the
  // backend set as soon as the runner connects, then fires gw-solved so
  // whatever is already rendered can refresh itself.
  function hydrate(){
    if(typeof gwLocalSolved !== 'function') return;
    gwLocalSolved().then(ids => {
      if(!ids || !ids.length) return;
      const before = solvedSet();
      const merged = new Set([...before, ...ids]);
      if(merged.size !== before.size)
        localStorage.setItem(KEY, JSON.stringify([...merged]));
      document.dispatchEvent(new CustomEvent('gw-solved', { detail: [...merged] }));
    });
  }
  document.addEventListener('gw-runner', e => {
    if(e.detail && e.detail.connected) hydrate();
  });

  window.GWProgress = { markSolved, isSolved, solvedSet, topics, overall,
                        streak, byLevel, all, hydrate };
})();
