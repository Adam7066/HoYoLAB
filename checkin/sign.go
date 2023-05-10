package checkin

import (
	"io"
	"net/http"

	"github.com/Adam7066/golang/log"
)

const lang = "zh-tw"
const referer = "https://act.hoyolab.com/"

func sign(endpoint, actID, cookie string) string {
	signUrl := endpoint + "/sign?lang=" + lang + "&act_id=" + actID

	req, err := http.NewRequest("POST", signUrl, nil)
	if err != nil {
		log.Error.Println(err)
		return ""
	}
	req.Header.Set("Referer", referer)
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Cookie", cookie)
	clt := http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		log.Error.Println(err)
		return ""
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error.Println(err)
		return ""
	}
	respBody := string(content)
	return respBody
}
