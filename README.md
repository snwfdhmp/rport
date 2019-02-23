[![Go Report Card](https://goreportcard.com/badge/github.com/snwfdhmp/rport)](https://goreportcard.com/report/github.com/snwfdhmp/rport) [![Documentation](https://godoc.org/github.com/snwfdhmp/rport?status.svg)](http://godoc.org/github.com/snwfdhmp/rport) [![GitHub issues](https://img.shields.io/github/issues/snwfdhmp/rport.svg)](https://github.com/snwfdhmp/rport/issues) [![license](https://img.shields.io/github/license/snwfdhmp/rport.svg?maxAge=6000)](https://github.com/snwfdhmp/rport/LICENSE)

# Easy plugged reporting service

## Getting started

### Start server

First, install with

```
$ go get github.com/snwfdhmp/rport/rportctl
```

Then, start with

```
$ rportctl start
```

### Configuring project for reports

First, download API package

```
$ go get github.com/snwfdhmp/rport
```

Then, in code

```golang
r = rport.NewReporter("https://localhost:9123")

//...
scope := someContextObject{}
err := funcIWantToTrackErrFrom()
if err != nil {
    r.Report(err, "func-label", scope)
}
```

## Documentation

Documentation can be found here : [![Documentation](https://godoc.org/github.com/snwfdhmp/rport?status.svg)](http://godoc.org/github.com/snwfdhmp/rport)

## Feedback

Feel free to open an issue for any feedback or suggestion.

I fix process issues quickly.

## Contributions

PR are accepted as soon as they follow Golang common standards.
For more information: https://golang.org/doc/effective_go.html

## License information

[![license](https://img.shields.io/github/license/snwfdhmp/rport.svg?maxAge=60000)](https://github.com/snwfdhmp/rport/LICENSE)