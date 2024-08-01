package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// CSVファイルが保存されているディレクトリ
	dir := "./kadai1/csv_files"

	// レコードがあるテーブル名を格納するスライス
	var tablesWithRecords []string

	// ディレクトリ内のCSVファイルを読み込む
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// CSVファイルのみを対象とする
		if !info.IsDir() && filepath.Ext(path) == ".csv" {
			// ファイルをオープン
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// CSVリーダーを作成
			reader := csv.NewReader(file)

			// 全てのレコードを読み込む
			records, err := reader.ReadAll()
			if err != nil {
				return err
			}

			// レコードがカラム名のみでないかをチェック
			if len(records) > 1 {
				// ファイル名からテーブル名を取得（拡張子を除く）
				tableName := filepath.Base(path)
				tableName = tableName[:len(tableName)-len(filepath.Ext(tableName))]
				tablesWithRecords = append(tablesWithRecords, tableName)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	// レコードが入ったテーブル名を表示
	fmt.Println("レコードがあるテーブル名:")
	for _, table := range tablesWithRecords {
		fmt.Println(table)
	}
}
