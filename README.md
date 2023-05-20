# About Docker

其實從5/15晚上我就一直再嘗試可不可以在D driver裡面所有的folder他的**git init**情況，希望每個folder如果有git init就討論他**git status**,如果都可以的話就自動推上去，而且不只是自己的local computer，而是**任何的**電腦，只要利用docker，在不同的環境都能做搜尋的情形，以下是按每天的紀錄來看:

### 2023/05/15
- 本來想要找docker wsl data的size,但後來有問題就先暫停

### 2023/05/16
- 出現了Recovery, Config.Msi, System Volume Information這些權限較高的資料夾，如果碰到的話就中斷，沒辦法繼續往前
### 2023/05/17
- 延續昨天的問題，但還是沒辦法解決
### 2023/05/18
- 用```if os.IsPermission(err)```的方式來確認如果他的權限較高就直接跳過去看下一個folder(```filepath.SkipDir```)
- docker -v /path/to/<folder you want to use>:/<folder you want to use>其實就是把本機的<folder you want to use>掛載docker的<folder you want to use>的容器中，他就會好像放在虛擬機上執行，但其實是在執行被掛載的本機裡面的各種資料。
- 成功的話就先把searchgitinit 用docker tag 重新命名**wenlianghuang/searchgitinit**，接著把它image push 到 docker hub
- 刪除本機docker的image
- 用其他的電腦下載並執行看是否可以成功 ==> 成功，但花時間。
### 2023/05/19

