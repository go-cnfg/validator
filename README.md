[![Go Reference](https://pkg.go.dev/badge/github.com/go-cnfg/validator.svg)](https://pkg.go.dev/github.com/go-cnfg/validator) ![main](https://github.com/go-cnfg/validator/actions/workflows/go.yaml/badge.svg?branch=main) [![codecov](https://codecov.io/gh/go-cnfg/validator/branch/main/graph/badge.svg)](https://codecov.io/gh/go-cnfg/validator)

# validator

[go-playground/validator](https://github.com/go-playground/validator) shim for
[cnfg](https://github.com/go-cnfg/cnfg), to better validation logic for configs.

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

cfg, err := cnfg.Parse(defaults,
    cnfg.Env[Config]("APP"),
    cnfg.Flags[Config](),
    validator.Validate[Config](),
)
```

A field that does not pass stops the parse:

```
Key: 'Config.Workers' Error:Field validation for 'Workers' failed on the 'gte' tag
```

`Validate` uses a shared validator with the default settings. `With` takes one you set up
yourself, for your own rules, a different tag name or translations:

```go
import v10 "github.com/go-playground/validator/v10"

v := v10.New()
v.SetTagName("check")

cfg, err := cnfg.Parse(defaults, cnfg.Flags[Config](), validator.With[Config](v))
```

Put the check last so every source has been read. `Parse` returns the zero value of your
config together with the error when a parser fails, so a config that does not pass never
reaches your program.
