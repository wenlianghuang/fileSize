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
		if info.IsDir() && (info.Name() == "$RECYCLE.BIN" || info.Name() == "Config.Msi" || info.Name() == "Recovery" || info.Name() == "System Volume Information") {
			return filepath.SkipDir
		}
		if info.IsDir() {
			gitPath := filepath.Join(path, ".git")
			_, err := os.Stat(gitPath)
			if err == nil {
				fmt.Println(path)
				// 在這裡可以對已經被Git初始化的資料夾進行處理
			} else if os.IsPermission(err) {
				// 如果是權限不足的錯誤，就試著修改該資料夾的權限
				err = os.Chmod(path, 0777)
				if err != nil {
					fmt.Printf("Failed to change permission for %s: %v\n", path, err)
				}
				// 然後跳過該資料夾，繼續尋找下一個資料夾
				return filepath.SkipDir
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
