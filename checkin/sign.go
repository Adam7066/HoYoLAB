package checkin

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dghubble/sling"
)

const lang = "zh-tw"
const referer = "https://act.hoyolab.com/"

func sign(endpoint, actID, cookie string, extraHeader map[string]string) string {
	signUrl := endpoint + "/sign?lang=" + lang + "&act_id=" + actID

	s := sling.New().Post(signUrl).
		Set("Referer", referer).
		Set("Accept-Encoding", "gzip, deflate, br").
		Set("Cookie", cookie)
	for k, v := range extraHeader {
		s = s.Set(k, v)
	}

	req, err := s.Request()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	respBody := string(content)
	return respBody
}
