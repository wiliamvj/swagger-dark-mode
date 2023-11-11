package routes

import (
  "swagger-dark-mode/docs/custom"
  "swagger-dark-mode/internal/handler"

  _ "swagger-dark-mode/docs"

  "github.com/go-chi/chi/v5"
  httpSwagger "github.com/swaggo/http-swagger"
)

var (
  docsURL = "http://localhost:8080/docs/doc.json"
)

//	@title		Swagger Dark Mode
//	@version	1.0
func InitRoutes(r chi.Router) {
  r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL),
    httpSwagger.AfterScript(custom.CustomJS),
    httpSwagger.DocExpansion("none"),
    httpSwagger.UIConfig(map[string]string{
      "defaultModelsExpandDepth": `"-1"`,
    }),
  ))

  r.Get("/user", handler.GetUser)
}
