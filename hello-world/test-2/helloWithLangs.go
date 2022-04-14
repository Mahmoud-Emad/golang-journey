package main
import "fmt"


// Here We'll use if condation
// func Say(lang, name string) string {
// 	// This function writes to only say hello to [Name] -> You should pass it when you call it 
// 	// lang is language("FR", "SP", "EN") -> [Hello, Hola, Oui]
// 	word := ""
// 	if lang == "FR" {
// 		word = "Oui, "
// 	}else if lang == "SP" { 
// 		word = "Hola, "
// 	}else{
// 		word = "Hello, "
// 	}
// 	return word + name
// }
// Here We'll use Switch Case
func Say(lang, name string) string {
	// This function writes to only say hello to [Name] -> You should pass it when you call it 
	// lang is language("FR", "SP", "EN") -> [Hello, Hola, Oui]
	word := ""
	switch lang {
	case "FR":word = "Oui, "
	case "SP":word = "Hola, "
	default:word = "Hello, "
	} 
	return word + name
}


func helloWithLangs(name, lang string) string {
	// Name for name of user
	// Lang for Message Language 
	if name == "" {
		name = "Guest"
		fmt.Printf("You're %q\n", name)
	}
	return Say(lang, name)
}