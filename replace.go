package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 引数の存在チェック
	if len(os.Args) < 2 {
		fmt.Println("please specify target file path")
		return
	}

	// ファイルを開き、内容を変数に読み込む
	fileName := os.Args[1]
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("could not open the file", fileName)
		return
	}
	contents := string(bytes)

	// 置換する（上の方が優先順位高い）
	contents = strings.ReplaceAll(contents, "素晴らしい", "不思議だ")
	contents = strings.ReplaceAll(contents, "GO言語", "COME言語")
	contents = strings.ReplaceAll(contents, "GO", "頑張れ")

	// ファイルに書き戻す
	err = ioutil.WriteFile(fileName, []byte(contents), os.ModePerm)
	if err != nil {
		fmt.Println("could not write to the file", fileName)
		return
	}

	fmt.Println("done")
}
