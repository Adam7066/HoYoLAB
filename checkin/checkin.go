package checkin

import (
	"fmt"
	
	"hoyolab/db"
)

func CheckinUser(user db.User) {
	cookie := "ltmid_v2=" + user.LtmidV2 + "; ltuid_v2=" + user.LtuidV2 + "; ltoken_v2=" + user.LtokenV2
	if user.Genshin {
		genshinActID := "e202102251931481"
		genshinEndpoint := "https://sg-hk4e-api.hoyolab.com/event/sol"
		signJson := sign(genshinEndpoint, genshinActID, cookie, nil)
		fmt.Println("[Genshin]", user.Name, signJson)
	}
	if user.Starrail {
		starrailActID := "e202303301540311"
		starrailEndpoint := "https://sg-public-api.hoyolab.com/event/luna/os"
		signJson := sign(starrailEndpoint, starrailActID, cookie, nil)
		fmt.Println("[Starrail]", user.Name, signJson)
	}
	if user.ZZZ {
		zzzActID := "e202406031448091"
		zzzEndpoint := "https://sg-public-api.hoyolab.com/event/luna/zzz/os"
		signJson := sign(
			zzzEndpoint,
			zzzActID,
			cookie,
			map[string]string{"x-rpc-signgame": "zzz"},
		)
		fmt.Println("[ZZZ]", user.Name, signJson)
	}
}
