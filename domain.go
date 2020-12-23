package logicboxes

import (
	"strconv"

	"github.com/umtaktpe/logicboxes-go/models"
)

func (c *Client) DomainNameservers(domainName string) (*models.DomainNameservers, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"options":     "NsDetails",
	}

	response := models.DomainNameservers{}
	if err := c.GetRequest("domains/details-by-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainDetails(domainName string) (*models.DomainDetails, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"options":     "All",
	}

	response := models.DomainDetails{}
	if err := c.GetRequest("domains/details-by-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainStatus(domainName string) (*models.DomainStatus, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"options":     "DomainStatus",
	}

	response := models.DomainStatus{}
	if err := c.GetRequest("domains/details-by-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainOrderDetails(domainName string) (*models.DomainOrderDetails, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"options":     "OrderDetails",
	}

	response := models.DomainOrderDetails{}
	if err := c.GetRequest("domains/details-by-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainDNSSECDetails(domainName string) (*models.DomainDNSSECDetails, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"options":     "DNSSECDetails",
	}

	response := models.DomainDNSSECDetails{}
	if err := c.GetRequest("domains/details-by-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainIDSContact(domainName string) (*models.DomainIDSContact, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"options":     "ContactIds",
	}

	response := models.DomainIDSContact{}
	if err := c.GetRequest("domains/details-by-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainTechContact(domainName string) (*models.DomainTechContact, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"options":     "TechContactDetails",
	}

	response := models.DomainTechContact{}
	if err := c.GetRequest("domains/details-by-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainAdminContact(domainName string) (*models.DomainAdminContact, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"options":     "AdminContactDetails",
	}

	response := models.DomainAdminContact{}
	if err := c.GetRequest("domains/details-by-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainBillingContact(domainName string) (*models.DomainBillingContact, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"options":     "BillingContactDetails",
	}

	response := models.DomainBillingContact{}
	if err := c.GetRequest("domains/details-by-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainRegistrantContact(domainName string) (*models.DomainRegistrantContact, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"options":     "RegistrantContactDetails",
	}

	response := models.DomainRegistrantContact{}
	if err := c.GetRequest("domains/details-by-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainCheck(domainName, tlds []string) (map[string]models.DomainCheck, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
		"tlds":        tlds,
	}

	response := map[string]models.DomainCheck{}
	if err := c.GetRequest("domains/available.json", &response, query); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DomainOrderID(domainName string) (int, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
	}

	var response int
	if err := c.GetRequest("domains/orderid.json", &response, query); err != nil {
		return 0, err
	}

	return response, nil
}

func (c *Client) DomainEnableTheftProtection(domainName string) (*models.DomainTheftProtection, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"order-id": orderID,
	}

	response := models.DomainTheftProtection{}
	if err := c.PostRequest("domains/enable-theft-protection.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainDisableTheftProtection(domainName string) (*models.DomainTheftProtection, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"order-id": orderID,
	}

	response := models.DomainTheftProtection{}
	if err := c.PostRequest("domains/disable-theft-protection.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainModifyNameserver(domainName string, ns []string) (*models.DomainModifyNameserver, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"order-id": orderID,
		"ns":       ns,
	}

	response := models.DomainModifyNameserver{}
	if err := c.PostRequest("domains/modify-ns.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainDelete(domainName string) (*models.DomainDelete, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"order-id": orderID,
	}

	response := models.DomainDelete{}
	if err := c.PostRequest("domains/delete.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainCancelTransfer(domainName string) (string, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return "", err
	}

	query := map[string]interface{}{
		"order-id": orderID,
	}

	var response string
	if err := c.PostRequest("domains/cancel-transfer.json", &response, query); err != nil {
		return "", err
	}

	return response, nil
}

func (c *Client) DomainRegister(domainName string, query map[string]interface{}) (*models.DomainRegister, error) {
	query["domain-name"] = domainName

	response := models.DomainRegister{}
	if err := c.PostRequest("domains/register.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainTransfer(domainName string, query map[string]interface{}) (*models.DomainTransfer, error) {
	query["domain-name"] = domainName

	response := models.DomainTransfer{}
	if err := c.PostRequest("domains/transfer.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainAuthCode(domainName, authCode string) (string, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return "", err
	}

	query := map[string]interface{}{
		"order-id":  orderID,
		"auth-code": authCode,
	}

	var response string
	if err := c.PostRequest("domains/transfer/submit-auth-code.json", &response, query); err != nil {
		return "", err
	}

	return response, nil
}

func (c *Client) DomainModifyAuthCode(domainName, authCode string) (*models.DomainModifyAuthCode, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"order-id":  orderID,
		"auth-code": authCode,
	}

	response := models.DomainModifyAuthCode{}
	if err := c.PostRequest("domains/modify-auth-code.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainValidateTransferRequest(domainName string) (bool, error) {
	query := map[string]interface{}{
		"domain-name": domainName,
	}

	var response bool
	if err := c.GetRequest("domains/validate-transfer.json", &response, query); err != nil {
		return false, err
	}

	return response, nil
}

func (c *Client) DomainRenew(domainName string, query map[string]interface{}) (*models.DomainRenew, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query["order-id"] = orderID
	query["domain-name"] = domainName

	response := models.DomainRenew{}
	if err := c.PostRequest("domains/renew.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainCustomerDefaultNameserver(customerID int) ([]string, error) {
	query := map[string]interface{}{
		"customer-id": customerID,
	}

	var response []string
	if err := c.GetRequest("domains/customer-default-ns.json", &response, query); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DomainAddChildNameserver(domainName, cns, ip string) (*models.DomainAddChildNameserver, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"order-id": orderID,
		"cns":      cns,
		"ip":       ip,
	}

	response := models.DomainAddChildNameserver{}
	if err := c.PostRequest("domains/add-cns.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainModifyChildNameserver(domainName, oldCns, newCns string) (*models.DomainModifyChildNameserver, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"order-id": orderID,
		"old-cns":  oldCns,
		"new-cns":  newCns,
	}

	response := models.DomainModifyChildNameserver{}
	if err := c.PostRequest("domains/modify-cns-name.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainListLockApplied(domainName string) (*models.DomainListLockApplied, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"order-id": orderID,
	}

	response := models.DomainListLockApplied{}
	if err := c.GetRequest("domains/locks.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainSuspend(domainName, reason string) (*models.DomainSuspend, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"order-id": orderID,
		"reason":   reason,
	}

	response := models.DomainSuspend{}
	if err := c.PostRequest("orders/suspend.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainUnsuspend(domainName string) (*models.DomainUnsuspend, error) {
	orderID, err := c.DomainOrderID(domainName)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"order-id": orderID,
	}

	response := models.DomainUnsuspend{}
	if err := c.PostRequest("orders/unsuspend.json", &response, query); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DomainSuggestName(keyword, tldOnly string, exactMatch bool) (map[string]models.DomainSuggestName, error) {
	query := map[string]interface{}{
		"keyword":     keyword,
		"exact-match": strconv.FormatBool(exactMatch),
	}

	if tldOnly != "" {
		query["tld-only"] = tldOnly
	}

	response := map[string]models.DomainSuggestName{}
	if err := c.GetRequest("domains/v5/suggest-names.json", &response, query); err != nil {
		return nil, err
	}

	return response, nil
}
