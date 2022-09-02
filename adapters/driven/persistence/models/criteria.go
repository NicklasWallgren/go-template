package models

import "github.com/mitchellh/mapstructure"

type Criteria struct {
	Title string
}

func (c Criteria) ToWhereMap() (*map[string]interface{}, error) {
	result := &map[string]interface{}{}

	// TODO, does not support operators

	return result, mapstructure.Decode(c, &result)
}
