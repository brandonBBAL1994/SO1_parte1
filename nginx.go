package main

import(
	"encoding/json" //codificar y decodificar json
	"log" //ver errores en servidor
	"net/http" //peticiones, funcionalidad web
	"github.com/gorilla/mux"
	"fmt"  //esta es para imprimir en consola
	"bytes"
	//"strconv" //esta es para convesiones
)

type Caso struct{
	Nombre string `json:"nombre"`
	Departamento string `json:"departamento"`
	Edad int `json:"edad"`
	FormaContagio string `json:"forma_contagio"`
	Estado string `json:"estado"`
}

func CrearCaso(w http.ResponseWriter, req *http.Request){
	var nuevoCaso Caso
	_ = json.NewDecoder(req.Body).Decode(&nuevoCaso)
	json.NewEncoder(w).Encode(nuevoCaso)

	//aqui debo mandar a la url por post del servidor de colas de nats de go..... la url debe ser estática
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