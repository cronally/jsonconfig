package jsonconfig

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// A type that satisfies JsonConfigReader must be able to
// get a section by name, key by name, and the value of
// that key
type JsonConfigReader interface {
	GetSection(name string) JsonConfigReader
	GetKey(key string) JsonConfigReader
	GetString() (string, error)
	GetFloat() (float64, error)
	GetInt() (int64, error)
}

// JsonConfig holds the filename and parsed JSON data
type JsonConfig struct {
	Filename string
	Parsed   interface{}
}

// GetString gets the string value of the current level in the parsed
// JSON tree
func (jc *JsonConfig) GetString() (string, error) {
	switch t := jc.Parsed.(type) {
	default:
		return "", errors.New(fmt.Sprintf("Unknown type: %T", t))

	case json.Number:
		return t.String(), nil

	case string:
		return jc.Parsed.(string), nil
	}

}

// GetFloat gets the string value of the current level in the parsed
// JSON tree
func (jc *JsonConfig) GetFloat() (float64, error) {
	switch t := jc.Parsed.(type) {
	default:
		return 0, errors.New(fmt.Sprintf("Unknown type: %T", t))

	case json.Number:
		v, err := t.Float64()
		if err != nil {
			return 0, err
		}
		return v, nil

	}

}

// GetInt gets the string value of the current level in the parsed
// JSON tree
func (jc *JsonConfig) GetInt() (int64, error) {
	switch t := jc.Parsed.(type) {
	default:
		return 0, errors.New(fmt.Sprintf("Unknown type: %T", t))

	case json.Number:
		v, err := t.Int64()
		if err != nil {
			return 0, err
		}
		return v, nil

	}

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

	decoder := json.NewDecoder(bytes.NewBuffer(data))
	decoder.UseNumber()
	err = decoder.Decode(&j)
	if err != nil {
		return nil, err
	}

	return &JsonConfig{filename, j}, nil
}
