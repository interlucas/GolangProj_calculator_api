package tool
import(
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken:"123ABC",
		Username: "alex",
	},
	"jason": {
		AuthToken:"456DEF",
		Username: "jason",
	},
	"marie": {
		AuthToken:"789GHI",
		Username: "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:"100",
		Username: "alex",
	},
	"jason": {
		Coins:"200",
		Username: "jason",
	},
	"marie": {
		Coins:"300",
		Username: "marie",
	},
}