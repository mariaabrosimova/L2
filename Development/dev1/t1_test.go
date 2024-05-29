// package dev1

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func TestShowTime(t *testing.T) {
// 	currentTime, err := CurrentTime()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	time.Sleep(time.Second * 1)
// 	fmt.Printf("%d:%02d\n", currentTime.Hour(), currentTime.Minute())
// }

package dev1

import (
	"fmt"
	"testing"
	"time"
)

func TestGetNTPTime(t *testing.T) {
	cur_time, err := GetNTPTime()
	if err != nil {
		t.Fatalf("Ошибка при вызове GetNTPTime: %v", err)
		fmt.Printf("%d:%02d\n", cur_time.Hour(), cur_time.Minute())
	}

	if cur_time.IsZero() {
		t.Errorf("Ожидалось непустое время, но получено пустое")
	}
	time.Sleep(time.Second * 1)
	if cur_time.After(time.Now()) {
		t.Errorf("Полученное время (%v) находится в будущем", cur_time)
	} else {
		fmt.Printf("%d:%02d\n", cur_time.Hour(), cur_time.Minute())
	}
}
