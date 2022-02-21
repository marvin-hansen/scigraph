// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package dbg_utils

import (
	"fmt"
	"github.com/go-numb/go-ftx/rest/private/orders"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func DbgRequest(dbg bool, r http.Request) {
	if dbg {
		log.Println("JSON request")
		data, _ := ioutil.ReadAll(r.Body)
		log.Println(string(data))
	}
}

func DbgLeverageReq(dbg bool, sym, key string, leverage int) {
	if dbg {
		log.Println("Extracted symbol: " + sym)
		log.Println("Extracted key: " + key)
		log.Println("Extracted leverage: " + strconv.FormatInt(int64(leverage), 10))
	}
}

// These request don't implement stringer interface and as such require separate dbgPrint functions.

func DbgRequestForPlaceOrder(dbg bool, msg string, orderRequest *orders.RequestForPlaceOrder) {
	if orderRequest == nil {
		log.Println("orderRequest nil; return!")
		return
	}
	if dbg {
		log.Println()
		log.Println(msg)
		log.Println()
		log.Println("Market: ", orderRequest.Market)
		log.Println("Type: ", orderRequest.Type)
		log.Println("Side: ", orderRequest.Side)
		log.Println("Price: ", fmt.Sprintf("%.2f", orderRequest.Price))
	}
}

func DbgRequestForPlaceTriggerOrder(dbg bool, orderRequest *orders.RequestForPlaceTriggerOrder) {
	if orderRequest == nil {
		println("orderRequest nil; return!")
		return
	}
	if dbg {
		println("Market: ", orderRequest.Market)
		println("Type: ", orderRequest.Type)
		println("Side: ", orderRequest.Side)
		println("Trigger Price: ", fmt.Sprintf("%.4f", orderRequest.TriggerPrice))
		println("Order Price: ", fmt.Sprintf("%.4f", orderRequest.OrderPrice))
	}
}

func DbgRequestForPlaceTrailingStopOrder(dbg bool, orderRequest *orders.RequestForPlaceTriggerOrder) {
	if orderRequest == nil {
		println("orderRequest nil; return!")
		return
	}
	if dbg {
		println("Market: ", orderRequest.Market)
		println("Type: ", orderRequest.Type)
		println("Side: ", orderRequest.Side)
		println("TrailValue: ", fmt.Sprintf("%.4f", orderRequest.TrailValue))
	}
}
