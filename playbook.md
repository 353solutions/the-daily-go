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
- mkdir nlp
- download nlp.go
- go mod init
    - go.mod
- go mod tidy
- Go over the code
    - Tokenize
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
    - go generate
- cmd/nlpd
    - gorilla
    - `_check`
    - handler
- logging & metrics
    - expvar
- Dockerfile
    - Version
- Development workflow & code reviews
    - GitHub account
    - feature branches
    - PR
- Continuous integration
    - Circle CI?
    - Jenkins
- Deployment strategies
    - green/blue
    - canary
    - HAProxy?
