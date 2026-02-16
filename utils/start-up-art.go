package utils

import (
	"fmt"
	"strconv"
	"time"
	"net"
	"github.com/fatih/color"
)

var (
	green = color.RGB(0, 255, 0).SprintFunc()
	blue  = color.RGB(70, 120, 255).SprintFunc()
)

const letterDelay time.Duration = 10

func DisplayStartupVisuals(port int) {

	line1 := "\n--------------------   Wifi File Transfer   --------------------\n"
	line2 := "________________________________________________________________\n"

	// for _, char := range (line1 + line2) {
	// 	fmt.Printf("%c", char)
	// 	time.Sleep(letterDelay * time.Millisecond)
	// }

	time.Sleep(100 * time.Millisecond)

	fmt.Print(line1)

	time.Sleep(100 * time.Millisecond)

	fmt.Print(line2)

	fmt.Println()

	animateInfo("127.0.0.1", "Running locally at:  ", port)
	animateInfo(GetLocalIP(), "Accessible on LAN at:", port)
	
	fmt.Println()
}

func animateInfo(ip string, msg string, port int) {

	wordInfo := "INFO"
	wordUrl := "http://" + ip + ":" + strconv.Itoa(port) + "\n"

	for _, char := range (wordInfo) {
		fmt.Printf("%s", green(string(char)))
		time.Sleep(letterDelay * time.Millisecond)
	}

	for _, char := range ("   " + msg + "   ") {
		fmt.Printf("%c", char)
		time.Sleep(letterDelay * time.Millisecond)
	}

	for _, char := range (wordUrl) {
		fmt.Printf("%s", blue(string(char)))
		time.Sleep(letterDelay * time.Millisecond)
	}
}

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// fmt.Printf("\n%s   App running on:   %s%s\n\n",
	// 	green("INFO"),
	// 	blue("http://127.0.0.1:"),
	// 	blue(port),
	// )

/*
░█░█░▀█▀░█▀▀░▀█▀░░░█▀▀░▀█▀░█░░░█▀▀░░░▀█▀░█▀▄░█▀█░█▀█░█▀▀░█▀▀░█▀▀░█▀▄
░█▄█░░█░░█▀▀░░█░░░░█▀▀░░█░░█░░░█▀▀░░░░█░░█▀▄░█▀█░█░█░▀▀█░█▀▀░█▀▀░█▀▄
░▀░▀░▀▀▀░▀░░░▀▀▀░░░▀░░░▀▀▀░▀▀▀░▀▀▀░░░░▀░░▀░▀░▀░▀░▀░▀░▀▀▀░▀░░░▀▀▀░▀░▀
*/
