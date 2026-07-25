#!/usr/bin/env bash
#
# reconstruct.sh — rebuild the challenges/ skeleton from the Go roadmap.
#
# Structure:  challenges/<level>/<NN-topic>/<NN-subtopic>/<puzzle>/
#
# Level is the TOP dimension and is assigned per subtopic by COMPETENCY, not by
# ordinal quartile. Each subtopic is tagged J/M/S/T (junior/middle/senior/staff)
# from the level profiles in challenges/GENERATION.md. A topic therefore spans
# only the levels its subtopics actually belong to (ragged grid) — foundational
# topics collapse into junior, deep topics sit at senior/staff. This keeps the
# scope rule sound: junior is self-contained fundamentals.
#
# Subtopics are renumbered from 01 within each (level, topic), roadmap order
# preserved. Empty subtopics get a .keep; a subtopic holding a puzzle (dir with
# go.mod) loses its .keep. Idempotent: existing puzzles preserved.
#
set -euo pipefail

ROADMAP="${ROADMAP:-/Users/mrb/Desktop/SeniorProject/Roadmap/Programming/languages/golang}"
CH="${CH:-$(cd "$(dirname "$0")/../challenges" && pwd)}"

declare -A LNAME=([J]=junior [M]=middle [S]=senior [T]=staff)

# DATA lines:
#   T <NN> <local-slug> <roadmap-dir>   -> start a topic
#   <L> <sub-slug>                       -> a subtopic at level L, under current topic
# Order within a topic = roadmap order (drives per-(level,topic) renumbering).
# 07-concurrency/00-introduction dropped (not a coding puzzle).
DATA='
T 01 language-basics 02-language-basics
J variables-and-constants
J data-types
J composite-types
J conditionals
J loops
J functions
J pointers
J arrays
J strings
J maps
J structs
T 02 methods-and-interfaces 03-methods-and-interfaces
J methods-vs-functions
J pointer-receivers
J value-receivers
J interfaces-basics
J empty-interfaces
M embedding-interfaces
M type-assertions
M type-switch
M method-sets-deep
S interface-internals
S method-dispatch
M common-interfaces
M interface-best-practices
S interface-anti-patterns
M method-values-and-expressions
M methods-on-defined-types
S sealed-interfaces
M cross-package-methods
M struct-method-promotion
T 03 error-handling 05-error-handling
J error-handling-basics
J error-interface
J errors-new
J fmt-errorf
M wrapping-unwrapping-errors
M sentinel-errors
M panic-and-recover
S stack-traces-debugging
S errors-is-vs-as-deep
S custom-error-types
S errors-join
T error-design-best-practices
T handle-dont-just-check
T 04 standard-library 08-standard-library
J io-and-file-handling
J flag
J time
J encoding-json
J os
M bufio
M slog
M regexp
M go-embed
M net
S net-http-internals
M encoding
S crypto
S io-fs
M templates
M sort-slices-maps
S container
J fmt
J strings-bytes
J strconv
M path-filepath
T 05 code-organization 06-code-organization
J modules-and-dependencies
J packages
J package-import-rules
J project-layout
M publishing-modules
M internal-packages
M workspaces
M dependency-injection
S architecture-patterns
M module-versioning
S private-modules
T 06 go-toolchain 10-go-toolchain
J core-go-commands
M code-generation-build-tags
J code-quality-and-analysis
M security
S performance-and-debugging
M deployment-and-tooling
M go-work
S debugging-with-delve
M go-tool-suite
M live-reload
S build-tools
T 07 testing-and-benchmarking 09-testing-and-benchmarking
J testing-basics
J table-driven-tests
M mocks-and-stubs
M httptest
M benchmarks
J coverage
J subtests
M testmain
S parallel-tests
M test-helpers
M golden-files
S fuzzing
S integration-tests
T e2e-tests
T mocking-libraries
T property-based-testing
T benchmark-deep
T 08 modern-language-features 18-modern-language-features
M iterators-and-range-over-func
J loopvar-semantics
J min-max-clear-builtins
S generic-type-aliases
M modern-stdlib-additions
T 09 generics 04-generics
M why-generics
M generic-functions
M generic-types-interfaces
M type-constraints
M type-inference
S generic-constraints-deep
S generic-performance
S generics-vs-interfaces
S generic-data-structures
S generic-limitations
S methods-on-generic-types
S stdlib-generic-packages
S comparable-and-ordered
T generic-type-aliases
T recursive-type-constraints
T generic-pitfalls
T generic-testing-helpers
T 10 concurrency 07-concurrency
M goroutines
M channels
M select-and-buffering
M worker-pools
M sync-package
M context-package
S race-detection
S concurrency-patterns
S errgroup-x-sync
S goroutine-lifecycle-leaks
S deadlock-livelock-starvation
S channel-internals
T scheduler-deep-dive
S advanced-channel-patterns
T lock-free-programming
S testing-concurrent-code
T performance-tuning
S concurrency-anti-patterns
M time-based-concurrency
S goroutine-pools-3rd-party
T production-patterns
T pipeline-production-patterns
S cancellation-deep
T concurrent-data-structures
T memory-ordering-barriers
S concurrency-in-stdlib
S primitives-decision-guide
S modern-features
T 11 web-development 19-web-development
M building-clis
M net-http-server
M routing-and-handlers
M rest-api-design
M orms-and-db-access
S middleware-and-context
S grpc-and-protobuf
S realtime-communication
S web-frameworks-optional
T 12 design-patterns-in-go 13-design-patterns-in-go
M functional-options
M builder-pattern
M strategy-pattern
M decorator-pattern
M adapter-pattern
M factory-pattern
M observer-pattern
M singleton-pattern
M iterator-pattern
M facade-pattern
S proxy-pattern
S chain-of-responsibility-pattern
S command-pattern
S state-pattern
S object-pool-pattern
S pubsub-pattern
S futures-promises-pattern
S registry-pattern
S composite-pattern
S fail-fast-pattern
T 13 observability 17-observability-and-runtime-introspection
M runtime-metrics-package
M expvar
S runtime-trace-application-tracing
S opentelemetry-in-go
S godebug-and-runtime-debug
T 14 performance-engineering 12-performance-engineering
S cpu-profiling
S memory-profiling
S mutex-block-profiling
S benchmarking-strategy
S optimization-workflow
T pprof-deep
T trace-tool
T 15 advanced-topics 11-advanced-topics
S memory-management-in-depth
S escape-analysis
S reflection
S unsafe-package
S build-constraints
S cgo-basics
T compiler-linker-flags
T plugins-dynamic-loading
T assembly
T linkname-directive
T pgo
T runtime-hooks
T plugin-package
T controller-runtime
S serverless-go
T compilation-pipeline
T 16 runtime-and-internals 14-runtime-and-internals
S runtime-source-dive
T scheduler-source
T gc-source
T memory-allocator
S runtime-package-deep
T go-runtime-architecture
T 17 go-source-reading 15-go-source-reading
S net-http-source
S sync-source
T runtime-source
S context-source
S database-sql-source
T encoding-json-source
T 18 webassembly 16-webassembly-and-alternative-targets
S goos-js-wasm-browser
S wasi-and-wasip1
T tinygo-for-wasm-and-embedded
T wasm-interop-and-performance
T wasm-in-production
'

subtopic_has_puzzle() {  # dir with a go.mod one level down
  find "$1" -mindepth 2 -maxdepth 2 -name go.mod -print -quit 2>/dev/null | grep -q .
}

gslug=''; rmdir=''
declare -A cnt=()     # per-(level,topic) subtopic counter
declare -A tnum=()    # topic dir number, keyed "level|slug"
declare -A levt=()    # per-level topic counter (contiguous, no gaps)

while read -r a b c d; do
  [ -z "${a:-}" ] && continue
  if [ "$a" = "T" ] && [[ "$b" =~ ^[0-9]+$ ]]; then
    gslug="$c"; rmdir="$d"      # DATA order = global learning-path order
    src="$ROADMAP/$rmdir"
    [ -d "$src" ] || echo "WARN: missing roadmap dir $src" >&2
    # reset per-(level,topic) subtopic counters
    cnt[junior]=0; cnt[middle]=0; cnt[senior]=0; cnt[staff]=0
    printf '  %-26s -> %s\n' "$b-$gslug" "$rmdir"
    continue
  fi
  lvl="${LNAME[$a]:-}"
  [ -z "$lvl" ] && { echo "WARN: bad level tag '$a' for $b" >&2; continue; }
  # assign this topic a contiguous number within its level on first use
  key="$lvl|$gslug"
  if [ -z "${tnum[$key]:-}" ]; then
    levt[$lvl]=$(( ${levt[$lvl]:-0} + 1 ))
    tnum[$key]=$(printf '%02d' "${levt[$lvl]}")
  fi
  cnt[$lvl]=$(( cnt[$lvl] + 1 ))
  nn=$(printf '%02d' "${cnt[$lvl]}")
  dir="$CH/$lvl/${tnum[$key]}-$gslug/$nn-$b"
  mkdir -p "$dir"
  if subtopic_has_puzzle "$dir"; then rm -f "$dir/.keep"; else : > "$dir/.keep"; fi
done <<< "$DATA"

echo
for l in junior middle senior staff; do
  printf 'level %-7s topics=%s subtopics=%s\n' "$l" \
    "$(find "$CH/$l" -mindepth 1 -maxdepth 1 -type d 2>/dev/null | wc -l | tr -d ' ')" \
    "$(find "$CH/$l" -mindepth 2 -maxdepth 2 -type d 2>/dev/null | wc -l | tr -d ' ')"
done
echo "puzzles=$(find "$CH" -name go.mod | wc -l | tr -d ' ')"
