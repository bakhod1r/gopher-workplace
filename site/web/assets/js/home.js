/* Gopher Workplace — LeetCode-style home dashboard. Everything below is
   derived from real data: window.CATALOG + solved/streak state (GWProgress). */
(function(){
function render(){
  const esc = s => String(s==null?'':s).replace(/[&<>"]/g,m=>({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;'}[m]));
  const P = window.GWProgress;
  if(!window.CATALOG || !P) return;

  const LV = { J:{k:'Junior',c:'#00af9b'}, M:{k:'Middle',c:'#1a90ff'},
               S:{k:'Senior',c:'#8b5cf6'}, ST:{k:'Staff',c:'#f59e0b'} };
  const lvOf = l => LV[l] || LV.J;

  const iconFor = (name, sub, topic) => {
    const t = ((name||'')+' '+(sub||'')+' '+(topic||'')).toLowerCase();
    if(/race|concurren|goroutine|collect/.test(t)) return 'conc';
    if(/channel/.test(t))         return 'channels';
    if(/map/.test(t))             return 'maps';
    if(/leak|gc|pointer/.test(t)) return 'runtime';
    if(/chunk|batch/.test(t))     return 'stdlib';
    if(/slice|dedupe/.test(t))    return 'slices';
    if(/error/.test(t))           return 'errors';
    if(/generic/.test(t))         return 'generics';
    if(/test|bench/.test(t))      return 'testing';
    if(/perf|alloc/.test(t))      return 'perf';
    if(/const|variab|basic|type|limit/.test(t)) return 'basics';
    return 'slices';
  };

  const ov = P.overall(), lvs = P.byLevel(), groups = P.topics();

  /* ---- progress ring (SVG donut) ---- */
  (function(){
    const r=52, C=2*Math.PI*r, off=C*(1-ov.pct/100);
    document.getElementById('ring').innerHTML =
      '<svg viewBox="0 0 130 130" class="donut">'
      + '<circle class="track" cx="65" cy="65" r="'+r+'"/>'
      + '<circle class="fill" cx="65" cy="65" r="'+r+'" '
      +   'stroke-dasharray="'+C.toFixed(1)+'" stroke-dashoffset="'+off.toFixed(1)+'"/>'
      + '<text x="65" y="60" class="pct">'+ov.pct+'%</text>'
      + '<text x="65" y="80" class="sub">solved</text>'
      + '</svg>';
    document.getElementById('ringmeta').innerHTML =
        '<div class="rm-big">'+ov.done+' <span>/ '+ov.total+'</span></div>'
      + '<div class="rm-row"><span class="dot ok"></span>'+ov.done+' solved</div>'
      + '<div class="rm-row"><span class="dot open"></span>'+(ov.open-ov.done)+' playable</div>'
      + '<div class="rm-row"><span class="dot lock"></span>'+(ov.total-ov.open)+' locked</div>';
  })();

  /* ---- difficulty breakdown bars ---- */
  document.getElementById('diffs').innerHTML = ['J','M','S','ST'].map(l=>{
    const b=lvs[l], pct=b.total?Math.round(b.done/b.total*100):0, c=lvOf(l).c;
    return '<div class="diff">'
      + '<div class="d-top"><span class="d-lbl" style="color:'+c+'">'+lvOf(l).k+'</span>'
      +   '<span class="d-n">'+b.done+' / '+b.total+'</span></div>'
      + '<div class="d-bar"><i style="width:'+pct+'%;background:'+c+'"></i></div></div>';
  }).join('');

  /* ---- streak + continue ---- */
  const strk = P.streak();
  document.getElementById('streakNum').textContent = strk;
  let next=null, nextTopic='';
  for(const g of groups){ const c=g.items.find(it=>!it.locked && !P.isSolved(it.id));
    if(c){ next=c; nextTopic=g.topic; break; } }
  if(!next) for(const g of groups){ const c=g.items.find(it=>!it.locked);
    if(c){ next=c; nextTopic=g.topic; break; } }
  const cbtn=document.getElementById('continueBtn');
  if(next){ cbtn.href='playground.html?p='+encodeURIComponent(next.id);
    cbtn.textContent='▸ '+(P.isSolved(next.id)?'revisit':'continue'); }

  /* ---- study-plan / topic cards ---- */
  document.getElementById('plans').innerHTML = groups.map(g=>{
    const first=g.items[0]||{}, ico=iconFor(first.name,first.sub,g.topic);
    const lvSet=[...new Set(g.items.map(i=>i.lv))];
    const c = lvOf(lvSet[0]||'J').c;
    return '<a class="plan" href="problemset.html" style="--ac:'+c+'">'
      + '<div class="plan-ic" data-ico="'+ico+'"></div>'
      + '<div class="plan-b">'
      +   '<div class="plan-t">'+esc(g.topic)+'</div>'
      +   '<div class="plan-m">'+g.done+' / '+g.total+' solved · '+g.open+' playable</div>'
      +   '<div class="plan-bar"><i style="width:'+g.pct+'%"></i></div>'
      + '</div></a>';
  }).join('');

  /* ---- problems table preview (first 8) ---- */
  const rows = P.all().slice(0,8);
  const stCell = s => s==='solved'?'<span class="st-ico ok" title="Solved">✓</span>'
                    : s==='locked'?'<span class="st-ico lock" title="Locked">🔒</span>'
                    : '<span class="st-ico todo" title="Todo">▫</span>';
  document.getElementById('ptbody').innerHTML = rows.map(r=>{
    const lv=lvOf(r.lv), open=r.status!=='locked';
    const title = open
      ? '<a class="ti-link" href="playground.html?p='+encodeURIComponent(r.id)+'">'+r.num+'. '+esc(r.name)+'</a>'
      : '<span class="ti-link locked">'+r.num+'. '+esc(r.name)+'</span>';
    return '<tr class="'+r.status+'">'
      + '<td class="c-st">'+stCell(r.status)+'</td>'
      + '<td class="c-ti">'+title+(r.sub?'<span class="ti-sub">'+esc(r.sub)+'</span>':'')+'</td>'
      + '<td class="c-tp">'+esc(r.topic)+'</td>'
      + '<td class="c-df"><span class="df" style="color:'+lv.c+';border-color:'+lv.c+'">'+lv.k+'</span></td>'
      + '</tr>';
  }).join('');

  /* paint icons (once per element — render() runs again on a solved-set sync) */
  document.querySelectorAll('[data-ico]').forEach(e=>{
    if(e.querySelector('svg')) return;
    e.insertAdjacentHTML('afterbegin', window.iconHTML(e.dataset.ico,'cover-ico'));
  });
}
render();
// The runner's solved set arrives after load; repaint with the real numbers.
document.addEventListener('gw-solved', render);
})();
