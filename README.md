# The Daily Go
GopherCon Israel ∴  2021<br />
URL: [353solutions.com/c/tdg21/](http://353solutions.com/c/tdg21/) ([Download Zip](353-tdg.zip) - Unzip and open `README.html`)
GitHub Repo: [https://github.com/353solutions/nlp](https://github.com/353solutions/nlp) (we're on the `gcil2021`  branch)

<!--
-->

Miki Tebeka
<i class="far fa-envelope"></i> [miki@353solutions.com](mailto:miki@353solutions.com), <i class="fab fa-twitter"></i> [@tebeka](https://twitter.com/tebeka), <i class="fab fa-linkedin-in"></i> [mikitebeka](https://www.linkedin.com/in/mikitebeka/)


#### Shameless Plugs

- [Books](https://www.353solutions.com/books)
- [LinkedIn Learning classes](https://www.linkedin.com/learning/instructors/miki-tebeka)

# Code

<pre>
├── <a href="html/README.md.html">REAMDE.md</a> - Landing page
├── <a href="html/go.mod.html">go.mod</a> - Dependencies
├── <a href="html/nlp.go.html">nlp.go</a> - Main package code
├── <a href="html/doc.go.html">doc.go</a> - Package level documentation
├── <a href="html/example_test.go.html">example_test.go</a> - Testable example
├── <a href="html/gen_stop.go.html">gen_stop.go</a> - Generate `stop_words.go`
├── <a href="html/stop_words.txt.html">stop_words.txt</a> - Stop word list
├── <a href="html/nlp_test.go.html">nlp_test.go</a> - Test
├── testdata - Test data files
│   └── <a href="html/testdata/tokenizer_cases.json.html">tokenizer_cases.json</a> - Test cases
├── stemmer - Sub package
│   ├── <a href="html/stemmer/stemmer.html">stemmer.go</a>
│   └── <a href="html/stemmer/stemmer_test.html">stemmer_test.go</a>
└── cmd - Front ends
    └── nlpd
        └── <a href="html/cmd/nlpd/nlpd.html">nlpd.go</a> - HTTP (REST) front end
</pre>

[Console Log](console.log)

# Possible Exercises

- Test nlpd `tokenizeHandler`
- Validate that `Port` in configuration is above 8000
- Add metrics
    - Last time `/health` was hit
    - Number of bytes sent
    - Median runtime of last 10 calls to `Tokenize`
    - Consider using a [middleware](https://www.alexedwards.net/blog/making-and-using-middleware) for the last two
- Use `gorilla/mux` for routing instead of `http.HandleFunc`
    - See that only `POST` requests get to `Tokenize`
- Add a gRPC back end

# Data
- [nlp.go](data/nlp.go)
- [nlpd.go](data/nlpd.go)
- [stemmer.go](data/stemmer.go)
- [gen_stop.go](data/gen_stop.go)
- [stop_words.txt](data/stop_words.txt)
- [tokenizer_cases.json](data/tokenizer_cases.json)

# Links

- [Standard Project Layout](https://github.com/golang-standards/project-layout)
- [Go 1.16 "embed"](https://golangtutorial.dev/tips/embed-files-in-go/)
- [Conway's Law](https://en.wikipedia.org/wiki/Conway%27s_law)
- [Use internal packages to reduce your public API surface](https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface) by Dave Cheney
- [Crash only software](https://en.wikipedia.org/wiki/Crash-only_software)
    - See also [here](https://lwn.net/Articles/191059/)
- Linters
    - [go vet](https://golang.org/cmd/vet/)
    - [golangci-lint](https://github.com/golangci/golangci-lint)
    - [staticcheck](https://staticcheck.io/)
    - [gosec](https://github.com/securego/gosec)
    - Miki's [recheck](https://github.com/tebeka/recheck)
- Dependency Management
    - [Our Software Depedency Problem](https://research.swtch.com/deps) by Russ Cox
    - [Using Go Modules](https://blog.golang.org/using-go-modules)
    - [Using Go modules with vendor support on Travis CI](https://arslan.io/2018/08/26/using-go-modules-with-vendor-support-on-travis-ci/)
    - [Go Modules](https://github.com/golang/go/wiki/Modules)
    - [Go Modules for Package Mainainers](https://www.youtube.com/watch?v=ms5l0zxC-uM)
    - Also [this summary](modules.html)
- Logging & Metrics
    - [zap](https://godoc.org/go.uber.org/zap) - Logging library
    - [logrus](https://godoc.org/github.com/sirupsen/logrus)
    - Built-in [log](https://golang.org/pkg/log/) (See [here](https://dave.cheney.net/2015/11/05/lets-talk-about-logging) why)
    - [expvar](https://golang.org/pkg/expvar/)
- [Fallacies of distributed computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing#The_fallacies)
- [The Twelve-Factor App](https://12factor.net/)
- [Who's on First?'](https://www.youtube.com/watch?v=kTcRRaXV-fg)
- [Systems that run forever and self heal](https://www.infoq.com/presentations/self-heal-scalable-system/) by Joe Armstrong
- Fuzzing
    - [testing/quick](https://golang.org/pkg/testing/quick/) - In the standard library
    - [go-fuzz](https://github.com/dvyukov/go-fuzz) - Discovery driven
    - [gofuzz](https://github.com/google/gofuzz) - By Google
    - [Fuzzing Design Draft](https://go.googlesource.com/proposal/+/master/design/draft-fuzzing.md)
- Testing
    - [testing](https://golang.org/pkg/testing/)
    - [http/httptest](https://golang.org/pkg/net/http/httptest/)
    - [testify](https://godoc.org/github.com/stretchr/testify) for more testing frills
    - [gocheck](https://labix.org/gocheck)
    - [ginko](https://onsi.github.io/ginkgo/)
    - [Testable examples](https://blog.golang.org/examples)
    - [Using sub tests](https://blog.golang.org/subtests)
    - [How SQLite is Tested](https://www.sqlite.org/testing.html)
- Documentation
    - [Documenting Go Code](https://blog.golang.org/godoc-documenting-go-code)
    - [godock-tricks](https://pkg.go.dev/github.com/fluhus/godoc-tricks)
    - `gob` package [doc.go](https://golang.org/src/encoding/gob/doc.go)
- [Open Source Licences](https://opensource.org/licenses/)
- [How to Write Go Code](https://golang.org/doc/code.html)
- [Effective Go](https://golang.org/doc/effective_go.html) - Read this!

<!--
- Debugging
    - [dlv](https://github.com/go-delve/delve)
    - [Debug a Go Application running on Kubernetes cluster](https://www.youtube.com/watch?v=YXu2box7z9k)
    - [gdb](https://golang.org/doc/gdb)
    - [VSCode](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code)
- Benchmarking & Profiling
    - [High Performance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html)
    - [Performance](https://github.com/golang/go/wiki/Performance) in the Go Wiki
    - [benchcmp](https://godoc.org/golang.org/x/tools/cmd/benchcmp) - Compare benchmarks
    - [pprof](https://golang.org/pkg/pprof/) & [net/http/pprof](https://golang.org/pkg/net/http/pprof/)
    - [Optimization Tips](optimize.html)
    - [Latency numbers](https://twitter.com/piecalculus/status/459485747842523136?lang=en)
- [Go Proverbs](https://go-proverbs.github.io/) - Think about them ☺
- [Semantic versioning](https://semver.org/)
- [Miki's .vimrc](vimrc)



| Format                       | Type    | Schema | Package |
|------------------------------+---------+--------+---------|
| [JSON](http://www.json.org/) | Textual | No     | [encoding/json](https://golang.org/pkg/encoding/json) |
|------------------------------+---------+--------+---------|
|[XML](http://www.w3schools.com/xml/) | Textual | External | [encoding/xml](https://golang.org/pkg/encoding/xml) |
|------------------------------+---------+--------+---------|
| CSV                          | Textual | No     | [encoding/csv](https://golang.org/pkg/encoding/csv) |
|------------------------------+---------+--------+---------|
| [YAML](http://yaml.org/)     | Textual | No     | [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) |
|------------------------------+---------+--------+---------|
| [TOML](https://github.com/toml-lang/toml) | Textual | No | [BurntSushi/toml](https://github.com/BurntSushi/toml) |
|------------------------------+---------+--------+---------|
| [msgpack](http://msgpack.org/index.html) | Binary | No | [vmihailenco/msgpack](https://github.com/vmihailenco/msgpack) |
|------------------------------+---------+--------+---------|
| [bson](http://bsonspec.org/) | Binary  | No     | [mgo/bson](https://godoc.org/labix.org/v2/mgo/bson) |
|------------------------------+---------+--------+---------|
| [Protocol Buffers](https://developers.google.com/protocol-buffers/?hl=en) | Binary | Yes | [golang/protobuf](https://github.com/golang/protobuf/) |
|------------------------------+---------+--------+---------|
| [Cap'n Proto](https://capnproto.org/)  | Binary | Yes | [capnproto/go-capnproto2](https://github.com/capnproto/go-capnproto2) |
|------------------------------+---------+--------+---------|
| [flatbuffers](https://google.github.io/flatbuffers/) | Binary | Yes | [flatbuffers/go](github.com/google/flatbuffers/go) |
|------------------------------+---------+--------+---------|

- [Testable Examples](https://blog.golang.org/examples)
- Internal module repository/registry
    - https://godoc.org/rsc.io/go-import-redirector
    - `replace` in `go.mod`
	~~~
		module prj

		require (
			github.com/att/calc v0.0.0
			gopkg.in/yaml.v2 v2.2.2 // indirect
		)

		replace github.com/att/calc => /path/to/calc
	~~~
    - [jFrog](https://www.jfrog.com/confluence/display/RTF/Go+Registry)


- Serialization formats
    -[JSON](http://www.json.org/) - Textual, no schema ([encoding/json](https://golang.org/pkg/encoding/json))
- JSON
    - [encoding/json](https://golang.org/pkg/encoding/json/) in the standard library
    - [mapstructure](https://godoc.org/github.com/mitchellh/mapstructure#example-Decode) to fill struct from `map[string]interface{}`
- HTTP
    - [net/http](https://golang.org/pkg/net/http/) - Built-in HTTP client & server
    - [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) - More flexible mux
    - [chi](https://github.com/go-chi/chi) - Web framework
    - [fasthttp](https://godoc.org/github.com/valyala/fasthttp) - Faster HTTP server, use *only* if you really need it
    - [Making & Using HTTP Middleware](https://www.alexedwards.net/blog/making-and-using-middleware)
- [build tags](https://dave.cheney.net/2013/10/12/how-to-use-conditional-compilation-with-the-go-build-tool)
- Configuration
    - [envconfig](https://github.com/kelseyhightower/envconfig) - Configure using environment
    - Built in [flag](https://golang.org/pkg/flag/)
    - [Cobra](https://github.com/spf13/cobra) & [Viper](https://github.com/spf13/viper)
- [Conway's Law](https://en.wikipedia.org/wiki/Conway%27s_law)
- [Robustness Principle](https://en.wikipedia.org/wiki/Robustness_principle)
- [Rules of Optimization](http://wiki.c2.com/?RulesOfOptimization)
- Options & configuration
    - Built in [flag](https://golang.org/pkg/flag/)
    - [Cobra](https://github.com/spf13/cobra)
	- Works well with [Viper](https://github.com/spf13/viper)
    - [envconfig](https://github.com/kelseyhightower/envconfig)
- [Knight Capital](https://en.wikipedia.org/wiki/Knight_Capital_Group#2012_stock_trading_disruption) - the price of bugs in deployment
- [HTTP cats](https://www.flickr.com/photos/girliemac/sets/72157628409467125/)
    - Has [an API](https://http.cat/)
- [Error handling and Go](https://blog.golang.org/error-handling-and-go) blog post
- [Go standard library](https://golang.org/pkg/)
- [Falsehoods Programmers Believe about Time](https://infiniteundo.com/post/25326999628/falsehoods-programmers-believe-about-time)
-->
