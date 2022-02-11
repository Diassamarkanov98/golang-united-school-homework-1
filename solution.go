package solution
package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	fmt.Println(emoji.Sprint("Hello :world_map:!"))
}

func main() {
	GetMessage()
}
