package potato_test

import (
	"encoding/json"
	"fmt"

	potato "github.com/rise-worlds/potato-go"
)

func ExampleAPI_GetAccount() {
	api := potato.New(getAPIURL())

	account := potato.AccountName("poc.rex")
	info, err := api.GetAccount(account)
	if err != nil {
		if err == potato.ErrNotFound {
			fmt.Printf("unknown account: %s", account)
			return
		}

		panic(fmt.Errorf("get account: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}
