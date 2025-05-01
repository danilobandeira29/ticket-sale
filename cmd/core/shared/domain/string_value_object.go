package domain

type StringValueObject struct {
	value string
}

func (v StringValueObject) Value() string {
	return v.value
}

func (v StringValueObject) Equal(o StringValueObject) bool {
	return v.value == o.value
}
