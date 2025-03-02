package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//scan do input
	
	password := generatePassword(5)
	fmt.Println(password)
}

func generatePassword(length int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"

	allChars := lowerCase + upperCase + numbers + special

	mandatory := []byte{
		upperCase[rand.Intn(len(upperCase))],
		numbers[rand.Intn(len(numbers))],

	}
	password := make([]byte, length-len(mandatory))

	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	password = append(password, mandatory...)

	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

// Obrigatoriedades do sistema
// Gerar uma senha aleatória - ok
// Gerar uma senha aleatória MAS com obrigatoriedades: letra maiúscula, números... - ok
// Pegar o input do usuário de quantos caracteres ele quer na senha. E então usar o input para o tamanho da senha.