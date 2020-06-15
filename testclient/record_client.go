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

func HttpPostFormGetRecordHoneycomb(roleID, operatorUid, targetUid, recordType, startTime, endTime string) string {
	resp, err := tools.HoneycombPostForm(constant.HoneycombAddress+"/purchase/get",
		url.Values{"roleid": {roleID}, "operatorUid": {operatorUid}, "targetUid": {targetUid}, "type": {recordType}, "startTime": {startTime}, "endTime": {endTime}})

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	operateRecordReturn := new(returns.OperateRecordReturn)
	result := serialize.JSONDeserialize(operateRecordReturn, string(body))
	fmt.Println(roleID, result.(*returns.OperateRecordReturn).ErrCode+result.(*returns.OperateRecordReturn).ErrMsg)
	for index, record := range result.(*returns.OperateRecordReturn).OperateRecords {
		fmt.Println(index, record)
	}
	return string(body)
}
