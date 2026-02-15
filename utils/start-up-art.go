package utils

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	green = color.RGB(0, 255, 0).SprintFunc()
	blue  = color.RGB(70, 120, 255).SprintFunc()
)

func DisplayAsciiLogo(port int) {
	fmt.Println(`
---------------------   Wifi File Transfer   -----------------------
____________________________________________________________________
	`)
	fmt.Printf("%s   App running on:   %s%s\n\n",
		green("INFO"),
		blue("http://127.0.0.1:"),
		blue(port),
	)
}

/*
░█░█░▀█▀░█▀▀░▀█▀░░░█▀▀░▀█▀░█░░░█▀▀░░░▀█▀░█▀▄░█▀█░█▀█░█▀▀░█▀▀░█▀▀░█▀▄
░█▄█░░█░░█▀▀░░█░░░░█▀▀░░█░░█░░░█▀▀░░░░█░░█▀▄░█▀█░█░█░▀▀█░█▀▀░█▀▀░█▀▄
░▀░▀░▀▀▀░▀░░░▀▀▀░░░▀░░░▀▀▀░▀▀▀░▀▀▀░░░░▀░░▀░▀░▀░▀░▀░▀░▀▀▀░▀░░░▀▀▀░▀░▀
*/
