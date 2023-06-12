package figure

import (
	_ "github.com/szkiba/xk6-g0"
	"github.com/szkiba/xk6-g0/g0"
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

var Symbols = interp.Exports{}

//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract -name figure github.com/common-nighthawk/go-figure

func exports(vu modules.VU) interp.Exports {
	return Symbols
}

func init() {
	g0.RegisterExports(exports)
}
