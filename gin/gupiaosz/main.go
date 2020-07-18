package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"
)

//基于GO的股票行情查询api调用代码实例
//代码描述：基于GO的股票行情查询api调用代码实例 代码平台：聚合数据
//---------------------------------- // 股票查询调用示例代码 － 聚合数据

const APPKEY = "db21317cca0c494f83fa3110d8f269fe" //您申请的APPKEY

//定义结构体，并序列化
type Alldata struct {
	Error_code int `json:"error_code"` //错误代码，0成功
	Reason string `json:"reason"` //错误信息，SUCCESSED!
	Result Resultdata `json:"result"`
}

type Resultdata struct {
	TotalCount string `json:"totalCount"` //股票数量
	Page string  `json:"page"` //页数
	Num string  `json:"num"` //每页数量
	Data []Data `json:"data"`
}

type Data struct {
	Symbol string  `json:"symbol"` //聚合股票代码
	Name string  `json:"name"` //股票名称
	Trade  string  `json:"trade"` //当前价格
	Pricechange string  `json:"pricechange"` //比昨日收盘价的涨幅
	Changepercent string  `json:"changepercent"` //涨跌幅
	Buy string  `json:"buy"` //当前最高买价
	Sell string  `json:"sell"` //当前最低卖价
	Settlement string  `json:"settlement"` //最近一次成交价
	Open string  `json:"open"` //开盘价
	High string  `json:"high"` //当日最高价
	Low string  `json:"low"` //当日最低价
	Volume int  `json:"volume"` //成交总手数
	Amount int  `json:"amount"` //成交额
	Code string  `json:"code"` //股票代码
	Ticktime string  `json:"ticktime"` //统计时间
}
func main() {

	//1.股票查询
	Request1()

}

//1.股票查询
func Request1() {
	//深A请求示例： http://web.juhe.cn:8080/finance/stock/szall?key=您申请的APPKEY&page=1
	juheURL := "http://web.juhe.cn:8080/finance/stock/szall"

	//初始化参数
	param := url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("key", APPKEY) //您申请的APPKEY
	param.Set("type", "20") //每页返回条数,1(20条默认),2(40条),3(60条),4(80条)
	param.Set("stock", "a") //a表示A股，b表示B股,默认所有
	param.Set("page", "2") //1代表第一页

	//发送请求
	resultdata, err := Get(juheURL, param)
	if err != nil {
		fmt.Errorf("请求失败,错误信息:\r\n%v", err)
	} else {
		//var netReturn map[string]interface{}
		//json.Unmarshal(resultdata, &netReturn)      //二进制转存map
		//if netReturn["error_code"].(float64) == 0 { //类型转换为数字后再比较
		//	//fmt.Printf("接口返回result字段是:\r\n%v", netReturn["result"])
		//	result:=netReturn["result"]
		//	mymapresult :=result.(map[string]interface{}) //interface转其它类型map
		//	mydata:= mymapresult["resultdata"]
		//	myarr:=mydata.([]interface{}) //转换为gp类型的数组
		//	arr1:=myarr[0]
		//	var mymapstring map[string]string
		//	mymapstring =arr1.(map[string]string)
		//	//mygp:= structdata{}
		//	//json.Unmarshal([]byte(string(arr1)),&mygp)
		//	//var mymapgp map[string]resultdata
		//	//myapgpkey:=mymapinterface["name"]
		//	fmt.Printf("接口返回result字段是:\r\n%v", mymapstring)
		//}

		mydata :=Alldata{}
		json.Unmarshal([]byte(resultdata), &mydata)
		fmt.Println(fmt.Sprintf("%+v",mydata))
	}
}

// get 网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	//判断参数是否为URL
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery=params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	data,err2 :=ioutil.ReadAll(resp.Body)
	//fmt.Println("----2",string(structdata),err2) // 打印接口返回的所有数据json格式
	return data,err2
}

// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
