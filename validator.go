// Package validator checks the parsed config with github.com/go-playground/validator.
package validator

import (
	"github.com/go-cnfg/cnfg"
	"github.com/go-playground/validator/v10"
)

var std = validator.New()

// Validate checks the config against the validate tags of its fields. Add it as the
// last source, once every other one has been read:
//
//	type Config struct {
//		Addr    string `validate:"required,hostname_port"`
//		Workers int    `validate:"gte=1,lte=100"`
//	}
//
//	cfg, err := cnfg.Parse(defaults, cnfg.Flags(), validator.Validate())
//
// The error is the validator.ValidationErrors of the underlying library, so
// errors.As gives you the failing fields one by one.
func Validate() cnfg.Source {
	return With(std)
}

// With is [Validate] with a validator of your own, for the custom rules, tag name or
// translations you set up yourself.
func With(v *validator.Validate) cnfg.Source {
	return cnfg.Func(func(cfg any) (any, error) {
		return cfg, v.Struct(cfg)
	})
}
