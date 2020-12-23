package models

type ContactDetails struct {
	Entityid       string        `json:"entityid"`
	JumpConditions []interface{} `json:"jumpConditions"`
	Telnocc        string        `json:"telnocc"`
	Telno          string        `json:"telno"`
	Type           string        `json:"type"`
	Contactstatus  string        `json:"contactstatus"`
	Classname      string        `json:"classname"`
	Country        string        `json:"country"`
	WhoisValidity  struct {
		Valid       string        `json:"valid"`
		InvalidData []interface{} `json:"invalidData"`
	} `json:"whoisValidity"`
	Emailaddr       string        `json:"emailaddr"`
	Classkey        string        `json:"classkey"`
	Parentkey       string        `json:"parentkey"`
	Address1        string        `json:"address1"`
	City            string        `json:"city"`
	Currentstatus   string        `json:"currentstatus"`
	Contacttype     []interface{} `json:"contacttype"`
	Eaqid           string        `json:"eaqid"`
	Contactid       string        `json:"contactid"`
	Company         string        `json:"company"`
	Name            string        `json:"name"`
	Customerid      string        `json:"customerid"`
	Actioncompleted string        `json:"actioncompleted"`
	Zip             string        `json:"zip"`
	Entitytypeid    string        `json:"entitytypeid"`
	Description     string        `json:"description"`
}

type ContactDetailsCustomerID struct {
	TechContactDetails struct {
		ContactCompany      string `json:"contact.company"`
		EntityCurrentstatus string `json:"entity.currentstatus"`
		ContactTelno        string `json:"contact.telno"`
		WhoisValidity       struct {
			Valid       string        `json:"valid"`
			InvalidData []interface{} `json:"invalidData"`
		} `json:"whoisValidity"`
		EntityEntityid    string `json:"entity.entityid"`
		ContactCreationdt string `json:"contact.creationdt"`
		DesignatedAgent   string `json:"designated-agent"`
		EntityDescription string `json:"entity.description"`
		ContactContactid  string `json:"contact.contactid"`
		ContactZip        string `json:"contact.zip"`
		ContactTimestamp  string `json:"contact.timestamp"`
		ContactCity       string `json:"contact.city"`
		ContactCountry    string `json:"contact.country"`
		ContactAddress1   string `json:"contact.address1"`
		ContactTelnocc    string `json:"contact.telnocc"`
		ContactName       string `json:"contact.name"`
		EntityCustomerid  string `json:"entity.customerid"`
		ContactState      string `json:"contact.state"`
		ContactEmailaddr  string `json:"contact.emailaddr"`
		ContactType       string `json:"contact.type"`
	} `json:"techContactDetails"`
	RegistrantContactDetails struct {
		ContactCompany      string `json:"contact.company"`
		EntityCurrentstatus string `json:"entity.currentstatus"`
		ContactTelno        string `json:"contact.telno"`
		WhoisValidity       struct {
			Valid       string        `json:"valid"`
			InvalidData []interface{} `json:"invalidData"`
		} `json:"whoisValidity"`
		EntityEntityid    string `json:"entity.entityid"`
		ContactCreationdt string `json:"contact.creationdt"`
		DesignatedAgent   string `json:"designated-agent"`
		EntityDescription string `json:"entity.description"`
		ContactContactid  string `json:"contact.contactid"`
		ContactZip        string `json:"contact.zip"`
		ContactTimestamp  string `json:"contact.timestamp"`
		ContactCity       string `json:"contact.city"`
		ContactCountry    string `json:"contact.country"`
		ContactAddress1   string `json:"contact.address1"`
		ContactTelnocc    string `json:"contact.telnocc"`
		ContactName       string `json:"contact.name"`
		EntityCustomerid  string `json:"entity.customerid"`
		ContactState      string `json:"contact.state"`
		ContactEmailaddr  string `json:"contact.emailaddr"`
		ContactType       string `json:"contact.type"`
	} `json:"registrantContactDetails"`
	AdminContactDetails struct {
		ContactCompany      string `json:"contact.company"`
		EntityCurrentstatus string `json:"entity.currentstatus"`
		ContactTelno        string `json:"contact.telno"`
		WhoisValidity       struct {
			Valid       string        `json:"valid"`
			InvalidData []interface{} `json:"invalidData"`
		} `json:"whoisValidity"`
		EntityEntityid    string `json:"entity.entityid"`
		ContactCreationdt string `json:"contact.creationdt"`
		DesignatedAgent   string `json:"designated-agent"`
		EntityDescription string `json:"entity.description"`
		ContactContactid  string `json:"contact.contactid"`
		ContactZip        string `json:"contact.zip"`
		ContactTimestamp  string `json:"contact.timestamp"`
		ContactCity       string `json:"contact.city"`
		ContactCountry    string `json:"contact.country"`
		ContactAddress1   string `json:"contact.address1"`
		ContactTelnocc    string `json:"contact.telnocc"`
		ContactName       string `json:"contact.name"`
		EntityCustomerid  string `json:"entity.customerid"`
		ContactState      string `json:"contact.state"`
		ContactEmailaddr  string `json:"contact.emailaddr"`
		ContactType       string `json:"contact.type"`
	} `json:"adminContactDetails"`
	BillingContactDetails struct {
		ContactCompany      string `json:"contact.company"`
		EntityCurrentstatus string `json:"entity.currentstatus"`
		ContactTelno        string `json:"contact.telno"`
		WhoisValidity       struct {
			Valid       string        `json:"valid"`
			InvalidData []interface{} `json:"invalidData"`
		} `json:"whoisValidity"`
		EntityEntityid    string `json:"entity.entityid"`
		ContactCreationdt string `json:"contact.creationdt"`
		DesignatedAgent   string `json:"designated-agent"`
		EntityDescription string `json:"entity.description"`
		ContactContactid  string `json:"contact.contactid"`
		ContactZip        string `json:"contact.zip"`
		ContactTimestamp  string `json:"contact.timestamp"`
		ContactCity       string `json:"contact.city"`
		ContactCountry    string `json:"contact.country"`
		ContactAddress1   string `json:"contact.address1"`
		ContactTelnocc    string `json:"contact.telnocc"`
		ContactName       string `json:"contact.name"`
		EntityCustomerid  string `json:"entity.customerid"`
		ContactState      string `json:"contact.state"`
		ContactEmailaddr  string `json:"contact.emailaddr"`
		ContactType       string `json:"contact.type"`
	} `json:"billingContactDetails"`
	Registrant string `json:"registrant"`
	Type       string `json:"type"`
	Tech       string `json:"tech"`
	Billing    string `json:"billing"`
	Admin      string `json:"admin"`
}

type ContactModify struct {
	Entityid         string `json:"entityid"`
	Actiontype       string `json:"actiontype"`
	Actiontypedesc   string `json:"actiontypedesc"`
	Eaqid            string `json:"eaqid"`
	Actionstatus     string `json:"actionstatus"`
	Actionstatusdesc string `json:"actionstatusdesc"`
}

type ContactDefault struct {
	Admin      string `json:"admin"`
	Tech       string `json:"tech"`
	Type       string `json:"type"`
	Billing    string `json:"billing"`
	Registrant string `json:"registrant"`
}

type ContactDelete struct {
	Entityid         string `json:"entityid"`
	Actiontype       string `json:"actiontype"`
	Actiontypedesc   string `json:"actiontypedesc"`
	Eaqid            string `json:"eaqid"`
	Actionstatus     string `json:"actionstatus"`
	Actionstatusdesc string `json:"actionstatusdesc"`
}
