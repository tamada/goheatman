package heatman

/*
Context represents parameters for building heat map.
*/
type Context struct {
	/*
	   SizeOfAPixel shows a pixel size of a cell.
	*/
	SizeOfAPixel int
	/*
	   GapOfAdditionalLine shows to need assistnace lines by some cells.
	   If this value is negative value, no assistance lines are drawn.
	*/
	GapOfAdditionalLine int

	/*
	   Destination represents the destination file name.
	*/
	Destination string

	/*
	   WithHeader shows header type.
	*/
	WithHeader HeaderModel

	/*
	   Converter is value to color mapper.
	*/
	Converter HeatmapConverter
}

/*
HeaderModel shows the header type of given csv file.
*/
type HeaderModel int

const (
	/*
		RowHeader shows the given csv file has row header only.
	*/
	RowHeader HeaderModel = iota
	/*
		ColumnHeader shows the given csv file has column header only.
	*/
	ColumnHeader
	/*
		RowColumnHeader shows the given csv file has row and column header.
	*/
	RowColumnHeader
	/*
		NoHeaders shows the given csv file has no headers.
	*/
	NoHeaders
	/*
		InvalidHeaderModel shows some error on parsing header model.
	*/
	InvalidHeaderModel
)

/*
NewContext create an instance of Context.
*/
func NewContext(sizeOfAPixel, gapOfAdditionalLine int) *Context {
	return &Context{
		SizeOfAPixel:        sizeOfAPixel,
		GapOfAdditionalLine: gapOfAdditionalLine,
	}
}
