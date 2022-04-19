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

func getTestReplaceData(key string) []byte {
	str := fmt.Sprintf(` 
		[
		  {
			"_key": "%v",
			"item1": "dataReplaced"
		  }
	]
	`, key)
	return []byte(str)
}

func getKeysToDelete(key1, key2 string) []byte {
	str := fmt.Sprintf(` 
		[
		"%v",
		"%v"
		]
	`, key1, key2)
	return []byte(str)
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
