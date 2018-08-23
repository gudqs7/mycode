package main

import (
	"fmt"
    "io/ioutil"
	"net/http"
	"strconv"
)

type R struct {
	count int
	url string
	param map[string]string
}

func main(){

	var r R
	r.count = 300
	r.url = "http://localhost:9090/wechat/bargain/helpBargain?"
	r.param = make(map[string]string)
	r.param["myBargainId"] = "8"
	r.param["shopClientId"] = "550"
	channel := make(chan int)
	
	// go getReq(r, channel, 1)
	goBargain(r, channel);

	<- channel

}

func goBargain(r R, c chan int){
	count := 0
	for count < r.count {
		go getReq2(r, c, count)
		count++
	}
}

func getReq(r R, c chan int, num int){
	count := 0
	for count < r.count {
		url := r.url
		scid := r.param["shopClientId"]
		id , _ :=strconv.Atoi(scid)
		scid =strconv.Itoa(id+1)
		r.param["shopClientId"] = scid
		// fmt.Printf("%s : %d\n", scid, id)
		for k, v := range r.param {
			url += k + "=" + v + "&"
		}
		url = url[0:len(url)-1]
		
		count++

		client := &http.Client{}
		//提交请求
		reqest, err := http.NewRequest("GET", url, nil)

		//增加header选项
		reqest.Header.Add("Cookie", "JSESSIONId=fidjkasfjejfjdsg")
		reqest.Header.Add("User-Agent", "Apple IOS Wechat")
		reqest.Header.Add("X-Requested-With", "XMLHttpRequest")

		if err != nil {
			fmt.Println(err)
		}   
		//处理返回结果
		resp, e := client.Do(reqest)

		if e != nil {
			fmt.Println(e)
		}  

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Printf("No.%d: Times:%d Info:%s : %s\n", num, count, url, string(body))
	}
	c <- 0
}

func getReq2(r R, c chan int, num int){
	url := r.url
	scid := r.param["shopClientId"]
	id , _ :=strconv.Atoi(scid)
	scid =strconv.Itoa(id+1)
	r.param["shopClientId"] = scid
	// fmt.Printf("%s : %d\n", scid, id)
	for k, v := range r.param {
		url += k + "=" + v + "&"
	}
	url = url[0:len(url)-1]

	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Cookie", "JSESSIONId=fidjkasfjejfjdsg")
	reqest.Header.Add("User-Agent", "Apple IOS Wechat")
	reqest.Header.Add("X-Requested-With", "XMLHttpRequest")

	if err != nil {
		fmt.Println(err)
	}   
	//处理返回结果
	resp, e := client.Do(reqest)

	if e != nil {
		fmt.Println(e)
	}  

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Printf("No.%d: Times:%d Info:%s : %s\n", num, count, url, string(body))
	c <- 0
}