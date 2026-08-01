package main

import (
 "bytes"
 "encoding/json"
 "fmt"
 "io"
 "net/http"
 "os"
 "strings"
 "time"
)

// DİKKAT: Render adresini buraya yazıp sonuna /api/chat ekle
const proxyURL = "https://garlic-proxy.onrender.com/api/chat"

type RequestBody struct {
 Prompt string `json:"prompt"`
}

type ResponseBody struct {
 Result string `json:"result"`
}

func main() {
 if len(os.Args) < 2 {
  fmt.Println("Kullanım: garlic \"Sorunuz\"")
  os.Exit(1)
 }

 userPrompt := strings.Join(os.Args[1:], " ")

 reqData := RequestBody{Prompt: userPrompt}
 jsonData, _ := json.Marshal(reqData)

 client := http.Client{Timeout: 20 * time.Second}
 resp, err := client.Post(proxyURL, "application/json", bytes.NewBuffer(jsonData))
 if err != nil {
  fmt.Println("\nBağlantı Hatası: Sunucuya ulaşılamadı.")
  os.Exit(1)
 }
 defer resp.Body.Close()

 body, _ := io.ReadAll(resp.Body)

 var resData ResponseBody
 _ = json.Unmarshal(body, &resData)

 if resp.StatusCode == 200 {
  fmt.Printf("\n\033[1;32m[GarlicAI]:\033[0m\n%s\n\n", resData.Result)
 } else {
  fmt.Printf("\n[Hata]: Sunucu yanıt vermedi (Kod: %d)\n\n", resp.StatusCode)
 }
}
