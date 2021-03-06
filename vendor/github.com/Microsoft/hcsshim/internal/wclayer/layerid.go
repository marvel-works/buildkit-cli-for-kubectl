package wclayer

import (
	"path/filepath"

	"github.com/Microsoft/go-winio/pkg/guid"
)

// LayerID returns the layer ID of a layer on disk.
func LayerID(path string) (guid.GUID, error) {
	_, file := filepath.Split(path)
	return NameToGuid(file)
}
