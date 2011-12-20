package tripit

// This file was generated by a tool. Do not edit.

import (
	"os"
	"json"
)

// A specialization of Vector for Warning objects
type WarningVector []Warning

func (p *WarningVector) UnmarshalJSON(b []byte) os.Error {
	var arr *[]Warning
	arr = (*[]Warning)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Warning, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}
