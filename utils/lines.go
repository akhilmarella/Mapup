package utils

import (
	"encoding/json"
	"mapup/api"
	"os"

	"github.com/golang/geo/s2"
	"github.com/rs/zerolog/log"
)

func GetLines() ([]s2.Shape, error) {
	var shapes []s2.Shape

	// Read the JSON file
	data, err := os.ReadFile("/home/akhil/Github/Mapup/utils/lines.json")
	if err != nil {
		log.Error().Err(err).Any("lines", data).Any("action", "utils_lines.go_GetLines").
			Msg("error in reading lines from json")
		return nil, err
	}

	// Unmarshal the JSON data
	var geojson []api.GeoJSON
	err = json.Unmarshal(data, &geojson)
	if err != nil {
		log.Error().Err(err).Any("lines_string_data", geojson).Any("action", "utils_lines.go_GetLines").
			Msg("error in unmarshalling ")
		return nil, err
	}

	// Convert GeoJSON data into s2.Shape slice
	for _, feature := range geojson {
		latlngs := make([]s2.LatLng, len(feature.Line.Coordinates))
		for i, coord := range feature.Line.Coordinates {
			latlngs[i] = s2.LatLngFromDegrees(coord[1], coord[0])
		}
		shapes = append(shapes, s2.PolylineFromLatLngs(latlngs))
	}
	return shapes, nil
}
