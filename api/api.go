package api

import (
	_ "embed"
	"net/http"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@v0.81.0 -package api -clean -target ../internal/gen/api openapi.yaml

//go:embed openapi.yaml
var openApiDoc []byte

func OapiSchemaHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(openApiDoc)
}
