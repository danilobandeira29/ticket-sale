package domain

import "fmt"

type Name struct {
	vo StringValueObject
}

func NewName(v string) *Name {
	return &Name{
		vo: StringValueObject{value: v},
	}
}

func (n *Name) String() string {
	return n.vo.Value()
}

func (n *Name) Equal(o *Name) bool {
	return n.vo.Equal(o.vo)
}

func (n *Name) Scan(src any) error {
	switch source := src.(type) {
	case string:
		n.vo.value = source
		return nil
	default:
		return fmt.Errorf("cannot scan Name from %T", src)
	}
}
