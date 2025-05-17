package entity

import "testing"

func TestOrder_String(t *testing.T) {
	o := &Order{}
	s := o.String()
	t.Log(s)
}
