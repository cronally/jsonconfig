package jsonconfig

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	section := "meta"
	nextSection := "home"
	key := "title"
	value := "this is a title"

	jcfg, err := NewJsonConfig("test.json")
	if err != nil {
		t.Error(err)
	}

	v, err := jcfg.GetSection(section).GetSection(nextSection).GetKey(key).GetValue()
	if err != nil {
		t.Error(err)
	}

	if v != value {
		t.Errorf("Expected '%s'", value)
	}
}

func TestTypes(t *testing.T) {
	section := "various"

	jcfg, err := NewJsonConfig("test.json")
	if err != nil {
		t.Error(err)
	}

	var numbers float64
	numbers = 10
	v, err := jcfg.GetSection(section).GetKey("numbers").GetValue()
	if err != nil {
		t.Error(err)
	}

	if v != numbers {
		t.Errorf("Expected '%f'", numbers)
	}

	var money float64
	money = 10.50
	v, err = jcfg.GetSection(section).GetKey("money").GetValue()
	if err != nil {
		t.Error(err)
	}

	if v != money {
		t.Errorf("Expected '%f'", money)
	}
}
