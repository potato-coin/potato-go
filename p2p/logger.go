package p2p

import (
	"fmt"

	potato "github.com/rise-worlds/potato-go"
	"go.uber.org/zap"
)

// Just Use same patterns to potato-go/logger.go
// TODO be improved in terms of external package integration for logger by potato-go

// logger default use nil zap logger
var p2pLog = zap.NewNop()

// EnableP2PLogging enable p2p package to log by zap
func EnableP2PLogging() {
	p2pLog = potato.NewLogger(false)
}

// logErr log err msg by p2pLog
func logErr(msg string, err error) {
	p2pLog.Error(msg, zap.Error(err))
}

// SyncLogger sync logger, should `defer SyncLogger()` when use p2p package
func SyncLogger() {
	err := p2pLog.Sync()
	if err != nil {
		// logger is err, log by fmt
		fmt.Printf("err by sync p2p logger %s", err.Error())
	}
}
