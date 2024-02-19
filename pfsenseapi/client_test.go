package pfsenseapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientErrors(t *testing.T) {
	for code, expectedErr := range responseCodeErrorMap {
		t.Run(strconv.Itoa(code), func(t *testing.T) {
			apiRes := new(apiResponse)
			apiRes.Code = code
			apiRes.Message = "test message"

			response, err := json.Marshal(apiRes)
			require.NoError(t, err)

			ctx := context.Background()

			handler := func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(code)
				_, err := io.WriteString(w, string(response))
				require.NoError(t, err)
			}

			server := httptest.NewServer(http.HandlerFunc(handler))
			defer server.Close()

			client := NewClientWithNoAuth(server.URL)
			for _, method := range []string{"GET", "POST", "PUT", "DELETE"} {
				var res []byte

				switch method {
				case "GET":
					res, err = client.get(ctx, "/test", nil)
				case "POST":
					res, err = client.post(ctx, "/test", nil, nil)
				case "PUT":
					res, err = client.put(ctx, "/test", nil, nil)
				case "DELETE":
					res, err = client.delete(ctx, "/test", nil)
				}
				require.ErrorIs(t, err, expectedErr)
				require.Nil(t, res)
			}
		})
	}
}
