package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

func main() {
	args := os.Args
	println(args[1])
	url, _ := url.Parse(args[1])
	req := http.Request{
		Method: "GET",
		URL:    url,
	}
	client := http.Client{}
	resp, err := client.Do(&req)
	if err != nil {
		println("err1:", err.Error())
		cmd := exec.Command("systemctl", "restart", "network")
		stdout, err := cmd.StdoutPipe()
		if err != nil { //获取输出对象，可以从该对象中读取输出结果
			println("err3:", err.Error())
			return
		}
		err = cmd.Run()
		if err != nil {
			println("err4:", err.Error())
			return
		}
		if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
			println("err5:", err.Error())
			return
		} else {
			log.Println(string(opBytes))
			return
		}
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			println("err2:", err.Error())
		}
		println(string(body))
		return
	}
}
