package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ContactHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	var m ContactMessage

	config, err := readConf("conf.yaml")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&m)

	if err != nil || m.Validate() == false {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	err, dest := config.GetDestinationById(m.Destination)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	mail := MailMessage{
		dest:    dest,
		message: m,
	}

	err = mail.Send(config.Mailjet)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		fmt.Fprintf(w, "Failed To Send Message")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
