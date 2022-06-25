package ipsource

import (
	"encoding/json"
	"gostars/models/common"
	"io/ioutil"
	"net/http"
)

func OnlineIpInfo(ip string) *common.IPInfo {
	url := "http://ip-api.com/json/" + ip + "?lang=zh-CN"
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var result common.IPInfo
	if err := json.Unmarshal(out, &result); err != nil {
		return nil
	}
	return &result
}
