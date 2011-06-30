package tripit

// This file was generated by a tool. Do not edit.

import (
	"os"
	"json"
	"container/vector"
)

// A specialization of Vector for Warning objects
type WarningVector struct {
	vector.Vector
}

// AppendVector appends the entire vector x to the end of this vector.
func (p *WarningVector) AppendVector(x *WarningVector) {
	p.Vector.AppendVector(&x.Vector)
}

// At returns the i'th element of the vector.
func (p *WarningVector) At(i int) Warning {
	return p.Vector.At(i).(Warning)
}

// Copy makes a copy of the vector and returns it.
func (p *WarningVector) Copy() WarningVector {
	return WarningVector{p.Vector.Copy()}
}

// Do calls function f for each element of the vector, in order. The behavior of Do is undefined if f changes *p.
func (p *WarningVector) Do(f func(elem Warning)) {
	p.Vector.Do(func(e interface{}) { f(e.(Warning)) })
}

// Insert inserts into the vector an element of value x before the current element at index i.
func (p *WarningVector) Insert(i int, x Warning) {
	p.Vector.Insert(i, x)
}

// InsertVector inserts into the vector the contents of the vector x such that the 0th element of x appears at
// index i after insertion.
func (p *WarningVector) InsertVector(i int, x *WarningVector) {
	p.Vector.InsertVector(i, &x.Vector)
}

// Last returns the element in the vector of highest index.
func (p *WarningVector) Last() Warning {
	return p.Vector.Last().(Warning)
}

// Pop deletes the last element of the vector.
func (p *WarningVector) Pop() Warning {
	return p.Vector.Pop().(Warning)
}

// Push appends x to the end of the vector.
func (p *WarningVector) Push(x Warning) {
	p.Vector.Push(x)
}

// Resize changes the length and capacity of a vector. If the new length is shorter than the current length,
// Resize discards trailing elements. If the new length is longer than the current length, Resize adds the
// respective zero values for the additional elements. The capacity parameter is ignored unless the new length
// or capacity is longer than the current capacity. The resized vector's capacity may be larger than the
// requested capacity.
func (p *WarningVector) Resize(length, capacity int) *WarningVector {
	p.Vector = *p.Vector.Resize(length, capacity)
	return p
}

// Set sets the i'th element of the vector to value x.
func (p *WarningVector) Set(i int, x Warning) {
	p.Vector.Set(i, x)
}

// Slice returns a new sub-vector by slicing the old one to extract slice [i:j]. The elements are copied.
// The original vector is unchanged.
func (p *WarningVector) Slice(i, j int) *WarningVector {
	v := p.Vector.Slice(i, j)
	return &WarningVector{*v}
}

// UnmarshalJSON customizes the JSON unmarshalling by accepting single elements or arrays of elements.
func (p *WarningVector) UnmarshalJSON(b []byte) os.Error {
	var arr []Warning
	err := json.Unmarshal(b, &arr)
	if err != nil {
		arr = make([]Warning, 1)
		err := json.Unmarshal(b, &arr[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				arr = arr[0:0]
			} else {
				return err
			}
		}
	}
	p.Cut(0, p.Len())
	for _, v := range arr {
		p.Push(v)
	}
	return nil
}

// MarshalJSON customizes the JSON output for Vectors.
func (p *WarningVector) MarshalJSON() ([]byte, os.Error) {
	var a []Warning
	if p == nil {
		a = make([]Warning, 0)
	} else {
		a = make([]Warning, p.Len())
		for i := 0; i < p.Len(); i++ {
			a[i] = p.At(i)
		}
	}
	return json.Marshal(a)
}

// Data returns all the elements as a slice.
func (p *WarningVector) Data() []Warning {
	arr := make([]Warning, p.Len())
	var i int
	i = 0
	p.Do(func(v Warning) {
		arr[i] = v
		i++
	})
	return arr
}
