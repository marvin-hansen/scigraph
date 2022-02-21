package printing_utils

import (
	"log"
	"time"
)

func PrintStartHeader(serviceName string, port string, elapsed time.Duration) {
	log.Println()
	log.Println(serviceName)
	log.Printf("Service start time (Milliseconds): %d", elapsed.Milliseconds())
	log.Println("========================================== ")
	log.Println(" Health check at: 	host" + port + "/health")
	log.Println("========================================== ")
	log.Println()
}

func PrintStopHeader(elapsed time.Duration) {
	log.Println()
	log.Printf("Service shutdown took (Milliseconds): %d", elapsed.Milliseconds())
	log.Println("========================================== ")
	log.Println(" Shutdown complete - Switch off now!")
	log.Println("========================================== ")
	log.Println()
}
