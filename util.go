package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Parameterize(filter interface{}, hasFirstParam bool) string {
	var result string

	v := reflect.ValueOf(filter)
    typeOfS := v.Type()

    for i := 0; i< v.NumField(); i++ {
		field := v.Field(i)
		if field.IsZero() { continue }
		// if field.IsNil() { continue }

		val := field.Interface()
		name := strings.ToLower(typeOfS.Field(i).Name)
		joinChar := "&"
		if hasFirstParam && Empty(result) { joinChar = "?" }
		result += fmt.Sprintf("%s%s=%s", joinChar, name, val)
    }

	return result
}

func Empty(str string) bool { return str == "" }

func ParameterizePagination(pagination Pagination) string {
	type PaginationStr struct {
		Page string
		Page_Size string
	}
	page := "1" 
	pageSize := "50" 
	if pagination.Page != 0 { page = string(rune(pagination.Page)) }
	if pagination.PageSize != 0 { pageSize = string(rune(pagination.PageSize) ) }

	return Parameterize(PaginationStr{
		Page: page,
		Page_Size: pageSize,
	}, false)
}