package jtrans

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransformtion(t *testing.T) {
	assert := require.New(t)

	input := `{"first_name":"Dean","last_name":"Karn","nested":{"data":"nested data value"}}`
	var m map[string]interface{}
	err := json.Unmarshal([]byte(input), &m)
	assert.NoError(err)

	tf := Transformation{
		M.Build(Mapping, "first_name", "fname"),
		M.Build(Mapping, "last_name", "lname"),
		M.Build(Mapping, "nested.data", "nested_data"),
		M.Build(Constant, "track", "type"),
	}

	results := tf.Transform(m)
	b, err := json.Marshal(results)
	assert.NoError(err)
	assert.Equal(`{"fname":"Dean","lname":"Karn","nested_data":"nested data value","type":"track"}`, string(b))
}
