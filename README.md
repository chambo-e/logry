# `logry` [![GoDoc](https://godoc.org/github.com/chambo-e/logry?status.svg)](https://godoc.org/github.com/chambo-e/logry)

## Description

`logry` is a painless logging package with sentry integration.  

It's a simple [`logrus`](https://github.com/sirupsen/logrus) with a bundled [Sentry](https://sentry.io/) integration thanks to [`logrus_sentry`](https://github.com/evalphobia/logrus_sentry).

## Usage

```go
package main

import (
    "github.com/chambo-e/logry"
)

func main() {
    logry.Debugln("Will be printed on stdout except if APP_ENV=production")
    logry.Println("Will be printed on stdout")
    logry.Errorln("Will be printed on stderr and sent to Sentry")
    logry.Fatalln("Will be printed on stderr and sent to Sentry")
    logry.Panicln("Will be printed on stderr and sent to Sentry")
}
```

```bash
$ SENTRY_DSN=https://xxx:yyy@domain.com/XXX go run main.go
```

## Difference with `logrus`

You can do everything you were doing with `logrus` with `logry`, they expose the exact same interface.

There is only one main difference, the default log level is logrus.DebugLevel. You can change it with SetLevel or by setting env var `APP_ENV=production`.

`logry` reads you project DSN from app env var `SENTRY_DSN`. Default is an empty DSN and just disable the sentry exception logging.

These message levels are sent to sentry:
```go
[]logrus.Level{
    logrus.PanicLevel,
    logrus.FatalLevel,
    logrus.ErrorLevel,
}
```

`logry` also expose two additional methods:
```go
// To set sentry environment tag
// Default value is "development" and can be changed with APP_ENV
logry.SetEnvironment("lala")

// To set sentry release tag
logry.SetRelease("1.0.0")
```

### Contributions

Contributions are more than welcome as usual, feel free to open an issue :)

### Tests

There is no tests on this project because everything this package does is inside an hard to test init() function. But if this is an issue for some of you I'll be open to discuss on how it could be made.
