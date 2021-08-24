package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"strings"

	"net/http"
	"net/url"

	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	dumpRaw = false
	api     = " "
)

func isapikeyexist(api string) error {
	if api == " " {
		return errors.New("invalid api key")
	}
	return nil
}

var city = ""
var str = ""

func telegrambot() {
	bot, err := tgbotapi.NewBotAPI("1990027676:AAGkEFcn1FDDPtKY0iBI9zGNuyYH7tTipLs")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	for {
		select {
		case update := <-upd:
			ChatID := update.Message.Chat.ID
			Text := ""
			switch update.Message.Command() {
			case "city":
				Text = update.Message.Text
				out := strings.Replace(Text, "/city", "", -1)
				city = out
				urlString := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, api)
				u, err := url.Parse(urlString)
				res, err := http.Get(u.String())
				if err != nil {
					log.Fatal(err)
				}
				jsonBlob, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					log.Fatal(err)
				}

				var data map[string]interface{}
				if dumpRaw {
					fmt.Printf("blob = %s\n\n", jsonBlob)
				}
				err = json.Unmarshal(jsonBlob, &data)
				if err != nil {
					fmt.Println("error:", err)
				}
				data = deletuselessinf(data)

				//log.Printf("[%s] %d %s", UserName, ChatID, Text)
				data2, err := json.MarshalIndent(data, "", "    ")

				str = string(data2)

				reply := replyfunction(str)

				if os.Getenv("DB_SWITCH") == "on" {

					if err := collectData(update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Text, reply); err != nil {

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Database error, but bot still working.")
						bot.Send(msg)
					}
				}

				msg := tgbotapi.NewMessage(ChatID, reply)
				bot.Send(msg)

			default:
				{
					str = " wrong input try /city cityname"
					reply := replyfunction(str)
					msg := tgbotapi.NewMessage(ChatID, reply)
					bot.Send(msg)
				}
			}
		}

	}
}
func deletuselessinf(data map[string]interface{}) map[string]interface{} {
	delete(data, "sys")
	delete(data, "weather")
	delete(data, "timezone")
	delete(data, "cod")
	delete(data, "base")
	delete(data, "dt")
	delete(data, "id")
	return data
}

func replyfunction(str string) string {

	currenttime := time.Now().Format(time.RFC850)
	return currenttime + str
}

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")
var sslmode = os.Getenv("SSLMODE")

var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

func collectData(username string, chatid int64, message string, answer string) error {

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	data := `INSERT INTO users(username, chat_id, message, answer) VALUES($1, $2, $3, $4);`

	if _, err = db.Exec(data, `@`+username, chatid, message, answer); err != nil {
		return err
	}

	return nil
}

func createTable() error {

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err = db.Exec(`CREATE TABLE users(ID SERIAL PRIMARY KEY, TIMESTAMP TIMESTAMP DEFAULT CURRENT_TIMESTAMP, USERNAME TEXT, CHAT_ID INT, MESSAGE TEXT, ANSWER TEXT);`); err != nil {
		return err
	}

	return nil
}

func main() {

	time.Sleep(1 * time.Minute)

	if os.Getenv("CREATE_TABLE") == "yes" {

		if os.Getenv("DB_SWITCH") == "on" {

			if err := createTable(); err != nil {

				panic(err)
			}
		}
	}

	time.Sleep(1 * time.Minute)
	isapikeyexist(api)
	telegrambot()
}
