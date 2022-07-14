package main

//Создать программу печатающую точное время с использованием
//NTP -библиотеки. Инициализировать как go module. Использовать
//библиотеку github.com/beevik/ntp. Написать программу
//печатающую текущее время / точное время с использованием этой
//библиотеки.
import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func main() {
	time := GetCurrentTime()
	fmt.Println(time)
}

func GetCurrentTime() time.Time {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Println(err)
	}
	return time
}
