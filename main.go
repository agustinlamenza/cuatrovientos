package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(`
	Ingrese la edad mínima, máxima y sexo separando el numero por un espacio.
	Para el sexo femenimo (1) masculino (2)
	Ejemplo: 30 40 2
	`)
	var minStr, maxStr, sexStr string
	if _, err := fmt.Scanln(&minStr, &maxStr, &sexStr); err != nil {
		log.Fatalf("Error al leer el dato, %s", err)
	}

	edadMinima, err := strconv.Atoi(minStr)
	if err != nil {
		log.Fatalf("Dato invalido, %s", err)
	}

	edadMaxima, err := strconv.Atoi(maxStr)
	if err != nil {
		log.Fatalf("Dato invalido, %s", err)
	}

	sexo, err := strconv.Atoi(sexStr)
	if err != nil {
		log.Fatalf("Dato invalido, %s", err)
	}

	fd, err := os.Open("personas.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()

	people := []Person{}

	fileReader := csv.NewReader(fd)
	fileReader.Comma = ';'
	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range records {
		age, err := strconv.Atoi(v[4])
		if err != nil {
			continue
		}

		sex, err := strconv.Atoi(v[3])
		if err != nil {
			continue
		}

		p := Person{
			FirsName: v[1],
			LastName: v[2],
			Age:      age,
			Sex:      sex,
		}

		people = append(people, p)
	}

	for _, v := range people {
		if v.Age >= edadMinima && v.Age <= edadMaxima && v.Sex == sexo {
			fmt.Printf("Edad: %v - Nombre completo: %v, %v\n", v.Age, v.FirsName, v.LastName)
		}
	}
}

type Person struct {
	FirsName string
	LastName string
	Age      int
	Sex      int
}
