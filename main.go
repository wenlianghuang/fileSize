package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "D:\\"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && (info.Name() == "$RECYCLE.BIN" || info.Name() == "Config.Msi" || info.Name() == "Recovery") {
			return filepath.SkipDir
		}
		if info.IsDir() {
			gitPath := filepath.Join(path, ".git")
			_, err := os.Stat(gitPath)
			if err == nil {
				fmt.Println(path)
				// 在這裡可以對已經被Git初始化的資料夾進行處理
			} else if !os.IsNotExist(err) {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
