package main

import (
	"fmt"
	//"time"
	"math/rand"
	"io/ioutil"
	//"log"
	"net/http"
	"bytes"
	"encoding/json"
)

type Caso struct { //estructura para la leída de archivos
	Nombre string `json:"Nombre"`
	Departamento string `json:"Departamento"`
	Edad int `json:"Edad"`
	FormaContagio string `json:"Forma de contagio"`
	Estado string `json:"Estado"`
}

type CasoConvertido struct{//estructura para enviar el json
	Nombre string `json:"nombre"`
	Departamento string `json:"departamento"`
	Edad int `json:"edad"`
	FormaContagio string `json:"forma_contagio"`
	Estado string `json:"estado"`
}

func main() {
	
	var opcion int
	var url string
	var num_hilos int
	var num_solicitudes int
	var ruta_archivo string 

	fmt.Println("Proyecto # 2 - Sistemas Operativos 1")
	fmt.Println("Yoselin Lemus   - 201403819")
	fmt.Println("Ruben Osorio    - 201403703")
	fmt.Println("Brandon Alvarez - 201403862")

	for {
		url="";
		num_hilos=0;
		num_solicitudes=0;
		ruta_archivo="";
		opcion = 0;
		fmt.Println("---------- MENU -----------")
		fmt.Println("1. Envío de datos")
		fmt.Println("2. Salir")
		fmt.Println("---------------------------")
		fmt.Println("Ingrese opcion: ")

		fmt.Scanln(&opcion)

		switch {
			case opcion == 2: return
			case opcion != 1: continue
		}

		fmt.Println("Completar la siguiente información")
		/*
		fmt.Println("URL balanceador:")
		fmt.Scanln(&url)
		fmt.Println("Cantidad de hilos:")
		fmt.Scanln(&num_hilos)
		fmt.Println("Cantidad de solicitudes:")
		fmt.Scanln(&num_solicitudes)
		fmt.Println("Ruta del Archivo:")
		fmt.Scanln(&ruta_archivo)
		*/
		url = "http://localhost:3000/"
		num_hilos = 5
		num_solicitudes = 10
		ruta_archivo = "a.json"

		fmt.Printf("URL: %s, Hilos: %d, Solicitudes: %d, Archivo: %s\n",url,num_hilos,num_solicitudes,ruta_archivo);	
		
		if url==""{
			fmt.Println("NO se ha ingresado url, intente de nuevo")
			continue;
		}		
		if num_hilos < 1{
			fmt.Println("NO se ha ingresado numero de hilos, intente de nuevo")
			continue;
		}
		if num_solicitudes < 1{
			fmt.Println("NO se ha ingresado numero de solicitudes, intente de nuevo")
			continue;
		}

		if ruta_archivo == ""{
			fmt.Println("NO se ha ingresado la ruta del archivo, intente de nuevo")
			continue;
		}

		bytesLeidos, err := ioutil.ReadFile(ruta_archivo)
		if err != nil {
			fmt.Printf("Error leyendo archivo: %v", err)
			continue
		}

		contenido := string(bytesLeidos)
		var misCasos []Caso
		json.Unmarshal([]byte(contenido),&misCasos)

		for x := 0; x < num_hilos; x++{
			go miHilo(x, num_solicitudes/num_hilos, misCasos, url)
		}
	}	
}

func miHilo(hiloActual int, casosAEnviar int, arregloCasos []Caso, url string){
	for contador := 0; contador < casosAEnviar; contador++{
		var indice = rand.Intn(len(arregloCasos))
		peticion(arregloCasos[indice], url)
	}
}

func peticion(caso Caso, url string){
	casoEnviar := CasoConvertido { caso.Nombre, caso.Departamento, caso.Edad, caso.FormaContagio, caso.Estado }
	jsonValue, _ := json.Marshal(casoEnviar)
	resp, err := http.Post(url,"application/json",bytes.NewBuffer(jsonValue))
    if(err != nil){
		fmt.Println(err)
    }
    defer resp.Body.Close()
}
