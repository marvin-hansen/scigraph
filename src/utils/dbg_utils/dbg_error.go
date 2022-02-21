package dbg_utils

import (
	"log"
	"os"
)

func CheckError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func CheckPrintErr(err error, errorMsg string) {
	if err != nil {
		log.Println("Error:", err.Error())
		log.Println("Error Message: ", errorMsg)
	}
}

func CheckPrintErrStop(err error, errorMsg string) {
	if err != nil {
		log.Println("Error:", err.Error())
		log.Println("Error Message: ", errorMsg)
		os.Exit(42)
	}
}
