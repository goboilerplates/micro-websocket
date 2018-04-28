package api_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/goboilerplates/core"

	"github.com/goboilerplates/micro-websocket/api"
	"github.com/stretchr/testify/assert"
)

// SetupSample
func SetupSample() core.SampleEntity {
	return core.SampleEntity{Name: "Sample"}
}

// TestConvertInterfaceToJSONBytes .
func TestConvertInterfaceToJSONBytes(test *testing.T) {
	result := api.ConvertInterfaceToJSONBytes("Test", nil)
	assert.Equal(test, len(result) > 0, true)

	result = api.ConvertInterfaceToJSONBytes("Hello", errors.New("Error"))
	assert.Nil(test, result, nil)

	sample := SetupSample()

	result = api.ConvertInterfaceToJSONBytes(sample, nil)
	assert.Equal(test, len(result) > 0, true)

	result = api.ConvertInterfaceToJSONBytes(sample, errors.New("Error"))
	assert.Nil(test, result)
}

// TestConvertBytesToMap .
func TestConvertBytesToMap(test *testing.T) {
	sampleBytes, err := json.Marshal(SetupSample())
	assert.Nil(test, err)

	result, err := api.ConvertBytesToMap(sampleBytes)
	assert.Nil(test, err)
	assert.NotNil(test, result)
	assert.Equal(test, result["Name"], "Sample")
}
