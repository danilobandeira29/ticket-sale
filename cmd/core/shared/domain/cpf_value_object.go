package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type CPF struct {
	value string
}

var (
	ErrCPFLength        = errors.New("invalid length, must be 11")
	ErrCPFInvalidFormat = errors.New("invalid format")
	ErrCPFInvalid       = errors.New("invalid cpf")
	xp                  = regexp.MustCompile(`(\d{3})\.?(\d{3})\.?(\d{3})-?(\d{2})`)
	cpfFirstDigitTable  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

func NewCPF(v string) (*CPF, error) {
	cpfSanitized := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, v)
	if len(cpfSanitized) != 11 {
		return &CPF{}, ErrCPFLength
	}
	if !xp.MatchString(v) {
		return &CPF{}, fmt.Errorf("%w: %s", ErrCPFInvalidFormat, v)
	}
	digits := xp.FindStringSubmatch(v)
	allDigitsEquals := strings.Count(cpfSanitized, string(cpfSanitized[0])) == len(cpfSanitized)
	if allDigitsEquals {
		return &CPF{}, ErrCPFInvalid
	}
	if !isValidCPF(cpfSanitized) {
		return &CPF{}, ErrCPFInvalid
	}
	return &CPF{value: fmt.Sprintf("%s.%s.%s-%s", digits[1], digits[2], digits[3], digits[4])}, nil
}

func (v *CPF) Value() string {
	return v.value
}

func (v *CPF) Equal(o *CPF) bool {
	return v.value == o.value
}

func (v *CPF) Scan(src any) error {
	switch source := src.(type) {
	case string:
		v.value = source
		return nil
	default:
		return fmt.Errorf("cannot scan CPF from %T", source)
	}
}

func isValidCPF(cpf string) bool {
	firstPart := cpf[0:9]
	sum := sumDigit(firstPart, cpfFirstDigitTable)
	r1 := sum % 11
	d1 := 0
	if r1 >= 2 {
		d1 = 11 - r1
	}
	secondPart := firstPart + strconv.Itoa(d1)
	dsum := sumDigit(secondPart, cpfSecondDigitTable)
	r2 := dsum % 11
	d2 := 0
	if r2 >= 2 {
		d2 = 11 - r2
	}
	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)
	return finalPart == cpf
}

func sumDigit(s string, table []int) int {
	if len(s) != len(table) {
		return 0
	}
	sum := 0
	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}
	return sum
}
