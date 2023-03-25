package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Hello, World! \n Testando Arquivo \n Escrevendo no arquivo."))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes \n", size)

	f.Close()

	// leitura

	file, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	//leitura de pouco em pouco abrindo o arquivo

	file2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)

	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
