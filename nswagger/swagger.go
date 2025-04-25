package nswagger

import (
	"embed"
	"net/http"
)

//go:embed index.html api.swagger.json
var swaggerFiles embed.FS

func SwaggerHandler() http.Handler {
	fs := http.FS(swaggerFiles)
	return http.StripPrefix("/swagger/", http.FileServer(fs))
}
