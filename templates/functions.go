package templates

import (
	"text/template"

	"github.com/leonelquinteros/gorand"
)

// FuncMap is passed to the template object that renders every view
var FuncMap = template.FuncMap{
	"ID":      ID,
	"BuildID": BuildID,
}

// ID returns a random [8]byte value encoded as hex string (16).
func ID() string {
	id, err := gorand.GetHex(8)
	if err != nil {
		return "0123456789abcdef"
	}
	return id
}

var buildID string

func init() {
	// Generate buildID
	id, err := gorand.GetHex(8)
	if err != nil {
		buildID = "0123456789abcdef"
	} else {
		buildID = id
	}
}

// BuildID returns a randomly-generated-at-runtime constant value
func BuildID() string {
	return buildID
}
