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

	if jcfg.GetSection(section).GetSection(nextSection).GetKey(key).GetValue() != value {
		t.Errorf("Expected '%s'", value)
	}
}
