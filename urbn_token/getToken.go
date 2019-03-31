package main

import (
	"fmt"

	"github.com/urbn/gourbnkit/token"
)

var payloadHmacKey = "ac85c9715a7ddb57bbeb5914cb79856bec7f4afceb363b14a7c9739d6f0360dbc3c7b13a7593a927f81e37630ab3573172b8d3cb2be9404481b8a4e8a65399be213b77a09cf88d7d63277308cbd6258f01ae003ce903297172b7629082085703f9d053a68c93ad6f0fe0427f374fdb14bfad63fec238babb976379f9b203397c"
var payloadKey = []byte("675ce83c423d6ac1c94442a3444aabcb")
var signingSecret = []byte("e20b64eebf8aef525feb6682193efd8cbee8b3eec5d737d916728a6994fc4edf")

func main() {
	//fmt.Println("Get token")
	//cs := token.ClaimData{
	//	DataCenterID: "US-PA",
	//	ProfileID:    "0b9cdf5c-fa3d-4bc0-b877-cdf40fa0a5db",
	//	CartID:       "b1d37cc9-8962-46d7-be09-182ee025965a",
	//	BrandID:      "AN",
	//	Scope:        []string{"CALL_CENTER", "GUEST"},
	//	SiteID:       "an-us",
	//	CreatedTime:  float64(time.Now().Unix()),
	//	SiteGroup:    "an-us",
	//	Employee:     false,
	//	Tracer:       "Tarce-01",
	//	WebID:        "",
	//}
	//tok, err := token.GenerateToken(signingSecret, payloadKey, payloadHmacKey, &cs, 8760*time.Hour, "an")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(tok)

	tok := "eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjoie1wicHJvZmlsZUlkXCI6IFwiSlByRVg5UkRkNGxjVlBaeTBURGg3dz09MjU3NTBkNTFhNDc0OGI4NzY3MmFmMjIyMTdjY2FhZjg5MmU0MTIxMDA3ZDY3MzRkZjc1MzQ5YTMyYWMzZmFmN1wiLFwic2NvcGVcIjogW1wiVElCQ09cIl0sXCJicmFuZElkXCI6IFwiYW5cIixcImRhdGFDZW50ZXJJZFwiOiBudWxsLFwiY3JlYXRlZFRpbWVcIjogMTU1MTg5NjEwMy41MDY5OTk5Njk0ODI0MjE4NzUsXCJ0cmFjZXJcIjogXCJHVVVSUVJDTUxKXCJ9IiwiaXNzIjoiYW4iLCJpYXQiOjE1NTE4OTYxMDQsImV4cCI6MTU1MTg5NjEzNH0.WXhNYp5s-iGkSTZ_8qiyLlnFG0oKT8_QVmp1hArJKk4"
	urbnToken := token.ValidateToken(tok, signingSecret, payloadKey)
	fmt.Printf("%+v", urbnToken)

}
