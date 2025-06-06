package main

import (
	"fmt"
)

func main() {
	var name1, name2, name3 string
	fmt.Scan(&name1, &name2, & name3)
	animal := " "

	if name1 == "vertebrado" {
		if name2 == "ave" {
			if name3 == "carnivoro" {
				animal = "aguia"
			}
			if name3 == "onivoro" {
				animal = "pomba"
			}
		}
		if name2 == "mamifero" {
			if name3 == "onivoro"{
				animal = "homem"
			}
			if name3 == "herbivoro" {
				animal = "vaca"
			}
		}
		
	}

	if name1 == "invertebrado" {
		if name2 == "inseto" {
			if name3 == "hematofago" {
				animal = "pulga"
			}
			if name3 == "herbivoro" {
				animal = "lagarta"
			}
		}
		if name2 == "anelideo" {
			if name3 == "hematofago" {
				animal = "sanguessuga"
			}
			if name3 == "onivoro" {
				animal = "minhoca"
			}
		}
	}

	fmt.Println(animal)
}