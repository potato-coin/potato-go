package potato_test

import (
	"encoding/json"
	"fmt"

	potato "github.com/rise-worlds/potato-go"
)

func ExampleAPI_GetInfo() {
	api := potato.New(getAPIURL())

	info, err := api.GetInfo()
	if err != nil {
		panic(fmt.Errorf("get info: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}
