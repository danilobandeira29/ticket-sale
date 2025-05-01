package domain

import (
	"errors"
	"testing"
)

func TestNewCPF_WithMask(t *testing.T) {
	value := "360.747.500-84"
	cpf, err := NewCPF(value)
	if err != nil {
		t.Errorf("error creating cpf\nexpect: %s\ngot: %v\n", value, err)
		return
	}
	if cpf.Value() != value {
		t.Errorf("error cpf value\nexpect: %s\ngot: %s\n", value, cpf.Value())
		return
	}
}

func TestNewCPF_WithoutMask(t *testing.T) {
	value := "36074750084"
	cpf, err := NewCPF(value)
	if err != nil {
		t.Errorf("without mask creating cpf\nexpect: 360.747.500-84\ngot: %v\n", err)
		return
	}
	if cpf.Value() != "360.747.500-84" {
		t.Errorf("without mask cpf value\nexpect: 360.747.500-84\ngot: %s\n", cpf.Value())
		return
	}
}

func TestNewCPF_InvalidLenNotEqualToEleven(t *testing.T) {
	value := "360747500844"
	_, err := NewCPF(value)
	if err == nil {
		t.Error("expected error when creating CPF with invalid length, but got nil")
		return
	}
	if !errors.Is(err, ErrCPFLength) {
		t.Errorf("expected: %v\ngot: %v\n", ErrCPFLength, err)
	}

}

func TestNewCPF_InvalidFormat(t *testing.T) {
	value := "123456789.00"
	_, err := NewCPF(value)
	if err == nil {
		t.Error("expected error when creating CPF with invalid length, but got nil")
		return
	}
	if !errors.Is(err, ErrCPFInvalidFormat) {
		t.Errorf("expected: %v\ngot: %v\n", ErrCPFInvalidFormat, err)
	}

}

func TestNewCPF_Algorithm(t *testing.T) {
	value := "000.000.000-00"
	_, err := NewCPF(value)
	if err == nil {
		t.Error("expected error because same digits is not permitted in cpf algorithm")
		return
	}
	if !errors.Is(err, ErrCPFInvalid) {
		t.Errorf("expected error: %v\ngot: %v\n", ErrCPFInvalid, err)
	}
}
