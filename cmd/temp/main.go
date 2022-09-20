package main

import (
	"fmt"
	"log"
	"random-string-generator-service/internal/app/nanoid/config"
	"random-string-generator-service/internal/app/nanoid/repository"
	"random-string-generator-service/internal/app/nanoid/service"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	limit := config.LimitCharacter
	prefix := config.Prefix
	repoNano := repository.NewCmdNanoidRepository(limit, prefix)
	serviceNano := service.NewNanoidService(repoNano)

	fmt.Print("Enter how many random strings to generate: ")
	var input int
	fmt.Scanf("%d", &input)

	for i := 0; i < input; i++ {
		randomObj, ok := serviceNano.RunGenerateRandomString()
		if ok {
			response := fmt.Sprintf("%d : %s", (i + 1), randomObj.Value)
			fmt.Println(response)
		} else {
			fmt.Println("failed generate random string")
			break
		}
	}
}
