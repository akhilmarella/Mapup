package service

import (
	"mapup/utils"

	"github.com/golang/geo/s2"
	"github.com/rs/zerolog/log"
)

func GetIntersectionPoint(lineString *s2.Polyline) ([]int, []s2.Point, error) {
	// Get lines from json file to
	// find intersection point
	lines, err := utils.GetLines()
	if err != nil {
		log.Error().Err(err).Any("lines", lines).Any("action", "service_intersection.go_GetIntersectionpoint").
			Msg("error in fetching from  lines data ")
		return nil, nil, err
	}

	// Create an S2 loop for the linestring
	loop := s2.LoopFromPoints(*lineString)

	// Create an S2 index for the lines
	index := s2.NewShapeIndex()

	// Add each line to the index
	for _, line := range lines {
		index.Add(line)
	}

	// Add a variable for storing intersecting lines number,
	// slice for storing intersecting points
	intersectingLines := make([]int, 0)
	intersectionPoints := make([]s2.Point, 0)

	for i, line := range lines {
		if line.NumEdges() == 0 {
			continue
		}
		for j := 0; j < line.NumEdges(); j++ {
			if loop.ContainsPoint(line.Edge(j).V0) || loop.ContainsPoint(line.Edge(j).V1) {
				intersectingLines = append(intersectingLines, i)
				intersectionPoints = append(intersectionPoints, line.Edge(j).V0, line.Edge(j).V1)
				break
			}
		}
	}
	return intersectingLines, intersectionPoints, nil
}
