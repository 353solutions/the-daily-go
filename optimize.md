# Miki's Optimization Guide

## 3 Rules of Optimization

1. Don't
: You first need to have measurable performance goals ("as fast as you can't"
is not an acceptable goal). If you hit these goals go do something with better
business value.

2. Don't ... yet
: It's much easier and cheaper to fix the problem with hardware. Get a faster
CPU, faster network ... Developer time & money are the most expensive resources
in software development. Also note that optimized code is much harder to
maintain.

3. Profile before optimizing
: Bottlenecks will surprise you. Don't *guess* where the code spends it's time,
use a profiler and see.

See more [here](https://users.ece.utexas.edu/~adnan/pike.html) and
[here](http://wiki.c2.com/?RulesOfOptimization).

## General

1. Algorithms & Data structures Rule
: They will usually give you much better performance than any other trick.

2. Know thy Hardware
: CPU affinity, CPU cache, memory, [latency numbers
...](https://twitter.com/piecalculus/status/459485747842523136?lang=en).
For example: [Cache-oblivious
algorithms](https://en.wikipedia.org/wiki/Cache-oblivious_algorithm)

3. Include performance in your process
: Design & code reviews, run & compare benchmarks on CI ...

## Go Specific

1. Memory Allocation
: Avoid allocations as possible (see the design of
[io.Reader](https://golang.org/pkg/io/#Reader)). Pre-allocate if you already
know the size. Be careful of slices keep large amounts of memory 
(`s := make([]int, 1000000)[:3]`)

2. `defer` might slow you Down
: However consider the advantages.

3. strings are immutable
: Use [bytes.Buffer](https://golang.org/pkg/bytes/#Buffer) or [strings.Builder](https://golang.org/pkg/strings/#Builder)

4. Know when a goroutine is going to stop
: Avoid goroutine leaks. Use [context](https://golang.org/pkg/context/) for cancellation/timeouts.

5. Cgo calls are expensive
: Group them together in one `cgo` call.

6. Channel can be slower than `sync.Mutex`
: However they are much easier to work with

7. Interface calls are more expensive the struct calls
: You can extract the value from the interface first. However it's less 
generic code.

8. Use `go run -gcflags=-m -l`
: You'll see what escapes to the heap.

### Reading
- [So you wanna go fast](https://www.slideshare.net/TylerTreat/so-you-wanna-go-fast-80300458)
- [Performance tuning workshop](https://github.com/davecheney/gophercon2018-performance-tuning-workshop/blob/master/6-tips-and-tricks/1-tips-and-tricks.md).
- [Quick look at some compiler optimization](http://www.golangbootcamp.com/book/tricks_and_tips#sec-compiler_optimizations)
- [Performance Mantras](http://www.brendangregg.com/blog/2018-06-30/benchmarking-checklist.html)


{::comment}
## Performance Mantras

By [Craig Hanson and Pat
Crain](http://www.brendangregg.com/blog/2018-06-30/benchmarking-checklist.html)

1. Don't do it
: Can we avoid doing the calculation at all? For example: Do we need to parse
the input or just pass it as-is?

2. Do it, but don't do it again
: Can we use [memoization](https://en.wikipedia.org/wiki/Memoization)/caching?
Parse objects once at the "edges" and use the parsed objects internally.

3. Do it less
: Do we need to run this every millisecond? Can every second work? Can we use
only a subset of the data?

4. Do it later
: Can we make this API call async?

5. Do it when they're not looking
: Can we run the calculation in the background while doing another task?

6. Do it concurrently
: Will concurrency help here? Consider [Amdhal's law](https://en.wikipedia.org/wiki/Amdahl%27s_law).

7. Do it cheaper
: Can we use a map here instead of a slice? Research available algorithms and
data structures and know their complexity. Test them on *your* data

{:/comment}

