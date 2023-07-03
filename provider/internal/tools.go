package internal

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

func SafeDeleteKey(value interface{}, key string) error {
	if reflect.ValueOf(value).Kind() != reflect.Map {
		return errors.New("Not a map")
	}

	mapValue := value.(map[string]interface{})

	if strings.Contains(key, ".") {
		keys := strings.SplitN(key, ".", 2)
		topKey := keys[0]
		tailKey := keys[1]
		val, ok := mapValue[topKey]
		if !ok {
			return errors.New("No key")
		}
		return SafeDeleteKey(val, tailKey)
	} else {
		if _, ok := mapValue[key]; !ok {
			return errors.New("No key")
		}
		delete(mapValue, key)
	}

	return nil
}

func Convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			if v != nil {
				m2[k.(string)] = Convert(v)
			}
		}
		return m2
	case map[string]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			if v != nil {
				m2[k] = Convert(v)
			}
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = Convert(v)
		}
	}
	return i
}
