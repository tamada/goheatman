package internal

import (
	"fmt"
	"strings"

	heatman "github.com/tamada/goheatman"
)

/*
Version represents the version of the product.
*/
const Version = "1.0.0"

func ParseHeaderModel(model string) (heatman.HeaderModel, error) {
	switch strings.ToLower(model) {
	case "both":
		return heatman.RowColumnHeader, nil
	case "column":
		return heatman.ColumnHeader, nil
	case "row":
		return heatman.RowHeader, nil
	case "no":
		return heatman.NoHeaders, nil
	}
	return heatman.InvalidHeaderModel, fmt.Errorf("%s: unknown header model", model)
}

func ParseColorType(colorType string) (heatman.HeatmapConverter, error) {
	switch strings.ToLower(colorType) {
	case "default":
		return &heatman.DefaultHeatmapConverter{}, nil
	case "gray":
		return &heatman.GraymapConverter{}, nil
	}
	return nil, fmt.Errorf("%s: unknown color type", colorType)
}
