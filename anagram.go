package main

import (
        "bufio"
        "fmt"
        "os"
		"strings"
)

// テキストファイルの中身を1行ずつ読み込んで配列化して返す
func readLines(path string) []string{
    f, err := os.Open(path)
	var textArr = []string{}
    if err != nil {
        fmt.Println(err)
        return textArr
    }go
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

// 1つの文字列を並び替える
func sortOneLine(oneLine string) {
	var oneArr = []string{}
	oneArr = strings.Split(oneLine, "")
	fmt.Println(oneArr)
	//return oneArr
}

// 文字列の順列を返す
func perm(word string) []string{
	retWord := []string{}
	//文字数が1文字の場合はそのまま返す
	if len(word) == 1 {
		return []string{word}
	}
	for i:=0; i < len(word); i++ {
		// foreach (string t in Perms(s.Remove(i, 1))) {
		// i文字目を削除した文字列を引数に再帰
		substr := perm(word[0:i] + word[i+1:len(word)])
		// 1要素ずつ処理する場合はこういう書き方らしい…
		for _, c := range(substr) {
			retWord = append(retWord, word[i:i+1] + c)
		}
	}
	return retWord
}

// テキストの中にretwordが存在する場合は戻りようの配列に格納する
func searchAnagram (textArr []string){
	for {
		// アナグラムリストの生成
		firstLine := textArr[0]
		textArr = removeElement(textArr, 0)
		oneAnagram := perm(firstLine)

		// 結果用配列
		result := []string{firstLine}
		for _, anagramElement := range oneAnagram {
			for i, textArrElement := range textArr {
				if anagramElement == textArrElement {
					result = append(result, textArrElement)
					textArr = removeElement(textArr, i)
				}
			}
		}
		
		if len(result) >= 2 {
			fmt.Println(result)
		}
		if len(textArr) == 0 {
			break
		}
	}
}

func removeElement(trimArr []string, index int) []string{
	trimedArr := []string{}
	i := 0
	for _, e := range trimArr {
		if i != index {
			trimedArr = append(trimedArr, e)
		}
		i += 1
	}
	return trimedArr
}

func main() {
	// テキストファイルを1行ずつ読み込む
	textArr := readLines("text.txt")
	searchAnagram(textArr)

}