package main

import "os"

func main() {
	f, err := os.Create("created.file")
	if err != nil {
		panic(err)

	f.Close()
}



}