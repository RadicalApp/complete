package complete

import "reflect"

// Result is a generic struct used with completion handlers.
type Result struct {
	DataType string
	Data     interface{}
}

// NewResult returns a new result struct.
func NewResult(data interface{}) Result {
	return Result{Data: data, DataType: reflect.TypeOf(data).String()}
}

// Iteratable is used to iterate through a list of Results.
type Iteratable interface {
	Next() bool
	Value() Result
}

// ResultArray implements iteratable to allow mobile-compatible array passing.
type ResultArray struct {
	isLocked bool
	index    int
	data     []Result
}

// Next returns true is there is additional results in the array.
func (d *ResultArray) Next() bool {
	d.index++
	if d.index >= len(d.data) {
		return false
	}
	return true
}

// Value gets the current value of the iterator
func (d *ResultArray) Value() Result {
	return d.data[d.index]
}

// NewResultIterator returns a new iterator for Result arrays.
func NewResultIterator(data []Result) *ResultArray {
	return &ResultArray{data: data, index: -1}
}

// GetResults is a Go-compatible way to get an array/slice of Result objects.
func GetResults(results ResultArray) []Result {
	return results.data
}
