package main

import (
	file "L2/Development/dev3/file"
	dev3 "L2/Development/dev3/types"
	"flag"
	"log"
	"os"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func ParseFlagOptions() dev3.SortOptions {
	column := flag.Int("k", 1, "number of word to sort by")
	numeric := flag.Bool("n", false, "numeric sort")
	reverse := flag.Bool("r", false, "reverse sort")
	omitDuplicates := flag.Bool("u", false, "omit duplicates")
	month := flag.Bool("M", false, "month sort")
	trim := flag.Bool("b", false, "trim left spaces")
	checkSorted := flag.Bool("c", false, "check if sorted")
	suffixNumeric := flag.Bool("h", false, "suffixes supported numeric sort")

	flag.Parse()

	options := dev3.SortOptions{
		Column:         *column,
		Numeric:        *numeric,
		Reverse:        *reverse,
		OmitDuplicates: *omitDuplicates,
		Month:          *month,
		Trim:           *trim,
		CheckSorted:    *checkSorted,
		SuffixNumeric:  *suffixNumeric,
	}

	return options
}

func main() {
	from := "Development/dev3/input"

	//	Init logger
	errLogger := log.New(os.Stderr, "sort: ", log.Ldate)

	//	Init sort options
	options := ParseFlagOptions()
	if options.IsIncompatible() {
		errLogger.Printf("options \"%s\" incompatible", options.String())
		return
	}

	//	Get data from file
	raw, err := file.Lines(from)
	if err != nil {
		errLogger.Fatalf(err.Error())
	}
	for _, s := range raw {
		log.Println(s)
	}

	//	Sort extracted data
	sorter := dev3.StraightSorter{
		Options:   options,
		ErrLogger: errLogger,
	}
	sorter.InitComparer()

	sorted, _ := sorter.Sort(raw)

	//	Write it to the new file
	log.Println("____SORTED____")
	for _, s := range sorted {
		log.Println(s)
	}
}
