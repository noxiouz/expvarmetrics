package expvarmetrics

import (
	"bytes"
	"encoding/json"
	"expvar"

	"github.com/rcrowley/go-metrics"
)

var (
	_ expvar.Var = &MeterVar{}
)

// MeterVar adds expvar.Var interface to go-metrics.Meter
type MeterVar struct {
	metrics.Meter
}

// NewMeterVar returns new MeterVar with go-metrics.StandartMeter inside
func NewMeterVar() MeterVar {
	return MeterVar{
		Meter: metrics.NewMeter(),
	}
}

type meterStats struct {
	Rate  rateStats `json:"rate"`
	Count int64     `json:"count"`
}

func (m MeterVar) String() string {
	ss := m.Snapshot()
	stats := meterStats{
		Count: ss.Count(),
		Rate: rateStats{
			Rate1:    ss.Rate1(),
			Rate5:    ss.Rate5(),
			Rate15:   ss.Rate15(),
			RateMean: ss.RateMean(),
		},
	}

	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(&stats)
	return buff.String()
}
