// Code generated by 'yaegi extract github.com/common-nighthawk/go-figure'. DO NOT EDIT.

package figure

import (
	"github.com/common-nighthawk/go-figure"
	"reflect"
)

func init() {
	Symbols["github.com/common-nighthawk/go-figure/figure"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Asset":             reflect.ValueOf(figure.Asset),
		"AssetDir":          reflect.ValueOf(figure.AssetDir),
		"AssetInfo":         reflect.ValueOf(figure.AssetInfo),
		"AssetNames":        reflect.ValueOf(figure.AssetNames),
		"MustAsset":         reflect.ValueOf(figure.MustAsset),
		"NewColorFigure":    reflect.ValueOf(figure.NewColorFigure),
		"NewFigure":         reflect.ValueOf(figure.NewFigure),
		"NewFigureWithFont": reflect.ValueOf(figure.NewFigureWithFont),
		"RestoreAsset":      reflect.ValueOf(figure.RestoreAsset),
		"RestoreAssets":     reflect.ValueOf(figure.RestoreAssets),
		"Write":             reflect.ValueOf(figure.Write),
	}
}
