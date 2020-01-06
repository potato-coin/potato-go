package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	potato "github.com/rise-worlds/potato-go"
)

func NewSetContract(account potato.AccountName, wasmPath, abiPath string) (out []*potato.Action, err error) {
	codeAction, err := NewSetCode(account, wasmPath)
	if err != nil {
		return nil, err
	}

	abiAction, err := NewSetABI(account, abiPath)
	if err != nil {
		return nil, err
	}

	return []*potato.Action{codeAction, abiAction}, nil
}

func NewSetCode(account potato.AccountName, wasmPath string) (out *potato.Action, err error) {
	codeContent, err := ioutil.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}

	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("setcode"),
		Authorization: []potato.PermissionLevel{
			{
				Actor:      account,
				Permission: potato.PermissionName("active"),
			},
		},
		ActionData: potato.NewActionData(SetCode{
			Account:   account,
			VMType:    0,
			VMVersion: 0,
			Code:      potato.HexBytes(codeContent),
		}),
	}, nil
}

func NewSetABI(account potato.AccountName, abiPath string) (out *potato.Action, err error) {
	abiContent, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}

	var abiPacked []byte
	if len(abiContent) > 0 {
		var abiDef potato.ABI
		if err := json.Unmarshal(abiContent, &abiDef); err != nil {
			return nil, fmt.Errorf("unmarshal ABI file: %s", err)
		}

		abiPacked, err = potato.MarshalBinary(abiDef)
		if err != nil {
			return nil, fmt.Errorf("packing ABI: %s", err)
		}
	}

	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("setabi"),
		Authorization: []potato.PermissionLevel{
			{
				Actor:      account,
				Permission: potato.PermissionName("active"),
			},
		},
		ActionData: potato.NewActionData(SetABI{
			Account: account,
			ABI:     potato.HexBytes(abiPacked),
		}),
	}, nil
}

// NewSetCodeTx is _deprecated_. Use NewSetContract instead, and build
// your transaction yourself.
func NewSetCodeTx(account potato.AccountName, wasmPath, abiPath string) (out *potato.Transaction, err error) {
	actions, err := NewSetContract(account, wasmPath, abiPath)
	if err != nil {
		return nil, err
	}
	return &potato.Transaction{Actions: actions}, nil
}

// SetCode represents the hard-coded `setcode` action.
type SetCode struct {
	Account   potato.AccountName `json:"account"`
	VMType    byte               `json:"vmtype"`
	VMVersion byte               `json:"vmversion"`
	Code      potato.HexBytes    `json:"code"`
}

// SetABI represents the hard-coded `setabi` action.
type SetABI struct {
	Account potato.AccountName `json:"account"`
	ABI     potato.HexBytes    `json:"abi"`
}
