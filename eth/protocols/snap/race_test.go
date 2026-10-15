// Copyright 2026 The go-ethereum Authors
// This file is part of the go-ethereum library.

//go:build race

package snap

func init() { raceEnabled = true }
