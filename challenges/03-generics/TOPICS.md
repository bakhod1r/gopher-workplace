# 03 Generics — Topics to Master

Layout: `03-generics/<level>/<puzzle>/`, **50 puzzles per level**
(junior, middle, senior, staff) — 200 in total.

The roadmap's 17 generics subtopics are covered inside each level, in
learning-path order, as four blocks.

| Block | Covers | Roadmap subtopics |
|-------|--------|-------------------|
| 1 | Generic functions | why-generics, generic-functions, type-inference |
| 2 | Type constraints | type-constraints, comparable-and-ordered, generic-constraints-deep, recursive-type-constraints |
| 3 | Generic types and methods | generic-types-interfaces, methods-on-generic-types, generic-data-structures, generic-type-aliases |
| 4 | Generics in practice | stdlib-generic-packages, generics-vs-interfaces, generic-performance, generic-limitations, generic-pitfalls, generic-testing-helpers |

Junior and middle puzzles are **implement-from-scratch** stubs
(`panic("not implemented")`). Senior and staff puzzles are
**planted-bug** fixes: exactly one bug between the `CHANGE CODE` markers.
Staff additionally carries the difficulty dial — many of its tests are
graded on scale, with a wall-clock or allocation ceiling that only the
correct implementation meets.

## Junior (50)

absgen, aliasgen, assertequalhelper, avggen, boxgen, cachegen, chunkgen,
clampgen, cmpcomparegen, comparegen, containsgen, filtergen, firstof,
groupbygen, identityfn, indexofgen, inferenceexplicit, issortedgen, lastof,
listgen, longesttext, mapfn, mapsclonegen, matrixgen, maxof, maxslice,
minmaxbuiltin, minof, optionalgen, pairgen, queuegen, reducegen, reversegen,
ringgen, setgen, slicesbinsearch, slicesclonegen, slicescompact,
slicescontains, slicesmaxmin, slicessortfunc, sortedcopy, stackgen, sumifgen,
sumnum, swapgen, tildeint, treegen, uniquegen, zipwithgen

## Middle (50)

bfsgen, comparableembed, composegen, currygen, divmodgen, eventbusgen,
fieldaccessgen, flatmapgen, gcdgen, graphgen, heapgen, lessergen, lrugen,
mapsdeletefunc, mediangen, memoizegen, mergesortedgen, methodtypeparamgen,
minbygen, minstackgen, nansafegen, numberfull, orderedmapgen, percentilegen,
powintgen, pqgen, preallocgen, queuetwostacks, rotategen, satadd, scangen,
signedonly, slicesbsearchfunc, slicescompactfunc, slicesconcatstd,
slicesdeletefunc, slicesgrowgen, slicesminmaxfunc, tabletestgen, takewhilegen,
tempconv, topngen, transposegen, triegen, typeswitchgen, unionfindgen,
uniquebygen, variancegen, windowgen, zipmapgen

## Senior (50, planted bugs)

aliasbug, bagzerobug, batcherbug, bfsmarkbug, binsearchbug, buildcapbug,
chainmergebug, clipbug, compactbug, concatalias, deepclonebug,
defaultmapzerobug, divbug, dropwhilebug, everynthzerobug, filterinplacebug,
groupbybug, growbug, heapsiftbug, insertboundbug, interleaveremainderbug,
lastindexbug, lazyzerobug, lrupromotebug, mapiterbug, memoizezerobug,
mergetiebug, minstacksyncbug, nanbug, orderedmapdupbug, partitionorderbug,
percentilebug, pqtiebug, queuetwostacksbug, ringstartbug, rotatebug, satbug,
scanbug, seedbug, setopsmutatebug, stablebug, tiebug, topnstablebug,
transposeraggedbug, triesetbug, unionfindbug, uniquekeybug, windowbug,
windowmaxexpirebug, zipmapbug

## Staff (50, planted bugs, graded on scale)

anyboxhotloopbug, asserthelperpassbug, binheapchildbug, bucketsplitcopybug,
clipheaderbug, clone2dsharebug, cmpoverflowbug, dequewrapmodbug,
dsucompressbug, ifacekeyguardbug, intervalboundbug, lesserselfbug,
lowerboundbug, lruo1bug, memokeynormbug, multisetremovebug, mustswallowbug,
nanminmaxbug, orderleakbug, ordmapunlinkbug, percentilerankbug,
pipelinerebuildbug, prealloclenbug, prependquadbug, ptrmethodsetbug,
radixlongestbug, rankvalueonlybug, retainbug, ringcapbug, satclampbug,
scaleroundbug, shallowclonebug, sparsegridkeybug, stdbinsearchcmpbug,
stdcompactadjbug, stdcomparelenbug, stddeletediscardbug, stdequalfuncfieldbug,
stdinsertdriftbug, stdmapsclonealiasbug, stdmapsdeletewalkbug,
stdrepeatsharedbug, stdsortstableweakbug, tablecaseptrbug, tallycopybug,
tildeconvbug, typeswitchdefaultbug, unionprecisionbug, unsigneddeltabug,
versionsnapaliasbug

## Progress

| Level | Authored | Target | Mode |
|-------|----------|--------|------|
| junior | 50 | 50 | implement stub |
| middle | 50 | 50 | implement stub |
| senior | 50 | 50 | planted bug |
| staff | 50 | 50 | planted bug + scale |
