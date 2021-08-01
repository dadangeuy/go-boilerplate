package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("cron: started at %s\n", time.Now().Format(time.RFC3339))
}
