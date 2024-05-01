package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/YounesBouchbouk/natours-api-golang/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func addAuthorization(
	request *http.Request,
	t *testing.T,
	tokenMaker token.Maker,
	authrizationType string,
	email string,
	duration time.Duration,
) {
	token, _, err := tokenMaker.CreateToken(email, "admin", duration)

	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authrizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestMiddlware(t *testing.T) {
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		checkResponse func(t *testing.T, record *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(request, t, tokenMaker, authorizationTypeBearer, "user", time.Minute)
			},
			checkResponse: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, record.Code)
			},
		},
		{
			name: "Unsupported Authorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(request, t, tokenMaker, "unsupported", "user", time.Minute)
			},
			checkResponse: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, record.Code)
			},
		},
		{
			name: "Header Not Sent",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			checkResponse: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, record.Code)
			},
		},
		{
			name: "InvalidAutorizationFormat",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(request, t, tokenMaker, "", "user", time.Minute)
			},
			checkResponse: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, record.Code)
			},
		},
		{
			name: "InvalidAutorizationFormat",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(request, t, tokenMaker, authorizationTypeBearer, "user", -time.Minute)
			},
			checkResponse: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, record.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, nil)
			authPath := "/auth"
			server.router.GET(
				authPath,
				AuthenticationMiddlware(*server),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}
