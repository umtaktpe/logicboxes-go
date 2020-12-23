package logicboxes

import (
	"github.com/umtaktpe/logicboxes-go/models"
)

func (c *Client) ResellerDetails(resellerID int) (*models.ResellerDetails, error) {
	query := map[string]interface{}{
		"reseller-id": resellerID,
	}

	response := models.ResellerDetails{}
	if err := c.GetRequest("resellers/details.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ResellerLoginToken(resellerID int, IP string) (string, error) {
	query := map[string]interface{}{
		"reseller-id": resellerID,
		"ip":          IP,
	}

	var response string
	if err := c.GetRequest("resellers/generate-login-token.json", &response, query); err != nil {
		return "", err
	}

	return response, nil
}

func (c *Client) ResellerAddFund(resellerID int, query map[string]interface{}) (int, error) {
	query["reseller-id"] = resellerID

	var response int
	if err := c.PostRequest("billing/add-reseller-fund.json", &response, query); err != nil {
		return 0, err
	}

	return response, nil
}

func (c *Client) ResellerBalance(resellerID int) (*models.ResellerBalance, error) {
	query := map[string]interface{}{
		"reseller-id": resellerID,
	}

	response := models.ResellerBalance{}
	if err := c.GetRequest("billing/reseller-balance.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ResellerGetPrice(resellerID int) (map[string]interface{}, error) {
	query := map[string]interface{}{
		"reseller-id": resellerID,
	}

	response := make(map[string]interface{})
	if err := c.GetRequest("products/reseller-price.json", &response, query); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ResellerGetCustomerPrice() (map[string]interface{}, error) {
	response := make(map[string]interface{})
	if err := c.GetRequest("products/customer-price.json", &response, nil); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ResellerCostPrice() (map[string]interface{}, error) {
	response := make(map[string]interface{})
	if err := c.GetRequest("products/reseller-cost-price.json", &response, nil); err != nil {
		return nil, err
	}

	return response, nil
}
