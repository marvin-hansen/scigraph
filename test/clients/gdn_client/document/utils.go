package document

import (
	"fmt"
	"scigraph/src/clients/gdn_client"
)

func getTestInsertData() []byte {
	return []byte(` 
		[
		  {
			"item1": "data1"
		  },
		  {
			"item2": "data2"
		  }
		]
	`)
}

func getTestUpdateData(key string) []byte {
	s := fmt.Sprintf("[ {\n   \"_key\": \"%v\",\n   \"item1\": \"data42\"\n} ]", key)
	return []byte(s)
}

func printRes(res gdn_client.Responder, verbose bool) {
	if verbose {
		println(res.String())
	}
}
