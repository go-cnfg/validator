[![Go Reference](https://pkg.go.dev/badge/github.com/go-cnfg/validator.svg)](https://pkg.go.dev/github.com/go-cnfg/validator) ![main](https://github.com/go-cnfg/validator/actions/workflows/go.yaml/badge.svg?branch=main) [![codecov](https://codecov.io/gh/go-cnfg/validator/branch/main/graph/badge.svg)](https://codecov.io/gh/go-cnfg/validator)

# validator

[go-playground/validator](https://github.com/go-playground/validator) as a parser for
[cnfg](https://github.com/go-cnfg/cnfg), so the rules of a config field live in its struct tag
and are checked once every source has been read.

```sh
go get github.com/go-cnfg/validator
```

```go
import (
    "time"

    "github.com/go-cnfg/cnfg"
    "github.com/go-cnfg/validator"
)

type Config struct {
    Addr    string        `validate:"required,hostname_port"`
    Workers int           `validate:"gte=1,lte=100"`
    Level   string        `validate:"oneof=debug info warn"`
    Timeout time.Duration `validate:"min=1s"`
}

cfg, err := cnfg.Parse(Config{
    Addr:    ":8080",
    Workers: 4,
    Level:   "info",
    Timeout: 30 * time.Second,
},
    cnfg.Env[Config]("APP"),
    cnfg.Flags[Config](),
    validator.Validate[Config](),
)
```

A field that does not pass stops the parse:

```
Key: 'Config.Workers' Error:Field validation for 'Workers' failed on the 'gte' tag
```

The error is go-playground's `ValidationErrors`, so `errors.As` gets you every field that
failed and the tag it failed on, for a message of your own:

```go
import v10 "github.com/go-playground/validator/v10"

var fields v10.ValidationErrors
if errors.As(err, &fields) {
    for _, f := range fields {
        log.Printf("%s: %s", f.Field(), f.Tag())
    }
}
```

`Validate` uses a shared validator with the default settings. `With` takes one you set up
yourself, for your own rules, a different tag name or translations:

```go
import v10 "github.com/go-playground/validator/v10"

v := v10.New()
v.SetTagName("check")

cfg, err := cnfg.Parse(Config{Addr: ":8080"}, cnfg.Flags[Config](), validator.With[Config](v))
```

Put the check last so every source has been read. `Parse` returns the zero value of your
config together with the error when a parser fails, so a config that does not pass never
reaches your program.
