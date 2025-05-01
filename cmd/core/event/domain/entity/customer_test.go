package entity

import "testing"

func TestCreateCustomer(t *testing.T) {
	c, err := CreateCustomer("354.756.560-02", "Danilo Bandeira")
	if err != nil {
		t.Errorf("expect to not return err\ngot: %v\n", err)
		return
	}
	if c.Name.String() != "Danilo Bandeira" {
		t.Errorf("expected name: Danilo Bandeira\ngot: %s\n", c.Name.String())
		return
	}
	if c.CPF.Value() != "354.756.560-02" {
		t.Errorf("expected cpf: 354.756.560-02\ngot: %s\n", c.CPF.Value())
		return
	}
	if c.ID.String() == "" {
		t.Errorf("expected to generate id\ngot empty id\n")
		return
	}
}

func TestNewCustomer(t *testing.T) {
	c, err := NewCustomer("0196885b-ec78-7377-aac4-e9fd1c17cc71", "354.756.560-02", "Danilo Bandeira")
	if err != nil {
		t.Errorf("expect to not return err\n got: %v\n", err)
		return
	}
	if c.ID.String() == "" {
		t.Errorf("expected to generate id\ngot empty id\n")
		return
	}
	if c.Name.String() != "Danilo Bandeira" {
		t.Errorf("expected name: Danilo Bandeira\ngot: %s\n", c.Name.String())
		return
	}
	if c.CPF.Value() != "354.756.560-02" {
		t.Errorf("expected cpf: 354.756.560-02\ngot: %s\n", c.CPF.Value())
		return
	}
}

func TestCustomerEqual(t *testing.T) {
	customer1, _ := CreateCustomer("354.756.560-02", "Danilo Bandeira")
	customer2, _ := NewCustomer(customer1.ID.String(), "354.756.560-02", "Danilo Bandeira")
	if !customer1.Equal(customer2) {
		t.Errorf("expected customers to be equal:\ncustomer 1: %v\ncustomer 2: %v\n", customer1, customer2)
		return
	}
}
