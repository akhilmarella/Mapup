package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"mapup/api"
	"mapup/service"
	"net/http"

	"github.com/golang/geo/s2"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetIntersection(c *gin.Context) {
	Id := c.Writer.Header().Get("id")
	if Id == "" {
		log.Error().Any("action", "handlers_intersection.go_GetIntersection").
			Msg("id is empty")
		c.JSON(http.StatusNoContent, gin.H{"message": "id is empty"})
		return
	}

	log.Info().Any("id:", Id).Msg("new request for id")

	// Validate request & unmarshall the linestring into
	// s2.Polyline
	lineString, err := validateRequest(c.Request.Body)
	if err != nil {
		log.Error().Err(err).Any("action", "handlers_intersection.go_GetIntersection").
			Msg("error in unmarshalling ")
		c.JSON(http.StatusBadGateway, gin.H{"message": "error in unmarshalling the request linestring"})
		return
	}

	// Get intersection points from linestring
	intersectingLines, intersectingPoints, err := service.GetIntersectionPoint(lineString)
	if err != nil {
		log.Error().Err(err).Any("line_string", lineString).Any("action", "handlers_intersection.go_getIntersection").
			Msg("error found in fetching  data from intersecting points ")
		c.JSON(http.StatusBadRequest, gin.H{"message": "error in fetching from intersection point from linestring"})
		return
	}

	if len(intersectingLines) == 0 {
		log.Error().Any("action", "handlers_intersection.go_GetIntersection").
			Msg("no intersection line")
		c.JSON(http.StatusNoContent, gin.H{"intersecting_lines_index": intersectingLines})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"intersecting_lines_index": intersectingLines,
		"intersecting_points": intersectingPoints})
}

func validateRequest(req io.ReadCloser) (*s2.Polyline, error) {
	byteData, err := ioutil.ReadAll(req)
	if err != nil {
		log.Error().Err(err).Any("action", "handlers_intersectio.go_validateRequest").
			Msg("error in converting req body into byte ")
		return nil, err
	}

	var lineStringData api.LineStringRequest
	// Unmashalling the lineString data
	err = json.Unmarshal(byteData, &lineStringData)
	if err != nil {
		log.Error().Err(err).Any("byte_data", byteData).Any("line_string_data", lineStringData).
			Any("action", "handlers_intersectio.go_validateRequest").Msg("error in unmarshalling ")
		return nil, err
	}

	// Convert the JSON data into the appropriate format
	// for finding the intersection point
	var latLngs []s2.LatLng
	for _, coord := range lineStringData.Coordinates {
		latLngs = append(latLngs, s2.LatLngFromDegrees(coord[1], coord[0]))
	}

	linestring := s2.PolylineFromLatLngs(latLngs)

	return linestring, nil
}
