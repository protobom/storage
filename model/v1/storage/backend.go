package storage

import (
	"github.com/bom-squad/protobom/pkg/sbom"
	"github.com/protobom/storage/pkg/options"
)

type Backend interface {
	Store(*sbom.Document, *options.StoreOptions) error
	Retrieve(string, *options.RetrieveOptions) (*sbom.Document, error)
}
