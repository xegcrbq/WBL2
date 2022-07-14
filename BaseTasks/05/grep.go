package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной
утилитой (man grep — смотрим описание и основные параметры).
Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/

type flags struct {
	A, B, c    int
	i, v, F, n bool
	p, regular string
}

type founded struct {
	lineNumber int
	value      string
}

func main() {
	//пример аргументов
	//-n -C 1 -c 7 -v -p test.txt со
	A := flag.Int("A", 0, "печатать +N строк после совпадения")
	B := flag.Int("B", 0, "печатать +N строк до совпадения")
	C := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	c := flag.Int("c", -1, "(количество строк)")

	i := flag.Bool("i", false, "(игнорировать регистр)")
	v := flag.Bool("v", false, "(вместо совпадения, исключать)")
	F := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "напечатать номер строки")
	p := flag.String("p", "", "путь до файла с входными данными")
	flag.Parse()
	flags := flags{int(math.Max(float64(*A), float64(*C))), int(math.Max(float64(*B), float64(*C))), *c, *i, *v, *F, *n, *p, os.Args[len(os.Args)-1]}
	var data []byte
	var err error
	if flags.p != "" { //чтение данных из файла
		data, err = ioutil.ReadFile(flags.p)
		if err != nil {
			log.Fatal(err)
		}
	} else { //чтение данных из консоли пока не встречен |
		data, err = bufio.NewReader(os.Stdin).ReadBytes('|')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
	if len(flags.regular) > 0 {
		flags.grep(data)
	}
}

func (f flags) grep(data []byte) {
	var result []founded
	if f.i { //игнорирование регистра
		data = bytes.ToLower(data)
		f.regular = strings.ToLower(f.regular)
	}
	text := strings.Split(string(data), "\r\n")

	for i, line := range text {
		if f.c != -1 { //проверка первых c строк
			if f.c == i {
				break
			}
		}
		if f.F { //только точное совпадение
			if strings.EqualFold(line, f.regular) {
				result = append(result, founded{i, line})
			}
		} else {
			check, err := regexp.MatchString(f.regular, line)
			if err != nil {
				log.Fatal(err)
			}
			if f.v {
				if !check { //исключение вместо совпадения
					result = append(result, founded{i, line})
				}
			} else {
				if check { //совпадение
					result = append(result, founded{i, line})
				}
			}
		}
	}

	for _, val := range result {
		//вывод соседних строк
		for i := int(math.Max(0, float64(val.lineNumber-f.B))); i <= int(math.Min(float64(len(text))-1, float64(val.lineNumber+f.A))); i++ {
			if f.n { //вывод номера строки
				fmt.Printf("%v %s\n", i, text[i])
			} else { //вывод только строки
				fmt.Println(text[i])
			}
		}
		fmt.Println()
	}
}
