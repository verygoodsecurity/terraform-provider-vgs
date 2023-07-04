package internal

import (
	"encoding/json"
	asserting "github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSafeDeleteKey_KeyExists(t *testing.T) {
	assert := asserting.New(t)

	var value interface{}
	err := json.Unmarshal([]byte(`{"attributes":{"updated_at": "2023-07-03T14:54:57", "id": "eccf1542-820b-4e7e-8b3f-03c4d8639f9d"}}`), &value)
	assert.Nil(err)

	var expectedValue interface{}
	err = json.Unmarshal([]byte(`{"attributes":{"id": "eccf1542-820b-4e7e-8b3f-03c4d8639f9d"}}`), &expectedValue)
	assert.Nil(err)

	err = SafeDeleteKey(value, "attributes.updated_at")
	assert.Nil(err)
	assert.True(reflect.DeepEqual(value, expectedValue))
}

func TestSafeDeleteKey_KeyNotExists(t *testing.T) {
	assert := asserting.New(t)

	var value interface{}
	err := json.Unmarshal([]byte(`{"attributes":{"updated_at": "2023-07-03T14:54:57", "id": "eccf1542-820b-4e7e-8b3f-03c4d8639f9d"}}`), &value)
	assert.Nil(err)

	err = SafeDeleteKey(value, "attributes.updated_at.no_key")
	assert.NotNil(err)
}

func TestSafeDeleteKey_KeyNotExists2(t *testing.T) {
	assert := asserting.New(t)

	var value interface{}
	err := json.Unmarshal([]byte(`{"attributes":{"id": "eccf1542-820b-4e7e-8b3f-03c4d8639f9d"}}`), &value)
	assert.Nil(err)

	err = SafeDeleteKey(value, "attributes.updated_at")
	assert.NotNil(err)
}

func TestSafeDeleteKey_NotMap(t *testing.T) {
	assert := asserting.New(t)

	err := SafeDeleteKey("string", "attributes.updated_at")
	assert.NotNil(err)

}

func TestConvert(t *testing.T) {
	assert := asserting.New(t)

	var value interface{}
	err := json.Unmarshal([]byte(`{"attributes":{"id": "eccf1542-820b-4e7e-8b3f-03c4d8639f9d", "value": null}}`), &value)
	assert.Nil(err)

	var expectedValue interface{}
	err = json.Unmarshal([]byte(`{"attributes":{"id": "eccf1542-820b-4e7e-8b3f-03c4d8639f9d"}}`), &expectedValue)
	assert.Nil(err)

	value = Convert(value)

	assert.True(reflect.DeepEqual(value, expectedValue))
}
