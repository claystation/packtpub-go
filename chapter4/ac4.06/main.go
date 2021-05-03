package main

import "fmt"

func typeNames(v interface{}) string {
	switch v.(type) {
	case int, int32, int64:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func main() {
	data := []interface{}{1, 3.14, "hello", true, struct{}{}}

	for i := 0; i < len(data); i++ {
		fmt.Printf("%v is %v\n", data[i], typeNames(data[i]))
	}
}
