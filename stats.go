package expvarmetrics

import (
	"bytes"
	"encoding/json"
)

type rateStats struct {
	Rate1    float64 `json:"1min"`
	Rate5    float64 `json:"5min"`
	Rate15   float64 `json:"15min"`
	RateMean float64 `json:"mean"`
}

func toString(stats interface{}) string {
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(stats)
	return buff.String()
}
