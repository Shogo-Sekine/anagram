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

// テキストファイルを読み込んでマップを作る
func makeMap() {
	textArr := readLines("text.txt")
	placeMap := make(map[string]string)
	for i:= 0; i < len(textArr); i++ {
		joinedArr := sortOneLine(textArr[i])
		placeMap[textArr[i]] = joinedArr
	}
	sortAndSearchAnagram(placeMap)
}

// マップのvalueを辞書順に並び替え、アナグラムが存在するindexのみを出力する
func sortAndSearchAnagram(placeMap map[string]string) {
	// マップを辞書順に並び替える
	sortedList := List{}
	for i, v := range placeMap {
		sm := sortedMap{i, v}
		sortedList = append(sortedList, sm)
	}
	sort.Sort(sortedList)

	// アナグラムが存在するindexのみを出力する
	endFlg := false
	result := []string{}
	for{
		result = []string{}
		tempValue := sortedList[0].value
		for i, _ := range sortedList {
			if len(result) == 0 {
				tempValue = sortedList[i].value
				result = append(result, sortedList[i].index)
			}
			// iがリストの最後まで来たら終了フラグを立ててループを抜ける
			i++
			if i == (len(sortedList) - 1) {
				endFlg = true
				break
			}
			// valueの値が異なるまで処理を続ける
			if tempValue != sortedList[i].value {
				// アナグラムが存在するときのみ出力
				if len(result) > 1 {
					fmt.Println(result)
					result = []string{}
					continue
				}
				result = []string{}
				continue
			}
			result = append(result, sortedList[i].index)
		}

		if endFlg == true {
			break
		}
	}
}

func main() {
	start := time.Now();
	makeMap()
	end := time.Now();
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())	
}