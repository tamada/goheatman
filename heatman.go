package heatman

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"strconv"
)

/*
Table represents the data for heat map.
*/
type Table struct {
	width   int
	height  int
	data    [][]*number
	context *Context
}

type number struct {
	value float64
}

/*
HeatmapConverter is an interface for converting the value to color.
*/
type HeatmapConverter interface {
	Convert(value float64) color.Color
}

/*
DefaultHeatmapConverter is default heatmap generator.
*/
type DefaultHeatmapConverter struct {
}

/*
Convert converts the given value to color for heatmap.
*/
func (dc *DefaultHeatmapConverter) Convert(value float64) color.Color {
	var hsb = createHSB(value)
	var r, g, b, a = hsb.RGBA()
	return color.RGBA{R: uint8(r & 0xff), G: uint8(g & 0xff), B: uint8(b & 0xff), A: uint8(a & 0xff)}
}

/*
GraymapConverter is grayscaled heatmap generator.
*/
type GraymapConverter struct {
}

/*
Convert converts the given value to color for graymap.
*/
func (gc *GraymapConverter) Convert(value float64) color.Color {
	var gray = uint8((1-value)*255 + 0.5)
	return color.RGBA{R: gray, G: gray, B: gray, A: 0xff}
}

/*
ScalerImage generates the image of scaler for heat map.
*/
func ScalerImage(context *Context) *Table {
	var data = [][]*number{}
	for i := 0; i < 10; i++ {
		var subd = []*number{}
		for j := 0; j < 255; j++ {
			subd = append(subd, &number{float64(j) / 255.0})
		}
		data = append(data, subd)
	}
	return &Table{width: 255, height: 10, context: context, data: data}
}

/*
ColorModel shows the color model of the heat map.
Always returns RGBAModel
*/
func (table *Table) ColorModel() color.Model {
	return color.RGBAModel
}

func (table *Table) calculateGap() (int, int) {
	var xCount = 0
	var yCount = 0
	if table.context.GapOfAdditionalLine > 0 {
		var gap = table.context.GapOfAdditionalLine
		xCount = table.width / gap
		yCount = table.height / gap
		if table.width%gap == 0 {
			xCount--
		}
		if table.height%gap == 0 {
			yCount--
		}
	}
	return xCount, yCount
}

/*
Bounds returns the size of the heatmap image.
*/
func (table *Table) Bounds() image.Rectangle {
	var scale = table.context.SizeOfAPixel
	var xCount, yCount = table.calculateGap()
	var r = image.Rect(0, 0, table.width*scale+xCount, table.height*scale+yCount)
	return r
}

func createHSB(value float64) *HSB {
	var hue = (1 - value) * 240 / 360
	return &HSB{
		Hue:        hue,
		Saturation: 1.0,
		Brightness: 1.0,
	}
}

/*
At returns color of heatmap image at given point.
*/
func (table *Table) At(x, y int) color.Color {
	var scale = table.context.SizeOfAPixel
	var gap = table.context.GapOfAdditionalLine
	var xx = x / scale
	var yy = y / scale
	var line = scale*gap + 1
	if gap > 0 {
		xx = (x - x/line) / scale
		yy = (y - y/line) / scale
	}
	if gap > 0 && (x != 0 && x%line == (line-1) || y != 0 && y%line == (line-1)) {
		return color.RGBA{0xff, 0xff, 0xff, 0x00}
	}
	if len(table.data[yy]) < xx || table.data[yy][xx] == nil {
		return color.RGBA{R: 0, G: 0, B: 0, A: 0}
	}
	return table.context.Converter.Convert(table.data[yy][xx].value)
}

/*
HSB shows a color for HSB (Hue, Saturation, and Brightness).
The range of each variable is [0, 1].
*/
type HSB struct {
	Hue, Saturation, Brightness float64
}

func (hsb *HSB) String() string {
	return fmt.Sprintf("hue: %f, saturation: %f, brightness: %f", hsb.Hue, hsb.Saturation, hsb.Brightness)
}

/*
RGBA converts HSB color to RGBA color.
This routine is refered from java.awt.Color#HSBtoRGB in amazon-corretto8
*/
func (hsb *HSB) RGBA() (uint32, uint32, uint32, uint32) {
	var r, g, b uint32 = 0, 0, 0
	if hsb.Saturation == 0.0 {
		r = uint32(hsb.Brightness*255.0 + 0.5)
		return r, r, r, 255
	}
	var h = (hsb.Hue - math.Floor(hsb.Hue)) * 6.0
	var f = h - math.Floor(h)
	var p = hsb.Brightness * (1.0 - hsb.Saturation)
	var q = hsb.Brightness * (1.0 - hsb.Saturation*f)
	var t = hsb.Brightness * (1.0 - (hsb.Saturation * (1.0 - f)))
	switch int(h) {
	case 0:
		r = uint32(hsb.Brightness*255 + 0.5)
		g = uint32(t*255.0 + 0.5)
		b = uint32(p*255.0 + 0.5)
	case 1:
		r = uint32(q*255.0 + 0.5)
		g = uint32(hsb.Brightness*255.0 + 0.5)
		b = uint32(p*255.0 + 0.5)
	case 2:
		r = uint32(p*255.0 + 0.5)
		g = uint32(hsb.Brightness*255.0 + 0.5)
		b = uint32(t*255.0 + 0.5)
	case 3:
		r = uint32(p*255.0 + 0.5)
		g = uint32(q*255.0 + 0.5)
		b = uint32(hsb.Brightness*255.0 + 0.5)
	case 4:
		r = uint32(t*255.0 + 0.5)
		g = uint32(p*255.0 + 0.5)
		b = uint32(hsb.Brightness*255.0 + 0.5)
	case 5:
		r = uint32(hsb.Brightness*255.0 + 0.5)
		g = uint32(p*255.0 + 0.5)
		b = uint32(q*255.0 + 0.5)
	}
	return r, g, b, 0xff
}
func (model *HeaderModel) hasColumnHeader() bool {
	return *model == ColumnHeader || *model == RowColumnHeader
}

func (model *HeaderModel) hasRowHeader() bool {
	return *model == RowHeader || *model == RowColumnHeader
}

/*
NewTable generates an instance of Table.
The range of each value on given csv file must be [0, 1].
*/
func NewTable(reader *csv.Reader, context *Context) *Table {
	var data = [][]*number{}
	var width = 0
	var first = true

	for height := 0; ; height++ {
		var row, err = reader.Read()
		if first && context.WithHeader.hasColumnHeader() {
			first = false
			height--
			continue
		}
		if err == io.EOF {
			return &Table{width: width, height: height, data: data, context: context}
		}
		var rowNums = convert(row, context)
		if width < len(rowNums) {
			width = len(rowNums)
		}
		data = append(data, rowNums)
		first = false
	}
}

func convert(values []string, context *Context) []*number {
	var results = []*number{}
	var first = true
	for _, value := range values {
		if first && context.WithHeader.hasRowHeader() {
			first = false
			continue
		}
		var val, err = strconv.ParseFloat(value, 64)
		if err == nil {
			results = append(results, &number{val})
		} else {
			results = append(results, nil)
		}
		first = false
	}
	return results
}
