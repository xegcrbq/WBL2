package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

/*
Отсортировать строки в файле по аналогии с консольной
утилитой sort (man sort — смотрим описание и основные
параметры): на входе подается файл из несортированными
строками, на выходе — файл с отсортированными.
Реализовать поддержку утилитой следующих ключей:
-k — указание колонки для сортировки (слова в строке могут
выступать в качестве колонок, по умолчанию разделитель —
пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
*/
type flags struct {
	k       int
	n, r, u bool
}

func main() {
	//для создания файла testS.txt использовались аргументы
	// -u -k 1 -n test.txt testS.txt
	k := flag.Int("k", -1, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	flags := flags{*k, *n, *r, *u}
	inputFilePath := os.Args[len(os.Args)-2]
	outputFilePath := os.Args[len(os.Args)-1]
	data, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	data = flags.sort(data)
	ioutil.WriteFile(outputFilePath, data, os.FileMode(0777))
}

func (f flags) sort(data []byte) (result []byte) {
	rows := strings.Split(string(data), "\r\n")

	if f.u {
		rows = makeUnique(rows)
	}

	if f.k > -1 {
		var fun func(i, j int) bool
		words := make([][]string, len(rows))
		for i, _ := range rows {
			words[i] = strings.Split(rows[i], " ")
		}
		switch true {
		case f.r && f.n:
			fun = func(i, j int) bool {
				return words[j][f.k] < words[i][f.k]
			}
		case f.n:
			fun = func(i, j int) bool {
				return words[j][f.k] > words[i][f.k]
			}
		case f.r:
			fun = func(i, j int) bool {
				return len(words[j][f.k]) > len(words[i][f.k])
			}
		default:
			fun = func(i, j int) bool {
				return len(words[j][f.k]) < len(words[i][f.k])
			}
		}
		sort.SliceStable(words, fun)
		for i := range words {
			rows[i] = strings.Join(words[i], " ")
		}
	} else {
		var fun func(i, j int) bool
		switch true {
		case f.r && f.n:
			fun = func(i, j int) bool {
				return rows[j] < rows[i]
			}
		case f.n:
			fun = func(i, j int) bool {
				return rows[j] > rows[i]
			}
		case f.r:
			fun = func(i, j int) bool {
				return len(rows[j]) > len(rows[i])
			}
		default:
			fun = func(i, j int) bool {
				return len(rows[j]) < len(rows[i])
			}
		}
		sort.SliceStable(rows, fun)
	}

	result = []byte(strings.Join(rows, "\n"))
	return
}

func makeUnique(s []string) (result []string) {
	tempMap := map[string]struct{}{}
	for _, r := range s {
		tempMap[r] = struct{}{}
	}
	result = make([]string, 0, len(tempMap))
	for k := range tempMap {
		result = append(result, k)
	}
	return
}
