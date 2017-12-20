package smartisanFunc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"smartisan/Data"
	"strconv"
	"strings"
	"time"
)

func BuyGood(goodId int) {
	postDataStr := "address_id=14153&district_code=440106&telphone=&district_name=%E5%A4%A9%E6%B2%B3%E5%8C%BA&telephone=&accept_name=%E9%99%88%E6%B8%8A&city_code=440100&city_name=%E5%B9%BF%E5%B7%9E%E5%B8%82&province_code=440000&street=%E4%BA%94%E5%B1%B1%E8%B7%AF+371+%E5%8F%B7%E4%B9%8B%E4%B8%80%E4%B8%BB%E6%A5%BC%E4%B8%AD%E5%85%AC%E6%95%99%E8%82%B2%E5%A4%A7%E5%8E%A6+2301&area_code=&mobile=15111048113&province_name=%E5%B9%BF%E4%B8%9C%E7%9C%81&default=1&pay_channel_type=1&invoice_style=2&invoice_type=0&discount_amount=0&goods_infos=%5B%7B%22group_id%22%3A0%2C%22rule_id%22%3A0%2C%22sku%22%3A" + strconv.Itoa(goodId) + "%2C%22pgoods_id%22%3A0%2C%22num%22%3A1%2C%22is_gift%22%3A0%7D%5D&source=11"
	body := ioutil.NopCloser(strings.NewReader(postDataStr)) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "http://store.smartisan.com/index.php?r=Order/Create", body)
	req.Header.Set("Referer", "http://store.smartisan.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时                                                       //看下发送的结构

	resp, _ := client.Do(req) //发送
	fmt.Println(req)
	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data) + time.Now().String())
}
