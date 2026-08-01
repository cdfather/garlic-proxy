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

// Render adresin
const proxyURL = "https://garlic-proxy.onrender.com/api/chat"

type RequestBody struct {
 Prompt string `json:"prompt"`
}

type ResponseBody struct {
 Result string `json:"result"`
}

// ANSI Color Codes
const (
 ColorPurple = "\033[1;35m"
 ColorCyan = "\033[1;36m"
 ColorYellow = "\033[1;33m"
 ColorReset = "\033[0m"
)

func printBanner() {
 banner := `
   ______ ___ ____ __ ____ ______ ___ ____
  / ____/ / | / __ \ / / / _// ____/ / | / _/
 / / __ / /| | / /_/ // / / / / / / /| | / /  
/ /_/ / / ___ |/ _, _// /____/ / / /___ / ___ |_/ /   
\____/ /_/ |_/_/ |_/_____/___/ \____/ /_/ |_/___/   
`
 fmt.Println(ColorPurple + banner + ColorReset)
 fmt.Println(ColorCyan + " 🧄 GarlicAI — Terminal AI Assistant" + ColorReset)
 fmt.Println(ColorYellow + " ─────────────────────────────────────────────────────────" + ColorReset)
 fmt.Println(" Usage: garlic \"Your prompt goes here\"")
 fmt.Println(" Example: garlic \"How do I check memory usage in Ubuntu?\"")
 fmt.Println(ColorYellow + " ─────────────────────────────────────────────────────────\n" + ColorReset)
}

func main() {
 if len(os.Args) < 2 {
  printBanner()
  os.Exit(0)
 }

 userPrompt := strings.Join(os.Args[1:], " ")

 reqData := RequestBody{Prompt: userPrompt}
 jsonData, _ := json.Marshal(reqData)

 // HTTP Request oluşturma
 req, err := http.NewRequest("POST", proxyURL, bytes.NewBuffer(jsonData))
 if err != nil {
  fmt.Printf("\n%s[GarlicAI Error]:%s Request creation failed: %v\n\n", ColorPurple, ColorReset, err)
  os.Exit(1)
 }

 req.Header.Set("Content-Type", "application/json")
 req.Header.Set("User-Agent", "GarlicAI-CLI/1.0")

 client := http.Client{Timeout: 30 * time.Second}
 resp, err := client.Do(req)
 if err != nil {
  // Bağlantı hatasının tam detayını ekrana basacak
  fmt.Printf("\n%s[GarlicAI Network Error]:%s %v\n\n", ColorPurple, ColorReset, err)
  os.Exit(1)
 }
 defer resp.Body.Close()

 body, _ := io.ReadAll(resp.Body)

 if resp.StatusCode == 200 {
  var resData ResponseBody
  _ = json.Unmarshal(body, &resData)
  fmt.Printf("\n%s[GarlicAI]:%s\n%s\n\n", ColorPurple, ColorReset, resData.Result)
 } else {
  // Sunucudan 200 dışında bir kod dönerse dönen ham cevabı basacak
  fmt.Printf("\n%s[GarlicAI Server Error (%d)]:%s %s\n\n", ColorPurple, ColorReset, resp.StatusCode, string(body))
 }
}
