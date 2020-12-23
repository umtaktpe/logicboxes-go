package tests

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"syreclabs.com/go/faker"
)

func TestDomainRegister(t *testing.T) {
	client := baseClient()

	customerID, err := client.CustomerCreate(customerData())
	if err != nil {
		t.Error(err.Error())
	}

	contactID, err := client.ContactCreate(int(customerID), contactData())
	if err != nil {
		t.Error(err.Error())
	}

	domain := domainData(customerID, contactID)
	domainName := faker.Internet().DomainWord() + ".com"
	if _, err := client.DomainRegister(domainName, domain); err != nil {
		t.Error(err.Error())
	}

	details, err := client.DomainDetails(domainName)
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotEqual(t, details, nil, "Response shouldn't be nil.")
	if details != nil {
		assert.Equal(t, details.Domainname, domainName, "Domain name should be same.")
		assert.Equal(t, details.Customerid, strconv.Itoa(customerID), "Customer id should be same.")
	}
}

func TestDomainCustomerDefaultNameserver(t *testing.T) {
	client := baseClient()

	data := customerData()
	id, err := client.CustomerCreate(data)
	if err != nil {
		t.Error(err.Error())
	}

	ns, err := client.DomainCustomerDefaultNameserver(int(id))
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotEqual(t, len(ns), 0, "Nameservers length shouldn't equal to zero.")
}

func TestDomainEnableTheftProtection(t *testing.T) {
	client := baseClient()

	customerID, err := client.CustomerCreate(customerData())
	if err != nil {
		t.Error(err.Error())
	}

	contactID, err := client.ContactCreate(int(customerID), contactData())
	if err != nil {
		t.Error(err.Error())
	}

	domain := domainData(customerID, contactID)
	domainName := faker.Internet().DomainWord() + ".com"
	if _, err := client.DomainRegister(domainName, domain); err != nil {
		t.Error(err.Error())
	}

	resp, err := client.DomainEnableTheftProtection(domainName)
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotEqual(t, resp, nil, "Response shouldn't be nil.")
	if resp != nil {
		assert.Equal(t, resp.Status, "Success", "Response status should equal to success.")
	}
}

func TestDomainDisableTheftProtection(t *testing.T) {
	client := baseClient()

	customerID, err := client.CustomerCreate(customerData())
	if err != nil {
		t.Error(err.Error())
	}

	contactID, err := client.ContactCreate(int(customerID), contactData())
	if err != nil {
		t.Error(err.Error())
	}

	domain := domainData(customerID, contactID)
	domainName := faker.Internet().DomainWord() + ".com"
	if _, err := client.DomainRegister(domainName, domain); err != nil {
		t.Error(err.Error())
	}

	resp, err := client.DomainDisableTheftProtection(domainName)
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotEqual(t, resp, nil, "Response shouldn't be nil.")
	if resp != nil {
		assert.Equal(t, resp.Status, "Success", "Response status should equal to success.")
	}
}

func TestDomainModifyNameserver(t *testing.T) {
	client := baseClient()

	customerID, err := client.CustomerCreate(customerData())
	if err != nil {
		t.Error(err.Error())
	}

	contactID, err := client.ContactCreate(int(customerID), contactData())
	if err != nil {
		t.Error(err.Error())
	}

	domain := domainData(customerID, contactID)
	domainName := faker.Internet().DomainWord() + ".com"
	if _, err := client.DomainRegister(domainName, domain); err != nil {
		t.Error(err.Error())
	}

	ns := []string{"ns1.ni.net.tr", "ns2.ni.net.tr"}
	resp, err := client.DomainModifyNameserver(domainName, ns)
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotEqual(t, resp, nil, "Response shouldn't be nil.")
	if resp != nil {
		assert.Equal(t, resp.Status, "Success", "Response status should equal to success.")
	}
}

func TestDomainSuspend(t *testing.T) {
	client := baseClient()

	customerID, err := client.CustomerCreate(customerData())
	if err != nil {
		t.Error(err.Error())
	}

	contactID, err := client.ContactCreate(int(customerID), contactData())
	if err != nil {
		t.Error(err.Error())
	}

	domain := domainData(customerID, contactID)
	domainName := faker.Internet().DomainWord() + ".com"
	if _, err := client.DomainRegister(domainName, domain); err != nil {
		t.Error(err.Error())
	}

	resp, err := client.DomainSuspend(domainName, "reason")
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotEqual(t, resp, nil, "Response shouldn't be nil.")
	if resp != nil {
		assert.Equal(t, resp.Status, "Success", "Response status should equal to success.")
	}
}

func TestDomainUnsuspend(t *testing.T) {
	client := baseClient()

	customerID, err := client.CustomerCreate(customerData())
	if err != nil {
		t.Error(err.Error())
	}

	contactID, err := client.ContactCreate(int(customerID), contactData())
	if err != nil {
		t.Error(err.Error())
	}

	domain := domainData(customerID, contactID)
	domainName := faker.Internet().DomainWord() + ".com"
	if _, err := client.DomainRegister(domainName, domain); err != nil {
		t.Error(err.Error())
	}

	resp, err := client.DomainUnsuspend(domainName)
	if err != nil {
		t.Error(err.Error())
	}
	assert.NotEqual(t, resp, nil, "Response shouldn't be nil.")
	if resp != nil {
		assert.Equal(t, resp.Status, "Success", "Response status should equal to success.")
	}
}
