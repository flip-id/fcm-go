package fcm

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

func getNewAppParams() (s *httptest.Server, cfg *jwt.Config, err error) {
	s = initMockTokenServer()

	b, err := mockServiceAcct(s.URL)
	if err != nil {
		return
	}

	cfg, err = google.JWTConfigFromJSON(b)
	return
}

func initMockTokenServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"access_token": "mock-token",
			"scope": "user",
			"token_type": "bearer",
			"expires_in": 3600
		}`))
	}))
}

func mockServiceAcct(tokenURL string) ([]byte, error) {
	b, err := os.ReadFile("test/data/service_account.json")
	if err != nil {
		return nil, err
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(b, &parsed); err != nil {
		return nil, err
	}
	parsed["token_uri"] = tokenURL
	return json.Marshal(parsed)
}
