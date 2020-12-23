package models

type CustomerDetails struct {
	Indiangstenabled                 string `json:"indiangstenabled"`
	Country                          string `json:"country"`
	Langpref                         string `json:"langpref"`
	Australiagstenabled              string `json:"australiagstenabled"`
	City                             string `json:"city"`
	Creationdt                       string `json:"creationdt"`
	OtherState                       string `json:"other_state"`
	Zip                              string `json:"zip"`
	Telno                            string `json:"telno"`
	Russiavatenabled                 string `json:"russiavatenabled"`
	IsDominicanTaxConfiguredByParent bool   `json:"isDominicanTaxConfiguredByParent"`
	Customerid                       string `json:"customerid"`
	TwofactorsmsauthEnabled          string `json:"twofactorsmsauth_enabled"`
	TwofactorgoogleauthEnabled       string `json:"twofactorgoogleauth_enabled"`
	Parentid                         string `json:"parentid"`
	Name                             string `json:"name"`
	Newzealandgstenabled             string `json:"newzealandgstenabled"`
	Singaporegstenabled              string `json:"singaporegstenabled"`
	Pin                              string `json:"pin"`
	Stateid                          string `json:"stateid"`
	Address1                         string `json:"address1"`
	Resellerid                       string `json:"resellerid"`
	Company                          string `json:"company"`
	TwofactorauthEnabled             string `json:"twofactorauth_enabled"`
	Telnocc                          string `json:"telnocc"`
	Username                         string `json:"username"`
	State                            string `json:"state"`
	Salespersonshash                 struct {
	} `json:"salespersonshash"`
	Totalreceipts           string `json:"totalreceipts"`
	Salescontactid          string `json:"salescontactid"`
	Useremail               string `json:"useremail"`
	Indianservicetaxenabled string `json:"indianservicetaxenabled"`
	Customerstatus          string `json:"customerstatus"`
}

type CustomerMoveProduct struct {
	Status string `json:"status"`
}
