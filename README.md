 # envx [![Build Status](https://travis-ci.org/bojanz/envx.png?branch=master)](https://travis-ci.org/bojanz/envx) [![Go Report Card](https://goreportcard.com/badge/github.com/bojanz/envx)](https://goreportcard.com/report/github.com/bojanz/envx) [![GoDoc](https://godoc.org/github.com/bojanz/envx?status.svg)](https://godoc.org/github.com/bojanz/envx)

Allows expanding env variables with defaults: ${var:default}.

 ```go
    // "0.0.0.0:80" if $PORT is missing or empty.
    addr := envx.Expand("0.0.0.0:${PORT:80}")
    // "/srv" if $WORKDIR is empty.
    dir := envx.Expand("${WORKDIR:/srv}")
 ```
