package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/axgle/mahonia"

	"500wan/spider"

	"github.com/robfig/cron"
)

var msglogger *log.Logger

func init() {
	logpath := os.Getenv("LOGPATH")
	logfile := os.Getenv("LOGFILE")
	logfilename := fmt.Sprintf("%s/%s", logpath, logfile)
	fd, err := os.OpenFile(logfilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Panic("failed to Openfile")
	}
	msglogger = log.New(fd, "500wan:", log.LUTC)
	msglogger.SetFlags(log.LUTC | log.Lmicroseconds | log.Ldate)
}

func MakeHtmlName() string {
	htmlPath := os.Getenv("DOWNPATH")
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	filepath := fmt.Sprintf("%s/%s/%s", htmlPath, year, month)
	os.MkdirAll(filepath, 0777)
	return fmt.Sprintf("%s/%s/%s/%s-%s-%s.html", htmlPath, year, month, year, month, day)
}

func DownloadHtml() {
	msglogger.Println("begin down html ")
	sp := spider.NewSpider()
	body := sp.Fetch("http://live.500.com")

	decoder := mahonia.NewDecoder("GB18030")
	utf8_body := decoder.ConvertString(string(body))

	filename := MakeHtmlName()
	msglogger.Println("file =", filename)
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		msglogger.Panic("Open file err ", filename, err)
	}
	fd.WriteString(utf8_body)
	msglogger.Println("end down html ")
}

func main() {
	msglogger.Println("begin 500man's dream!")
	cc := cron.New()
	cc.AddFunc("0 30 0 * * *", DownloadHtml)
	cc.Start()
	defer cc.Stop()
	for {
		time.Sleep(1)
	}
}
