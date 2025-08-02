//go:build !solution

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const LOG_FILENAME = "async_log.txt"

var logMutex sync.Mutex
var logFile *os.File

func InitLogger() error {
	var err error
	logMutex.Lock()
	defer logMutex.Unlock()

	logFile, err = os.OpenFile(LOG_FILENAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return err
}

func CloseLogger() {
	logMutex.Lock()
	defer logMutex.Unlock()

	if logFile != nil {
		logFile.Close()
		logFile = nil
	}
}

func MakeLog(msg string, duration ...time.Duration) {
	logMutex.Lock()
	defer logMutex.Unlock()

	if logFile == nil {
		var err error
		logFile, err = os.OpenFile(LOG_FILENAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("Failed to initialize logger: %v", err)
			return
		}
	}

	// Базовое сообщение с timestamp
	logMsg := fmt.Sprintf(
		"[%s] %s",
		time.Now().Format("2006-01-02 15:04:05"), msg,
	)

	// Добавляем duration, если он был передан
	if len(duration) > 0 {
		logMsg += fmt.Sprintf("\t%f", duration[0].Seconds())
	}

	// Добавляем перевод строки
	logMsg += "\n"

	_, writeErr := logFile.WriteString(logMsg)

	if writeErr != nil {
		log.Printf("Failed to write to log file: %v", writeErr)
	}
}

func checkErr(err error, ignore bool) {
	if err != nil {
		tag := "[ERROR] "
		msg := "fetch: " + err.Error() + "\n"
		if !ignore {
			log.SetPrefix("")
			log.SetFlags(0)
			log.Fatal(msg)
		} else {
			MakeLog(tag + msg)
		}
	}
}

func CheckUrl(url string) {
	start_time := time.Now()
	response, err := http.Get(url)
	checkErr(err, true)

	if response == nil {
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	checkErr(err, true)
	bodySize := len(body)
	duration := time.Since(start_time)
	fmt.Printf("%f %d %s  \n", duration.Seconds(), bodySize, url)
	msg_body := "[INFO] " + url
	MakeLog(msg_body, duration)
}

func main() {
	defer CloseLogger()
	err := InitLogger()
	if err != nil {
		panic(err)
	}

	var urls = os.Args[1:]
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			CheckUrl(u)
		}(url)
	}
	wg.Wait()
}
