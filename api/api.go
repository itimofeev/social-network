package api

import (
	_ "embed"
	"net/http"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@v0.76.0 --debug.ignoreNotImplemented "oauth2 security" -package api -clean -no-webhook-client -no-client -target ../internal/gen/api openapi.yaml

//go:embed openapi.yaml
var openApiDoc []byte

func OapiSchemaHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(openApiDoc)
}
