# About
* Professional work with Go for about 8 years, 20 years in general
* Worked
    - Big companies: Intel, Qualcomm, Applied Materials
    - Small: EWT/FATTOC, SauceLabs, ...
* Things I wrote in Go
    * HTTP Proxy
    * In memory database
    * go2xunit
    * selenium
    * Serverless (nuclio), close to 500K RPS
    * Data streaming
* Contributed to many open source projects
    * Including Go
* Still code every day
    * Mostly Go, Python, bash, Clojure, C, ...
* Still learn every day
    - ןלוכמ רתוי ידימלתמו יתוברמ רתוי ירבחמו יתוברמ יתדמל הברה 

# Workshop
- create project on github
- git clone project
- cd project, git add ...
- go mod init
    - go.mod
- go mod tidy
- Go over the code
    - Tokenize (document)
    - Must in init/var
- testing
    - `go build` ignores files ending with `_test.go`
    - nlp_test.go
    - fail vs fatal
    - table
    - [] vs nil
    - testify (assert vs require)
    - go get github.com/stretchr/testify
	- modules.md, our software dependency problem
	- vendor
    - quick
    - example_test.go
	- show on godoc.org
- Performance tuning
    - BenchmarkToeknizer
    - go test -bench . -run '^$' .
	- GOMAXPROCS
    - go test -bench . -run '^$' . -cpuprofile=prof.out
	- nlp.test
    - go tool pprof
	- top 20
	- lots of memory
    - change to `make([]string, 0, 20)`
    - Much of the doubling of speed for core Python that has occurred over the
      last ten decade has occurred one little step at a time, none of the them
      being individually dramatic.
	- Raymond Hettinger
    - optimize.md
- debugging
    - The most effective debugging tool is still careful thought, coupled with
      judiciously placed print statements. - Brian Kernighan
    - sleep (hammock driven development)
    - Feynman algorithm?
    - IDE
    - dlv test .
    - mention gdb
    - mention logs
- add stemmer
    - Conway's law
- add stop words
    - go generate
    - build tags
    - go generate
- cmd/nlpd
    - grpc?
    - gorilla
    - `_check`
    - handler
- logging & metrics
    - expvar
    - https://medium.com/netflix-techblog/sps-the-pulse-of-netflix-streaming-ae4db0e05f8a
- Dockerfile
    - Version
    - `CGO_ENABLED=0`
    - runtime.Version, runtime.GOOS ...
- https://memegenerator.net/img/instances/22605665/worked-fine-in-dev-ops-problem-now.jpg
- Publishing/workflow & code reviews
    - GitHub account
    - feature branches
    - PR
    - tag version (semver)
    - go get in testing
    - license
    - solutions/nlp/doc.go
    - README.md
    - GOOS=windows go build
- Continuous integration
    - Circle CI?
    - solutions/nlp/.circleci/, solutions/nlp/Dockerfile.test
    - Jenkins
    - chatops
- Deployment strategies
    - green/blue
    - canary
    - HAProxy?
