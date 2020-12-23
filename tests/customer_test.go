package tests

import (
	"fmt"
	"reflect"
	"testing"

	"syreclabs.com/go/faker"

	"github.com/stretchr/testify/assert"
)

func TestCustomerCreate(t *testing.T) {
	client := baseClient()

	data := customerData()
	id, err := client.CustomerCreate(data)

	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, reflect.TypeOf(id).String(), "int", "Return id should be integer.")
	assert.NotEqual(t, id, 0, "Id shouldn't equal to zero.")
}

func TestCustomerUpdate(t *testing.T) {
	client := baseClient()

	data := customerData()
	id, err := client.CustomerCreate(data)
	if err != nil {
		t.Error(err.Error())
	}

	resp, err := client.CustomerUpdate(int(id), data)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, resp, true, "Response should be true.")
}

func TestCustomerDetails(t *testing.T) {
	client := baseClient()

	data := customerData()
	id, err := client.CustomerCreate(data)
	if err != nil {
		t.Error(err.Error())
	}

	resp, err := client.CustomerDetails(data["username"].(string))
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotEqual(t, resp, nil, "Response shouldn't be nil.")
	if resp != nil {
		assert.Equal(t, resp.Customerid, fmt.Sprintf("%d", id), "Customer id should be same.")
		assert.Equal(t, resp.Username, data["username"].(string), "Username should be same.")
	}
}

func TestCustomerChangePassword(t *testing.T) {
	client := baseClient()

	data := customerData()
	id, err := client.CustomerCreate(data)
	if err != nil {
		t.Error(err.Error())
	}

	resp, err := client.CustomerChangePassword(int(id), "Fxvavfv41v2zx*")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, resp, true, "Response should be true.")
}

func TestCustomerDelete(t *testing.T) {
	client := baseClient()

	data := customerData()
	id, err := client.CustomerCreate(data)
	if err != nil {
		t.Error(err.Error())
	}

	resp, err := client.CustomerDelete(int(id))
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, resp, true, "Response should be true.")
}

func TestCustomerMoveProduct(t *testing.T) {
	client := baseClient()

	customer := customerData()
	customerID, err := client.CustomerCreate(customer)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, reflect.TypeOf(customerID).String(), "int", "Return id should be integer.")
	assert.NotEqual(t, customerID, 0, "Id shouldn't equal to zero.")

	contact := contactData()
	contactID, err := client.ContactCreate(int(customerID), contact)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, reflect.TypeOf(contactID).String(), "int", "Return id should be integer.")
	assert.NotEqual(t, contactID, 0, "Id shouldn't equal to zero.")

	domain := domainData(customerID, contactID)
	domainName := faker.Internet().DomainWord() + ".com"

	if _, err := client.DomainRegister(domainName, domain); err != nil {
		t.Error(err.Error())
	}

	customer2 := customerData()
	customerID2, err := client.CustomerCreate(customer2)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, reflect.TypeOf(customerID2).String(), "int", "Return id should be integer.")
	assert.NotEqual(t, customerID2, 0, "Id shouldn't equal to zero.")

	resp, err := client.CustomerMoveProduct(domainName, "oldcontact", customerID, customerID2)
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotEqual(t, resp, nil, "Response shouldn't be nil.")
	if resp != nil {
		assert.Equal(t, resp.Status, "Success", "Status should be success.")
	}
}
