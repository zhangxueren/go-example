package asset

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DataItem struct {
	WxID   string `json:"wxId"`
	TeamID int    `json:"teamId"`
}

func ResetWxidTeamid() {
	data := []DataItem{
		{"ww7809deb62d31506e_wkwx_BJ1-202106170121", 2392},
		{"ww7809deb62d31506e_wkwx_BJ1-202106170131", 2361},
		{"ww7809deb62d31506e_wkwx_BJ1-202106170132", 2360},
		{"ww7809deb62d31506e_wkwx_BJ1-202106170136", 3120},
		{"ww7809deb62d31506e_wkwx_BJ1-202106170139", 2359},
		{"ww7809deb62d31506e_wkwx_BJ1-202106170558", 2358},
		{"ww7809deb62d31506e_wkwx_BJ1-202106170565", 2357},
		{"ww7809deb62d31506e_wkwx_BJ1-202106170568", 2355},
		{"ww7809deb62d31506e_wkwx_BJ1-202106170569", 2353},
		{"ww7809deb62d31506e_wkwx_BJ1-202106170588", 2394},
		{"ww7809deb62d31506e_wkwx_BJ1-202106171210", 2396},
		{"ww7809deb62d31506e_wkwx_BJ1-202106171216", 2924},
		{"ww7809deb62d31506e_wkwx_BJ1-202106172796", 2361},
		{"wwb5f29f00ef255b16_wkwx_qw210629001434", 2333},
		{"wwb5f29f00ef255b16_wkwx_qw210629001437", 2289},
		{"wwb5f29f00ef255b16_wkwx_qw210629001440", 2288},
		{"wwb5f29f00ef255b16_wkwx_qw210629001442", 2287},
		{"wwb5f29f00ef255b16_wkwx_qw210629001443", 2286},
		{"wwb5f29f00ef255b16_wkwx_qw210629001553", 2285},
		{"wwb5f29f00ef255b16_wkwx_qw210629001558", 2284},
		{"wwb5f29f00ef255b16_wkwx_qw210629001561", 2283},
		{"wwb5f29f00ef255b16_wkwx_qw210629001565", 2276},
		{"wwb5f29f00ef255b16_wkwx_qw210629001566", 2275},
		{"wwb5f29f00ef255b16_wkwx_qw210629001569", 2271},
		{"wwb5f29f00ef255b16_wkwx_qw210629001572", 2270},
		{"wwb5f29f00ef255b16_wkwx_qw210629001573", 2269},
		{"wwb5f29f00ef255b16_wkwx_qw210629001574", 2268},
		{"wwb5f29f00ef255b16_wkwx_qw210629001580", 2267},
		{"wwb5f29f00ef255b16_wkwx_qw210629001587", 2266},
		{"wwb5f29f00ef255b16_wkwx_qw210629001588", 2265},
		{"wwb5f29f00ef255b16_wkwx_qw210629001592", 2264},
		{"wwb5f29f00ef255b16_wkwx_qw210629001595", 2263},
		{"wwb5f29f00ef255b16_wkwx_qw210629001596", 2262},
		{"wwb5f29f00ef255b16_wkwx_qw210629001597", 2261},
		{"wwb5f29f00ef255b16_wkwx_qw210629001600", 2260},
		{"wwb5f29f00ef255b16_wkwx_qw210629001604", 2251},
		{"wwb5f29f00ef255b16_wkwx_qw210629001606", 2250},
		{"wwb5f29f00ef255b16_wkwx_qw210629001608", 2249},
		{"wwb5f29f00ef255b16_wkwx_qw210629001610", 2248},
		{"wwb5f29f00ef255b16_wkwx_qw210629001613", 2247},
		{"wwb5f29f00ef255b16_wkwx_qw210629001617", 2246},
		{"wwb5f29f00ef255b16_wkwx_qw210629001619", 2245},
		{"wwb5f29f00ef255b16_wkwx_qw210629001620", 2244},
		{"wwb5f29f00ef255b16_wkwx_qw210629001621", 2243},
		{"wwb5f29f00ef255b16_wkwx_qw210629001623", 2242},
		{"wwb5f29f00ef255b16_wkwx_qw210629001627", 2241},
		{"wwb5f29f00ef255b16_wkwx_qw210629001631", 2240},
		{"wwb5f29f00ef255b16_wkwx_qw210629001635", 2239},
		{"wwb5f29f00ef255b16_wkwx_qw210629001643", 2238},
		{"wwb5f29f00ef255b16_wkwx_qw210629001648", 2237},
		{"wwb5f29f00ef255b16_wkwx_qw210629001651", 2236},
		{"wwb5f29f00ef255b16_wkwx_qw210629001654", 2235},
		{"wwb5f29f00ef255b16_wkwx_qw210629001663", 2234},
		{"wwb5f29f00ef255b16_wkwx_qw210629001665", 2233},
		{"wwb5f29f00ef255b16_wkwx_qw210629001666", 2232},
		{"wwb5f29f00ef255b16_wkwx_qw210629001670", 2211},
		{"wwb5f29f00ef255b16_wkwx_qw210629001673", 2210},
		{"wwb5f29f00ef255b16_wkwx_qw210629001675", 2209},
		{"wwb5f29f00ef255b16_wkwx_qw210629001677", 2208},
		{"wwb5f29f00ef255b16_wkwx_qw210629001678", 2207},
		{"wwb5f29f00ef255b16_wkwx_qw210629001679", 2206},
		{"wwb5f29f00ef255b16_wkwx_qw210629001681", 2205},
		{"wwb5f29f00ef255b16_wkwx_qw210629001682", 2204},
		{"wwb5f29f00ef255b16_wkwx_qw210630000121", 2203},
		{"ww7809deb62d31506e_wkwx_qw210714000177", 2399},
		{"ww7809deb62d31506e_wkwx_qw210723000342", 2398},
		{"ww7809deb62d31506e_wkwx_qw210723000414", 2396},
		{"ww7809deb62d31506e_wkwx_qw210723000426", 2395},
		{"ww7809deb62d31506e_wkwx_qw210723000437", 2394},
		{"ww7809deb62d31506e_wkwx_qw210723000559", 2393},
		{"wwb5f29f00ef255b16_wkwx_qw210727000054", 2202},
		{"wwb5f29f00ef255b16_wkwx_qw210727000056", 2160},
		{"wwb5f29f00ef255b16_wkwx_qw210727000058", 2159},
		{"wwb5f29f00ef255b16_wkwx_qw210825000505", 2158},
		{"wwb5f29f00ef255b16_wkwx_qw210923001303", 3045},
		{"wwb5f29f00ef255b16_wkwx_qw210923001307", 2157},
		{"wwb5f29f00ef255b16_wkwx_qw210923001321", 2156},
		{"wwb5f29f00ef255b16_wkwx_qw210924000023", 3044},
		{"wwb5f29f00ef255b16_wkwx_qw210929000609", 2155},
		{"wwb5f29f00ef255b16_wkwx_qw211018001274", 3043},
		{"wwb5f29f00ef255b16_wkwx_qw211018001275", 3042},
		{"wwb5f29f00ef255b16_wkwx_qw211018001291", 3041},
		{"wwb5f29f00ef255b16_wkwx_qw211018001293", 3019},
		{"wwb5f29f00ef255b16_wkwx_qw211018001354", 3018},
		{"wwb5f29f00ef255b16_wkwx_qw211019000111", 3017},
		{"wwb5f29f00ef255b16_wkwx_qw211019000114", 3016},
		{"wwb5f29f00ef255b16_wkwx_qw211019000117", 2154},
		{"wwb5f29f00ef255b16_wkwx_qw211019000123", 3015},
		{"wwb5f29f00ef255b16_wkwx_qw211019000125", 3014},
		{"wwb5f29f00ef255b16_wkwx_qw211019000127", 3013},
		{"wwb5f29f00ef255b16_wkwx_qw211019000151", 3012},
		{"wwb5f29f00ef255b16_wkwx_qw211019000168", 2528},
		{"wwb5f29f00ef255b16_wkwx_qw211019000169", 2527},
		{"wwb5f29f00ef255b16_wkwx_qw211019000179", 2526},
		{"wwb5f29f00ef255b16_wkwx_qw211022000031", 2525},
		{"wwb5f29f00ef255b16_wkwx_qw211029000097", 2524},
		{"wwb5f29f00ef255b16_wkwx_qw211029000098", 2351},
		{"wwb5f29f00ef255b16_wkwx_qw211029000100", 2350},
		{"wwb5f29f00ef255b16_wkwx_qw211029000110", 2349},
		{"wwb5f29f00ef255b16_wkwx_qw211029000111", 2347},
		{"wwb5f29f00ef255b16_wkwx_qw211029000113", 2344},
		{"wwb5f29f00ef255b16_wkwx_qw211029000129", 2343},
		{"wwb5f29f00ef255b16_wkwx_qw211029000135", 2342},
		{"wwb5f29f00ef255b16_wkwx_qw211029000140", 2330},
	}

	targetURL := "http://10.109.51.148:8099/wxtb/returnRobot/updateTeamForce"

	for _, item := range data {
		payload := map[string]interface{}{
			"wxIdList": []string{item.WxID},
			"teamId":   item.TeamID,
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			continue
		}

		resp, err := http.Post(targetURL, "application/json", bytes.NewBuffer(jsonPayload))
		if err != nil {
			fmt.Println("Error making request:", err)
			continue
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Response for wxID %s: %s\n", item.WxID, body)
	}
}
