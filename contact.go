package logicboxes

import (
	"github.com/umtaktpe/logicboxes-go/models"
)

func (c *Client) ContactCreate(customerID int, query map[string]interface{}) (int, error) {
	query["customer-id"] = customerID

	var response int
	if err := c.PostRequest("contacts/add.json", &response, query); err != nil {
		return 0, nil
	}

	return response, nil
}

func (c *Client) ContactDetails(contactID int) (*models.ContactDetails, error) {
	query := map[string]interface{}{
		"contact-id": contactID,
	}

	response := models.ContactDetails{}
	if err := c.GetRequest("contacts/details.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ContactDetailsCustomerID(customerID int, contactType string) (map[string]models.ContactDetailsCustomerID, error) {
	query := map[string]interface{}{
		"customer-id": customerID,
	}

	if len(contactType) != 0 {
		query["type"] = contactType
	} else {
		query["type"] = "Contact"
	}

	response := map[string]models.ContactDetailsCustomerID{}
	if err := c.PostRequest("contacts/default.json", &response, query); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ContactModify(contactID int, query map[string]interface{}) (map[string]models.ContactModify, error) {
	query["contact-id"] = contactID

	response := map[string]models.ContactModify{}
	if err := c.PostRequest("contacts/modify.json", &response, query); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ContactDefaultSet(customerID int, query map[string]interface{}) (*models.ContactDefault, error) {
	query["customer-id"] = customerID

	response := models.ContactDefault{}
	if err := c.PostRequest("contacts/modDefault.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ContactDelete(contactID int) (*models.ContactDelete, error) {
	query := map[string]interface{}{
		"contact-id": contactID,
	}

	response := models.ContactDelete{}
	if err := c.PostRequest("contacts/delete.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}
