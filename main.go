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

// BURAYI KENDİ RENDER URL'İN İLE DEĞİŞTİR:
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
 ColorGreen = "\033[1;32m"
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

 client := http.Client{Timeout: 30 * time.Second}
 resp, err := client.Post(proxyURL, "application/json", bytes.NewBuffer(jsonData))
 if err != nil {
  fmt.Println("\n" + ColorPurple + "[GarlicAI]: " + ColorReset + "Connection Error! Could not reach the server.")
  os.Exit(1)
 }
 defer resp.Body.Close()

 body, _ := io.ReadAll(resp.Body)

 var resData ResponseBody
 _ = json.Unmarshal(body, &resData)

 if resp.StatusCode == 200 {
  fmt.Printf("\n%s[GarlicAI]:%s\n%s\n\n", ColorPurple, ColorReset, resData.Result)
 } else {
  fmt.Printf("\n%s[GarlicAI Error]:%s Server failed to respond (Code: %d)\n\n", ColorPurple, ColorReset, resp.StatusCode)
 }
}
