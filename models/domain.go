package models

type DomainNameservers struct {
	Gdpr struct {
		Enabled  string `json:"enabled"`
		Eligible string `json:"eligible"`
	} `json:"gdpr"`
	AutoRenewAttemptDuration   string `json:"autoRenewAttemptDuration"`
	IsImmediateReseller        string `json:"isImmediateReseller"`
	Privacyprotectedallowed    string `json:"privacyprotectedallowed"`
	NoOfNameServers            string `json:"noOfNameServers"`
	Isprivacyprotected         string `json:"isprivacyprotected"`
	Classkey                   string `json:"classkey"`
	IsOrderSuspendedUponExpiry string `json:"isOrderSuspendedUponExpiry"`
	Classname                  string `json:"classname"`
	Ns2                        string `json:"ns2"`
	Ns1                        string `json:"ns1"`
	AutoRenewTermType          string `json:"autoRenewTermType"`
}

type DomainDetails struct {
	Entityid          string `json:"entityid"`
	Domsecret         string `json:"domsecret"`
	NoOfNameServers   string `json:"noOfNameServers"`
	Multilingualflag  string `json:"multilingualflag"`
	Registrantcontact struct {
		Company       string        `json:"company"`
		Address1      string        `json:"address1"`
		Telno         string        `json:"telno"`
		Telnocc       string        `json:"telnocc"`
		Contactid     string        `json:"contactid"`
		Type          string        `json:"type"`
		Contacttype   []interface{} `json:"contacttype"`
		Customerid    string        `json:"customerid"`
		Country       string        `json:"country"`
		Parentkey     string        `json:"parentkey"`
		Contactstatus string        `json:"contactstatus"`
		State         string        `json:"state"`
		Emailaddr     string        `json:"emailaddr"`
		City          string        `json:"city"`
		Name          string        `json:"name"`
		Zip           string        `json:"zip"`
	} `json:"registrantcontact"`
	Classname string `json:"classname"`
	Cns       struct {
	} `json:"cns"`
	Entitytypeid           string        `json:"entitytypeid"`
	Paused                 string        `json:"paused"`
	OrderSuspendedByParent string        `json:"orderSuspendedByParent"`
	Description            string        `json:"description"`
	Customercost           string        `json:"customercost"`
	Registrantcontactid    string        `json:"registrantcontactid"`
	JumpConditions         []interface{} `json:"jumpConditions"`
	Actioncompleted        string        `json:"actioncompleted"`
	Creationtime           string        `json:"creationtime"`
	Customerid             string        `json:"customerid"`
	Orderstatus            []string      `json:"orderstatus"`
	Moneybackperiod        string        `json:"moneybackperiod"`
	Classkey               string        `json:"classkey"`
	Resellercost           string        `json:"resellercost"`
	Recurring              string        `json:"recurring"`
	Billingcontact         struct {
		Company       string        `json:"company"`
		Address1      string        `json:"address1"`
		Telno         string        `json:"telno"`
		Telnocc       string        `json:"telnocc"`
		Contactid     string        `json:"contactid"`
		Type          string        `json:"type"`
		Contacttype   []interface{} `json:"contacttype"`
		Customerid    string        `json:"customerid"`
		Country       string        `json:"country"`
		Parentkey     string        `json:"parentkey"`
		Contactstatus string        `json:"contactstatus"`
		State         string        `json:"state"`
		Emailaddr     string        `json:"emailaddr"`
		City          string        `json:"city"`
		Name          string        `json:"name"`
		Zip           string        `json:"zip"`
	} `json:"billingcontact"`
	Admincontactid    string        `json:"admincontactid"`
	AutoRenewTermType string        `json:"autoRenewTermType"`
	Addons            []interface{} `json:"addons"`
	Techcontact       struct {
		Company       string        `json:"company"`
		Address1      string        `json:"address1"`
		Telno         string        `json:"telno"`
		Telnocc       string        `json:"telnocc"`
		Contactid     string        `json:"contactid"`
		Type          string        `json:"type"`
		Contacttype   []interface{} `json:"contacttype"`
		Customerid    string        `json:"customerid"`
		Country       string        `json:"country"`
		Parentkey     string        `json:"parentkey"`
		Contactstatus string        `json:"contactstatus"`
		State         string        `json:"state"`
		Emailaddr     string        `json:"emailaddr"`
		City          string        `json:"city"`
		Name          string        `json:"name"`
		Zip           string        `json:"zip"`
	} `json:"techcontact"`
	RaaVerificationStartTime   string        `json:"raaVerificationStartTime"`
	RaaVerificationStatus      string        `json:"raaVerificationStatus"`
	Dnssec                     []interface{} `json:"dnssec"`
	Isprivacyprotected         string        `json:"isprivacyprotected"`
	Allowdeletion              string        `json:"allowdeletion"`
	Techcontactid              string        `json:"techcontactid"`
	IsOrderSuspendedUponExpiry string        `json:"isOrderSuspendedUponExpiry"`
	Bulkwhoisoptout            string        `json:"bulkwhoisoptout"`
	Endtime                    string        `json:"endtime"`
	Eaqid                      string        `json:"eaqid"`
	Productkey                 string        `json:"productkey"`
	Domainname                 string        `json:"domainname"`
	TncRequired                string        `json:"tnc_required"`
	Ns2                        string        `json:"ns2"`
	Ns1                        string        `json:"ns1"`
	AutoRenewAttemptDuration   string        `json:"autoRenewAttemptDuration"`
	Admincontact               struct {
		Company       string        `json:"company"`
		Address1      string        `json:"address1"`
		Telno         string        `json:"telno"`
		Telnocc       string        `json:"telnocc"`
		Contactid     string        `json:"contactid"`
		Type          string        `json:"type"`
		Contacttype   []interface{} `json:"contacttype"`
		Customerid    string        `json:"customerid"`
		Country       string        `json:"country"`
		Parentkey     string        `json:"parentkey"`
		Contactstatus string        `json:"contactstatus"`
		State         string        `json:"state"`
		Emailaddr     string        `json:"emailaddr"`
		City          string        `json:"city"`
		Name          string        `json:"name"`
		Zip           string        `json:"zip"`
	} `json:"admincontact"`
	IsImmediateReseller     string   `json:"isImmediateReseller"`
	Billingcontactid        string   `json:"billingcontactid"`
	Domainstatus            []string `json:"domainstatus"`
	Parentkey               string   `json:"parentkey"`
	Privacyprotectedallowed string   `json:"privacyprotectedallowed"`
	Orderid                 string   `json:"orderid"`
	Serviceproviderid       string   `json:"serviceproviderid"`
	Gdpr                    struct {
		Enabled  string `json:"enabled"`
		Eligible string `json:"eligible"`
	} `json:"gdpr"`
	Productcategory string `json:"productcategory"`
	Currentstatus   string `json:"currentstatus"`
}

type DomainStatus struct {
	Gdpr struct {
		Enabled  string `json:"enabled"`
		Eligible string `json:"eligible"`
	} `json:"gdpr"`
	AutoRenewAttemptDuration   string   `json:"autoRenewAttemptDuration"`
	IsImmediateReseller        string   `json:"isImmediateReseller"`
	Privacyprotectedallowed    string   `json:"privacyprotectedallowed"`
	Isprivacyprotected         string   `json:"isprivacyprotected"`
	Classkey                   string   `json:"classkey"`
	RaaVerificationStatus      string   `json:"raaVerificationStatus"`
	IsOrderSuspendedUponExpiry string   `json:"isOrderSuspendedUponExpiry"`
	Classname                  string   `json:"classname"`
	Domainstatus               []string `json:"domainstatus"`
	RaaVerificationStartTime   string   `json:"raaVerificationStartTime"`
	AutoRenewTermType          string   `json:"autoRenewTermType"`
}

type DomainOrderDetails struct {
	Orderid                  string   `json:"orderid"`
	TncRequired              string   `json:"tnc_required"`
	Parentkey                string   `json:"parentkey"`
	Creationtime             string   `json:"creationtime"`
	Orderstatus              []string `json:"orderstatus"`
	Recurring                string   `json:"recurring"`
	AutoRenewTermType        string   `json:"autoRenewTermType"`
	Serviceproviderid        string   `json:"serviceproviderid"`
	Classname                string   `json:"classname"`
	IsImmediateReseller      string   `json:"isImmediateReseller"`
	Productkey               string   `json:"productkey"`
	Customerid               string   `json:"customerid"`
	Classkey                 string   `json:"classkey"`
	OrderSuspendedByParent   string   `json:"orderSuspendedByParent"`
	AutoRenewAttemptDuration string   `json:"autoRenewAttemptDuration"`
	Endtime                  string   `json:"endtime"`
	Multilingualflag         string   `json:"multilingualflag"`
	Privacyprotectedallowed  string   `json:"privacyprotectedallowed"`
	Domsecret                string   `json:"domsecret"`
	Gdpr                     struct {
		Enabled  string `json:"enabled"`
		Eligible string `json:"eligible"`
	} `json:"gdpr"`
	IsOrderSuspendedUponExpiry string `json:"isOrderSuspendedUponExpiry"`
	Domainname                 string `json:"domainname"`
	Productcategory            string `json:"productcategory"`
	Allowdeletion              string `json:"allowdeletion"`
	Isprivacyprotected         string `json:"isprivacyprotected"`
	Bulkwhoisoptout            string `json:"bulkwhoisoptout"`
	Moneybackperiod            string `json:"moneybackperiod"`
}

type DomainDNSSECDetails struct {
	Gdpr struct {
		Enabled  string `json:"enabled"`
		Eligible string `json:"eligible"`
	} `json:"gdpr"`
	IsImmediateReseller        string        `json:"isImmediateReseller"`
	Dnssec                     []interface{} `json:"dnssec"`
	Classkey                   string        `json:"classkey"`
	IsOrderSuspendedUponExpiry string        `json:"isOrderSuspendedUponExpiry"`
	AutoRenewTermType          string        `json:"autoRenewTermType"`
	Classname                  string        `json:"classname"`
	AutoRenewAttemptDuration   string        `json:"autoRenewAttemptDuration"`
	Privacyprotectedallowed    string        `json:"privacyprotectedallowed"`
	Isprivacyprotected         string        `json:"isprivacyprotected"`
}

type DomainAdminContact struct {
	Classname             string `json:"classname"`
	PrivacyBillingcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-billingcontact"`
	Classkey                 string `json:"classkey"`
	AutoRenewAttemptDuration string `json:"autoRenewAttemptDuration"`
	AutoRenewTermType        string `json:"autoRenewTermType"`
	PrivacyTechcontact       struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-techcontact"`
	Admincontact struct {
		Company       string        `json:"company"`
		Address1      string        `json:"address1"`
		Telno         string        `json:"telno"`
		Telnocc       string        `json:"telnocc"`
		Contactid     string        `json:"contactid"`
		Type          string        `json:"type"`
		Contacttype   []interface{} `json:"contacttype"`
		Customerid    string        `json:"customerid"`
		Country       string        `json:"country"`
		Parentkey     string        `json:"parentkey"`
		Contactstatus string        `json:"contactstatus"`
		State         string        `json:"state"`
		Emailaddr     string        `json:"emailaddr"`
		City          string        `json:"city"`
		Name          string        `json:"name"`
		Zip           string        `json:"zip"`
	} `json:"admincontact"`
	IsImmediateReseller     string `json:"isImmediateReseller"`
	Privacyprotectedallowed string `json:"privacyprotectedallowed"`
	Gdpr                    struct {
		Enabled  string `json:"enabled"`
		Eligible string `json:"eligible"`
	} `json:"gdpr"`
	Privacyprotectendtime    string `json:"privacyprotectendtime"`
	PrivacyRegistrantcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-registrantcontact"`
	Isprivacyprotected         string `json:"isprivacyprotected"`
	IsOrderSuspendedUponExpiry string `json:"isOrderSuspendedUponExpiry"`
	PrivacyAdmincontact        struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-admincontact"`
}

type DomainIDSContact struct {
	Gdpr struct {
		Enabled  string `json:"enabled"`
		Eligible string `json:"eligible"`
	} `json:"gdpr"`
	IsImmediateReseller string `json:"isImmediateReseller"`
	PrivacyAdmincontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-admincontact"`
	PrivacyRegistrantcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-registrantcontact"`
	Registrantcontactid   string `json:"registrantcontactid"`
	Isprivacyprotected    string `json:"isprivacyprotected"`
	Admincontactid        string `json:"admincontactid"`
	PrivacyBillingcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-billingcontact"`
	Classkey           string `json:"classkey"`
	PrivacyTechcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-techcontact"`
	Techcontactid              string `json:"techcontactid"`
	AutoRenewTermType          string `json:"autoRenewTermType"`
	IsOrderSuspendedUponExpiry string `json:"isOrderSuspendedUponExpiry"`
	Classname                  string `json:"classname"`
	AutoRenewAttemptDuration   string `json:"autoRenewAttemptDuration"`
	Privacyprotectedallowed    string `json:"privacyprotectedallowed"`
	Privacyprotectendtime      string `json:"privacyprotectendtime"`
	Billingcontactid           string `json:"billingcontactid"`
}

type DomainTechContact struct {
	Classname             string `json:"classname"`
	PrivacyBillingcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-billingcontact"`
	Classkey                 string `json:"classkey"`
	AutoRenewAttemptDuration string `json:"autoRenewAttemptDuration"`
	AutoRenewTermType        string `json:"autoRenewTermType"`
	PrivacyTechcontact       struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-techcontact"`
	IsImmediateReseller string `json:"isImmediateReseller"`
	Techcontact         struct {
		Company       string        `json:"company"`
		Address1      string        `json:"address1"`
		Telno         string        `json:"telno"`
		Telnocc       string        `json:"telnocc"`
		Contactid     string        `json:"contactid"`
		Type          string        `json:"type"`
		Contacttype   []interface{} `json:"contacttype"`
		Customerid    string        `json:"customerid"`
		Country       string        `json:"country"`
		Parentkey     string        `json:"parentkey"`
		Contactstatus string        `json:"contactstatus"`
		State         string        `json:"state"`
		Emailaddr     string        `json:"emailaddr"`
		City          string        `json:"city"`
		Name          string        `json:"name"`
		Zip           string        `json:"zip"`
	} `json:"techcontact"`
	Privacyprotectedallowed string `json:"privacyprotectedallowed"`
	Gdpr                    struct {
		Enabled  string `json:"enabled"`
		Eligible string `json:"eligible"`
	} `json:"gdpr"`
	Privacyprotectendtime    string `json:"privacyprotectendtime"`
	PrivacyRegistrantcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-registrantcontact"`
	Isprivacyprotected         string `json:"isprivacyprotected"`
	IsOrderSuspendedUponExpiry string `json:"isOrderSuspendedUponExpiry"`
	PrivacyAdmincontact        struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-admincontact"`
}

type DomainBillingContact struct {
	Classname             string `json:"classname"`
	PrivacyBillingcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-billingcontact"`
	Classkey                 string `json:"classkey"`
	AutoRenewAttemptDuration string `json:"autoRenewAttemptDuration"`
	Billingcontact           struct {
		Company       string        `json:"company"`
		Address1      string        `json:"address1"`
		Telno         string        `json:"telno"`
		Telnocc       string        `json:"telnocc"`
		Contactid     string        `json:"contactid"`
		Type          string        `json:"type"`
		Contacttype   []interface{} `json:"contacttype"`
		Customerid    string        `json:"customerid"`
		Country       string        `json:"country"`
		Parentkey     string        `json:"parentkey"`
		Contactstatus string        `json:"contactstatus"`
		State         string        `json:"state"`
		Emailaddr     string        `json:"emailaddr"`
		City          string        `json:"city"`
		Name          string        `json:"name"`
		Zip           string        `json:"zip"`
	} `json:"billingcontact"`
	AutoRenewTermType  string `json:"autoRenewTermType"`
	PrivacyTechcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-techcontact"`
	IsImmediateReseller     string `json:"isImmediateReseller"`
	Privacyprotectedallowed string `json:"privacyprotectedallowed"`
	Gdpr                    struct {
		Enabled  string `json:"enabled"`
		Eligible string `json:"eligible"`
	} `json:"gdpr"`
	Privacyprotectendtime    string `json:"privacyprotectendtime"`
	PrivacyRegistrantcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-registrantcontact"`
	Isprivacyprotected         string `json:"isprivacyprotected"`
	IsOrderSuspendedUponExpiry string `json:"isOrderSuspendedUponExpiry"`
	PrivacyAdmincontact        struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-admincontact"`
}

type DomainRegistrantContact struct {
	Registrantcontact struct {
		Company       string        `json:"company"`
		Address1      string        `json:"address1"`
		Telno         string        `json:"telno"`
		Telnocc       string        `json:"telnocc"`
		Contactid     string        `json:"contactid"`
		Type          string        `json:"type"`
		Contacttype   []interface{} `json:"contacttype"`
		Customerid    string        `json:"customerid"`
		Country       string        `json:"country"`
		Parentkey     string        `json:"parentkey"`
		Contactstatus string        `json:"contactstatus"`
		State         string        `json:"state"`
		Emailaddr     string        `json:"emailaddr"`
		City          string        `json:"city"`
		Name          string        `json:"name"`
		Zip           string        `json:"zip"`
	} `json:"registrantcontact"`
	Classname             string `json:"classname"`
	PrivacyBillingcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-billingcontact"`
	Classkey                 string `json:"classkey"`
	AutoRenewAttemptDuration string `json:"autoRenewAttemptDuration"`
	AutoRenewTermType        string `json:"autoRenewTermType"`
	PrivacyTechcontact       struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-techcontact"`
	IsImmediateReseller     string `json:"isImmediateReseller"`
	Privacyprotectedallowed string `json:"privacyprotectedallowed"`
	Gdpr                    struct {
		Enabled  string `json:"enabled"`
		Eligible string `json:"eligible"`
	} `json:"gdpr"`
	Privacyprotectendtime    string `json:"privacyprotectendtime"`
	PrivacyRegistrantcontact struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-registrantcontact"`
	Isprivacyprotected         string `json:"isprivacyprotected"`
	IsOrderSuspendedUponExpiry string `json:"isOrderSuspendedUponExpiry"`
	PrivacyAdmincontact        struct {
		Telno                 string `json:"telno"`
		Telnocc               string `json:"telnocc"`
		Country               string `json:"country"`
		Company               string `json:"company"`
		Emailaddr             string `json:"emailaddr"`
		Name                  string `json:"name"`
		Zip                   string `json:"zip"`
		Privacyprotectionroid string `json:"privacyprotectionroid"`
		Address1              string `json:"address1"`
		State                 string `json:"state"`
		City                  string `json:"city"`
	} `json:"privacy-admincontact"`
}

type DomainCheck struct {
	Classkey string `json:"classkey"`
	Status   string `json:"status"`
}

type DomainTheftProtection struct {
	Actiontypedesc   string `json:"actiontypedesc"`
	Entityid         string `json:"entityid"`
	Actionstatus     string `json:"actionstatus"`
	Status           string `json:"status"`
	Eaqid            string `json:"eaqid"`
	Error            string `json:"error"`
	Description      string `json:"description"`
	Actiontype       string `json:"actiontype"`
	Actionstatusdesc string `json:"actionstatusdesc"`
}

type DomainModifyNameserver struct {
	Actiontypedesc   string `json:"actiontypedesc"`
	Entityid         string `json:"entityid"`
	Actionstatus     string `json:"actionstatus"`
	Status           string `json:"status"`
	Eaqid            string `json:"eaqid"`
	Currentaction    string `json:"currentaction"`
	Description      string `json:"description"`
	Actiontype       string `json:"actiontype"`
	Actionstatusdesc string `json:"actionstatusdesc"`
}

type DomainDelete struct {
	Status        string `json:"status"`
	Eaqid         string `json:"eaqid"`
	Currentaction string `json:"currentaction"`
}

type DomainRegister struct {
	Actiontypedesc string `json:"actiontypedesc"`
	Entityid       string `json:"entityid"`
	Actionstatus   string `json:"actionstatus"`
	Privacydetails struct {
		Actiontypedesc          string `json:"actiontypedesc"`
		Unutilisedsellingamount string `json:"unutilisedsellingamount"`
		Sellingamount           string `json:"sellingamount"`
		Entityid                string `json:"entityid"`
		Actionstatus            string `json:"actionstatus"`
		Status                  string `json:"status"`
		Eaqid                   string `json:"eaqid"`
		Customerid              string `json:"customerid"`
		Description             string `json:"description"`
		Actiontype              string `json:"actiontype"`
		Invoiceid               string `json:"invoiceid"`
		Sellingcurrencysymbol   string `json:"sellingcurrencysymbol"`
		Actionstatusdesc        string `json:"actionstatusdesc"`
	} `json:"privacydetails"`
	Status                  string `json:"status"`
	Eaqid                   string `json:"eaqid"`
	Description             string `json:"description"`
	Actiontype              string `json:"actiontype"`
	Actionstatusdesc        string `json:"actionstatusdesc"`
	Invoiceid               string `json:"invoiceid"`
	Sellingcurrencysymbol   string `json:"sellingcurrencysymbol"`
	Sellingamount           string `json:"sellingamount"`
	Unutilisedsellingamount string `json:"unutilisedsellingamount"`
	Customerid              string `json:"customerid"`
}

type DomainTransfer struct {
	Actiontypedesc          string `json:"actiontypedesc"`
	Unutilisedsellingamount string `json:"unutilisedsellingamount"`
	Sellingamount           string `json:"sellingamount"`
	Entityid                string `json:"entityid"`
	Actionstatus            string `json:"actionstatus"`
	Privacydetails          struct {
		Priority                string `json:"priority"`
		Customerid              string `json:"customerid"`
		Actionstarted           string `json:"actionstarted"`
		Entityid                string `json:"entityid"`
		Sellingcurrencysymbol   string `json:"sellingcurrencysymbol"`
		Invoiceid               string `json:"invoiceid"`
		Unutilisedsellingamount string `json:"unutilisedsellingamount"`
		Sellingamount           string `json:"sellingamount"`
		Actionstatusdesc        string `json:"actionstatusdesc"`
		Actionstatus            string `json:"actionstatus"`
		Statusupdatetime        string `json:"statusupdatetime"`
		Actionadded             string `json:"actionadded"`
		Actiontype              string `json:"actiontype"`
		Status                  string `json:"status"`
		Entitytypeid            string `json:"entitytypeid"`
		Actiontypedesc          string `json:"actiontypedesc"`
	} `json:"privacydetails"`
	Status                string `json:"status"`
	Eaqid                 string `json:"eaqid"`
	Error                 string `json:"error"`
	Customerid            string `json:"customerid"`
	Description           string `json:"description"`
	Actiontype            string `json:"actiontype"`
	Invoiceid             string `json:"invoiceid"`
	Sellingcurrencysymbol string `json:"sellingcurrencysymbol"`
	Actionstatusdesc      string `json:"actionstatusdesc"`
}

type DomainModifyAuthCode struct {
	Actiontypedesc   string `json:"actiontypedesc"`
	Entityid         string `json:"entityid"`
	Actionstatus     string `json:"actionstatus"`
	Status           string `json:"status"`
	Eaqid            string `json:"eaqid"`
	Currentaction    string `json:"currentaction"`
	Error            string `json:"error"`
	Description      string `json:"description"`
	Actiontype       string `json:"actiontype"`
	Actionstatusdesc string `json:"actionstatusdesc"`
}

type DomainRenew struct {
	Description           string `json:"description"`
	Entityid              string `json:"entityid"`
	Actiontype            string `json:"actiontype"`
	Actiontypedesc        string `json:"actiontypedesc"`
	Eaqid                 string `json:"eaqid"`
	Actionstatus          string `json:"actionstatus"`
	Actionstatusdesc      string `json:"actionstatusdesc"`
	Invoiceid             string `json:"invoiceid"`
	Sellingcurrencysymbol string `json:"sellingcurrencysymbol"`
	Sellingamount         string `json:"sellingamount"`
	Customerid            string `json:"customerid"`
	Privacydetails        struct {
		Description           string `json:"description"`
		Entityid              string `json:"entityid"`
		Actiontype            string `json:"actiontype"`
		Actiontypedesc        string `json:"actiontypedesc"`
		Eaqid                 string `json:"eaqid"`
		Actionstatus          string `json:"actionstatus"`
		Actionstatusdesc      string `json:"actionstatusdesc"`
		Invoiceid             string `json:"invoiceid"`
		Sellingcurrencysymbol string `json:"sellingcurrencysymbol"`
		Sellingamount         string `json:"sellingamount"`
		Customerid            string `json:"customerid"`
	} `json:"privacydetails"`
}

type DomainAddChildNameserver struct {
	Actiontypedesc   string `json:"actiontypedesc"`
	Entityid         string `json:"entityid"`
	Actionstatus     string `json:"actionstatus"`
	Status           string `json:"status"`
	Eaqid            string `json:"eaqid"`
	Currentaction    string `json:"currentaction"`
	Description      string `json:"description"`
	Actiontype       string `json:"actiontype"`
	Actionstatusdesc string `json:"actionstatusdesc"`
}

type DomainModifyChildNameserver struct {
	Actiontypedesc   string `json:"actiontypedesc"`
	Entityid         string `json:"entityid"`
	Actionstatus     string `json:"actionstatus"`
	Status           string `json:"status"`
	Eaqid            string `json:"eaqid"`
	Currentaction    string `json:"currentaction"`
	Description      string `json:"description"`
	Actiontype       string `json:"actiontype"`
	Actionstatusdesc string `json:"actionstatusdesc"`
}

type DomainListLockApplied struct {
	Transferlock bool `json:"transferlock"`
	Customerlock bool `json:"customerlock"`
}

type DomainSuspend struct {
	Actiontypedesc   string `json:"actiontypedesc"`
	Entityid         string `json:"entityid"`
	Actionstatus     string `json:"actionstatus"`
	Status           string `json:"status"`
	Eaqid            string `json:"eaqid"`
	Error            string `json:"error"`
	Description      string `json:"description"`
	Actiontype       string `json:"actiontype"`
	Actionstatusdesc string `json:"actionstatusdesc"`
}

type DomainUnsuspend struct {
	Actiontypedesc   string `json:"actiontypedesc"`
	Entityid         string `json:"entityid"`
	Actionstatus     string `json:"actionstatus"`
	Status           string `json:"status"`
	Eaqid            string `json:"eaqid"`
	Error            string `json:"error"`
	Description      string `json:"description"`
	Actiontype       string `json:"actiontype"`
	Actionstatusdesc string `json:"actionstatusdesc"`
}

type DomainSuggestName struct {
	Status string `json:"status"`
	InGa   string `json:"in_ga"`
	Score  string `json:"score"`
	Spin   string `json:"spin"`
}
