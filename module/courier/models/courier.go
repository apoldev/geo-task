package models

import "encoding/json"

type Courier struct {
	Score    int   `json:"score"`
	Location Point `json:"location"`
}

func (c Courier) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
