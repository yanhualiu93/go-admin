package custom

import (
	"go-admin/common/errors"
	"strings"
)

type Map map[string]interface{}

func (m Map) Set(key string, val interface{}) Map {
	if m == nil {
		m = make(Map)
	}
	m[key] = val
	return m
}
func (m Map) Get(key string) (interface{}, error) {
	if v, ok := m[key]; ok {
		return v, nil
	} else {
		return nil, errors.NewParamRequiredError(key)
	}
}
func (m Map) GetString(key string) (string, error) {
	val, err := m.Get(key)
	if err != nil {
		return "", err
	}
	switch v := val.(type) {
	case string:
		if len(strings.TrimSpace(v)) == 0 {
			return "", errors.NewParamRequiredError(key)
		}
		return v, nil
	default:
		return "", errors.NewParamNonstandardError(key)
	}
}
func (m Map) GetInt64(key string) (int64, error) {
	val, err := m.Get(key)
	if err != nil {
		return 0, err
	}
	switch v := val.(type) {
	case int64:
		return v, nil
	default:
		return 0, errors.NewParamNonstandardError(key)
	}
}
func (m Map) GetFloat64(key string) (float64, error) {
	val, err := m.Get(key)
	if err != nil {
		return 0, err
	}
	switch v := val.(type) {
	case float64:
		return v, errors.NewParamNonstandardError(key)
	default:
		return 0, errors.NewParamNonstandardError(key)
	}
}

func (m Map) Del(key string) Map {
	delete(m, key)
	return m
}
