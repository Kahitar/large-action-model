package main

import (
	"finances/cmd"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cmd.Execute()
}

