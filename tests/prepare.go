package tests

import (
	"github.com/umtaktpe/logicboxes-go"
	lib "github.com/umtaktpe/logicboxes-go/library"
	"syreclabs.com/go/faker"
)

func baseClient() *logicboxes.Client {
	client := logicboxes.NewClient(lib.GetConfig("auth-userid"), lib.GetConfig("api-key"))

	return client
}

func customerData() map[string]interface{} {
	data := map[string]interface{}{
		"username":       faker.Internet().Email(),
		"passwd":         "Fxejvcv5cv2zx*",
		"name":           faker.Name().Name(),
		"company":        faker.Company().Name(),
		"address-line-1": "Example address for register",
		"city":           faker.Address().City(),
		"state":          faker.Address().State(),
		"country":        "TR",
		"zipcode":        "35000",
		"phone-cc":       90,
		"phone":          5555555555,
		"lang-pref":      "en",
	}

	return data
}

func contactData() map[string]interface{} {
	data := map[string]interface{}{
		"name":           faker.Name().Name(),
		"company":        faker.Company().Name(),
		"email":          faker.Internet().Email(),
		"address-line-1": "Example address for register",
		"city":           faker.Address().City(),
		"country":        "TR",
		"zipcode":        "35000",
		"phone-cc":       90,
		"phone":          5555555555,
		"type":           "Contact",
	}

	return data
}

func domainData(customerID, contactID int) map[string]interface{} {
	data := map[string]interface{}{
		"years":              1,
		"ns":                 []string{"dns3.parkpage.foundationapi.com", "dns4.parkpage.foundationapi.com"},
		"customer-id":        customerID,
		"reg-contact-id":     contactID,
		"admin-contact-id":   contactID,
		"tech-contact-id":    contactID,
		"billing-contact-id": contactID,
		"invoice-option":     "NoInvoice",
		"purchase-privacy":   true,
		"protect-privacy":    true,
	}

	return data
}
