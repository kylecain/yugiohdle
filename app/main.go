package main

import (
    "yugiohdle/app/router"
)

func main() {
    r := router.New()

    r.Logger.Fatal(r.Start(":8080"))
}
