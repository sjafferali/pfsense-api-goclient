package pfsenseapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnbound_CreateHostOverride(t *testing.T) {
	data := mustReadFileString(t, "testdata/listhostoverrides.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			value := UnboundHostOverride{}

			body, err := io.ReadAll(r.Body)

			if err != nil {
				t.Fatalf("Unable to read body %v", err)
			}

			if err := json.Unmarshal(body, &value); err != nil {
				t.Fatalf("Encountered error while parsing body %v", err)
			}

			r := apiWriteResponse[UnboundHostOverride]{
				apiResponse: apiResponse{
					Status:  "ok",
					Code:    200,
					Return:  0,
					Message: "success",
				},
				Data: &value,
			}

			result, err := json.Marshal(r)

			if err != nil {
				t.Fatalf("Encountered error while marshalling response %v", err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)

			if _, err := w.Write(result); err != nil {
				t.Errorf("Encountered error while writing response: %v", err)
			}
		} else {
			w.WriteHeader(200)
			_, _ = io.WriteString(w, data)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Unbound.CreateHostOverride(context.Background(), &UnboundHostOverride{
		Domain: "test.com",
		Host:   "two",
	}, true)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestUnbound_UpdateHostOverride(t *testing.T) {
	data := mustReadFileString(t, "testdata/listhostoverrides.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			value := UnboundHostOverride{}

			body, err := io.ReadAll(r.Body)

			if err != nil {
				t.Fatalf("Unable to read body %v", err)
			}

			if err := json.Unmarshal(body, &value); err != nil {
				t.Fatalf("Encountered error while parsing body %v", err)
			}

			r := apiWriteResponse[UnboundHostOverride]{
				apiResponse: apiResponse{
					Status:  "ok",
					Code:    200,
					Return:  0,
					Message: "success",
				},
				Data: &value,
			}

			result, err := json.Marshal(r)

			if err != nil {
				t.Fatalf("Encountered error while marshalling response %v", err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)

			if _, err := w.Write(result); err != nil {
				t.Errorf("encountered error while writing response: %v", err)
			}
		} else {
			w.WriteHeader(200)
			_, _ = io.WriteString(w, data)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Unbound.UpdateHostOverride(context.Background(), &UnboundHostOverride{
		Domain: "test.com",
		Host:   "two",
	}, true)
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestUnbound_ListHostOverrides(t *testing.T) {
	data := makeResultList(t, mustReadFileString(t, "testdata/listhostoverrides.json"))

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(data.popStatus())
		_, _ = io.WriteString(w, data.popResult())
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	response, err := newClient.Unbound.ListHostOverrides(context.Background())
	require.NoError(t, err)
	require.Len(t, response, 2)

	response, err = newClient.Unbound.ListHostOverrides(context.Background())
	require.Error(t, err)
	require.Nil(t, response)

	response, err = newClient.Unbound.ListHostOverrides(context.Background())
	require.Error(t, err)
	require.Nil(t, response)

}

func TestUnbound_DeleteHostOverride(t *testing.T) {
	data := mustReadFileString(t, "testdata/listhostoverrides.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(200)
			_, _ = io.WriteString(w, data)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	err := newClient.Unbound.DeleteHostOverride(context.Background(), "two", "test.com", true)
	require.NoError(t, err)
}
