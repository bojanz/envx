 # envx [![Build Status](https://travis-ci.org/bojanz/envx.png?branch=master)](https://travis-ci.org/bojanz/envx) [![Coverage Status](https://coveralls.io/repos/github/bojanz/envx/badge.svg?branch=master)](https://coveralls.io/github/bojanz/envx?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/bojanz/envx)](https://goreportcard.com/report/github.com/bojanz/envx) [![GoDoc](https://godoc.org/github.com/bojanz/envx?status.svg)](https://godoc.org/github.com/bojanz/envx)

Allows retrieving environment variables with a fallback.

- Retrieve a single variable: `envx.Get(key, fallback string)`

- Replace all $var, ${var} and ${var:default} variables: `envx.Expand(s string)`

 ```go
    port := envx.Get("PORT", "80")
    // "0.0.0.0:80" if $HOST is "0.0.0.0" and $PORT is missing/empty.
    addr := envx.Expand("${HOST}:${PORT:80}")
    // "/srv/www" if $WORKDIR is missing/empty.
    dir := envx.Expand("${WORKDIR:/srv}/www")
 ```
