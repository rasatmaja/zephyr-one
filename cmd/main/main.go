package main

import (
	"fmt"

	"github.com/rasatmaja/zephyr-one/internal/server"
)

func main() {
	fmt.Println(`
  ______             _                     ____               
 |___  /            | |                   / __ \              
    / /  ___  _ __  | |__   _   _  _ __  | |  | | _ __    ___ 
   / /  / _ \| '_ \ | '_ \ | | | || '__| | |  | || '_ \  / _ \
  / /__|  __/| |_) || | | || |_| || |    | |__| || | | ||  __/
 /_____|\___|| .__/ |_| |_| \__, ||_|     \____/ |_| |_| \___|
             | |             __/ |                            
             |_|            |___/                             

	`)
	server.Start()
}
