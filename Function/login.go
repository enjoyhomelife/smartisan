package smartisanFunc

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"smartisan/Data"
	"strconv"
	"strings"
)

//日志初始化
func LogInit() {
	var logFile *os.File
	var err error
	// 定义一个文件
	fileName := smartisanData.UserName + ".log"
	if checkFileIsExist(fileName) { //如果文件存在
		logFile, err = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
	} else {
		logFile, err = os.Create(fileName) //创建文件
	}

	if err != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象
	smartisanData.Log = log.New(logFile, "[Log]", log.LstdFlags)
	//配置log的Flag参数
	smartisanData.Log.SetFlags(smartisanData.Log.Flags() | log.LstdFlags)

}

func LogStr(a ...interface{}) {
	fmt.Println(a)
	smartisanData.Log.Println(a)
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func Draw() {
	v := url.Values{}
	v.Set("batchId", "3555")
	v.Set("subkey", "1502971200")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "http://www.smartisan.com/store/index.php?r=lottery/winlist", body)
	req.Header.Set("Referer", "https://account.smartisan.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时                                                       //看下发送的结构

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}

func Hit() {
	v := url.Values{}
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "http://www.smartisan.com/store/index.php?r=activity/getJianGuoPro100DayCoupon&action=get", body)
	req.Header.Set("Referer", "https://account.smartisan.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时                                                       //看下发送的结构

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}

//登陆函数
func Login() {
	smartisanData.CookieJar, _ = cookiejar.New(nil)
	smartisanData.Tr = &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}

	v := url.Values{}
	//v.Set("username", "18520484208")
	//v.Set("password", "cy921103")
	v.Set("username", "15111048113")
	v.Set("password", "chenyuanjiayou")
	//v.Set("username", "13450281155")
	//v.Set("password", "yy110389")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "https://account.smartisan.com/v2/session/?m=post", body)
	req.Header.Set("Referer", "https://account.smartisan.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时                                                       //看下发送的结构

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
	var loginstruct smartisanData.LoginStruct

	if err := json.Unmarshal(data, &loginstruct); err == nil {
		if loginstruct.Data.Uid > 0 {
			LogStr("登陆成功,ID:" + strconv.Itoa(loginstruct.Data.Uid))
		}
	}

	//获取地址信息
	//getAddress()
}

func getAddress() {
	postDataStr := ""

	body := ioutil.NopCloser(strings.NewReader(postDataStr)) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "http://store.smartisan.com/serv/v1/address/list", body)

	req.Header.Set("Referer", "http://store.smartisan.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时                                                       //看下发送的结构

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(data, &smartisanData.AddressData)
	if err != nil {
		return
	}

	LogStr("获取收货地址成功:", smartisanData.AddressData.Data.List[0])
}
