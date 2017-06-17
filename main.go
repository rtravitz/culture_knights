package main

import (
	"os"
)

func main() {
	a := App{}
	a.Initialize(os.Getenv("CULTURE_DB"))

	a.Run(":" + os.Getenv("PORT"))
}
