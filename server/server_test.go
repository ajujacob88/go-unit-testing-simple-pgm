package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *Weather
	expectedError error
}

func TestGetWeather(t *testing.T) {
	tests := []Tests{
		{
			name: "basic-request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`"city":"Ekm,koc","forecast":"sunny"`))
			})),
			response: &Weather{
				City:     "Ekm,koc",
				Forecast: "sunny",
			},
			expectedError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()

			resp, err := GetWeather(test.server.URL)

			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
			}
			if !errors.Is(err, test.expectedError) {
				t.Errorf("FAILED: expected %v, got %v\n", test.expectedError, err)

			}

		})
	}
}
