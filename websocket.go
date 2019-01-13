package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"sync"
)

type CoinInfo struct {
	ID                  string `json:"Id"`
	URL                 string `json:"Url"`
	ImageURL            string `json:"ImageUrl"`
	Name                string `json:"Name"`
	Symbol              string `json:"Symbol"`
	CoinName            string `json:"CoinName"`
	FullName            string `json:"FullName"`
	Algorithm           string `json:"Algorithm"`
	ProofType           string `json:"ProofType"`
	FullyPremined       string `json:"FullyPremined"`
	TotalCoinSupply     string `json:"TotalCoinSupply"`
	PreMinedValue       string `json:"PreMinedValue"`
	TotalCoinsFreeFloat string `json:"TotalCoinsFreeFloat"`
	SortOrder           string `json:"SortOrder"`
	Sponsored           bool   `json:"Sponsored"`
}

type CoinListData struct {
	Response         string `json:"Response"`
	Message          string `json:"Message"`
	BaseImageURL     string `json:"BaseImageUrl"`
	BaseLinkURL      string `json:"BaseLinkUrl"`
	DefaultWatchlist struct {
		CoinIs    string `json:"CoinIs"`
		Sponsored string `json:"Sponsored"`
	} `json:"DefaultWatchlist"`
	Data struct {
		CoinList map[string]CoinInfo
	} `json:"Data"`
}

var clientMutex  sync.Mutex
var clients []*websocket.Conn

type apiPriceResponse struct {
	USD float64 `json:"USD"`
}

type HttpJsonResponse struct {
	Coin 		string 	`json:"coin"`
	Currency 	string 	`json:"currency"`
	Price 		string  `json:"price"`
}

func main() {
	apiUrl := "https://min-api.cryptocompare.com/data/all/coinlist"
	resp, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var coinlist CoinListData
	err = json.Unmarshal(body, &coinlist)
	if(err != nil){
		fmt.Println("whoops:", err)
	}
	fmt.Printf("%+v", coinlist.Data.CoinList)

	coins := []string { "BTC", "ETH", "LTC" }
	r := gin.Default()
	r.LoadHTMLFiles("./views/ws.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "ws.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	for _, coin := range coins {
		go interval(coin)
	}


	r.Run("localhost:12312")
}

func interval(coin string) {
	for {
		time.Sleep(1*time.Second)
		resp, err := http.Get(fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=USD", coin))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		var s apiPriceResponse
		err = json.Unmarshal(body, &s)
		if(err != nil){
			fmt.Println("whoops:", err)
		}

		httpResponse := &HttpJsonResponse{
			Coin: coin,
			Currency: "USD",
			Price: fmt.Sprintf("%.2f", s.USD),
		}
		for _, conn := range clients {
			clientMutex.Lock()
			conn.WriteJSON(httpResponse)
			clientMutex.Unlock()
		}
	}
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	clients = append(clients, conn)
	//for {
	//	t, msg, err := conn.ReadMessage()
	//	if err != nil {
	//		break
	//	}
	//	log.Println(t, msg)
	//	conn.WriteMessage(t, msg)
	//}
}
