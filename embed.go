// Package conductorr exists at the root level of the Conductorr source code in order
// to embed the different folders to be accessible in the binary.
package conductorr

import "embed"

//go:embed web/build
var WebDist embed.FS

//go:embed dist
var CSLDist embed.FS

//go:embed migrations
var Migrations embed.FS
