// cmd/nftmarketplaceengineplatformx/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplaceengineplatformx/internal/nftmarketplaceengineplatformx"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplaceengineplatformx.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
