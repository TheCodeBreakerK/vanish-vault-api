// Package migrations contains the SQL migration files embedded into the binary.
package migrations

import "embed"

// FS holds the embedded SQL migration files.
//
//go:embed *.sql
var FS embed.FS
