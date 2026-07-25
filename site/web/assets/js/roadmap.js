/* Gopher Workplace roadmap — a NeetCode-style dependency graph of Go topics. */

/* theme (shared behaviour with app.js) */
(function themeInit(){
  const sel = document.getElementById('theme');
  const saved = localStorage.getItem('gw-theme') || 'light';
  document.documentElement.setAttribute('data-theme', saved);
  if(sel){ sel.value = saved;
    sel.addEventListener('change', () => {
      document.documentElement.setAttribute('data-theme', sel.value);
      localStorage.setItem('gw-theme', sel.value);
    });
  }
})();

/* node: id,label,x,y (center),status(done|active|locked),link */
const N = [
  { id:"intro", label:"introduction", ic:"intro", x:450, y:56, status:"done" },
  { id:"language-basics", label:"language basics", ic:"basics", x:450, y:142, status:"active", link:"playground.html" },
  { id:"methods-and-interfaces", label:"methods & interfaces", ic:"methods", x:450, y:228, status:"locked" },
  { id:"error-handling", label:"error handling", ic:"errors", x:450, y:314, status:"locked" },
  { id:"standard-library", label:"standard library", ic:"stdlib", x:450, y:400, status:"locked" },
  { id:"code-organization", label:"code organization", ic:"codeorg", x:450, y:486, status:"locked" },
  { id:"go-toolchain", label:"go toolchain", ic:"toolchain", x:450, y:572, status:"locked" },
  { id:"testing-and-benchmarking", label:"testing", ic:"testing", x:450, y:658, status:"locked" },
  { id:"modern-language-features", label:"modern features", ic:"modern", x:450, y:744, status:"locked" },
  { id:"generics", label:"generics", ic:"generics", x:450, y:830, status:"locked" },
  { id:"concurrency", label:"concurrency", ic:"conc", x:450, y:916, status:"locked" },
  { id:"web-development", label:"web development", ic:"channels", x:450, y:1002, status:"locked" },
  { id:"design-patterns-in-go", label:"design patterns", ic:"patterns", x:450, y:1088, status:"locked" },
  { id:"observability", label:"observability", ic:"observ", x:450, y:1174, status:"locked" },
  { id:"performance-engineering", label:"performance", ic:"perf", x:450, y:1260, status:"locked" },
  { id:"advanced-topics", label:"advanced topics", ic:"advanced", x:450, y:1346, status:"locked" },
  { id:"runtime-and-internals", label:"runtime internals", ic:"runtime", x:450, y:1432, status:"locked" },
  { id:"go-source-reading", label:"source reading", ic:"source", x:450, y:1518, status:"locked" },
  { id:"webassembly", label:"webassembly", ic:"wasm", x:450, y:1604, status:"locked" },
];
const LESSONS = {
  "language-basics": [
    { name:"slices: dedupe", lv:"junior", status:"active", link:"playground.html" },
    { name:"slices: chunk", lv:"middle", status:"locked" },
    { name:"slices: in-place leak", lv:"senior", status:"locked" },
    { name:"slices: concurrent collect", lv:"staff", status:"locked" },
    { name:"plan limits", lv:"junior", status:"active", link:"playground.html" },
    { name:"variables and constants", lv:"junior", status:"locked" },
    { name:"data types", lv:"junior", status:"locked" },
    { name:"composite types", lv:"junior", status:"locked" },
    { name:"conditionals", lv:"junior", status:"locked" },
    { name:"loops", lv:"junior", status:"locked" },
    { name:"functions", lv:"junior", status:"locked" },
    { name:"pointers", lv:"junior", status:"locked" },
    { name:"arrays", lv:"junior", status:"locked" },
    { name:"+3 more subtopics", lv:"", status:"locked" }
  ],
  "methods-and-interfaces": [
    { name:"methods vs functions", lv:"junior", status:"locked" },
    { name:"pointer receivers", lv:"junior", status:"locked" },
    { name:"value receivers", lv:"junior", status:"locked" },
    { name:"interfaces basics", lv:"junior", status:"locked" },
    { name:"empty interfaces", lv:"junior", status:"locked" },
    { name:"embedding interfaces", lv:"middle", status:"locked" },
    { name:"type assertions", lv:"middle", status:"locked" },
    { name:"type switch", lv:"middle", status:"locked" },
    { name:"+11 more subtopics", lv:"", status:"locked" }
  ],
  "error-handling": [
    { name:"error handling basics", lv:"junior", status:"locked" },
    { name:"error interface", lv:"junior", status:"locked" },
    { name:"errors new", lv:"junior", status:"locked" },
    { name:"fmt errorf", lv:"junior", status:"locked" },
    { name:"wrapping unwrapping errors", lv:"middle", status:"locked" },
    { name:"sentinel errors", lv:"middle", status:"locked" },
    { name:"panic and recover", lv:"middle", status:"locked" },
    { name:"stack traces debugging", lv:"senior", status:"locked" },
    { name:"+5 more subtopics", lv:"", status:"locked" }
  ],
  "standard-library": [
    { name:"io and file handling", lv:"junior", status:"locked" },
    { name:"flag", lv:"junior", status:"locked" },
    { name:"time", lv:"junior", status:"locked" },
    { name:"encoding json", lv:"junior", status:"locked" },
    { name:"os", lv:"junior", status:"locked" },
    { name:"fmt", lv:"junior", status:"locked" },
    { name:"strings bytes", lv:"junior", status:"locked" },
    { name:"strconv", lv:"junior", status:"locked" },
    { name:"+13 more subtopics", lv:"", status:"locked" }
  ],
  "code-organization": [
    { name:"modules and dependencies", lv:"junior", status:"locked" },
    { name:"packages", lv:"junior", status:"locked" },
    { name:"package import rules", lv:"junior", status:"locked" },
    { name:"project layout", lv:"junior", status:"locked" },
    { name:"internal packages", lv:"middle", status:"locked" },
    { name:"publishing modules", lv:"middle", status:"locked" },
    { name:"workspaces", lv:"middle", status:"locked" },
    { name:"dependency injection", lv:"middle", status:"locked" },
    { name:"+3 more subtopics", lv:"", status:"locked" }
  ],
  "go-toolchain": [
    { name:"core go commands", lv:"junior", status:"locked" },
    { name:"code quality and analysis", lv:"junior", status:"locked" },
    { name:"code generation build tags", lv:"middle", status:"locked" },
    { name:"security", lv:"middle", status:"locked" },
    { name:"deployment and tooling", lv:"middle", status:"locked" },
    { name:"go work", lv:"middle", status:"locked" },
    { name:"go tool suite", lv:"middle", status:"locked" },
    { name:"live reload", lv:"middle", status:"locked" },
    { name:"+3 more subtopics", lv:"", status:"locked" }
  ],
  "testing-and-benchmarking": [
    { name:"testing basics", lv:"junior", status:"locked" },
    { name:"table driven tests", lv:"junior", status:"locked" },
    { name:"coverage", lv:"junior", status:"locked" },
    { name:"subtests", lv:"junior", status:"locked" },
    { name:"mocks and stubs", lv:"middle", status:"locked" },
    { name:"httptest", lv:"middle", status:"locked" },
    { name:"benchmarks", lv:"middle", status:"locked" },
    { name:"testmain", lv:"middle", status:"locked" },
    { name:"+9 more subtopics", lv:"", status:"locked" }
  ],
  "modern-language-features": [
    { name:"loopvar semantics", lv:"junior", status:"locked" },
    { name:"min max clear builtins", lv:"junior", status:"locked" },
    { name:"iterators and range over func", lv:"middle", status:"locked" },
    { name:"modern stdlib additions", lv:"middle", status:"locked" },
    { name:"generic type aliases", lv:"senior", status:"locked" }
  ],
  "generics": [
    { name:"why generics", lv:"middle", status:"locked" },
    { name:"generic functions", lv:"middle", status:"locked" },
    { name:"generic types interfaces", lv:"middle", status:"locked" },
    { name:"type constraints", lv:"middle", status:"locked" },
    { name:"type inference", lv:"middle", status:"locked" },
    { name:"generic constraints deep", lv:"senior", status:"locked" },
    { name:"generic performance", lv:"senior", status:"locked" },
    { name:"generics vs interfaces", lv:"senior", status:"locked" },
    { name:"+9 more subtopics", lv:"", status:"locked" }
  ],
  "concurrency": [
    { name:"goroutines", lv:"middle", status:"locked" },
    { name:"channels", lv:"middle", status:"locked" },
    { name:"select and buffering", lv:"middle", status:"locked" },
    { name:"sync package", lv:"middle", status:"locked" },
    { name:"context package", lv:"middle", status:"locked" },
    { name:"worker pools", lv:"middle", status:"locked" },
    { name:"time based concurrency", lv:"middle", status:"locked" },
    { name:"concurrency patterns", lv:"senior", status:"locked" },
    { name:"+20 more subtopics", lv:"", status:"locked" }
  ],
  "web-development": [
    { name:"building clis", lv:"middle", status:"locked" },
    { name:"net http server", lv:"middle", status:"locked" },
    { name:"routing and handlers", lv:"middle", status:"locked" },
    { name:"rest api design", lv:"middle", status:"locked" },
    { name:"orms and db access", lv:"middle", status:"locked" },
    { name:"middleware and context", lv:"senior", status:"locked" },
    { name:"grpc and protobuf", lv:"senior", status:"locked" },
    { name:"realtime communication", lv:"senior", status:"locked" },
    { name:"+1 more subtopics", lv:"", status:"locked" }
  ],
  "design-patterns-in-go": [
    { name:"functional options", lv:"middle", status:"locked" },
    { name:"builder pattern", lv:"middle", status:"locked" },
    { name:"strategy pattern", lv:"middle", status:"locked" },
    { name:"decorator pattern", lv:"middle", status:"locked" },
    { name:"adapter pattern", lv:"middle", status:"locked" },
    { name:"factory pattern", lv:"middle", status:"locked" },
    { name:"observer pattern", lv:"middle", status:"locked" },
    { name:"singleton pattern", lv:"middle", status:"locked" },
    { name:"+12 more subtopics", lv:"", status:"locked" }
  ],
  "observability": [
    { name:"runtime metrics package", lv:"middle", status:"locked" },
    { name:"expvar", lv:"middle", status:"locked" },
    { name:"runtime trace application tracing", lv:"senior", status:"locked" },
    { name:"opentelemetry in go", lv:"senior", status:"locked" },
    { name:"godebug and runtime debug", lv:"senior", status:"locked" }
  ],
  "performance-engineering": [
    { name:"cpu profiling", lv:"senior", status:"locked" },
    { name:"memory profiling", lv:"senior", status:"locked" },
    { name:"mutex block profiling", lv:"senior", status:"locked" },
    { name:"benchmarking strategy", lv:"senior", status:"locked" },
    { name:"optimization workflow", lv:"senior", status:"locked" },
    { name:"pprof deep", lv:"staff", status:"locked" },
    { name:"trace tool", lv:"staff", status:"locked" }
  ],
  "advanced-topics": [
    { name:"memory management in depth", lv:"senior", status:"locked" },
    { name:"escape analysis", lv:"senior", status:"locked" },
    { name:"reflection", lv:"senior", status:"locked" },
    { name:"unsafe package", lv:"senior", status:"locked" },
    { name:"build constraints", lv:"senior", status:"locked" },
    { name:"cgo basics", lv:"senior", status:"locked" },
    { name:"serverless go", lv:"senior", status:"locked" },
    { name:"compiler linker flags", lv:"staff", status:"locked" },
    { name:"+8 more subtopics", lv:"", status:"locked" }
  ],
  "runtime-and-internals": [
    { name:"runtime source dive", lv:"senior", status:"locked" },
    { name:"runtime package deep", lv:"senior", status:"locked" },
    { name:"scheduler source", lv:"staff", status:"locked" },
    { name:"gc source", lv:"staff", status:"locked" },
    { name:"memory allocator", lv:"staff", status:"locked" },
    { name:"go runtime architecture", lv:"staff", status:"locked" }
  ],
  "go-source-reading": [
    { name:"net http source", lv:"senior", status:"locked" },
    { name:"sync source", lv:"senior", status:"locked" },
    { name:"context source", lv:"senior", status:"locked" },
    { name:"database sql source", lv:"senior", status:"locked" },
    { name:"runtime source", lv:"staff", status:"locked" },
    { name:"encoding json source", lv:"staff", status:"locked" }
  ],
  "webassembly": [
    { name:"goos js wasm browser", lv:"senior", status:"locked" },
    { name:"wasi and wasip1", lv:"senior", status:"locked" },
    { name:"tinygo for wasm and embedded", lv:"staff", status:"locked" },
    { name:"wasm interop and performance", lv:"staff", status:"locked" },
    { name:"wasm in production", lv:"staff", status:"locked" }
  ],
};
const E = [
  ["intro","language-basics"], ["language-basics","methods-and-interfaces"], ["methods-and-interfaces","error-handling"], ["error-handling","standard-library"], ["standard-library","code-organization"], ["code-organization","go-toolchain"], ["go-toolchain","testing-and-benchmarking"], ["testing-and-benchmarking","modern-language-features"], ["modern-language-features","generics"], ["generics","concurrency"], ["concurrency","web-development"], ["web-development","design-patterns-in-go"], ["design-patterns-in-go","observability"], ["observability","performance-engineering"], ["performance-engineering","advanced-topics"], ["advanced-topics","runtime-and-internals"], ["runtime-and-internals","go-source-reading"], ["go-source-reading","webassembly"]
];
const W = 900, H = 1664, BW = 240, BH = 50;
const byId = Object.fromEntries(N.map(n => [n.id, n]));
const svgNS = "http://www.w3.org/2000/svg";
const el = (t, a={}) => { const e=document.createElementNS(svgNS,t); for(const k in a) e.setAttribute(k,a[k]); return e; };

function build(){
  const svg = el("svg", { class:"rm-svg", width:W, height:H, viewBox:`0 0 ${W} ${H}` });

  // edges first
  for(const [a,b] of E){
    const p=byId[a], c=byId[b];
    const x1=p.x, y1=p.y+BH/2, x2=c.x, y2=c.y-BH/2, my=(y1+y2)/2;
    const hot = (p.status!=="locked" && c.status!=="locked") || p.status==="active" || c.status==="active";
    svg.appendChild(el("path", { class:"edge"+(hot?" hot":""), d:`M${x1},${y1} C${x1},${my} ${x2},${my} ${x2},${y2}` }));
  }
  // nodes
  for(const n of N){
    const g = el("g", { class:"node "+n.status, "data-id":n.id });
    g.appendChild(el("rect", { x:n.x-BW/2, y:n.y-BH/2, width:BW, height:BH, rx:9 }));
    const ic = window.iconSVG(n.ic||n.id);
    ic.setAttribute("x", n.x-BW/2+13); ic.setAttribute("y", n.y-9);
    ic.setAttribute("width", 18); ic.setAttribute("height", 18);
    ic.setAttribute("class", "nico");
    g.appendChild(ic);
    const lbl = el("text", { class:"lbl", x:n.x-BW/2+39, y:n.y+5 }); lbl.textContent=n.label;
    g.appendChild(lbl);
    g.addEventListener("click", e => { e.stopPropagation(); hideTip(); openMenu(e, n); });
    g.addEventListener("mousemove", e => { if(!menu.classList.contains("open")) showTip(e, n); });
    g.addEventListener("mouseleave", hideTip);
    svg.appendChild(g);
  }
  document.getElementById("rm-graph").appendChild(svg);
}

/* ---- lesson popover ---- */
const menu = document.getElementById("rm-menu");
const esc = s => String(s).replace(/[&<>]/g, m => ({'&':'&amp;','<':'&lt;','>':'&gt;'}[m]));

function openMenu(e, n){
  const items = LESSONS[n.id] || [];
  let body = items.length
    ? '<ul class="rm-list">' + items.map((it,i) =>
        '<li class="'+it.status+'" data-link="'+(it.link||'')+'" data-i="'+i+'">'
        + '<span class="st">'+(it.status==="done"?"✓":it.status==="active"?"▸":"▫")+'</span>'
        + '<span class="nm">'+esc(it.name)+'</span>'
        + '<span class="lv">'+esc(it.lv||"")+'</span>'
        + (it.tag?'<span class="tag">'+esc(it.tag)+'</span>':'')
        + '</li>').join('') + '</ul>'
    : '<div class="rm-empty">lessons coming soon</div>';

  menu.innerHTML = '<div class="mh"><span class="tt">'+esc(n.label)+'</span><span class="x" data-close>✕</span></div>' + body;
  menu.classList.add("open");
  // position near cursor, clamped to viewport
  const mw = menu.offsetWidth, mh = menu.offsetHeight;
  let x = Math.min(e.clientX+10, window.innerWidth - mw - 12);
  let y = Math.min(e.clientY+10, window.innerHeight - mh - 12);
  menu.style.left = Math.max(12,x)+"px";
  menu.style.top  = Math.max(56,y)+"px";

  menu.querySelector("[data-close]").addEventListener("click", closeMenu);
  menu.querySelectorAll(".rm-list li").forEach(li => li.addEventListener("click", () => {
    if(li.classList.contains("locked")) return;
    const link = li.getAttribute("data-link");
    if(link) location.href = link;
  }));
}
function closeMenu(){ menu.classList.remove("open"); }
document.addEventListener("click", e => { if(!menu.contains(e.target)) closeMenu(); });
window.addEventListener("keydown", e => { if(e.key==="Escape") closeMenu(); });

const tip = document.getElementById("rm-tip");
function showTip(e, n){
  const msg = n.status==="active" ? "in progress — click to open"
            : n.status==="done"   ? "foundation"
            : "locked — finish earlier topics first";
  tip.innerHTML = '<div class="t">'+n.label+'</div><div class="m">'+msg+'</div>';
  tip.style.display="block";
  tip.style.left = Math.min(e.clientX+14, window.innerWidth-250)+"px";
  tip.style.top  = (e.clientY+14)+"px";
}
function hideTip(){ tip.style.display="none"; }

build();
