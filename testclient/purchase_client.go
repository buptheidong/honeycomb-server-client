package testclient

import (
	"fmt"
	"honeycomb-server-client/constant"
	"honeycomb-server-client/serialize"
	"honeycomb-server-client/tools"
	"honeycomb/model/returns"
	"io/ioutil"
	"net/url"
)

func HttpPostFormSubmitPurchase(roleID, purchaseID, name, desc, purchaseType, contentType, place, adjustModulus string) string {
	resp, err := tools.HoneycombPostForm(constant.HoneycombAddress+"/purchase/submit",
		url.Values{"roleid": {roleID}, "purchaseid": {purchaseID}, "name": {name}, "desc": {desc}, "purchaseType": {purchaseType}, "contentType": {contentType}, "place": {place}, "adjustModulus": {adjustModulus}})

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	purchaseReturn := new(returns.PurchaseReturn)
	result := serialize.JSONDeserialize(purchaseReturn, string(body))
	fmt.Println(result.(*returns.PurchaseReturn))
	return string(body)
}
