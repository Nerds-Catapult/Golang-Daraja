package darajaAuth

import "reflect"

func struct2Map(structure interface{}) map[string]interface{}{
	m := make(map[string]interface{})
	v := reflect.ValueOf(structure)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		m[v.Type().Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	return m
}