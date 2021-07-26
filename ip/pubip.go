// Package pubip gets your public IP address from several services
package ip

import (
	"io/ioutil"
	"net/http"
)

// APIURIs is the URIs of the services.
var APIURI = "http://ip.dhcp.cn/?ip"

func GetPublicIP(client *http.Client) (string, error) {
	c := http.DefaultClient
	if client != nil {
		c = client
	}
	// 获取外网 IP
	resp, err := c.Get(APIURI)
	if err != nil {
		panic(err)
	}
	// 程序在使用完 response 后必须关闭 response 的主体。
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), err
}
