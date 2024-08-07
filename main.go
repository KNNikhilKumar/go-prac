package main

import (
	"fmt"
	"go-prac/firstclass"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//1.FirstClass
	firstclass.ExecFC()
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	fmt.Println(port)

}
