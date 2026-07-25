/* Line-style topic icons (24x24, stroke=currentColor). Shared by the home
   explore cards and the roadmap nodes. */
window.ICONS = {
  intro:    '<rect x="3" y="4" width="18" height="16" rx="2"/><path d="M7 9l3 3-3 3M13 15h4"/>',
  basics:   '<path d="M12 2l8 4.5v9L12 20l-8-4.5v-9z"/><path d="M12 11l8-4.5M12 11v9M12 11L4 6.5"/>',
  slices:   '<path d="M8 4H5v16h3M16 4h3v16h-3"/><path d="M9 9h6M9 12h6M9 15h6"/>',
  flow:     '<circle cx="6" cy="6" r="2"/><circle cx="6" cy="18" r="2"/><circle cx="18" cy="12" r="2"/><path d="M6 8v8M6 12h8"/>',
  methods:  '<path d="M9 2v6M15 2v6"/><rect x="7" y="8" width="10" height="6" rx="2"/><path d="M12 14v4a2 2 0 01-2 2H8"/>',
  generics: '<path d="M8 7l-4 5 4 5M16 7l4 5-4 5M13 5l-2 14"/>',
  errors:   '<path d="M12 3l9 16H3z"/><path d="M12 10v4M12 17h.01"/>',
  codeorg:  '<path d="M3 6a2 2 0 012-2h4l2 2h8a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2z"/>',
  conc:     '<circle cx="12" cy="12" r="3"/><path d="M12 2v3M12 19v3M2 12h3M19 12h3M5 5l2 2M17 17l2 2M19 5l-2 2M7 17l-2 2"/>',
  stdlib:   '<path d="M4 5a2 2 0 012-2h13v16H6a2 2 0 00-2 2z"/><path d="M4 5v14"/>',
  testing:  '<path d="M9 3h6M10 3v6l-5 9a2 2 0 002 3h10a2 2 0 002-3l-5-9V3"/><path d="M8 15h8"/>',
  perf:     '<path d="M13 2L4 14h7l-1 8 9-12h-7z"/>',
  patterns: '<circle cx="7" cy="7" r="3"/><rect x="13" y="4" width="7" height="7" rx="1"/><path d="M7 14l4 6H3z"/>',
  runtime:  '<rect x="7" y="7" width="10" height="10" rx="1"/><path d="M9 2v3M15 2v3M9 19v3M15 19v3M2 9h3M2 15h3M19 9h3M19 15h3"/>',
  advanced: '<path d="M12 3l2.5 6H21l-5 4 2 7-6-4-6 4 2-7-5-4h6.5z"/>',
  maps:     '<circle cx="7" cy="8" r="3"/><path d="M9 10l7 7M13 14l3-3 2 2M4 20h16"/>',
  channels: '<rect x="3" y="8" width="18" height="8" rx="2"/><path d="M8 12h8M13 9l3 3-3 3"/>',
  toolchain:'<path d="M14 7a4 4 0 00-5 5l-6 6 2 2 6-6a4 4 0 005-5l-2 2-2-.5-.5-2z"/>',
  modern:   '<path d="M12 3l2 5 5 2-5 2-2 5-2-5-5-2 5-2z"/><path d="M19 15l1 2 2 1-2 1-1 2-1-2-2-1 2-1z"/>',
  observ:   '<circle cx="12" cy="12" r="3"/><path d="M2 12s3.5-7 10-7 10 7 10 7-3.5 7-10 7-10-7-10-7z"/>',
  source:   '<path d="M4 4h11l5 5v11H4z"/><path d="M15 4v5h5M8 13l-2 2 2 2M12 13l2 2-2 2"/>',
  wasm:     '<rect x="3" y="4" width="18" height="16" rx="2"/><path d="M7 10l-2 2 2 2M11 9l-1 6M14 10l2 2-2 2"/>',
};

window.iconHTML = (id, cls) =>
  '<svg viewBox="0 0 24 24" class="ico '+(cls||'')+'" fill="none" stroke="currentColor" '
  + 'stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">'
  + (window.ICONS[id] || window.ICONS.slices) + '</svg>';

/* real SVG element (for nesting inside the roadmap <svg>) */
window.iconSVG = id => {
  const doc = new DOMParser().parseFromString(window.iconHTML(id), 'image/svg+xml');
  return doc.documentElement;
};
