package modules

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Проверяем существует ли файл лога за сегодня
		checkFile, _ := os.Stat(fmt.Sprintf("data/log/%02d-%02d-%02d.log", time.Now().Day(), time.Now().Month(), time.Now().Year()))
		today := fmt.Sprintf("%02d-%02d-%02d", time.Now().Day(), time.Now().Month(), time.Now().Year())

		if checkFile == nil {
			// Создаём если нет
			createFile, _ := os.Create(fmt.Sprintf("data/log/%s.log", today))
			defer createFile.Close()
		}

		// Забираем тело запроса в переменную
		body, _ := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, req)
		openFile, _ := os.OpenFile(fmt.Sprintf("data/log/%s.log", today), os.O_APPEND|os.O_WRONLY, 0600)
		defer openFile.Close()
		statusCode := w.Header().Get("Status-Code")

		// Приводим тело запроса к читабельному виду
		bodyStr := string(body)
		bodyStr = strings.ReplaceAll(bodyStr, "\n", "")
		bodyStr = strings.ReplaceAll(bodyStr, "  ", "")
		bodyStr = strings.ReplaceAll(bodyStr, "%20", " ")
		// Делаем запись в лог
		log.Printf(today + " " + time.Now().Format("15:04:05") + " " + req.Method + " " + req.RequestURI + " " + statusCode + bodyStr + "\r")
		openFile.WriteString(today + " " + time.Now().Format("15:04:05") + " " + req.Method + " " + req.RequestURI + " " + statusCode + bodyStr + "\r")
	})
}

func ErrLogging(w http.ResponseWriter, r *http.Request, err, params string) {
	start := time.Now()
	// Проверяем существует ли файл лога ошибок за сегодня
	checkFile, _ := os.Stat(fmt.Sprintf("data/error/%02d-%02d-%02d.log", time.Now().Day(), time.Now().Month(), time.Now().Year()))
	// Создаём если нет
	if checkFile == nil {
		createFile, _ := os.Create(fmt.Sprintf("data/error/%02d-%02d-%02d.log", time.Now().Day(), time.Now().Month(), time.Now().Year()))
		defer createFile.Close()
	}
	// Делаем запись
	openFile, _ := os.OpenFile(fmt.Sprintf("data/error/%02d-%02d-%02d.log", time.Now().Day(), time.Now().Month(), time.Now().Year()), os.O_APPEND|os.O_WRONLY, 0600)
	defer openFile.Close()
	openFile.WriteString("==========\n" + "method: " + r.RequestURI + "\ndate: " + start.Format("15:04:05") + "\nerror: " + err + " \n" + params + "\n==========\n")
}

func ErrLoggingNotREST(method, err, params string) {
	start := time.Now()
	// Проверяем существует ли файл лога ошибок за сегодня
	checkFile, _ := os.Stat(fmt.Sprintf("data/error/%02d-%02d-%02d.log", time.Now().Day(), time.Now().Month(), time.Now().Year()))
	// Создаём если нет
	if checkFile == nil {
		createFile, _ := os.Create(fmt.Sprintf("data/error/%02d-%02d-%02d.log", time.Now().Day(), time.Now().Month(), time.Now().Year()))
		defer createFile.Close()
	}
	// Делаем запись
	openFile, _ := os.OpenFile(fmt.Sprintf("data/error/%02d-%02d-%02d.log", time.Now().Day(), time.Now().Month(), time.Now().Year()), os.O_APPEND|os.O_WRONLY, 0600)
	defer openFile.Close()
	openFile.WriteString("==========\n" + "method: " + method + "\n" + "ndate: " + start.Format("15:04:05") + "\nerror: " + err + " \n" + params + "\n==========\n")
}

func WarnLogging(w http.ResponseWriter, r *http.Request, warn, params string) {
	start := time.Now()
	// Проверяем существует ли файл лога ошибок за сегодня
	checkFile, _ := os.Stat(fmt.Sprintf("data/warning/%02d-%02d-%02d.log", time.Now().Day(), time.Now().Month(), time.Now().Year()))
	// Создаём если нет
	if checkFile == nil {
		createFile, _ := os.Create(fmt.Sprintf("data/warning/%02d-%02d-%02d.log", time.Now().Day(), time.Now().Month(), time.Now().Year()))
		defer createFile.Close()
	}
	// Делаем запись
	openFile, _ := os.OpenFile(fmt.Sprintf("data/warning/%02d-%02d-%02d.log", time.Now().Day(), time.Now().Month(), time.Now().Year()), os.O_APPEND|os.O_WRONLY, 0600)
	defer openFile.Close()
	openFile.WriteString("==========\n" + "method: " + r.RequestURI + "\ndate: " + start.Format("15:04:05") + "\nwarning: " + warn + " \n" + params + "\n==========\n")
}
