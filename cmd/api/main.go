package main


import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.handleFunc()
}