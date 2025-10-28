# HoYoLAB

## 簡介
- 用於 HoYoLAB 的腳本，目前支援
    - 原神 Genshin Impact
    - 崩壞星穹鐵道 Honkai: Star Rail
    - 絕區零 Zenless Zone Zero

## 功能
- [x] 每日簽到
- [ ] 兌換碼

## 使用方法
1. 互動式新增帳號：`go run main.go --add`
    - 依照提示輸入 Name、LtmidV2、LtuidV2、LtokenV2 及遊戲選擇。
2. 執行每日簽到：`go run main.go`
    - 已簽到過的帳號當天不會重複簽到。

> [!TIP]
> 可以搭配使用 crontab 定時執行每日簽到。

### Cookie 取得方式
- 登入 HoYoLAB 簽到網站，開啟瀏覽器開發者工具，於 Cookie 中：
- 取得 ltmid_v2、ltuid_v2、ltoken_v2 對應填入。
- Name 欄位僅用於辨識帳號。