package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	fmt.Sprintf("%T", client)

}
