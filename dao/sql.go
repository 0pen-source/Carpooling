package dao

import (
	"fmt"
	"reflect"
	"strings"
)

func getsql(sql string, paras ...interface{}) string {
	for _, para := range paras {
		newstr := paraToString(para)
		sql = strings.Replace(sql, "?", newstr, 1)
	}

	return sql
}



func paraToString(para interface{}) string {
	var newstr string

	switch reflect.TypeOf(para).Kind() {
	case reflect.String:
		newstr = fmt.Sprintf(`"%s"`, para.(string))
	case reflect.Int:
		newstr = fmt.Sprintf("%d", para.(int))
	case reflect.Int64:
		newstr = fmt.Sprintf("%d", para.(int64))
	case reflect.Float64:
		newstr = fmt.Sprintf("%f", para.(float64))
	case reflect.Slice:
		newstr_array := make([]string, 0)
		s := reflect.ValueOf(para)
		for i := 0; i < s.Len(); i++ {
			newstr_array = append(newstr_array, paraToString(s.Index(i)))
		}
		newstr = "(" + strings.Join(newstr_array, ",") + ")"
	default:
		newstr = fmt.Sprintf(`"%v"`, para)
	}

	return newstr
}
