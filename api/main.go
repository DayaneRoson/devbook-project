package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*func init() {
	key := make([]byte, 64)
	if _, error := rand.Read(key); error != nil {
		log.Fatal(error)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)
}*/

func main() {
	config.ChargeConfigs()
	r := router.Generate()
	fmt.Printf("Listening on port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
