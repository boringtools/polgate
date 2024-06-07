package main

import (
	"embed"
	"fmt"

	"github.com/boringtools/polgate/cmd"
	"github.com/boringtools/polgate/pkg/common"
)

//go:embed internal/policy
var RegoPolicyDir embed.FS

var banner string = `
 _______  _______  ___      _______  _______  _______  _______ 
|       ||       ||   |    |       ||   _   ||       ||       |
|    _  ||   _   ||   |    |    ___||  |_|  ||_     _||    ___|
|   |_| ||  | |  ||   |    |   | __ |       |  |   |  |   |___ 
|    ___||  |_|  ||   |___ |   ||  ||       |  |   |  |    ___|
|   |    |       ||       ||   |_| ||   _   |  |   |  |   |___ 
|___|    |_______||_______||_______||__| |__|  |___|  |_______|
`

func main() {
	common.RegoPolicyDir = RegoPolicyDir
	fmt.Println(banner)
	cmd.Execute()
}
