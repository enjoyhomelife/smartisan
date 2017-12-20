package smartisanData

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/cookiejar"
)

//用户名 密码
var UserName = "15111048113"
var PassWord = "chenyuanjiayou"

//登陆cookie
var CookieJar *cookiejar.Jar

var Tr = &http.Transport{
	TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	DisableCompression: true,
}

//登陆返回json结构
type LoginData struct {
	Uid int
}

type LoginStruct struct {
	Errno int
	Data  LoginData
}

type sigleAddress struct {
	District_code int
	Telphone      string
	District_name string
	Id            int
	Telephone     string
	Accept_name   string
	City_code     int
	City_name     string
	Province_code int
	Street        string
	Area_code     string
	Mobile        string
	Province_name string
	Default       int
}

type addressData struct {
	List  []sigleAddress
	Count int
}

type address struct {
	Extra   []int
	Data    addressData
	Code    int
	ErrInfo []int
}

//地址信息
var AddressData address

//日志
var Log *log.Logger
