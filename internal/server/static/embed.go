package static

import (
	"embed"
)

//go:embed *
var static embed.FS

// getStaticFS returns the static embed.FS
func getStaticFS() embed.FS {
	return static
}
