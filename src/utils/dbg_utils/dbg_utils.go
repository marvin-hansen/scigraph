// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com

package dbg_utils

func NilCheck(value interface{}, msg string) {
	if value == nil {
		println(msg)
		panic(value)
	}
}

func EmptyStringCheck(value string, msg string) {
	if value == "" {
		println(msg)
		panic(value)
	}
}

func BooleanCheck(actualValue, expectedValue bool, msg string) {
	if actualValue != expectedValue {
		println(msg)
		panic(msg)
	}
}
