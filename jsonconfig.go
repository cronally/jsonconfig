package jsonconfig

import (
	"encoding/json"
	"io/ioutil"
)

// A type that satisfies JsonConfigReader must be able to
// get a section by name, key by name, and the value of
// that key
type JsonConfigReader interface {
	GetSection(name string) JsonConfigReader
	GetKey(key string) JsonConfigReader
	GetValue() string
}

// JsonConfig holds the filename and parsed JSON data
type JsonConfig struct {
	Filename string
	Parsed   interface{}
}

// GetValue gets the value of the current level in the parsed
// JSON tree
func (jc *JsonConfig) GetValue() string {
	return jc.Parsed.(string)
}

// GetKey returns a JsonConfigReader containing the contents of
// the specified key
func (jc *JsonConfig) GetKey(key string) JsonConfigReader {
	return jc.GetSection(key)
}

// GetSection returns a JsonConfigReader fontaining the contents of
// a parent node in the JSON structure
func (jc *JsonConfig) GetSection(name string) JsonConfigReader {
	return &JsonConfig{
		jc.Filename,
		jc.Parsed.(map[string]interface{})[name],
	}

}

// NewJsonConfig loads a file containing JSON data by filename and
// returns a JsonConfig with the parsed data
func NewJsonConfig(filename string) (*JsonConfig, error) {
	var j interface{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &j)
	if err != nil {
		return nil, err
	}

	return &JsonConfig{filename, j}, nil
}
