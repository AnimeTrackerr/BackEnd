package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load(path string) {
	err := godotenv.Load(path)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("env variables loaded successfully\n")
	}

	_ , present1 := os.LookupEnv("MONGODB_URI")

	if !present1 {
		fmt.Printf("Error : Cannot find environment variable: \"MONGODB_URI\" \n")
		os.Exit(1)
	}

}