// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package dbg_utils

import (
	"fmt"
	"log"
)

// DbgStringerObject prints any struct or type that implements the stringer interface.
func DbgStringerObject(dbg bool, msg string, strObj fmt.Stringer) {
	if dbg {
		if strObj == nil {
			println("object nil; return!")
			return
		}
		log.Println(msg)
		log.Println(strObj.String())
		log.Println()
	}
}

func DbgOrderBook(dbg bool, orderBook [][]float64) {
	if dbg {
		if len(orderBook) == 0 {
			println("order book empty; return!")
			return
		}

		println("=====================")
		println("==== Order Book ====")
		println("=====================")

		for i := range orderBook {
			str := fmt.Sprintf("Entry: %v, Price: %v, Size: %v \n",
				i,
				orderBook[i][0],
				orderBook[i][1],
			)
			println(str)
			println("=====================")
		}
	}
}
