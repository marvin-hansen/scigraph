// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package dbg_utils

import (
	"fmt"
	"github.com/go-numb/go-ftx/rest/private/orders"
)

func DbgTriggerOrderResponse(dbg bool, orderResponse *orders.ResponseForPlaceTriggerOrder) {
	if dbg {
		if orderResponse == nil {
			println("ResponseForPlaceTriggerOrder nil; return!")
			return
		}

		println("Order ID: ", orderResponse.ID)
		println("Order Status: ", orderResponse.Status)
		println("Order Side: ", orderResponse.Side)
		println("Order Type: ", orderResponse.Type)
		println("Order Price: ", fmt.Sprintf("%.4f", orderResponse.OrderPrice))
		println("Order Size: ", fmt.Sprintf("%.4f", orderResponse.Size))
		println("Trigger Price: ", fmt.Sprintf("%.4f", orderResponse.TriggerPrice))
	}
}

func DbgOrderResponse(dbg bool, orderResponse *orders.ResponseForPlaceOrder) {
	if dbg {
		if orderResponse == nil {
			println("ResponseForPlaceOrder; return!")
			return
		}
		println("Order ID: ", orderResponse.ID)
		println("Order Status: ", orderResponse.Status)
		println("Order Side: ", orderResponse.Side)
		println("Order Type: ", orderResponse.Type)
		println("Order Price: ", fmt.Sprintf("%.4f", orderResponse.Price))
		println("Order Size: ", fmt.Sprintf("%.4f", orderResponse.Size))
	}
}
