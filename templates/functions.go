package templates

import (
	"text/template"

	"github.com/leonelquinteros/gorand"
)

// FuncMap is passed to the template object that renders every view
var FuncMap = template.FuncMap{
	"ID": ID,
}

// ID returns a random [8]byte value encoded as hex string.
func ID() string {
	id, err := gorand.GetHex(8)
	if err != nil {
		return "0"
	}
	return id
}
