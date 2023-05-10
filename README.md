# HOYOLab

## 簡介
- 用於 Genshin Impact 以及 Honkai: Star Rail 的腳本

## 功能
- [x] 每日簽到
- [ ] 兌換碼

## 使用方法
- 在根目錄建立一個 `persons.json` 檔案，內容如下：
    ```json
    [
        {
            "name": "",
            "ltuid": "",
            "ltoken": ""
        }
    ]
    ```
- `name` 用於區別多個玩家的帳號
- `ltuid`, `ltoken` 請前往簽到網站，登入後開啟開發者工具，到控制台輸入以下指令，來取得 cookie。
    ```javascript
    document.cookie
    ```