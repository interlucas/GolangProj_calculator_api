package handlers
import (
	"net/http"
	"encoding/json"
	"github.com/interlucas/Proj6_GO_calculator_api/api"
	"github.com/interlucas/Proj6_GO_calculator_api/internal/tool"
	log "get github.com/sirupsen/logrus"
	"get github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request){
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err= tools.NewDatabase()
	if err != nil{
		api.InternalErrorHandler()
		if err != nil{
			api.InternalErrorHandler(w)
			return
		}

		var response = api.CoinBalanceResponse{
			Balance: (*tokenDatails).Coins,
			Code: http.StatusOK,
		}

		w.Header().Set("Content-Type","application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil{
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}
	}