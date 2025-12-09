# giftapi
api for @giftfest_bot telegram bot The biggest gift festival on Telegram! Every day — new gifts, fun tasks, thousands of prizes, and tons of fun! 
# main
```go
package main
import (
	"fmt"
	"log"
)

func main() {
	token := ""
	
	result, err := Auth(token)
	if err != nil {
		log.Fatal("error in auth:", err)
	}
    fmt.Printf("%+v\n", result)
	req, err := GetInventoryResources()
	if err != nil {
		log.Fatal("error in Get Inventory Resources:", err)
	}
   fmt.Printf("%+v\n", req)
}
```

# Launch (your script)
```
go run giftapi.go main.go
```
