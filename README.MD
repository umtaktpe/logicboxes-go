# Logicboxes Golang Package

LogicBoxes package provides most of the LogicBoxes API functionality and compatible company apis like resellerclub, whois.com and many more.

## Installation 

```
go get github.com/umtaktpe/logicboxes-go
```

## Usage
If you do not have the .env file, you can set your keys after creating the client.

```go
package main

import (
	"fmt"
	"log"

	"github.com/umtaktpe/logicboxes-go"
)

func main() {
	client := logicboxes.NewClient("auth-userid", "api-key")
	
	resp, err := client.DomainDetails("test.com")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp)
}
```

## Avalaible Methods
### Domain
```go
// get domain nameservers
client.DomainNameservers("domain-name.tld")
// get domains all details.
client.DomainDetails("domain-name.tld")
// get domains status
client.DomainStatus("domain-name.tld")
// get domains order details
client.DomainOrderDetails("domain-name.tld")
// get domains dnssec details
client.DomainDNSSECDetails("domain-name.tld")
// get domains contact id list.
client.DomainIDSContact("domain-name.tld")
// get domains registrant contact details
client.DomainRegistrantContact("domain-name.tld")
// get domains admin contact details
client.DomainAdminContact("domain-name.tld")
// get domains tech contact details
client.DomainTechContact("domain-name.tld")
// get domains billing contact details
client.DomainBillingContact("domain-name.tld")
// check for domain
domainNames := []string{"domain-name1", "domain-name2"}
tlds := []string{"com", "net"}
client.DomainCheck(domainNames, tlds)
// get domain order id
client.DomainOrderID("domain-name.tld")
// set enable theft protection
client.DomainEnableTheftProtection("domain-name.tld")
// set disable theft protection
client.DomainDisableTheftProtection("domain-name.tld")
// modify domain nameservers
ns := []string{"ns1.domain.com", "ns2.domain.com"}
client.DomainModifyNameserver("domain-name.tld", ns)
// delete this domain
client.DomainDelete("domain-name.tld")
// cancel domain transfer process
client.DomainCancelTransfer("domain-name.tld")
// register domain
data := map[string]interface{}{
		"years":              1,
		"ns":                 []string{"ns1.domain.com", "ns2.domain.com"},
		"customer-id":        1,
		"reg-contact-id":     1,
		"admin-contact-id":   1,
		"tech-contact-id":    1,
		"billing-contact-id": 1,
		"invoice-option":     "OptionType",
		"purchase-privacy":   true,
		"protect-privacy":    true,
}
client.DomainRegister("domain-name.tld", data)
// transfer domain
data := map[string]interface{}{
        "auth-code":          1,
		"years":              1,
		"ns":                 []string{"ns1.domain.com", "ns2.domain.com"},
		"customer-id":        1,
		"reg-contact-id":     1,
		"admin-contact-id":   1,
		"tech-contact-id":    1,
		"billing-contact-id": 1,
		"invoice-option":     "OptionType",
		"purchase-privacy":   true,
		"protect-privacy":    true,
}
client.DomainTransfer("domain-name.tld", data)
// set Auth Code
client.DomainAuthCode("domain-name.tld", "auth-code")
// modify Auth Code
client.DomainModifyAuthCode("domain-name.tld", "auth-code")
// renew domain (exp-date => epochtime)
data := map[string]interface{}{
    "years": 1,
    "exp-date": 1111111,
    "invoice-option": "OptionType"
}
client.DomainRenew("domain-name.tld", data)
// deafult customer nameservers
client.DomainCustomerDefaultNameserver(customerID)
// get lock applied list
client.DomainListLockApplied("domain-name.tld")
// set suspend
client.DomainSuspend("domain-name.tld")
// set unsuspend
client.DomainUnsuspend("domain-name.tld")
```
### Customer
```go
// Creating a customer
data := map[string]interface{}{
	"username":       "test",
	"passwd":         "Fxejvcv5cv2zx*",
	"name":           "test",
	"company":        "test",
	"address-line-1": "Example address for register",
	"city":           "test",
	"state":          "test",
	"country":        "TR",
	"zipcode":        "35000",
	"phone-cc":       90,
	"phone":          5555555555,
	"lang-pref":      "en",
}
client.CustomerCreate(data)
// getting a customer by email
client.CustomerDetails("test@test.com")
// Changing customers password
client.CustomerChangePassword(customerID, "123456")
// Change customer
client.CustomerMoveProduct("domain-name.tld", "old-contact", oldCustomerID, newCustomerID)
// customer delete
client.CustomerDetails(customerID)
```
