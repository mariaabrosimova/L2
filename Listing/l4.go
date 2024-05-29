package main

func main() {
	ch := make(chan int)
	// анонимная горутина отправит только 0-9 значения
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	// ch не будет закрыта и будет жадть значения - deadlock
	for n := range ch {
		println(n)
	}
}
