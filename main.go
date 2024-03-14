package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func absPath(path string) (string, error) {
	// 現在の作業ディレクトリを取得
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 相対パスを絶対パスに変換
	absPath, err := filepath.Abs(filepath.Join(cwd, path))
	if err != nil {
		return "", err
	}

	return absPath, nil
}

func main() {
	// コマンドライン引数を取得
	args := os.Args

	// 引数が1つ以上の場合は、最初の引数をパスとみなす
	if len(args) > 1 {
		path := args[1]

		absPath, err := absPath(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("絶対パス:", absPath)
	} else {
		fmt.Println("有効な引数ではありません")
	}
}
