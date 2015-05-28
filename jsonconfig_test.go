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

	v, err := jcfg.GetSection(section).GetSection(nextSection).GetKey(key).GetString()
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

	var numbers int64
	numbers = 10
	vint, err := jcfg.GetSection(section).GetKey("numbers").GetInt()
	if err != nil {
		t.Error(err)
	}

	if vint != numbers {
		t.Errorf("Expected '%f'", vint)
	}

	var money float64
	money = 10.50
	vfloat, err := jcfg.GetSection(section).GetKey("money").GetFloat()
	if err != nil {
		t.Error(err)
	}

	if vfloat != money {
		t.Errorf("Expected '%f'", money)
	}
}
