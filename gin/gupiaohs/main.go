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
	Error_code int `json:"error_code"` //0成功，202101 参数错误 202102 查询不到结果 202103 网络异常 系统级 10001 错误的请求KEY
	Reason string  `json:"reason"` //"SUCCESSED!",
	Result Result `json:"result"` //详细信息
}

type Result struct {
	DealNum string `json:"dealNum"` /*成交量*/
	DealPri string `json:"dealPri"` /*成交额*/
	HighPri string `json:"highPri"` /*最高*/
	IncrePer string `json:"increPer"` /*涨跌百分比*/
	Increase string `json:"increase"` /*涨跌幅*/
	Lowpri string `json:"lowpri"` /*最低*/
	Name string `json:"name"` /*名称*/
	Nowpri string `json:"nowpri"` /*当前价格*/
	OpenPri string `json:"openPri"` /*今开*/
	Time string `json:"time"` /*时间*/
	YesPri string `json:"yesPri"` /*昨收*/
}
func main() {

	//1.股票查询
	Request1()

}

//1.股票查询
func Request1() {
	//沪深请求示例：http://web.juhe.cn:8080/finance/stock/hs?gid=sh601009&key=您申请的APPKEY
	juheURL := "http://web.juhe.cn:8080/finance/stock/hs"

	//初始化参数
	param := url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("key", APPKEY) //您申请的APPKEY
	param.Set("type", "0") //0代表上证综合指数，1代表深证成份指数(输入此字段时,gid字段不起作用)

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
