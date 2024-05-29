package main

import (
	grep2 "L2/Development/dev5/grep"
	"errors"
	"flag"
	"log"
	"os"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func ParseFlags() grep2.Options {
	A := flag.Int("A", 0, "print +N strings after match")
	B := flag.Int("B", 0, "print +N strings before match")
	C := flag.Int("C", 0, "print +N strings before and after match")
	c := flag.Bool("c", false, "count match strings")
	i := flag.Bool("i", false, "ignore register")
	v := flag.Bool("v", false, "ignore matches")
	F := flag.Bool("F", false, "exact match")
	n := flag.Bool("n", true, "print line num")

	flag.Parse()

	options := grep2.Options{
		AfterPrint:  *C,
		BeforePrint: *C,
		Count:       *c,
		IgnoreCase:  *i,
		Invert:      *v,
		Fixed:       *F,
		LineNum:     *n,
		//Target:      flag.Arg(0),
		//FileName:    flag.Arg(1),
		FileName: "Development/dev5/test",
	}

	if *A > 0 {
		options.AfterPrint = *A
	}

	if *B > 0 {
		options.BeforePrint = *B
	}

	return options
}

func main() {

	elg := log.New(os.Stderr, "", log.Llongfile)

	options := ParseFlags()
	counter := 0

	scanner := grep2.NewScanner(options, ChooseMatcher(options), ChoosePrinter(options, &counter))

	file, err := os.Open(options.FileName)
	if err != nil {
		elg.Printf("grep: error: %s", errors.New("can't open file"))
	}
	defer file.Close()

	scanner.Scan(file)

	PrintCounter(options, counter)

}

func ChooseMatcher(options grep2.Options) grep2.Matcher {
	if options.Fixed {
		return grep2.Exact{}
	}

	return grep2.Have{}
}

func ChoosePrinter(options grep2.Options, counter *int) func(line int, s string) {
	if options.Count {
		return MakeCounter(counter)
	}

	if options.LineNum {
		return func(line int, s string) {
			log.Printf("line %d: %s", line, s)
		}
	}

	return func(line int, s string) {
		log.Println(s)
	}
}

func MakeCounter(counter *int) func(line int, s string) {
	return func(line int, s string) {
		*counter += 1
	}
}

func PrintCounter(options grep2.Options, counter int) {
	if options.Count {
		log.Println(counter)
	}
}
