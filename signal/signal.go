// Package signal : Handling SIGNAL with context package
//
// These codes are licensed under CC0.
// http:// creativecommons.org/publicdomain/zero/1.0/
//
// Deprecated: This package is deprecated. Use [os/signal.NotifyContext] instead,
// which provides equivalent functionality as part of the Go standard library since Go 1.16.
package signal

import (
	"context"
	"os"
	signl "os/signal"
)

// Context returns context.Context with Cancel
//
// Deprecated: Use [os/signal.NotifyContext] instead.
func Context(parent context.Context, sig ...os.Signal) context.Context {
	cctx, cancel := context.WithCancel(parent)
	go func() {
		defer cancel()

		sigCh := make(chan os.Signal, 1)
		signl.Notify(sigCh, sig...)
		defer signl.Stop(sigCh)

		select {
		case <-cctx.Done(): // cancel event from parent context
			return
		case <-sigCh: // catch SIGNAL
			return
		}
	}()
	return cctx
}
