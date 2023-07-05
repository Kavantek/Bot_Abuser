package main

import (
	mod "MSB/modules"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var agrCount int
var timer int
var Counter int
var SpamTimer time.Time
var MsgFlag bool = false

func main() {
	config := mod.CreateParamServer("./config/config.json")
	mod.CheckParam(config, true)

	bot, err := tgbotapi.NewBotAPI(config.BotParams.Token)
	if err != nil {
		panic(err)
	}

	agrCount = 0

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	var chatID int64 = -1001673668774

	var R tgbotapi.KeyboardButton

	R.Text = "Рано"

	var U tgbotapi.KeyboardButton

	U.Text = "Мы уже"

	var Z tgbotapi.KeyboardButton

	Z.Text = "Занято"

	var Sm tgbotapi.KeyboardButton

	Sm.Text = "Передвинь таймер"

	var Sc tgbotapi.KeyboardButton

	Sc.Text = "Сколько"

	var P tgbotapi.KeyboardButton

	P.Text = "Потеря бойца"

	first := tgbotapi.NewKeyboardButtonRow(R, U, Z)
	second := tgbotapi.NewKeyboardButtonRow(P, Sc)
	third := tgbotapi.NewKeyboardButtonRow(Sm)

	butt := tgbotapi.NewReplyKeyboard(first, second, third)

	msg := tgbotapi.NewMessage(chatID, "Я снова с вами!")
	bot.Send(msg)
	stiker := tgbotapi.NewStickerShare(chatID, "CAACAgIAAxkBAAEIpvdkQTa5LWDNwb_e4qV6FVNAGaGRzAACNRIAAlbWCUhVwiQqqj_qfi8E")
	bot.Send(stiker)
	menu := tgbotapi.NewMessage(chatID, `Вот команды, чтоб вам - ленивым жопам было удобнее:
Рано - Откладывает таймер на 5 минут
Мы уже - Откладывает таймер на 30 минут
Занято - Откладывает таймер на 10 минут
Выходной - Откладывает таймер до следующего дня
Передвинь таймер на Х - Откладывает таймер на Х минут
Сколько - показывает сколько времени до следующего перекура`)
	menu.ReplyMarkup = butt
	bot.Send(menu)

	go Bye(chatID, bot)

	start := time.Hour * 8
	fmt.Printf("Start: %v\n", start)
	end := time.Hour * 17
	fmt.Printf("End: %v\n", end)

	go func() {
		for {
			now := time.Duration(time.Now().Hour()) * time.Hour
			weekday := time.Now().Weekday()
			fmt.Printf("Now: %v, %v\n", now, weekday)
			if now >= start && now < end && fmt.Sprintf("%v", weekday) != "Sunday" && fmt.Sprintf("%v", weekday) != "Saturday" {
				timer = 60
				for {
					time.Sleep(time.Minute)
					timer--
					now := time.Duration(time.Now().Hour()) * time.Hour
					if timer <= 0 && now < end {
						audio := tgbotapi.NewAudioUpload(chatID, "./audio.ogg")
						bot.Send(audio)
						for i := 0; i < 5; i++ {
							time.Sleep(time.Second * 30)
							if timer > 0 {
								break
							}
						}
					}
					if timer <= 0 {
						break
					}
				}
			} else {
				time.Sleep(time.Minute * 10)
			}
		}
	}()

	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID = update.Message.Chat.ID
		fmt.Printf("ChatId: %v, Message: %v\n", chatID, update.Message.Text)
		go func() {
			tr := Router(update, bot)
			if tr == 9999999 {
				msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Потерпите поцаны. Осталось %v минут.", timer))
				bot.Send(msg)
			} else if tr == 744444444 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Заебали спамить черти! Я спать нахуй!")
				bot.Send(msg)
				Counter = 0
				time.Sleep(time.Minute * 5)
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Я проснулся, суки.  Вам пизда")
				bot.Send(msg)
			} else if tr == 1337 {
				bot.Send(menu)
			} else {
				timer += tr
			}
			fmt.Printf("Timer: %v\n", timer)
		}()
	}
}

func Router(update tgbotapi.Update, bot *tgbotapi.BotAPI) int {
	if Counter == 0 {
		SpamTimer = time.Now()
	}

	if Counter >= 15 {

	} else {
		if time.Now().Sub(SpamTimer) < time.Minute {
			Counter++
			fmt.Printf("Counter: %v\n", Counter)
		} else {
			Counter = 0
		}
	}

	if MsgFlag {
		anStr := update.Message.Text
		if len(anStr) > 4 {
			anStr = anStr[:4]
		}

		an, _ := strconv.Atoi(anStr)

		if an == 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Чё за хуйню написал этот кретин?")
			bot.Send(msg)
			MsgFlag = false
			return 0
		}
		if an > 120 || an < -120 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы чё ебанулись? Я таких чисел то не знаю. Идите нахуй")
			bot.Send(msg)
			MsgFlag = false
			return 0
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Ок. До следующего перекура %v минут", timer+an))
		bot.Send(msg)

		stiker := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAAEItANkRjpXuNspodI9Z5drh-WTdc46tAACeyUAAp7OCwABmYOHg-RjJsQvBA")
		bot.Send(stiker)
		MsgFlag = false
		return an
	}

	g, _ := regexp.MatchString("бот|Бот", update.Message.Text)

	if g {
		o, _ := regexp.MatchString("сучара|Лев|даун|хуе|пидр|лох|чмо|тварь|пидор|хуй|пидорас|гей|тупой|идиот|мразь|дебил|Пидр|Лох|Чмо|Тварь|Пидор|Хуй|Пидорас|Гей|Тупой|Идиот|Мразь|Дебил", update.Message.Text)
		if o {
			switch agrCount {
			case 0:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ты чё пёс? Ебло закрой, воняет")
				bot.Send(msg)
				agrCount++
				return 0
			case 1:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Опять завоняло. Тебя же просили не отрывать свою гнилую пасть")
				bot.Send(msg)
				agrCount++
				return 0
			case 2:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Как ты заебал. Хули толку с таким тупорылым общаться. Я в ахуе")
				bot.Send(msg)
				agrCount++
				return 0
			case 3:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ещё хоть что-то пикнешь, и я солью твои гей фото во флудилку. Все увидят какой маленький орган у тебя, и какой огромный в тебе")
				bot.Send(msg)
				agrCount = 0
				return 0
			}
		}
	}

	switch update.Message.Text {
	case "Передвинь таймер":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "На сколько минут?")
		bot.Send(msg)
		MsgFlag = true
		return 0
	case "Да":
		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "./pizda.jpeg")
		bot.Send(msg)
		return 0
	case "Занято":
		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "./blat.jpg")
		bot.Send(msg)
		return 10
	case "Рано":
		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "./pon.jpg")
		bot.Send(msg)
		return 5
	case "Мы уже":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Блять. И нахуй я тогда нужен?")
		bot.Send(msg)
		return 35
	case "Иди нахуй":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сам иди нахуй. Сегодня больше не куришь")
		bot.Send(msg)
		return 0
	case "Иди на хуй":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сам иди нахуй. Сегодня больше не куришь")
		bot.Send(msg)
		return 0
	case "Выходной":
		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "./soran.jpg")
		bot.Send(msg)
		ret := 24 - time.Now().Hour()
		return ret * 60
	case "Слава Украине":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Осуждаю. (Донос отправлен товарищу майору)")
		bot.Send(msg)
		return 0
	case "Пизда нам":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Понял")
		bot.Send(msg)
		pic := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "./Hehe.jpg")
		bot.Send(pic)
		time.Sleep(time.Second * 20)
		for i := 0; i < 50; i++ {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Убить спамеров")
			bot.Send(msg)
			time.Sleep(time.Second)
		}
		time.Sleep(time.Minute)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Думали всё?")
		bot.Send(msg)
		pic = tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "./Kill.jpeg")
		bot.Send(pic)
		for i := 0; i < 50; i++ {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Убивать")
			bot.Send(msg)
			time.Sleep(time.Second)
		}
		audio := tgbotapi.NewAudioUpload(update.Message.Chat.ID, "./scream.ogg")
		bot.Send(audio)
		return 0

	case "Потеря бойца":
		n := rand.Intn(100)
		if n < 33 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Подождите его пять минут и нахуй")
			bot.Send(msg)
			return 5
		} else if n > 66 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ждём его до конца")
			bot.Send(msg)
			return 60
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Лол. Чё он не мог раньше закончить. Пацаны курить хотят так-то.")
			bot.Send(msg)
			return 0
		}
	case "Сколько":
		return 9999999
	case "Меню":
		return 1337
	case "Пидор":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "А может это ты пидор?")
		bot.Send(msg)
		return 0
	case "Нет":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пидора ответ")
		bot.Send(msg)
		return 0
	case "Шлюхи аргумент":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Аргумент не нужен. Пидор обнаружен")
		bot.Send(msg)
		return 0
	case "300":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отсоси у тракториста")
		bot.Send(msg)
		return 0
	case "Триста":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отсоси у тракториста")
		bot.Send(msg)
		return 0
	case "/help":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Как я тебе помогу? У меня рук нет!")
		bot.Send(msg)
		return 0
	case "/start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я сказал стартуем!")
		bot.Send(msg)
		return 0
	case "/stop":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ха-ха! Хорошая попытка, клоун")
		bot.Send(msg)
		return 0
	case "/run":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Why are u running? WHY ARE U RUNNING?!")
		bot.Send(msg)
		return 0
	default:
		return 0
	}
}

func Bye(chatId int64, bot *tgbotapi.BotAPI) {
	for {
		var end string
		fmt.Scan(&end)

		fmt.Printf("Command: %v\n", end)

		if end == "end" {
			msg := tgbotapi.NewMessage(chatId, "Всё, всем пока. Меня в дурку увозят. Ха-ха!")
			bot.Send(msg)
			os.Exit(0)
		} else if end == "go" {
			msg := tgbotapi.NewMessage(chatId, "Внезапный перекур. Кто не идёт - тот куколд из кружка кожевников!")
			bot.Send(msg)
		} else if end == "tak" {
			msg := tgbotapi.NewMessage(chatId, "@Kavantek @daniayar01 @Kobalt_KSPA и @Lev - чмоня, которого упоминать нельзя, Вы чё суки ещё здесь? Я СКАЗАЛ СЪЕБАЛИСЬ НАХУЙ!!!")
			bot.Send(msg)
		} else if end == "suk" {
			msg := tgbotapi.NewMessage(chatId, "@Kavantek @daniayar01 @Kobalt_KSPA и @Lev - чмоня, которого упоминать нельзя, если через 10 секунд вы всё ещё будете здесь, я скину Вере вашу историю браузеров. БЕГОМ ОТ СЮДА БЛЯТЬ!!!")
			bot.Send(msg)
		} else if end == "timer" {
			sc := bufio.NewScanner(os.Stdin)
			sc.Scan()

			var ti string
			ti = sc.Text()

			timer, _ = strconv.Atoi(ti)

			fmt.Printf("Timer: %v\n", timer)
		} else if end == "speek" {
			in := bufio.NewReader(os.Stdin)
			mess, err := in.ReadString('\n')
			fmt.Printf("String to speek: %v\n", mess)
			if err != nil {
				fmt.Println("Ошибка ввода: ", err)
			}

			msg := tgbotapi.NewMessage(chatId, mess)
			bot.Send(msg)
		}
	}
}
