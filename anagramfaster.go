package main

import (
        "bufio"
        "fmt"
        "os"
		"strings"
		"sort"
		"time"
)

type sortedMap struct {
	index string
	value string
}

type List []sortedMap

func (l List) Len() int {
	return len(l)
}

func (l List) Less(i, j int) bool {
	    if l[i].value == l[j].value {
        return (l[i].index < l[j].index)
    } else {
        return (l[i].value < l[j].value)
    }
}

func (l List) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}

// テキストファイルの中身を1行ずつ読み込んで配列化して返す
func readLines(path string) []string{
    f, err := os.Open(path)
	var textArr = []string{}
    if err != nil {
        fmt.Println(err)
        return textArr
    }
    s := bufio.NewScanner(f)
    for s.Scan() {
		textArr = append(textArr, s.Text())
    }

    if s.Err() != nil {
        // non-EOF error.
        fmt.Println(s.Err())
    }
	return textArr
}

// 1つの文字列を分割→辞書順にソート→再結合
func sortOneLine(oneLine string) string{
	var oneArr = []string{}
	oneArr = strings.Split(oneLine, "")
	sort.Strings(oneArr)
	joinedArr := strings.Join(oneArr, "")
	return joinedArr
}

func makeMap() {
	textArr := readLines("text.txt")
	placeMap := make(map[string]string)
	for i:= 0; i < len(textArr); i++ {
		joinedArr := sortOneLine(textArr[i])
		placeMap[textArr[i]] = joinedArr
	}
	sortMap(placeMap)
	
}

func sortMap(placeMap map[string]string) {
	a := List{}
	for i, v := range placeMap {
		sm := sortedMap{i, v}
		a = append(a, sm)
	}

	start := time.Now();
	sort.Sort(a)
	end := time.Now();
	fmt.Printf("mapのソート所要時間：%f秒\n",(end.Sub(start)).Seconds())

	start = time.Now();	
	for {
		result := []string{a[0].index}
		tempValue := a[0].value
		a = removeElement(a, 0)

		for _, v2 := range a {
			if v2.value != tempValue {
				break
			}
			result = append(result, v2.index)
			a = removeElement(a, 0)
		}

		if len(result) > 1 {
			fmt.Println(result)
		}
		if len(a) == 0 {
			break
		}
	}
	end = time.Now();
	fmt.Printf("探索所要時間：%f秒\n",(end.Sub(start)).Seconds())
}

func removeElement(trimArr List, index int) List{
	trimedArr := List{}
	i := 0
	for _, v := range trimArr {
		if i != index {
			trimedArr = append(trimedArr, v)
		}
		i += 1
	}
	return trimedArr
}

func main() {
	makeMap()

}