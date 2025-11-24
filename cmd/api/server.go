package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "schoolmanagement/internal/api/middlewares"
	"schoolmanagement/internal/api/router"
	"schoolmanagement/internal/repository/sqlconnect"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	_, err = sqlconnect.ConnectDB()
	if err != nil {
		log.Fatalln("Database connection error:", err)
	}

	port := ":" + os.Getenv("API_PORT")

	cert := "cert.pem"
	key := "key.pem"
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// hppOptions := mw.HPPOptions{
	// 	Whitelist:                []string{"name", "city"},
	// 	CheckQuery:               true,
	// 	CheckBody:                true,
	// 	CheckBodyOnlyContentType: "application/x-www-form-urlencoded",
	// }

	// rl := mw.NewRateLimiter(5, time.Minute)

	// secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTime, rl.MiddleWare, mw.Cors)
	// mw.Cors(rl.MiddleWare(mw.ResponseTime(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	router := router.NewRouter()
	secureMux := mw.SecurityHeaders(router)

	server := &http.Server{
		Addr:      port,
		TLSConfig: tlsConfig,
		Handler:   secureMux,
	}

	fmt.Println("Starting server on port:", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
