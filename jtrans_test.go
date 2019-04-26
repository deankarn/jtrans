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
		{Type: Mapping, From: []string{"first_name"}, To: []string{"fname"}},
		{Type: Mapping, From: []string{"last_name"}, To: []string{"lname"}},
		{Type: Mapping, From: []string{"nested", "data"}, To: []string{"nested_data"}},
		{Type: Constant, From: []string{"track"}, To: []string{"type"}},
	}

	results := Transform(tf, m)
	b, err := json.Marshal(results)
	assert.NoError(err)
	assert.Equal(`{"fname":"Dean","lname":"Karn","nested_data":"nested data value","type":"track"}`, string(b))
}
