// Copyright (c) 2022-2022. Marvin Hansen | marvin.hansen@gmail.com

package dbg_utils

import (
	"log"
)

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
