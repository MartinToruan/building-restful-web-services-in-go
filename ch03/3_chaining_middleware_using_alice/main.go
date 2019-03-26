package main

import (
	"encoding/json"
	"fmt"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string `json:"name"`
	Area uint64 `json:"area"`
}

func filterContentType(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("Checking Content Type...")
		if r.Header.Get("Content-type") != "application/json"{
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, _ = w.Write([]byte("415 - Unsupported Media Type. Please send JSON."))
		}
		handler.ServeHTTP(w, r)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("Setting Up Coockie...")
		handler.ServeHTTP(w, r)
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Finish Setting Up Cookie...")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Check Method
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&tempCity); err != nil {
			log.Fatalf("Can't decode message: %v", err)
		}
		defer r.Body.Close()

		fmt.Printf("Got %s city with area %d sq miles!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
	}
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	chain := alice.New(filterContentType, setServerTimeCookie).Then(mainLogicHandler)
	http.Handle("/city", chain)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
