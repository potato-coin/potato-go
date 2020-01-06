Potato API library for Go
=========================

[点击查看中文版](./README-cn.md)

[![GoDoc](https://godoc.org/github.com/rise-worlds/potato-go?status.svg)](https://godoc.org/github.com/rise-worlds/potato-go)

This library provides simple access to data structures (binary packing
and JSON interface) and API calls to an Potato RPC server, running
remotely or locally.  It provides wallet functionalities (KeyBag), or
can sign transaction through the `kpocd` wallet. It also knows about
the P2P protocol on port 9876.

As of before the June launch, this library is pretty much in
flux. Don't expect stability, as we're moving alongside the main
`potato` codebase, which changes very fast.

This library is the basis for the `potato-bios` launch orchestrator tool
at https://github.com/rise-worlds/potato-bios

Basic usage
-----------

```go
api := potato.New("http://potato.jocky123.com")

infoResp, _ := api.GetInfo()
accountResp, _ := api.GetAccount("initn")
fmt.Println("Permission for initn:", accountResp.Permissions[0].RequiredAuth.Keys)
```

`potato.system` and `poc.token` contract _Actions_ are respectively in:
* https://github.com/rise-worlds/potato-go/tree/master/system ([godocs](https://godoc.org/github.com/rise-worlds/potato-go/system))
* https://github.com/rise-worlds/potato-go/tree/master/token ([godocs](https://godoc.org/github.com/rise-worlds/potato-go/token))

Example
-------

See example usages of the library:

* https://github.com/rise-wrolds/potato-bios/blob/master/bios/bios.go
* https://github.com/rise-wrolds/potato-bios/blob/master/bios/ops.go
* Some other `main` packages under `cmd/`.

Contributing
------------

Any contributions are welcome, use your standard GitHub-fu to pitch in and improve.

License
-------

MIT
