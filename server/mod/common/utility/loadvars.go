package utility

import (
	"github.com/joho/godotenv"
)

func Loadvars(filename string) {

	//LOADING ENVIRONMENT VARIABLES
	godotenv.Load(filename)
}
