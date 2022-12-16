package helpers

import "reflect"

func NormalizeFieldNameByTag(myStruct any, fieldName string, tagName string) string {
	field, exists := reflect.TypeOf(myStruct).Elem().FieldByName(fieldName)
	if !exists {
		return fieldName
	}

	fieldNameByTag, exists := field.Tag.Lookup(tagName)
	if !exists {
		return fieldName
	}

	return fieldNameByTag
}
