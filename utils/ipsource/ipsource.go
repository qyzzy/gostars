package ipsource

import (
	"encoding/json"
	"gostars/models/common"
	"io/ioutil"
	"net/http"
	"strings"
)

func OnlineIpInfo(ip string) *common.IPInfo {
	if !isRemoteAddr(ip) {
		return &common.IPInfo{
			QueryIp: "localhost",
			Region:  "local",
		}
	}
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

// todo
func isRemoteAddr(ip string) bool {
	tmp := strings.Split(ip, ".")
	if len(tmp) != 4 {
		return false
	}
	return true
}
