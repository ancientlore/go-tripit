package tripit

// This file was generated by a tool. Do not edit.

import (
	"encoding/json"
)

// InviteeVector is a specialization of Vector for Invitee objects.
type InviteeVector []Invitee

// UnmarshalJSON builds the vector from the JSON in b.
func (p *InviteeVector) UnmarshalJSON(b []byte) error {
	var arr *[]Invitee
	arr = (*[]Invitee)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Invitee, 1)
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
