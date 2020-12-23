package logicboxes

import (
	"github.com/umtaktpe/logicboxes-go/models"
)

func (c *Client) CustomerCreate(query map[string]interface{}) (int, error) {
	var response int
	if err := c.PostRequest("customers/v2/signup.json", &response, query); err != nil {
		return 0, err
	}

	return response, nil
}

func (c *Client) CustomerUpdate(customerID int, query map[string]interface{}) (bool, error) {
	query["customer-id"] = customerID

	var response bool
	if err := c.PostRequest("customers/modify.json", &response, query); err != nil {
		return false, err
	}

	return response, nil
}

func (c *Client) CustomerDetails(customerEmail string) (*models.CustomerDetails, error) {
	query := map[string]interface{}{
		"username": customerEmail,
	}

	response := models.CustomerDetails{}
	if err := c.GetRequest("customers/details.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CustomerChangePassword(customerID int, newPasswd string) (bool, error) {
	query := map[string]interface{}{
		"customer-id": customerID,
		"new-passwd":  newPasswd,
	}

	var response bool
	if err := c.PostRequest("customers/v2/change-password.json", &response, query); err != nil {
		return false, err
	}

	return response, nil
}

func (c *Client) CustomerDelete(customerID int) (bool, error) {
	query := map[string]interface{}{
		"customer-id": customerID,
	}

	var response bool
	if err := c.PostRequest("customers/delete.json", &response, query); err != nil {
		return false, err
	}

	return response, nil
}

func (c *Client) CustomerMoveProduct(domainName, contact string, ExistCustomerID, NewCustomerID int) (*models.CustomerMoveProduct, error) {
	query := map[string]interface{}{
		"domain-name":          domainName,
		"existing-customer-id": ExistCustomerID,
		"new-customer-id":      NewCustomerID,
		"default-contact":      contact,
	}

	response := models.CustomerMoveProduct{}
	if err := c.PostRequest("products/move.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CustomerGetID(customerEmail string) (string, error) {
	query := map[string]interface{}{
		"username": customerEmail,
	}

	response := models.CustomerDetails{}
	if err := c.GetRequest("customers/details.json", &response, query); err != nil {
		return "", err
	}

	return response.Customerid, nil
}
