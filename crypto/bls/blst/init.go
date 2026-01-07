//go:build ((linux && amd64) || (linux && arm64) || (darwin && amd64) || (darwin && arm64) || (windows && amd64)) && !blst_disabled

// Package blst implements a go-wrapper around a library implementing the
// BLS12-381 curve and signature scheme. This package exposes a public API for
// verifying and aggregating BLS signatures used by Ethereum.
//
// This implementation uses the library written by Supranational, blst.
package blst

import (
	"runtime"

	blst "github.com/supranational/blst/bindings/go"
)

func init() {
	// Reserve 1 core for general application work
	maxProcs := runtime.GOMAXPROCS(0) - 1
	if maxProcs <= 0 {
		maxProcs = 1
	}
	blst.SetMaxProcs(maxProcs)
}
