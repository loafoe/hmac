package alerts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayloadParse(t *testing.T) {
	var payload Payload

	data := `
	{
		"status": "resolved",
		"alerts": [
		  {
			"labels": {
			  "alertname": "foo"
			}
		  }
		]
	  }
	`
	err := json.Unmarshal([]byte(data), &payload)
	assert.Nil(t, err)
	assert.Equal(t, "resolved", payload.Status)
}
