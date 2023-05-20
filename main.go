package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//docker run -v /d/:/d/ <docker-image>
func main() {

	//root := "D:\\"
	root := "/d/"
	gitRepos := []string{}
	gitNotRepos := []string{}
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	// To check the mother code in D:\\ that it is git initialized or not
	for _, file := range files {
		if file.IsDir() {
			path := root + file.Name()
			//fmt.Println(path)

			//if path == "D:\\$RECYCLE.BIN" || path == "D:\\Recovery" /*|| path == "D:\\System Volume Information"*/ {
			//if path == "/d/$RECYCLE.BIN" || path == "/d/Recovery" || path == "/d/System Volume Information" || path == "/d/xampp" {
			if path == "/d/$RECYCLE.BIN" {
				continue
			}

			_, err := os.Stat(path + "/.git")
			if err == nil {
				//fmt.Println(path + " is a Git repository")
				gitRepos = append(gitRepos, path)
			} else {
				//fmt.Println(path + " is not a Git repository")

				gitNotRepos = append(gitNotRepos, path)
			}
		}
	}

	// Recursive the mother folder to find all the folder which is not git initialized
	for _, norepo := range gitNotRepos {

		err := filepath.Walk(norepo, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				// If the directory permission is too high, show the error and pass this directory
				if os.IsPermission(err) {
					fmt.Printf("Permission denied: %v\n", err)
					return filepath.SkipDir
				}
				return err
			}
			if info.IsDir() {
				if isGitRepo(path) {
					//fmt.Println(path + " is a Git repository")
					gitRepos = append(gitRepos, path)
					return filepath.SkipDir
				} else {
					//fmt.Println(path + " is not a Git repository")
					gitNotRepos = append(gitNotRepos, path)
				}
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	//Print and "git status" all the folder which is git initialized
	for _, repo := range gitRepos {
		err := os.Chdir(repo)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(repo + " is Git repo")

		//cmd := exec.Command("git", "status")
		//output, err := cmd.Output()
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//fmt.Println(string(output))
	}
}

func isGitRepo(path string) bool {
	_, err := os.Stat(path + "/.git")
	if err == nil {
		return true
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return false
	}
	for _, file := range files {
		if file.IsDir() {
			if isGitRepo(path + "\\" + file.Name()) {
				return true
			}
		}
	}
	return false
}
