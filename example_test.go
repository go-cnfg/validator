package validator_test

import (
	"errors"
	"fmt"

	v10 "github.com/go-playground/validator/v10"

	"github.com/go-cnfg/cnfg"
	"github.com/go-cnfg/validator"
)

type AppConfig struct {
	Addr    string `validate:"required,hostname_port"`
	Workers int    `validate:"gte=1,lte=100"`
}

func ExampleValidate() {
	defaults := AppConfig{Addr: ":8080", Workers: 4}

	cfg, err := cnfg.Parse(defaults,
		cnfg.EnvFrom("APP", []string{"APP_WORKERS=8"}),
		validator.Validate(),
	)
	fmt.Println(cfg.Workers, err)

	_, err = cnfg.Parse(defaults,
		cnfg.EnvFrom("APP", []string{"APP_WORKERS=0"}),
		validator.Validate(),
	)
	fmt.Println(err)

	// Output:
	// 8 <nil>
	// Key: 'AppConfig.Workers' Error:Field validation for 'Workers' failed on the 'gte' tag
}

func ExampleValidate_fields() {
	_, err := cnfg.Parse(AppConfig{}, validator.Validate())

	var fields v10.ValidationErrors
	if errors.As(err, &fields) {
		for _, f := range fields {
			fmt.Println(f.Field(), f.Tag())
		}
	}

	// Output:
	// Addr required
	// Workers gte
}

func ExampleWith() {
	type Config struct {
		Addr string `check:"required"`
	}

	v := v10.New()
	v.SetTagName("check")

	_, err := cnfg.Parse(Config{}, validator.With(v))
	fmt.Println(err)

	// Output:
	// Key: 'Config.Addr' Error:Field validation for 'Addr' failed on the 'required' tag
}
