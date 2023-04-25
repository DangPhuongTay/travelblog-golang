package main

import (
	"net/http"

	"github.com/DangPhuongTay/travelblog-golang/database"
	"github.com/rs/cors"
)

func main() {
	database.Connect()

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// port := os.Getenv("PORT")
	// app := fiber.New()
	// routes.Setup(app)
	// app.Listen(":" + port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
