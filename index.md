# The Daily Go
GopherCon Tenerife ∴  2019 <br />
URL: [353solutions.com/c/tdg/](http://353solutions.com/c/tdg/)
{: .url}

{::comment}
([Download Zip](https://storage.googleapis.com/353solutions/c/tdg/tdg.zip) - Unzip and open `index.html`)
{:/comment}

Miki Tebeka <br />
<i class="far fa-envelope"></i> [miki@353solutions.com](mailto:miki@353solutions.com), 
<i class="fab fa-twitter"></i> [@tebeka](https://twitter.com/tebeka),
<i class="fab fa-linkedin-in"></i> [mikitebeka](https://www.linkedin.com/in/mikitebeka/),
<i class="fab fa-blogger-b"></i> [pythonwise blog](http://pythonwise.blogspot.com/)

#### Shameless Plugs

* [Go Essential Training](https://www.linkedin.com/learning/go-essential-training/) - LinkedIn Learning
    - [Rest of classes](https://www.linkedin.com/learning/instructors/miki-tebeka)
* [Forging Python](http://forging-python.com) - Miki's book<br />

# Code

TBD

# Links
- [Effective Go](https://golang.org/doc/effective_go.html) - Read this!
- [Go Proverbs](https://go-proverbs.github.io/) - Think about them ☺

{::comment}
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
* Serialization formats
    * [JSON](http://www.json.org/) - Textual, no schema ([encoding/json](https://golang.org/pkg/encoding/json))
    * [XML](http://www.w3schools.com/xml/) - Textual, optional external schema ([encoding/xml](https://golang.org/pkg/encoding/xml))
    * CSV - Textual, no schema ([encoding/csv](https://golang.org/pkg/encoding/csv))
    * [YAML](http://yaml.org/) - Textual, no schema ([gopkg.in/yaml.v2](https://gopkg.in/yaml.v2))
    * [TOML](https://github.com/toml-lang/toml) - Textual, no schema ([BurntSushi/toml](https://github.com/BurntSushi/toml))
    * [msgpack](http://msgpack.org/index.html) - Binary, no schema ([vmihailenco/msgpack](https://github.com/vmihailenco/msgpack))
    * [bson](http://bsonspec.org/) - Binary, no schema ([mgo/bson](https://godoc.org/labix.org/v2/mgo/bson))
    * [Protocol Buffers](https://developers.google.com/protocol-buffers/?hl=en) - Binary, schema ([golang/protobuf](https://github.com/golang/protobuf/))
    * [Cap'n Proto](https://capnproto.org/) - Binary, schema ([capnproto/go-capnproto2](https://github.com/capnproto/go-capnproto2))
    * [flatbuffers](https://google.github.io/flatbuffers/) - Binary, schema ([flatbuffers/go](github.com/google/flatbuffers/go))
* JSON
    - [encoding/json](https://golang.org/pkg/encoding/json/) in the standard library
    - [mapstructure](https://godoc.org/github.com/mitchellh/mapstructure#example-Decode) to fill struct from `map[string]interface{}`
- HTTP
    - [net/http](https://golang.org/pkg/net/http/) - Built-in HTTP client & server
    - [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) - More flexible mux
    - [chi](https://github.com/go-chi/chi) - Web framework
    - [fasthttp](https://godoc.org/github.com/valyala/fasthttp) - Faster HTTP server, use *only* if you really need it
    - [Making & Using HTTP Middleware](https://www.alexedwards.net/blog/making-and-using-middleware)
- [Our Software Depedency Problem](https://research.swtch.com/deps) by Russ Cox
- [build tags](https://dave.cheney.net/2013/10/12/how-to-use-conditional-compilation-with-the-go-build-tool)
* Command line parsing
    - Built in [flag](https://golang.org/pkg/flag/)
    - [Cobra](https://github.com/spf13/cobra)
	- Works well with [Viper](https://github.com/spf13/viper)
- Benchmarking & Profiling
    - [Performance](https://github.com/golang/go/wiki/Performance) in the Go Wiki
    - [Performance Tuning Workshop](https://github.com/davecheney/gophercon2018-performance-tuning-workshop)
    - [benchcmp](https://godoc.org/golang.org/x/tools/cmd/benchcmp) - Compare benchmarks
    - [pprof](https://golang.org/pkg/pprof/) & [net/http/pprof](https://golang.org/pkg/net/http/pprof/)
- Debugging
    - [dlv](https://github.com/go-delve/delve)
    - [VSCode](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code)
    - [gdb](https://golang.org/doc/gdb)
- Logging
    - [zap](https://godoc.org/go.uber.org/zap) - Logging library
    - [logrus](https://godoc.org/github.com/sirupsen/logrus)
    - Built-in [log](https://golang.org/pkg/log/) (See [here](https://dave.cheney.net/2015/11/05/lets-talk-about-logging) why)
- [Robustness Principle](https://en.wikipedia.org/wiki/Robustness_principle)
- [Documenting Go Code](https://blog.golang.org/godoc-documenting-go-code)
- Testing
    * [testify](https://godoc.org/github.com/stretchr/testify) for more testing frills
    * [gocheck](https://labix.org/gocheck)
    * [Testable examples](https://blog.golang.org/examples)
    * [Using sub tests](https://blog.golang.org/subtests)
- [Rules of Optimization](http://wiki.c2.com/?RulesOfOptimization)
- [Knight Capital](https://en.wikipedia.org/wiki/Knight_Capital_Group#2012_stock_trading_disruption) - the price of bugs in deployment
- [HTTP cats](https://www.flickr.com/photos/girliemac/sets/72157628409467125/)
    - Has [an API](https://http.cat/)
- [Crash only software](https://en.wikipedia.org/wiki/Crash-only_software)
    - See also [here](https://lwn.net/Articles/191059/)
* [gometalinter](https://github.com/alecthomas/gometalinter) - For all your linting needs
* [Error handling and Go](https://blog.golang.org/error-handling-and-go) blog post
- [Go standard library](https://golang.org/pkg/)
* [Fallacies of distributed computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing#The_fallacies)
* [Falsehoods Programmers Believe about Time](https://infiniteundo.com/post/25326999628/falsehoods-programmers-believe-about-time)
* [The Twelve-Factor App](https://12factor.net/)
{:/comment}


# Data & Other
* [nlp.go](data/nlp.go)
* [Console log](tdg.log)
