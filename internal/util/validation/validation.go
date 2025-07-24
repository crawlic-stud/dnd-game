package validation

import (
	"fmt"
)

type BaseModel interface {
	Validate() error
}

type validator struct {
	checks       []bool
	descriptions []string
}

func NewValidator(model any) validator {
	return validator{}
}

func (v validator) Add(check bool, desc string) validator {
	v.checks = append(v.checks, check)
	v.descriptions = append(v.descriptions, desc)
	return v
}

func (v validator) CheckError(err error) validator {
	if err != nil {
		v.checks = append(v.checks, false)
		v.descriptions = append(v.descriptions, err.Error())
	} else {
		v.checks = append(v.checks, true)
		v.descriptions = append(v.descriptions, "")
	}
	return v
}

func (v validator) Validate() error {
	for ix, check := range v.checks {
		desc := v.descriptions[ix]
		if !check {
			return fmt.Errorf("validation error: %s", desc)
		}
	}
	return nil
}
