// Code generated by "stringer -type=InvoiceType"; DO NOT EDIT.

package model

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[增值税专用发票-1]
	_ = x[增值税普通发票-4]
	_ = x[增值税专用电子发票-8]
	_ = x[增值税普通电子发票-10]
}

const (
	_InvoiceType_name_0 = "增值税专用发票"
	_InvoiceType_name_1 = "增值税普通发票"
	_InvoiceType_name_2 = "增值税专用电子发票"
	_InvoiceType_name_3 = "增值税普通电子发票"
)

func (i InvoiceType) String() string {
	switch {
	case i == 1:
		return _InvoiceType_name_0
	case i == 4:
		return _InvoiceType_name_1
	case i == 8:
		return _InvoiceType_name_2
	case i == 10:
		return _InvoiceType_name_3
	default:
		return "InvoiceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
