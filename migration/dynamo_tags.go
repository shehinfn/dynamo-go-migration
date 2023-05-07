package migration

import (
	"reflect"
	"strings"
)

type DynamoTag struct {
	AttributeName string
	AttributeType string
	HashKey       bool
}

func ParseDynamoTags(model interface{}) []DynamoTag {
	tags := []DynamoTag{}

	typ := reflect.TypeOf(model)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tagValue := field.Tag.Get("dynamo")

		if tagValue == "" {
			continue
		}

		tag := DynamoTag{}
		tag.AttributeName = field.Name

		tagParts := strings.Split(tagValue, ",")
		for _, tagPart := range tagParts {
			if tagPart == "hash" {
				tag.HashKey = true
			} else {
				tag.AttributeType = strings.TrimSpace(tagPart)
			}
		}

		tags = append(tags, tag)
	}

	return tags
}
