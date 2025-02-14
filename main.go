package main

import (
    "fmt"
    "log"
    "github.com/skip2/go-qrcode"
)

func main() {
    // Content to be encoded in the QR code
    content := "https://www.youtube.com/@GraziMarsicano"

    // Generate the QR code and save it to a file
    err := qrcode.WriteFile(content, qrcode.Medium, 256, "qrcode.png")
    if err != nil {
        log.Fatal(err) // If an error occurs, terminate the program
    }

    fmt.Println("QR Code successfully generated in 'qrcode.png'")
}