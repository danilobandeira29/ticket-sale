package domain

type Name struct {
	vo StringValueObject
}

func NewName(v string) Name {
	return Name{
		vo: StringValueObject{value: v},
	}
}

func (n Name) String() string {
	return n.vo.Value()
}

func (n Name) Equal(o Name) bool {
	return n.vo.Equal(o.vo)
}
