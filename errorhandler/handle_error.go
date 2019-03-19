package errorhandler

import "log"

func HandleErr(err error, codeblock uint) {
	log.Printf("Error at codeblock %d", codeblock)
	log.Fatalf("Error occurred while deploying contract: %s", err)
}