// package dev1

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/beevik/ntp"
// )

// func Main_1() {

// 	exactTime, err := ntp.Time("time.google.com")
// 	if err != nil {
// 		log.Printf("Ошибка получения точного времени: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("Точное время: %s\n", exactTime.Format(time.RFC1123))
// }

// package dev1

// import (
// 	"time"

// 	"github.com/beevik/ntp"
// )

// /*
// === Базовая задача ===

// Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
// Использовать библиотеку https://github.com/beevik/ntp.
// Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

// Программа должна быть оформлена с использованием как go module.
// Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
// Программа должна проходить проверки go vet и golint.
// */

// func CurrentTime() (time.Time, error) {
// 	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
// 	if err != nil {
// 		return time.Time{}, err
// 	}
// 	time := time.Now().Add(response.ClockOffset)
// 	return time, err
// }

package dev1

import (
	"time"

	"github.com/beevik/ntp"
)

// GetNTPTime возвращает текущее точное время с NTP-сервера
func GetNTPTime() (time.Time, error) {
	return ntp.Time("time.google.com")
}
