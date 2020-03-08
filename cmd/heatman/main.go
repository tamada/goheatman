package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/png"
	"os"

	flag "github.com/spf13/pflag"
	heatman "github.com/tamada/goheatman"
	"github.com/tamada/goheatman/internal"
)

const heatmanName = "heatman"

func printHelp(programName string) {
	fmt.Printf(`%s [OPTIONS] [CSVFILE]
OPTIONS
    -a, --additional-line-gap <GAP> specifies the gap of assistant lines per cells.
                                    if GAP is less equals than 0, no assistant lines are drawn.
    -c, --color <TYPE>              specifies heatmap color type (color or gray), default: color.
    -d, --dest <DEST>               specifies the destination file.
    -h, --headers <HEADER>          specifies header model (both, row, column, or no), default: no.
    -p, --pixel <SIZE>              specifies the pixel size per cell.
    -s, --scaler                    generates scaler of heatmap.  If this option was specified,
                                    additional-line-gap, headers, pixel, and CSVFILE are ignored.
    -H, --help                      print this message.
ARGUMENTS
    CSVFILE                         input csv files with no headers.
                                    if no csv files are specified, heatman read csv from stdin.
                                    The value of each cell must be 0.0 to 1.0.
`, programName)
}

type options struct {
	context     *heatman.Context
	heatmapType string
	headerType  string
	helpFlag    bool
	scalerFlag  bool
}

func buildFlagSet() (*flag.FlagSet, *options) {
	var options = options{context: &heatman.Context{}, headerType: "no", helpFlag: false}
	var flags = flag.NewFlagSet(heatmanName, flag.ContinueOnError)
	flags.Usage = func() { printHelp(heatmanName) }
	flags.StringVarP(&options.context.Destination, "dest", "d", "heatman.png", "specifies the destination file")
	flags.IntVarP(&options.context.SizeOfAPixel, "pixel", "p", 1, "pixel size per cell")
	flags.IntVarP(&options.context.GapOfAdditionalLine, "additional-line-gap", "a", 0, "if greater than 0, draw assistant lines per specified number of cells")
	flags.StringVarP(&options.headerType, "headers", "h", "no", "header model (both, row, column, or no), default: no")
	flags.StringVarP(&options.heatmapType, "color", "c", "color", "specifies heatmap color type (default or gray)")
	flags.BoolVarP(&options.scalerFlag, "scaler", "s", false, "generates scaler")
	flags.BoolVarP(&options.helpFlag, "help", "H", false, "print this message.")
	return flags, &options
}

func findInput(args []string) (*os.File, error) {
	if len(args) == 1 {
		return os.Stdin, nil
	}
	return os.Open(args[1])
}

func printError(err error) int {
	if err == nil {
		fmt.Println(err.Error())
	}
	return 2
}

func writeImage(image image.Image, context *heatman.Context) error {
	var to, err = os.OpenFile(context.Destination, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer to.Close()
	return png.Encode(to, image)
}

func perform(args []string, context *heatman.Context, scalerFlag bool) int {
	if len(args) > 2 {
		printHelp(heatmanName)
		return 1
	}
	var from, err = findInput(args)
	if err != nil {
		return printError(err)
	}
	defer from.Close()
	var image image.Image
	if scalerFlag {
		image = heatman.ScalerImage(context)
	} else {
		image = heatman.NewTable(csv.NewReader(from), context)

	}
	var err2 = writeImage(image, context)
	if err2 != nil {
		return printError(err2)
	}
	return 0
}

func parseOptions(flags *flag.FlagSet, opts *options, args []string) error {
	var err = flags.Parse(args)
	if err != nil {
		return err
	}
	opts.context.WithHeader, err = internal.ParseHeaderModel(opts.headerType)
	if err != nil {
		return err
	}
	opts.context.Converter, err = internal.ParseColorType(opts.heatmapType)
	if err != nil {
		return err
	}
	if opts.context.GapOfAdditionalLine < 0 {
		opts.context.GapOfAdditionalLine = 0
	}
	return nil
}

func goMain() int {
	var flags, opts = buildFlagSet()
	var err = parseOptions(flags, opts, os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
	if opts.helpFlag {
		printHelp(heatmanName)
		return 0
	}
	return perform(flags.Args(), opts.context, opts.scalerFlag)
}

func main() {
	var statusCode = goMain()
	os.Exit(statusCode)
}
