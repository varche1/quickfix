package linesoftextgrp

//NoLinesOfText is a repeating group in LinesOfTextGrp
type NoLinesOfText struct {
	//Text is a required field for NoLinesOfText.
	Text string `fix:"58"`
	//EncodedTextLen is a non-required field for NoLinesOfText.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoLinesOfText.
	EncodedText *string `fix:"355"`
}

//Component is a fix50sp2 LinesOfTextGrp Component
type Component struct {
	//NoLinesOfText is a required field for LinesOfTextGrp.
	NoLinesOfText []NoLinesOfText `fix:"33"`
}

func New() *Component { return new(Component) }
