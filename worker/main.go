package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
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

		time.Sleep(1 * time.Second)
	}
}

func getLoadAverage() float64 { //ロードアベレージを取得
	cmdstr := "uptime | awk '$0 = $9' | sed -e 's/,//g'" // load average取得して 9番目のフィールド取得して カンマを削除
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
