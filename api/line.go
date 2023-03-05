package api

type LineStringRequest struct {
	Coordinates [][]float64 `json:"coordinates"`
}

type GeoJSON struct {
	Line struct {
		Type        string      `json:"type"`
		Coordinates [][]float64 `json:"coordinates"`
	} `json:"line"`
}
