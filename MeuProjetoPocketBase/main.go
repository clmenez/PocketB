package main

import (
	"log"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	// Criar instância do PocketBase
	app := pocketbase.New()

	// Hook customizado - executado antes de servir
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		log.Println("PocketBase customizado iniciado!")

		// Adicionar rota customizada
		se.Router.GET("/api/custom/hello", func(re *core.RequestEvent) error {
			return re.JSON(http.StatusOK, map[string]interface{}{
				"message": "Hello from custom Go route!",
				"version": "1.0.0",
			})
		})

		// Adicionar middleware customizado
		se.Router.Use(func(re *core.RequestEvent) error {
			log.Printf("Request: %s %s", re.Request.Method, re.Request.URL.Path)
			return re.Next()
		})

		return se.Next()
	})

	// Hook para quando um registro é criado
	app.OnRecordCreate("users").BindFunc(func(e *core.RecordEvent) error {
		log.Printf("Novo usuário criado: %s", e.Record.GetString("email"))
		// Aqui você pode adicionar lógica customizada
		return e.Next()
	})

	// Iniciar o PocketBase
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
