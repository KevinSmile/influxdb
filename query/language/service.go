// Package language exposes the flux parser as an interface.
package language

import (
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/parser"
	"github.com/influxdata/influxdb/query"
)

// DefaultService is the default language service.
var DefaultService query.LanguageService = defaultService{}

type defaultService struct{}

func (d defaultService) Parse(source string) *ast.Package {
	return parser.ParseSource(source)
}
