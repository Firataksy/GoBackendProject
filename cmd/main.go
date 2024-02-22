package main

import (
	"fmt"
	"net/http"

	"github.com/my/repo/internal/api"
)

func main() {
	err := api.Handler()
	if err != nil {
		fmt.Println(err)
	}

	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
	}

}
