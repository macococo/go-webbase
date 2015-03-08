package io

import (
	"encoding/json"
	"github.com/macococo/go-webbase/utils"
)

func ToJsonBytes(v interface{}) []byte {
	content, err := json.Marshal(v)
	utils.HandleError(err)

	return content
}

func FromJsonBytes(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	utils.HandleError(err)
}
