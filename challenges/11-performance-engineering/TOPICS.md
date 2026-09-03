# 11 Performance Engineering — Topics to Master

Layout: `11-performance-engineering/<level>/<puzzle>/`, **50 puzzles per level**
(junior, middle, senior, staff) — 200 in total.

The roadmap's 7 performance subtopics are covered inside each level, in
learning-path order, as four blocks.

| Block | Covers | Roadmap subtopics |
|-------|--------|-------------------|
| 1 | Measuring | benchmarking-strategy, cpu-profiling |
| 2 | Memory | memory-profiling, pprof-deep |
| 3 | Contention | mutex-block-profiling, trace-tool |
| 4 | Optimizing | optimization-workflow |

Junior and middle puzzles are **implement-from-scratch** stubs
(`panic("not implemented")`). Senior and staff puzzles are **planted-bug**
fixes: exactly one bug between the `CHANGE CODE` markers. Staff additionally
carries the difficulty dial — many of its tests are graded on scale, with an
allocation ceiling, a wall-clock budget, or `-race` that only the correct
implementation meets.

## Determinism rule

Live profilers and wall-clock timing do not reproduce in CI, so these puzzles
teach performance mechanics through deterministic assertions:

- allocation counts via `testing.AllocsPerRun` with an exact ceiling;
- benchmark harnesses as the subject under test (`b.N` use, `b.ResetTimer`
  placement, sink variables that defeat dead-code elimination);
- pprof and trace puzzles operate on parsed profile *fixture data* — flat vs
  cumulative aggregation, stack folding, sample weighting — as pure functions,
  never a live profiler;
- contention puzzles assert observable behaviour (ordering, lost updates,
  `-race` clean), not timing.

Scope law (GENERATION.md §2): level 11 may rely on everything introduced in
topics 01–10, so concurrency, generics, and error handling are all fair game.

## Junior (50)

allocsonce, allocsperop, amdahlbound, arrayvsslice, benchbytesop,
benchcompare, benchfixedinput, benchloop, benchreportallocs, benchresettimer,
benchsink, benchsubrun, bestofn, boxedintcount, buildervsconcat, byteappend,
copyvsappend, cumsum, cyclebudget, escapebufarg, fieldpad, flatsum,
funcselftime, growonce, hotpathshare, ifaceallocount, joinvsloop, latencyavg,
mapclearreuse, nsperop, opsperiter, preallocmap, preallocslice,
profilefilter, profilelabelsum, profilepercent, profilesort, regressionflag,
sliceheadercopy, sliceresetreuse, speedupratio, stringbytesconv,
structsizeorder, samplecount, throughputcalc, timerbudget, topnhot,
valuevspointer, warmupdrop, zeroalloccheck

## Middle (50)

allocratecalc, arenaslice, atomiccounter, batchflush, blockprofagg,
bucketcounts, bufreuse, callgraphedge, chunkedreader, contentionrank,
fanoutbatch, foldcollapse, gcbudget, gcpausesum, goroutinestates,
heapgrowthrate, histmerge, inlinemark, latencyhist, latencytrim, movingavg,
mutexprofagg, mutexvsatomic, outlierdrop, p50p90p99, pipelinestage,
poolreset, poolwrapper, profilediff, profilemerge, profilenormalize,
profiletotalpct, quantileinterp, ratewindow, ringbufwriter, rwlockread,
schedlatency, scratchbuf, selfvscum, semaphoregate, sharedcounter,
smallmapcache, stackfold, striped, stringinterner, tdigestlite,
tracespanmerge, tracespanoverlap, workerpoolsize, samplesymbolize

## Senior (50, planted bugs)

allocratewindowbug, atomicreadmodifybug, avgofavgbug, batchflushlostbug,
benchsetupallocbug, blockratecountbug, bnassizebug, bnignoredbug,
bufaliasbug, bytesopwrongbug, cachestalebug, chunkreaderboundbug,
counterracebug, cumdoublecountbug, deferinloopbug, diffsignbug,
filterinplacebug, foldjoinbug, gcpauseoverlapbug, histbucketedgebug,
histmergewidthbug, internerkeybug, labelsumkeybug, latencyunitbug,
lockheldcallbug, medianevenbug, mergelostsamplebug, movingavgwindowbug,
outliertrimbug, percentileindexbug, poolgrowbug, poolnoresetbug,
profilepctbasebug, quantileinterpbug, ratewindowdivbug,
reportallocsscopebug, resetlenbug, resettimerskipbug, ringoverwritebug,
scratchsharebug, selftimeleakbug, semaphorereleasebug, sinkmissingbug,
spanendbeforestartbug, subbenchsharedstatebug, symbolizeoffbug,
timerbeforesetupbug, topntiebug, warmupcountedbug, workerleakbug

## Staff (50, planted bugs, graded on scale)

allocceilingbug, arenareusebug, atomicpaddingbug, batchsizescalebug,
blockprofscalebug, byteslicereusebug, callgraphcyclebug, chunkedcopybug,
contentionrankstablebug, copyonwritealiasbug, cumcyclecountbug,
diffbaselinealiasbug, falsesharingbug, fanoutcancelbug, flatmapresizebug,
foldquadraticbug, gcbudgetscalebug, growthdoublingbug, histdynamicbucketbug,
histoverflowbug, ifaceboxhotloopbug, labelcardinalitybug,
latencyringcapbug, lockgranularitybug, mapkeyallocbug, mergesamplescalebug,
mutexprofratebug, percentiletiebreakbug, pipelinebackpressurebug,
poolcapleakbug, poolputaliasbug, preallocignoredbug, profileindexrebuildbug,
quantilebignbug, quadraticappendbug, ratecounterwrapbug, rwlockstarvebug,
scratchescapebug, semaphorefairbug, sketchprecisionbug, spinbackoffbug,
streamingmedianbug, stringbuildergrowbug, stripedhashbug,
symbolcachekeybug, tdigestmergebug, timewindowevictbug, topnheapbug,
workerpoolresizebug, zerocopyslicebug

## Progress

| Level | Authored | Target | Mode |
|-------|----------|--------|------|
| junior | 50 | 50 | implement stub |
| middle | 50 | 50 | implement stub |
| senior | 10 | 50 | planted bug |
| staff | 0 | 50 | planted bug + dial |
