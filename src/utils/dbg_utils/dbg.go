// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package dbg_utils

import (
	"log"
)

func DbgCheck(dbg bool) {
	if dbg {
		log.Println("Debug mode on!")
	}
}

func DbgPrint(dbg bool, msg string) {
	if dbg {
		log.Println(msg)
	}
}

func DbgCheckPrintLog(dbg bool, main, mtd, msg string, err error) {
	if dbg {
		if err != nil {
			log.Print("Main: ", main)
			log.Print("Mtd: ", mtd)
			log.Print("Msg: ", msg)
			log.Print("Error: ", err.Error())
		}
	}
}
