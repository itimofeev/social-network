package api

import (
	_ "embed"
	"net/http"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@v0.81.0 -package api -clean -target ../internal/server/backend/gen/api backend.yaml
//go:generate go run github.com/ogen-go/ogen/cmd/ogen@v0.81.0 -package api -clean -target ../internal/server/dialogs/gen/api dialogs.yaml

//go:embed backend.yaml
var backendAPIDoc []byte

//go:embed dialogs.yaml
var dialogsAPIDoc []byte

func BackendOapiSchemaHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(backendAPIDoc)
}

func DialogsOapiSchemaHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(dialogsAPIDoc)
}
