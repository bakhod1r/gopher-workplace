# Challenges hierarchy

1:1 mirror of the Go roadmap source tree (roadmap `01-introduction-to-go`
excluded; topics renumbered from 01). Structure follows the roadmap's own
nesting. The **level** dimension (junior · middle · senior · staff) lives
*inside* a leaf, added at generation time.

```
challenges/<NN-topic>/…/<leaf>/<level>/<puzzle>/
```

- Leaf dirs hold a `.keep` until authored.
- Skeleton: `bash scripts/mirror-roadmap.sh`.

**Topics:** 17.

---

## 01-language-basics
*298 dirs · 278 leaves*

- 01-variables-and-constants
  - junior
    - byteunits  ·leaf
    - discard  ·leaf
    - endpoint  ·leaf
    - plan-limits  ·leaf
    - retries  ·leaf
    - shadow  ·leaf
    - swap  ·leaf
    - temperature  ·leaf
    - typedconst  ·leaf
    - zerovalues  ·leaf
  - middle
    - byte-shift  ·leaf
    - clamp  ·leaf
    - frame-budget  ·leaf
    - int-limits  ·leaf
    - iota-perms  ·leaf
    - pi-precision  ·leaf
    - revoke  ·leaf
    - score-tiers  ·leaf
    - temp-scale  ·leaf
    - weekday-iota  ·leaf
  - senior
    - iota-offset  ·leaf
    - kb-shift  ·leaf
    - maxint-signbit  ·leaf
    - percent-order  ·leaf
    - revoke-xor  ·leaf
    - rgb-pack  ·leaf
    - status-zero  ·leaf
    - sum-shadow  ·leaf
    - tier-threshold  ·leaf
    - uint-underflow  ·leaf
  - staff
    - abs-minint  ·leaf
    - fixed-cents  ·leaf
    - float32-precision  ·leaf
    - hex-byte  ·leaf
    - iota-reset  ·leaf
    - kib-scale  ·leaf
    - minint-shift  ·leaf
    - overflow-mul  ·leaf
    - signbit-mask  ·leaf
    - untyped-ratio  ·leaf
- 02-data-types
  - junior
    - almostequal  ·leaf
    - average  ·leaf
    - bytecat  ·leaf
    - c2f  ·leaf
    - checksum  ·leaf
    - clockmod  ·leaf
    - complexabs  ·leaf
    - escapes  ·leaf
    - evenodd  ·leaf
    - firstrune  ·leaf
    - hexdigit  ·leaf
    - majority  ·leaf
    - nancheck  ·leaf
    - narrowing  ·leaf
    - quote  ·leaf
    - rawpath  ·leaf
    - round2  ·leaf
    - runecount  ·leaf
    - truncate  ·leaf
    - vowelcount  ·leaf
  - middle
    - atoi  ·leaf
    - baseconv  ·leaf
    - caesar  ·leaf
    - clamp01  ·leaf
    - csvfield  ·leaf
    - epsiloneq  ·leaf
    - gcd  ·leaf
    - hexencode  ·leaf
    - itoa  ·leaf
    - lerp  ·leaf
    - luhn  ·leaf
    - moneyparse  ·leaf
    - parsehex  ·leaf
    - popcount  ·leaf
    - reverserunes  ·leaf
    - rotate  ·leaf
    - semverparse  ·leaf
    - slugify  ·leaf
    - titlecase  ·leaf
    - utf8truncate  ·leaf
  - senior
    - clampbyte  ·leaf
    - countwords  ·leaf
    - csvsplit  ·leaf
    - durationparse  ·leaf
    - fletcher16  ·leaf
    - hexdecode  ·leaf
    - ipv4parse  ·leaf
    - jsonescape  ·leaf
    - maskcard  ·leaf
    - normspaces  ·leaf
    - ordinal  ·leaf
    - padleft  ·leaf
    - parsebool  ·leaf
    - runeat  ·leaf
    - runewidth  ·leaf
    - splitlines  ·leaf
    - tolowerascii  ·leaf
    - wordwrap  ·leaf
  - staff
    - bankround  ·leaf
    - base64val  ·leaf
    - checkedadd  ·leaf
    - endianswap  ·leaf
    - floatbits  ·leaf
    - fnv1a  ·leaf
    - graycode  ·leaf
    - parity  ·leaf
    - rune2decode  ·leaf
    - utf8count  ·leaf
    - utf8valid  ·leaf
    - varint  ·leaf
- 03-composite-types
  - junior
    - arrayval  ·leaf
    - checkoutgrid  ·leaf
    - concat  ·leaf
    - contains  ·leaf
    - copyslice  ·leaf
    - counter  ·leaf
    - dedupe  ·leaf
    - filterpos  ·leaf
    - flatten  ·leaf
    - getdefault  ·leaf
    - gradebook  ·leaf
    - invert  ·leaf
    - keys  ·leaf
    - lastn  ·leaf
    - lookup  ·leaf
    - mapmerge  ·leaf
    - minmax  ·leaf
    - point  ·leaf
    - rectarea  ·leaf
    - removeindex  ·leaf
    - reverse  ·leaf
    - rowsum  ·leaf
    - structsum  ·leaf
    - structupdate  ·leaf
    - subslice  ·leaf
    - sumints  ·leaf
    - uniquecount  ·leaf
    - wordfreq  ·leaf
  - middle
    - anagramgroup  ·leaf
    - chunk  ·leaf
    - dedupesorted  ·leaf
    - groupby  ·leaf
    - groupstructs  ·leaf
    - histogram  ·leaf
    - insertat  ·leaf
    - joinmanual  ·leaf
    - matmul  ·leaf
    - mergesorted  ·leaf
    - mostcommon  ·leaf
    - movezeros  ·leaf
    - partition  ·leaf
    - reverseinplace  ·leaf
    - rotate  ·leaf
    - runlength  ·leaf
    - setdiff  ·leaf
    - setintersect  ·leaf
    - setunion  ·leaf
    - sortstructs  ·leaf
    - splitmanual  ·leaf
    - structfilter  ·leaf
    - takewhile  ·leaf
    - topk  ·leaf
    - transpose  ·leaf
    - twosum  ·leaf
    - window  ·leaf
    - zip  ·leaf
  - senior
    - appendcapshared  ·leaf
    - appendnotcaptured  ·leaf
    - arrayvalue  ·leaf
    - commaokzero  ·leaf
    - copyemptydst  ·leaf
    - deduporder  ·leaf
    - embeddingshadow  ·leaf
    - emptystructset  ·leaf
    - filterreuse  ·leaf
    - grid2dalias  ·leaf
    - jsontag  ·leaf
    - mapstructupdate  ·leaf
    - mergedirection  ·leaf
    - nestedmapinit  ·leaf
    - nilmap  ·leaf
    - nilvsempty  ·leaf
    - preallocindex  ·leaf
    - rangecopymutate  ·leaf
    - removeorder  ·leaf
    - shallowclone  ·leaf
    - sliceequal  ·leaf
    - sortmutates  ·leaf
    - stackpop  ·leaf
    - subsliceclamp  ·leaf
    - twodshallow  ·leaf
    - valuereceiver  ·leaf
  - staff
    - appendnilreturn  ·leaf
    - appendtwobranch  ·leaf
    - bufconcat  ·leaf
    - clipcap  ·leaf
    - copyshiftleft  ·leaf
    - gridcolsum  ·leaf
    - growthstale  ·leaf
    - insertshift  ·leaf
    - keyfromstruct  ·leaf
    - maparraykey  ·leaf
    - mapclonesharedslice  ·leaf
    - mapniljson  ·leaf
    - ringbuffer  ·leaf
    - rotate90  ·leaf
    - sliceleak  ·leaf
    - slicetoarray  ·leaf
    - stringimmutable  ·leaf
    - structpadding  ·leaf
- 04-functions
  - junior
    - absif  ·leaf
    - averagef  ·leaf
    - cascade  ·leaf
    - clampf  ·leaf
    - clamploop  ·leaf
    - countdown  ·leaf
    - countword  ·leaf
    - dayname  ·leaf
    - divmod  ·leaf
    - doublecopy  ·leaf
    - firstlast  ·leaf
    - fizzbuzz  ·leaf
    - gcdloop  ·leaf
    - grade  ·leaf
    - gradeif  ·leaf
    - gridsum  ·leaf
    - ifinit  ·leaf
    - joinv  ·leaf
    - leapyear  ·leaf
    - max3  ·leaf
    - minmax  ·leaf
    - nocopysliceval  ·leaf
    - reverseloop  ·leaf
    - signswitch  ·leaf
    - sumrange  ·leaf
    - sumv  ·leaf
    - swap2  ·leaf
    - tagless  ·leaf
    - tally  ·leaf
    - weekday  ·leaf
  - middle
    - adder  ·leaf
    - appendall  ·leaf
    - bumpptr  ·leaf
    - compose  ·leaf
    - counter  ·leaf
    - curry  ·leaf
    - deferbump  ·leaf
    - deferloop  ·leaf
    - deferorder  ·leaf
    - defersnapshot  ·leaf
    - dropwhile  ·leaf
    - fallthru  ·leaf
    - filterints  ·leaf
    - firstpair  ·leaf
    - flatten  ·leaf
    - gotoretry  ·leaf
    - groupby  ·leaf
    - mapints  ·leaf
    - maxvar  ·leaf
    - memoize  ·leaf
    - namedret  ·leaf
    - oncefn  ·leaf
    - perindex  ·leaf
    - pipeline  ·leaf
    - reduceints  ·leaf
    - repeatfn  ·leaf
    - safecall  ·leaf
    - takewhile  ·leaf
    - tracker  ·leaf
    - zipwith  ·leaf
  - senior  ·leaf
  - staff  ·leaf
- 05-pointers
  - 01-pointers-basics  ·leaf
  - 02-pointers-with-structs  ·leaf
  - 03-with-maps-and-slices  ·leaf
  - 04-memory-management
    - 01-garbage-collection  ·leaf
  - 05-unsafe-pointer  ·leaf
  - 06-nil-pointer-dereference  ·leaf

## 02-methods-and-interfaces
*19 dirs · 19 leaves*

- 01-methods-vs-functions  ·leaf
- 02-pointer-receivers  ·leaf
- 03-value-receivers  ·leaf
- 04-interfaces-basics  ·leaf
- 05-empty-interfaces  ·leaf
- 06-embedding-interfaces  ·leaf
- 07-type-assertions  ·leaf
- 08-type-switch  ·leaf
- 09-method-sets-deep  ·leaf
- 10-interface-internals  ·leaf
- 11-method-dispatch  ·leaf
- 12-common-interfaces  ·leaf
- 13-interface-best-practices  ·leaf
- 14-interface-anti-patterns  ·leaf
- 15-method-values-and-expressions  ·leaf
- 16-methods-on-defined-types  ·leaf
- 17-sealed-interfaces  ·leaf
- 18-cross-package-methods  ·leaf
- 19-struct-method-promotion  ·leaf

## 03-generics
*17 dirs · 17 leaves*

- 01-why-generics  ·leaf
- 02-generic-functions  ·leaf
- 03-generic-types-interfaces  ·leaf
- 04-type-constraints  ·leaf
- 05-type-inference  ·leaf
- 06-generic-constraints-deep  ·leaf
- 07-generic-performance  ·leaf
- 08-generics-vs-interfaces  ·leaf
- 09-generic-data-structures  ·leaf
- 10-generic-limitations  ·leaf
- 11-methods-on-generic-types  ·leaf
- 12-stdlib-generic-packages  ·leaf
- 13-comparable-and-ordered  ·leaf
- 14-generic-type-aliases  ·leaf
- 15-recursive-type-constraints  ·leaf
- 16-generic-pitfalls  ·leaf
- 17-generic-testing-helpers  ·leaf

## 04-error-handling
*13 dirs · 13 leaves*

- 01-error-handling-basics  ·leaf
- 02-error-interface  ·leaf
- 03-errors-new  ·leaf
- 04-fmt-errorf  ·leaf
- 05-wrapping-unwrapping-errors  ·leaf
- 06-sentinel-errors  ·leaf
- 07-panic-and-recover  ·leaf
- 08-stack-traces-debugging  ·leaf
- 09-errors-is-vs-as-deep  ·leaf
- 10-custom-error-types  ·leaf
- 11-errors-join  ·leaf
- 12-error-design-best-practices  ·leaf
- 13-handle-dont-just-check  ·leaf

## 05-code-organization
*19 dirs · 17 leaves*

- 01-modules-and-dependencies
  - 01-go-mod-init  ·leaf
  - 02-go-mod-tidy  ·leaf
  - 03-go-mod-vendor  ·leaf
  - 04-minimal-version-selection-mvs  ·leaf
  - 05-module-proxy-and-checksum-db  ·leaf
  - 06-module-graph-pruning  ·leaf
  - 07-supply-chain-integrity  ·leaf
- 02-packages
  - 01-package-import-rules  ·leaf
  - 02-using-3rd-party-packages  ·leaf
  - 03-publishing-modules  ·leaf
- 03-project-layout  ·leaf
- 04-internal-packages  ·leaf
- 05-workspaces  ·leaf
- 06-dependency-injection  ·leaf
- 07-architecture-patterns  ·leaf
- 08-module-versioning  ·leaf
- 09-private-modules  ·leaf

## 06-concurrency
*157 dirs · 131 leaves*

- 00-introduction
  - 01-what-is-concurrency  ·leaf
  - 02-csp-model  ·leaf
  - 03-go-runtime-gmp  ·leaf
  - 04-memory-model  ·leaf
  - 05-when-to-use-concurrency  ·leaf
- 01-goroutines
  - 01-overview  ·leaf
  - 02-vs-os-threads  ·leaf
  - 03-stack-growth  ·leaf
  - 04-runtime-management  ·leaf
  - 05-best-practices  ·leaf
  - 06-common-pitfalls  ·leaf
- 02-channels
  - 01-buffered-vs-unbuffered  ·leaf
  - 02-select-statement  ·leaf
  - 03-worker-pools  ·leaf
  - 04-channel-direction  ·leaf
  - 05-nil-channels  ·leaf
  - 06-closing-channels  ·leaf
  - 07-range-over-channels  ·leaf
- 03-sync-package
  - 01-mutexes  ·leaf
  - 02-waitgroups  ·leaf
  - 03-once  ·leaf
  - 04-cond  ·leaf
  - 05-pool  ·leaf
  - 06-map  ·leaf
  - 07-atomic  ·leaf
- 04-context-package
  - 01-deadlines-and-cancellations  ·leaf
  - 02-common-usecases  ·leaf
  - 03-context-values  ·leaf
  - 04-context-tree  ·leaf
  - 05-context-internals  ·leaf
- 05-concurrency-patterns
  - 01-fan-in  ·leaf
  - 02-fan-out  ·leaf
  - 03-pipeline  ·leaf
  - 04-race-detection  ·leaf
  - 05-future-promise  ·leaf
  - 06-broadcast-pattern  ·leaf
  - 07-n-barrier  ·leaf
  - 08-push-pull  ·leaf
- 06-errgroup-x-sync
  - 01-errgroup  ·leaf
  - 02-semaphore  ·leaf
  - 03-singleflight  ·leaf
- 07-goroutine-lifecycle-leaks
  - 01-lifecycle  ·leaf
  - 02-detecting-leaks  ·leaf
  - 03-preventing-leaks  ·leaf
  - 04-pprof-tools  ·leaf
- 08-deadlock-livelock-starvation
  - 01-deadlock  ·leaf
  - 02-livelock  ·leaf
  - 03-starvation  ·leaf
- 09-channel-internals
  - 01-hchan-struct  ·leaf
  - 02-runtime-behavior  ·leaf
  - 03-buffer-mechanics  ·leaf
  - 04-send-receive-flow  ·leaf
- 10-scheduler-deep-dive
  - 01-gmp-model  ·leaf
  - 02-preemption  ·leaf
  - 03-gomaxprocs-tuning  ·leaf
  - 04-work-stealing  ·leaf
  - 05-syscall-handling  ·leaf
- 11-advanced-channel-patterns
  - 01-or-done-channel  ·leaf
  - 02-tee-channel  ·leaf
  - 03-bridge-channel  ·leaf
  - 04-generator  ·leaf
  - 05-ratelimiter  ·leaf
  - 06-handshaking  ·leaf
- 12-lock-free-programming
  - 01-cas-algorithms  ·leaf
  - 02-aba-problem  ·leaf
  - 03-lock-free-data-structures  ·leaf
  - 04-memory-fences  ·leaf
  - 05-lock-free-vs-wait-free  ·leaf
- 13-testing-concurrent-code
  - 01-race-detector-deep  ·leaf
  - 02-deterministic-testing  ·leaf
  - 03-waitgroup-in-tests  ·leaf
  - 04-mocking-time  ·leaf
  - 05-concurrent-fuzzing  ·leaf
- 14-performance-tuning
  - 01-gomaxprocs  ·leaf
  - 02-gogc  ·leaf
  - 03-lockosthread  ·leaf
  - 04-profiling-concurrent  ·leaf
  - 05-scheduler-tracing  ·leaf
- 15-concurrency-anti-patterns
  - 01-unlimited-goroutines  ·leaf
  - 02-mutex-copying  ·leaf
  - 03-channel-close-violations  ·leaf
  - 04-premature-optimization  ·leaf
  - 05-wait-for-empty-channel  ·leaf
  - 06-sleep-for-sync  ·leaf
- 16-time-based-concurrency
  - 01-ticker  ·leaf
  - 02-afterfunc  ·leaf
  - 03-timer-leaks  ·leaf
  - 04-exponential-backoff  ·leaf
  - 05-debounce-throttle  ·leaf
- 17-goroutine-pools-3rd-party
  - 01-ants  ·leaf
  - 02-tunny  ·leaf
  - 03-workerpool  ·leaf
  - 04-when-to-use  ·leaf
- 18-production-patterns
  - 01-backpressure  ·leaf
  - 02-dynamic-worker-scaling  ·leaf
  - 03-batching  ·leaf
  - 04-graceful-shutdown  ·leaf
  - 05-drain-pattern  ·leaf
  - 06-steady-state  ·leaf
- 19-pipeline-production-patterns
  - 01-error-propagation  ·leaf
  - 02-cancellation-propagation  ·leaf
  - 03-fan-out-within-pipeline  ·leaf
  - 04-batching-stages  ·leaf
  - 05-fan-in-fan-out-within  ·leaf
- 20-cancellation-deep
  - 01-cooperative-vs-force  ·leaf
  - 02-partial-cancellation  ·leaf
  - 03-cleanup-ordering  ·leaf
- 21-concurrent-data-structures
  - 01-ttl-caches  ·leaf
  - 02-lru-concurrent  ·leaf
  - 03-concurrent-skip-list  ·leaf
  - 04-concurrent-trees  ·leaf
  - 05-copy-on-write  ·leaf
  - 06-concurrent-counters  ·leaf
  - 07-concurrent-bloom-filter  ·leaf
- 22-memory-ordering-barriers
  - 01-hardware-barriers  ·leaf
  - 02-acquire-release  ·leaf
  - 03-sequential-consistency  ·leaf
  - 04-cache-coherence  ·leaf
  - 05-false-sharing  ·leaf
- 23-concurrency-in-stdlib
  - 01-net-http-server  ·leaf
  - 02-database-sql-pool  ·leaf
  - 03-sync-pool-internals  ·leaf
  - 04-runtime-internals  ·leaf
  - 05-time-package-concurrency  ·leaf
- 24-primitives-decision-guide
  - 01-channel-vs-mutex  ·leaf
  - 02-mutex-vs-atomic  ·leaf
  - 03-when-to-use-cond  ·leaf
  - 04-decision-tree  ·leaf
- 25-modern-features
  - 01-sync-oncefunc  ·leaf
  - 02-structured-concurrency  ·leaf
  - 03-future-proposals  ·leaf

## 07-standard-library
*21 dirs · 21 leaves*

- 01-io-and-file-handling  ·leaf
- 02-flag  ·leaf
- 03-time  ·leaf
- 04-encoding-json  ·leaf
- 05-os  ·leaf
- 06-bufio  ·leaf
- 07-slog  ·leaf
- 08-regexp  ·leaf
- 09-go-embed  ·leaf
- 10-net  ·leaf
- 11-net-http-internals  ·leaf
- 12-encoding  ·leaf
- 13-crypto  ·leaf
- 14-io-fs  ·leaf
- 15-templates  ·leaf
- 16-sort-slices-maps  ·leaf
- 17-container  ·leaf
- 18-fmt  ·leaf
- 19-strings-bytes  ·leaf
- 20-strconv  ·leaf
- 21-path-filepath  ·leaf

## 08-testing-and-benchmarking
*17 dirs · 17 leaves*

- 01-testing-basics  ·leaf
- 02-table-driven-tests  ·leaf
- 03-mocks-and-stubs  ·leaf
- 04-httptest  ·leaf
- 05-benchmarks  ·leaf
- 06-coverage  ·leaf
- 07-subtests  ·leaf
- 08-testmain  ·leaf
- 09-parallel-tests  ·leaf
- 10-test-helpers  ·leaf
- 11-golden-files  ·leaf
- 12-fuzzing  ·leaf
- 13-integration-tests  ·leaf
- 14-e2e-tests  ·leaf
- 15-mocking-libraries  ·leaf
- 16-property-based-testing  ·leaf
- 17-benchmark-deep  ·leaf

## 09-go-toolchain
*34 dirs · 27 leaves*

- 01-core-go-commands
  - 01-go-run  ·leaf
  - 02-go-build  ·leaf
  - 03-go-install  ·leaf
  - 04-go-fmt  ·leaf
  - 05-go-mod  ·leaf
  - 06-go-test  ·leaf
  - 07-go-clean  ·leaf
  - 08-go-doc  ·leaf
  - 09-go-version  ·leaf
- 02-code-generation-build-tags
  - 01-go-generate  ·leaf
  - 02-build-tags  ·leaf
- 03-code-quality-and-analysis
  - 01-go-vet  ·leaf
  - 02-goimports  ·leaf
  - 03-linters
    - 01-revive  ·leaf
    - 02-staticcheck  ·leaf
    - 03-golangci-lint  ·leaf
- 04-security
  - 01-govulncheck  ·leaf
- 05-performance-and-debugging
  - 01-pprof  ·leaf
  - 02-trace  ·leaf
  - 03-race-detector  ·leaf
- 06-deployment-and-tooling
  - 01-cross-compilation  ·leaf
  - 02-building-executables  ·leaf
- 07-go-work  ·leaf
- 08-debugging-with-delve  ·leaf
- 09-go-tool-suite  ·leaf
- 10-live-reload  ·leaf
- 11-build-tools  ·leaf

## 10-advanced-topics
*25 dirs · 24 leaves*

- 01-memory-management-in-depth  ·leaf
- 02-escape-analysis  ·leaf
- 03-reflection  ·leaf
- 04-unsafe-package  ·leaf
- 05-build-constraints  ·leaf
- 06-cgo-basics  ·leaf
- 07-compiler-linker-flags  ·leaf
- 08-plugins-dynamic-loading  ·leaf
- 09-assembly  ·leaf
- 10-linkname-directive  ·leaf
- 11-pgo  ·leaf
- 12-runtime-hooks  ·leaf
- 13-plugin-package  ·leaf
- 14-controller-runtime  ·leaf
- 15-serverless-go  ·leaf
- 16-compilation-pipeline
  - 01-lexer-scanner  ·leaf
  - 02-parser-ast  ·leaf
  - 03-type-checking  ·leaf
  - 04-ir-and-middle-end  ·leaf
  - 05-ssa-backend  ·leaf
  - 06-code-generation  ·leaf
  - 07-assembler-object-files  ·leaf
  - 08-linker  ·leaf
  - 09-build-orchestration-cache  ·leaf

## 11-performance-engineering
*7 dirs · 7 leaves*

- 01-cpu-profiling  ·leaf
- 02-memory-profiling  ·leaf
- 03-mutex-block-profiling  ·leaf
- 04-benchmarking-strategy  ·leaf
- 05-optimization-workflow  ·leaf
- 06-pprof-deep  ·leaf
- 07-trace-tool  ·leaf

## 12-design-patterns-in-go
*20 dirs · 20 leaves*

- 01-functional-options  ·leaf
- 02-builder-pattern  ·leaf
- 03-strategy-pattern  ·leaf
- 04-decorator-pattern  ·leaf
- 05-adapter-pattern  ·leaf
- 06-factory-pattern  ·leaf
- 07-observer-pattern  ·leaf
- 08-singleton-pattern  ·leaf
- 09-iterator-pattern  ·leaf
- 10-facade-pattern  ·leaf
- 11-proxy-pattern  ·leaf
- 12-chain-of-responsibility-pattern  ·leaf
- 13-command-pattern  ·leaf
- 14-state-pattern  ·leaf
- 15-object-pool-pattern  ·leaf
- 16-pubsub-pattern  ·leaf
- 17-futures-promises-pattern  ·leaf
- 18-registry-pattern  ·leaf
- 19-composite-pattern  ·leaf
- 20-fail-fast-pattern  ·leaf

## 13-runtime-and-internals
*6 dirs · 6 leaves*

- 01-runtime-source-dive  ·leaf
- 02-scheduler-source  ·leaf
- 03-gc-source  ·leaf
- 04-memory-allocator  ·leaf
- 05-runtime-package-deep  ·leaf
- 06-go-runtime-architecture  ·leaf

## 14-go-source-reading
*6 dirs · 6 leaves*

- 01-net-http-source  ·leaf
- 02-sync-source  ·leaf
- 03-runtime-source  ·leaf
- 04-context-source  ·leaf
- 05-database-sql-source  ·leaf
- 06-encoding-json-source  ·leaf

## 15-webassembly-and-alternative-targets
*5 dirs · 5 leaves*

- 01-goos-js-wasm-browser  ·leaf
- 02-wasi-and-wasip1  ·leaf
- 03-tinygo-for-wasm-and-embedded  ·leaf
- 04-wasm-interop-and-performance  ·leaf
- 05-wasm-in-production  ·leaf

## 16-observability-and-runtime-introspection
*5 dirs · 5 leaves*

- 01-runtime-metrics-package  ·leaf
- 02-expvar  ·leaf
- 03-runtime-trace-application-tracing  ·leaf
- 04-opentelemetry-in-go  ·leaf
- 05-godebug-and-runtime-debug  ·leaf

## 17-modern-language-features
*5 dirs · 5 leaves*

- 01-iterators-and-range-over-func  ·leaf
- 02-loopvar-semantics  ·leaf
- 03-min-max-clear-builtins  ·leaf
- 04-generic-type-aliases  ·leaf
- 05-modern-stdlib-additions  ·leaf
