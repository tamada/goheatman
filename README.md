[![Build Status](https://travis-ci.com/tamada/goheatman.svg?branch=master)](https://travis-ci.com/tamada/goheatman)
[![Coverage Status](https://coveralls.io/repos/github/tamada/goheatman/badge.svg?branch=master)](https://coveralls.io/github/tamada/goheatman?branch=master)

# goheatman

Create heat map image from given csv file.
Go language implementation of [tamada/heatman](https://github.com/tamada/heatman).

## Usage

```sh
$ heatman -h both testdata/sample.csv
# generate heatman.png.
$ heatman -H # print the following help message.
heatman [OPTIONS] [CSVFILE]
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
$ heatman -p 15 -a 4 -h both -d heatman2.png testdata/sample.csv
# generate heatmap2.png.
$ heatman -p 15 -h both --color gray -d heatman3.png testdata/sample.csv
# generate heatmap3.png in gray scale.
$ heatman --scaler -d heatman_scaler.png
# generate heatmap_scaler.png which shows scaler of colors.
```

### The Results

Each image is generated by above command.  The sources of below images are same.

#### heatman.png

![heatman.png](https://github.com/tamada/goheatman/blob/master/testdata/heatman.png)

It was too small.
The one cell in csv is shown in one pixel square.

#### heatman2.png

The one cell in csv is shown in 15-pixel square, and the auxiliary lines are drawn by 4 cells.

![heatman2.png](https://github.com/tamada/goheatman/blob/master/testdata/heatman2.png)

#### heatman_gray.png

The image is gray-scaled heatmap.

![heatman3.png](https://github.com/tamada/goheatman/blob/master/testdata/heatman3.png)

#### heatman_scaler.png

This is generated scaler image.
The most right side shows 1.0, and the most left side represents 0.0.

![heatman_scaler.png](https://github.com/tamada/goheatman/blob/master/testdata/heatman_scaler.png)

## Install

### Go lang

```sh
$ go get github.com/tamada/goheatman
```

### Homebrew

```sh
$ brew tap tamada/brew
$ brew install heatman
```

