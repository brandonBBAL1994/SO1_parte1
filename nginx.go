package main

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"bytes"
)

type Caso struct{
	Nombre string `json:"nombre"`
	Departamento string `json:"departamento"`
	Edad string `json:"edad"`
	FormaContagio string `json:"forma"`
	Estado string `json:"estado"`
}

func CrearCaso(w http.ResponseWriter, req *http.Request){
	var nuevoCaso Caso
	_ = json.NewDecoder(req.Body).Decode(&nuevoCaso)
	json.NewEncoder(w).Encode(nuevoCaso)

	jsonValue, _ := json.Marshal(nuevoCaso)
	resp, err := http.Post("http://35.223.179.117:30029/ingreso","application/json",bytes.NewBuffer(jsonValue))
    if(err != nil){
		fmt.Println(err)
    }
    defer resp.Body.Close()
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", CrearCaso).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}