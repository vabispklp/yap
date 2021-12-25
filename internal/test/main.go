package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

var cars = map[string]string{
	"id1": "Renault",
	"id2": "BMW",
	"id3": "VW",
	"id4": "Audi",
}

func main() {
	r := chi.NewRouter()
	r.Get("/cars", func(rw http.ResponseWriter, r *http.Request) {
		carsList := carsListFunc()
		_, err := io.WriteString(rw, strings.Join(carsList, ","))
		if err != nil {
			panic(err)
		}
	})

	r.Get("/car", func(rw http.ResponseWriter, r *http.Request) {
		carID := r.URL.Query().Get("id")
		if carID == "" {
			http.Error(rw, "carID param is missed", http.StatusBadRequest)
			return
		}
		rw.Write([]byte(carFunc(carID)))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// carsListFunc — вспомогательная функция для вывода всех машин.
func carsListFunc() []string {
	var list []string
	for _, c := range cars {
		list = append(list, c)
	}
	return list
}

// carFunc — вспомогательная функция для вывода определённой машины.
func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return ""
}
