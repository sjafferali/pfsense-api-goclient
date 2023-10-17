package pfsenseapi

import (
	"net/http"
	"testing"
)

func popSlice[K comparable](slice []K, s int) []K {
	return append(slice[:s], slice[s+1:]...)
}

type resultList struct {
	resultsData   []string
	resultsStatus []int
}

func (s *resultList) popResult() string {
	response := s.resultsData[0]

	s.resultsData = popSlice(s.resultsData, 0)
	return response
}

func (s *resultList) popStatus() int {
	response := s.resultsStatus[0]

	s.resultsStatus = popSlice(s.resultsStatus, 0)
	return response
}

func makeResultList(t *testing.T, data string) *resultList {
	return &resultList{
		resultsData: []string{
			data,
			mustReadFileString(t, "testdata/error.json"),
			mustReadFileString(t, "testdata/badjson.json"),
		},
		resultsStatus: []int{
			http.StatusOK,
			http.StatusBadRequest,
			http.StatusOK,
		},
	}
}
