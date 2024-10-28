package client

import (
	"math"
)

type Statistics struct {
	Mean  float64
	Std   float64
	Count int
	Sum   float64
	Sumsq float64
}

func (s *Statistics) DetectAnomaly(value float64, k float64) bool {
	if s.Count < 2 {
		return false // Not enough data to determine anomaly
	}

	lowerBound := s.Mean - k*s.Std
	upperBound := s.Mean + k*s.Std
	return value < lowerBound || value > upperBound
}

func (s *Statistics) Update(newValue float64) {
	s.Count++
	s.Sum += newValue
	s.Sumsq += newValue * newValue
	s.Mean = s.Sum / float64(s.Count)
	if s.Count > 1 {
		variance := (s.Sumsq / float64(s.Count)) - (s.Mean * s.Mean)
		if variance < 0 {
			variance = 0
		}
		s.Std = math.Sqrt(variance)
	}
}
