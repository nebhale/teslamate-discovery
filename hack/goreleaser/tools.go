// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

//go:build tools

// This package imports things required by build scripts, to force `go mod` to
// see them as dependencies
package tools

import (
	_ "github.com/goreleaser/goreleaser"
)
