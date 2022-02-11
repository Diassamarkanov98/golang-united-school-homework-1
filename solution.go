package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	fmt.Println(emoji.Sprint("Hello :world_map:!"))
}

GetMessage()
