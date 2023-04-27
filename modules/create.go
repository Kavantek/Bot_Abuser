package modules

import (
	// "MSB/cmd/DB"
	con "MSB/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	// "time"
)

var haveSQL bool = false
var connDB *sql.DB
var data con.Config

// Функция чтения ./config/config.json
func CreateParamServer(configPath string) con.Config {
	file, _ := ioutil.ReadFile(configPath)

	json.Unmarshal(file, &data)

	return data
}

// Функция проверки и создание директорий
func CheckParam(data con.Config, dirsCreate bool) {
	errorCreate := false
	fmt.Printf(`
========== ПАРАМЕТРЫ ==========

========== Версия ==========
Версия: %v
============================

========== Сервер ==========
Хост: %v
Порт: %v
============================

======= База данных ========
DB_PGSQL: %v
============================

======== Директории ========
	`,
		data.Version,
		data.Server.Host,
		data.Server.Port,
		data.DB_PGSQL,
	)

	if dirsCreate {
		errData := os.Mkdir("data", 0755)
		if errData != nil && !os.IsExist(errData) {
			fmt.Println(`
data:`, errData)
			errorCreate = true
		} else {
			fmt.Println(`
data: OK`)
		}

		errLog := os.Mkdir("data/log", 0755)
		if errLog != nil && !os.IsExist(errLog) {
			fmt.Println(`
log:`, errLog)
			errorCreate = true
		} else {
			fmt.Println(`
log: OK`)
		}

		errFalse := os.Mkdir("data/error", 0755)
		if errFalse != nil && !os.IsExist(errFalse) {
			fmt.Println(`
error:`, errFalse)
			errorCreate = true
		} else {
			fmt.Println(`
error: OK`)
		}

		warnLog := os.Mkdir("data/warning", 0755)
		if warnLog != nil && !os.IsExist(warnLog) {
			fmt.Println(`
Warning:`, warnLog)
			errorCreate = true
		} else {
			fmt.Println(`
Warning: OK`)
		}

		errUploads := os.Mkdir("data/upload", 0755)
		if errUploads != nil && !os.IsExist(errUploads) {
			fmt.Println(`
Uploads:`, errUploads)
			errorCreate = true
		} else {
			fmt.Println(`
Uploads: OK`)
		}
	}
	if errorCreate {
		fmt.Println(`
==== ПОДГОТОВКА ЗАВЕРШЕНА С ОШИБКАМИ ====`)
	} else {
		fmt.Println(`
==== ПОДГОТОВКА ЗАВЕРШЕНА БЕЗ ОШИБОК ====`)
	}
	// go SetConnectSQL()
}

// func SetConnectSQL() {
// 	for {
// 		var err error
// 		connDB, err = sql.Open("postgres", data.DB_PGSQL)
// 		if err != nil {
// 			SendMail("SQL connection Error", "Не удалось подключиться к серверу SQL", data)
// 		} else {
// 			err = DB.SqlPing(data, connDB)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 			} else {
// 				fmt.Printf("====== Подлючение к БД установлено ======\n")
// 				haveSQL = true
// 				break
// 			}
// 		}
// 		time.Sleep(time.Second * 5)
// 	}
// }

// func GetHaveSQL() (bool, *sql.DB) {
// 	return haveSQL, connDB
// }
