package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по
разделителю (TAB) на колонки и выводить запрошенные.
Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

*/

func main() {
	f := flag.String("f", "", "Выбрать поля (колонки). Перечислить значения через запятую")
	d := flag.String("d", "\t", "Использовать другой разделитель")
	s := flag.Bool("s", false, "Выводить только строки c разделителем")

	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите пустую строку, чтобы завершить ввод:\n")
	var data [][]string
	for {
		ok := scanner.Scan()
		if !ok || scanner.Err() != nil {
			break
		}
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if *s && !strings.Contains(line, *d) {
			continue
		}
		data = append(data, strings.Split(line, *d))
	}
	if *f != "" {
		columns := strings.Split(*f, ",")
		selectedColumns := make(map[int]bool)
		for _, number := range columns {
			num, err := strconv.Atoi(number)
			if err != nil {
				log.Println(err)
				continue
			}
			selectedColumns[num] = true
		}
		for _, line := range data {
			var dataForPrint []string
			for j, field := range line {
				if selectedColumns[j] {
					dataForPrint = append(dataForPrint, field)
				}
			}
			fmt.Println(strings.Join(dataForPrint, *d) + "\n")
		}

	} else {
		for _, line := range data {
			fmt.Println(strings.Join(line, *d) + "\n")
		}
	}
}
