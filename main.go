package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Penpen7/shellcheck-actions/shellcheck"
)

func main() {
	v := shellcheck.Response{}
	err := json.NewDecoder(os.Stdin).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

}
