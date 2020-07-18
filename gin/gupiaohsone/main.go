package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
)

//基于GO的股票行情查询api调用代码实例
//代码描述：基于GO的股票行情查询api调用代码实例 代码平台：聚合数据
//---------------------------------- // 股票查询调用示例代码 － 聚合数据

const APPKEY = "db21317cca0c494f83fa3110d8f269fe" //您申请的APPKEY

//定义结构体，并序列化
type Alldata struct {
	Resultcode string   `json:"resultcode"` ///*返回码，200:正常*/
	Reason     string   `json:"reason"`     //"SUCCESSED!",
	Result     []Result `json:"result"`     //详细信息
	Error_code int      `json:"error_code"` //0
}

type Result struct {
	Data      Data      `json:"data"`      //单股信息
	Dapandata Dapandata `json:"dapandata"` //大盘显示的信息
	Gopicture Gopicture `json:"gopicture"` //图片
}

type Dapandata struct {
	Dot       string `json:"dot"`       /*当前价格*/
	Name      string `json:"name"`      //股票名
	NowPic    string `json:"nowPic"`    /*涨量*/
	Rate      string `json:"rate"`      /*涨幅(%)*/
	TraAmount string `json:"traAmount"` /*成交额(万)*/
	TraNumber string `json:"traNumber"` /*成交量*/
}
type Gopicture struct {
	Minurl   string `json:"Minurl"`   /*分时K线图*/
	Dayurl   string `json:"Dayurl"`   /*日K线图*/
	Weekurl  string `json:"Weekurl"`  /*周K线图*/
	Monthurl string `json:"Monthurl"` /*月K线图*/
}

type Data struct {
	Gid            string `json:"gid"`            /*股票编号*/
	IncrePer       string `json:"increPer"`       /*涨跌百分比*/
	Increase       string `json:"increase"`       /*涨跌额*/
	Name           string `json:"name"`           /*股票名称*/
	TodayStartPri  string `json:"todayStartPri"`  /*今日开盘价*/
	YestodEndPri   string `json:"yestodEndPri"`   /*昨日收盘价*/
	NowPri         string `json:"nowPri"`         /*当前价格*/
	TodayMax       string `json:"todayMax"`       /*今日最高价*/
	TodayMin       string `json:"todayMin"`       /*今日最低价*/
	CompetitivePri string `json:"competitivePri"` /*竞买价*/
	ReservePri     string `json:"reservePri"`     /*竞卖价*/
	TraNumber      string `json:"traNumber"`      /*成交量*/
	TraAmount      string `json:"traAmount"`      /*成交金额*/
	BuyOne         string `json:"buyOne"`         /*买一*/
	BuyOnePri      string `json:"buyOnePri"`      /*买一报价*/
	BuyTwo         string `json:"buyTwo"`         /*买二*/
	BuyTwoPri      string `json:"buyTwoPri"`      /*买二报价*/
	BuyThree       string `json:"buyThree"`       /*买三*/
	BuyThreePri    string `json:"buyThreePri"`    /*买三报价*/
	BuyFour        string `json:"buyFour"`        /*买四*/
	BuyFourPri     string `json:"buyFourPri"`     /*买四报价*/
	BuyFive        string `json:"buyFive"`        /*买五*/
	BuyFivePri     string `json:"buyFivePri"`     /*买五报价*/
	SellOne        string `json:"sellOne"`        /*卖一*/
	SellOnePri     string `json:"sellOnePri"`     /*卖一报价*/
	SellTwo        string `json:"sellTwo"`        /*卖二*/
	SellTwoPri     string `json:"sellTwoPri"`     /*卖二报价*/
	SellThree      string `json:"sellThree"`      /*卖三*/
	SellThreePri   string `json:"sellThreePri"`   /*卖三报价*/
	SellFour       string `json:"sellFour"`       /*卖四*/
	SellFourPri    string `json:"sellFourPri"`    /*卖四报价*/
	SellFive       string `json:"sellFive"`       /*卖五*/
	SellFivePri    string `json:"sellFivePri"`    /*卖五报价*/
	Date           string `json:"date"`           /*日期*/
	Time           string `json:"time"`           /*时间*/
}

func Getaction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./showgupiao.tmpl")
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
	//1.股票查询
	mygupiao := Request1()
	err = t.Execute(w, *mygupiao)
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}

}
func main() {
	http.HandleFunc("/", Getaction)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
	}

}

//1.股票查询
func Request1() *Alldata {
	//沪深请求示例：http://web.juhe.cn:8080/finance/stock/hs?gid=sh601009&key=您申请的APPKEY
	juheURL := "http://web.juhe.cn:8080/finance/stock/hs"

	//初始化参数
	param := url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("key", APPKEY)     //您申请的APPKEY
	param.Set("gid", "sh601009") //股票编号，上海股市以sh开头，深圳股市以sz开头如：sh601009（type为0或者1时gid不传）

	//发送请求
	resultdata, err := Get(juheURL, param)
	mydata := Alldata{}
	if err != nil {
		fmt.Errorf("请求失败,错误信息:\r\n%v", err)
	} else {
		json.Unmarshal([]byte(resultdata), &mydata)
		fmt.Println(fmt.Sprintf("%+v", mydata))
	}
	return &mydata
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
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	//fmt.Println("----2",string(structdata),err2) // 打印接口返回的所有数据json格式
	return data, err2
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
