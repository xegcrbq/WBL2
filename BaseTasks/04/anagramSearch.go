package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Требования:
1. Входные данные для функции: ссылка на массив, каждый
элемент которого - слово на русском языке в кодировке
utf8
2. Выходные данные: ссылка на мапу множеств анаграмм
3. Ключ - первое встретившееся в словаре слово из
множества. Значение - ссылка на массив, каждый элемент
которого,
слово из множества.
4. Массив должен быть отсортирован по возрастанию.
5. Множества из одного элемента не должны попасть в
результат.
6. Все слова должны быть приведены к нижнему регистру.
7. В результате каждое слово должно встречаться только один
раз.
*/

func main() {
	fmt.Println(groupAnagram([]string{"сон", "нос", "пятак", "пятка", "тяпка", "листок", "слиток", "столик", "тест", "Слиток"}))
}
func groupAnagram(words []string) (result map[string][]string) {
	for i := range words { //Все слова должны быть приведены к нижнему регистру.
		words[i] = strings.ToLower(words[i])
	}
	result = make(map[string][]string)
	for _, word := range words {
		wordTemp := []rune(word)
		notFound := true
		sort.SliceStable(wordTemp, func(i, j int) bool {
			return wordTemp[i] < wordTemp[j]
		})
		for key := range result {
			keyTemp := []rune(key)
			sort.SliceStable(keyTemp, func(i, j int) bool {
				return keyTemp[i] < keyTemp[j]
			})
			if strings.EqualFold(string(keyTemp), string(wordTemp)) {
				result[key] = append(result[key], word)
				notFound = false
				break
			}
		}
		if notFound {
			result[word] = []string{word}
		}
	}
	for key := range result { //В результате каждое слово должно встречаться только один раз.
		result[key] = RemoveDuplicate(result[key])
	}

	for key := range result { //Множества из одного элемента не должны попасть в результат.
		if len(result[key]) == 1 {
			delete(result, key)
		}
	}

	return
}

func RemoveDuplicate(array []string) []string {
	m := make(map[string]string)
	for _, x := range array {
		m[x] = x
	}
	var ClearedArr []string
	for x, _ := range m {
		ClearedArr = append(ClearedArr, x)
	}
	return ClearedArr
}
