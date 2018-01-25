package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/youtangai/Optima/conductor/model"
	"github.com/youtangai/Optima/monitor/config"
)

func main() {
	conductorHost := flag.String("conductor_ip", "192.168.64.12", "conductor's IP")
	flag.Parse()
	config.SetConductorHost(*conductorHost)

	log.Println("Start Logging Load Average")
	cpuNum := runtime.NumCPU()     //cpuの数を取得
	hostname, err := os.Hostname() //ホストネームを取得
	if err != nil {
		log.Fatal(err)
	}

	addr := getNetIP()
	log.Printf("network interface = %s\n", addr)

	for {
		loadAverage := getLoadAverage()
		loadIndicator := loadAverage / float64(cpuNum) //負荷指標を計算

		log.Printf("Load Average =%f\n", loadAverage)
		log.Printf("CPU num =%d\n", cpuNum)
		log.Printf("Load Average / CPU num =%f\n", loadIndicator)
		log.Printf("Host name = %s\n", hostname)

		sendLoadIndicator(hostname, addr, loadIndicator)

		time.Sleep(1 * time.Minute)
	}
}

func getLoadAverage() float64 { //ロードアベレージを取得
	cmdstr := "uptime | awk '$0 = $12' | sed -e 's/,//g'" // load average取得して 9番目のフィールド取得して カンマを削除
	loadAverageByte, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
	}
	loadAverageStr := strings.TrimRight(string(loadAverageByte), "\n") //バイトをstringに変換して，改行を削除
	loadAverageFloat, err := strconv.ParseFloat(loadAverageStr, 32)    //Floatに変換
	if err != nil {
		log.Fatal(err)
	}
	return loadAverageFloat
}

func getNetIP() string { //IPを取得
	addrs, err := net.InterfaceAddrs() //ネットワークインタフェースを複数取得
	if err != nil {
		log.Fatal(err)
	}

	addr := strings.Split(addrs[1].String(), "/")[0] //2つ目のインタフェースを取得して， /で分割して，ipだけ取得
	return addr
}

func sendLoadIndicator(hostname string, addr string, load float64) {
	conductorURL := "http://" + config.GetConductorHost() + ":62070"
	reqBody := new(model.LoadIndicatorJson)
	reqBody.HostIP = addr
	reqBody.HostName = hostname
	reqBody.LoadIndicator = load
	json, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}
	_, err = http.Post(conductorURL+"/load_indicator", "application/json", bytes.NewBuffer(json))
	if err != nil {
		log.Fatal(err)
	}
}
