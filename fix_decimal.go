package quickfix

import "git.cryptology.com/lib/go/fixed"

//FIXDecimal is a FIX Float Value that implements an arbitrary precision fixed-point decimal.  Implements FieldValue
type FIXDecimal struct {
	fixed.Fixed

	//Scale is the number of digits after the decimal point when Writing the field value as a FIX value
	Scale int32
}

func (d FIXDecimal) Write() []byte {
	return []byte(d.Fixed.FormatToPrecision(int(d.Scale)))
}

func (d *FIXDecimal) Read(bytes []byte) (err error) {
	d.Fixed, err = fixed.NewFromString(string(bytes))
	return
}
