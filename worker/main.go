package main

import (
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.Println("Start Logging Load Average")
	for {
		cmdstr := "uptime | awk '$0 = $9' | sed -e 's/,//g'" // load average取得して 9番目のフィールド取得して カンマを削除
		loadAverageByte, err := exec.Command("sh", "-c", cmdstr).Output()
		if err != nil {
			log.Fatal(err)
		}
		loadAverageStr := strings.TrimRight(string(loadAverageByte), "\n")
		loadAverageFloat, err := strconv.ParseFloat(loadAverageStr, 32)
		if err != nil {
			log.Fatal(err)
		}

		cpuNum := runtime.NumCPU()

		loadIndicator := loadAverageFloat / float64(cpuNum)

		log.Printf("Load Average =%f\n", loadAverageFloat)
		log.Printf("CPU num =%d\n", cpuNum)
		log.Printf("Load Average / CPU num =%f\n", loadIndicator)
		time.Sleep(1 * time.Second)
	}
}
