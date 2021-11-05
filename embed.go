package conductorr

import "embed"

//go:embed web/build
var WebDist embed.FS
//go:embed dist
var CSLDist embed.FS
//go:embed migrations
var Migrations embed.FS
