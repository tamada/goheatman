package heatman

import (
	"encoding/csv"
	"image/color"
	"math"
	"strings"
	"testing"
)

func TestCsvToTable(t *testing.T) {
	var csvData = `,c1,c2,c3,c4
r1,0.0,0.1,0.2,0.3
r2,,0.2,0.3,0.4
r3,,,0.4,0.5
r4,,,,0.6
r5,,,,0.7`
	var context = NewContext(3, 0)
	context.WithHeader = RowColumnHeader
	var table = NewTable(csv.NewReader(strings.NewReader(csvData)), context)
	if table.width != 4 || table.height != 5 {
		t.Errorf("table size did not match, wont (4, 4), got (%d, %d)", table.width, table.height)
	}
	var bounds = table.Bounds()
	if bounds.Min.X != 0 || bounds.Min.Y != 0 || bounds.Max.X != 4*3 || bounds.Max.Y != 5*3 {
		t.Errorf("image size did not match, wont (12, 15), got (%d, %d)", bounds.Max.X, bounds.Max.Y)
	}
	if table.ColorModel() != color.RGBAModel {
		t.Errorf("color model did not match, wont %v, got %v", color.RGBAModel, table.ColorModel())
	}
}

func TestHSBtoRGB(t *testing.T) {
	var testcases = []struct {
		givesHSB *HSB
		wontR    uint32
		wontG    uint32
		wontB    uint32
	}{
		{&HSB{0.0, 1.0, 1.0}, 255, 0, 0},
		{&HSB{0.333334, 1.0, 1.0}, 0, 255, 0},
		{&HSB{0.666667, 1.0, 1.0}, 0, 0, 255},
	}
	for _, tc := range testcases {
		var r, g, b, a = tc.givesHSB.RGBA()
		if r != tc.wontR || b != tc.wontB || g != tc.wontG || a != 255 {
			t.Errorf("%v: converted rgba did not match, wont (%d, %d, %d, %d), got (%d, %d, %d, %d)",
				tc.givesHSB, tc.wontR, tc.wontG, tc.wontB, 255, r, g, b, a)
		}
	}
}

func TestCreateHSB(t *testing.T) {
	var testcases = []struct {
		givesValue float64
		wontHSB    *HSB
	}{
		{1.0, &HSB{0.0, 1.0, 1.0}},
		{0.0, &HSB{0.666667, 1.0, 1.0}},
		{0.5, &HSB{0.333334, 1.0, 1.0}},
	}

	for _, tc := range testcases {
		var hsb = createHSB(tc.givesValue)
		if !isNear(hsb, tc.wontHSB) {
			t.Errorf("hsb from %f did not match, wont %v, got %v", tc.givesValue, tc.wontHSB, hsb)
		}
	}
}

func TestHSBString(t *testing.T) {
	var testcases = []struct {
		hsb *HSB
		str string
	}{
		{&HSB{0.0, 1.0, 1.0}, "hue: 0.000000, saturation: 1.000000, brightness: 1.000000"},
		{&HSB{0.666667, 1.0, 1.0}, "hue: 0.666667, saturation: 1.000000, brightness: 1.000000"},
		{&HSB{0.333334, 1.0, 1.0}, "hue: 0.333334, saturation: 1.000000, brightness: 1.000000"},
	}
	for _, tc := range testcases {
		var msg = tc.hsb.String()
		if msg != tc.str {
			t.Errorf("string representation did not match, wont: %s, got: %s", tc.str, msg)
		}
	}
}

func isNear(hsb1, hsb2 *HSB) bool {
	var dH = hsb1.Hue - hsb2.Hue
	var dS = hsb1.Saturation - hsb2.Saturation
	var dB = hsb1.Brightness - hsb2.Brightness
	return math.Abs(dH) < 0.0001 && math.Abs(dS) < 0.0001 && math.Abs(dB) < 0.0001
}
