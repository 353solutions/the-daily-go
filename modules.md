# Working with Modules

Go [modules](https://github.com/golang/go/wiki/Modules) are the latest and
official way of managing dependencies. Please read
[this](https://research.swtch.com/deps) before installing dependencies.

Once per project, create `go.mod` (assume our project name is `nlp`)

    go mod init 353solutions.com/go/nlp

Note: If you're under `GOPATH` (run `go env GOPATH` to check), you'll need to
set `GO111MODULE` environment variable.

    export GO111MODULE=on

To install a package run:

    go get github.com/pkg/errors


To install a specific version run:

    go get github.com/pkg/errors@v0.8.0

To have `go mod` scan your code and update `go.mod` run:

    go mod tidy

`go get` uses `git` (or `hg`, `svn` ...) to fetch packages. This means you can
have an internal company git server and install packages from it. You can use
[Athens](https://docs.gomods.io/),
[go-import-redirect](https://godoc.org/rsc.io/go-import-redirector), [Go
Registry](https://www.jfrog.com/confluence/display/RTF/Go+Registry) and others
to have custom URLs. If you don't have a git server, you can use the `replace`
directive in `go.mod` to install packages from a local directory.

	module nlp

	require (
	    github.com/pkg/errors v0.8.1
	)

	replace github.com/pkg/errors => /path/to/errors

You can also choose to `vendor` your dependencies. Run `go mod vendor` and then a
`vendor` directory will be created with all your dependencies. This mean you're
not relying on internet connection to build. To use the vendor directory, pass
`-mod=vendor` to all go commands (e.g. `go build -mod=vendor`) or set it in
`GOFLAGS` environment variable (e.g. `export GOFLAGS="-mod=vendor"`)

## Extra Reading
- [Go Modules](https://github.com/golang/go/wiki/Modules)
- [Using Go modules with vendor support on Travis CI](https://arslan.io/2018/08/26/using-go-modules-with-vendor-support-on-travis-ci/)
- [Using Go Modules](https://blog.golang.org/using-go-modules)
- [Go Modules for Package Maintainers](https://www.youtube.com/watch?v=ms5l0zxC-uM) video
