package checkin

import (
	"fmt"
	"hoyolab/person"
	"sync"
)

func checkin(ltuid, ltoken string, wg *sync.WaitGroup) {
	defer wg.Done()

	genshinActID := "e202102251931481"
	genshinEndpoint := "https://sg-hk4e-api.hoyolab.com/event/sol"

	starrailActID := "e202303301540311"
	starrailEndpoint := "https://sg-public-api.hoyolab.com/event/luna/os"

	zzzActID := "e202406031448091"
	zzzEndpoint := "https://sg-public-api.hoyolab.com/event/luna/zzz/os"

	cookie := "ltuid=" + ltuid + "; ltoken=" + ltoken

	signJson := sign(genshinEndpoint, genshinActID, cookie, nil)
	fmt.Println(signJson)
	signJson = sign(starrailEndpoint, starrailActID, cookie, nil)
	fmt.Println(signJson)
	signJson = sign(
		zzzEndpoint,
		zzzActID,
		cookie,
		map[string]string{
			"x-rpc-signgame": "zzz",
		},
	)
	fmt.Println(signJson)
}

func Checkin(persons person.Persons) {
	var wg sync.WaitGroup
	for _, p := range persons {
		wg.Add(1)
		go checkin(p.Ltuid, p.Ltoken, &wg)
	}
	wg.Wait()
}
