// Slow is a simple program to load stdin into a buffer and slowly output it to stdout
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Version is replaced with a commit hash when compiled using GNU make.
var VERSION = "undefined" // make(1) will change this.
var delay = time.Microsecond * 300
var err error

// MaxScanTokenSize is the max bytes allowed in one line.
const MaxScanTokenSize = 128 * 4906

var helptext = `
NAME
	slow, version ` + VERSION + ` - Non-interactive stdout delay.
	Pipe fast things (such as lsmod or dmesg) into it.

AUTHOR
	Copyright (c) 2016 aerth [aerth@sdf.org]

SYNOPSIS
	[FILE] | slow -d 300

USAGE
	lsmod | slow 
	tree / | slow
	
TIPS
	Custom Delay: cat main.go | slow -d 100
	Delay is in microseconds. Default is 300. 
	The exception to '-d' flag is setting to 1-5 which are in seconds.
`

func showversion() {
	fmt.Println(os.Args[0], "Version #", VERSION)
	os.Exit(0)
}

func showhelp() {
	fmt.Println(helptext)
	os.Exit(0)
}

func main() {
	delay = time.Second / 3
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v":
			showversion()
		case "-h":
			showhelp()
		case "-d":
			switch os.Args[2] {
			case "1":
				delay = time.Second * 1
			case "2":
				delay = time.Second * 2
			case "3":
				delay = time.Second * 3
			case "4":
				delay = time.Second * 4
			case "5":
				delay = time.Second * 5
			default:
				delay, err = time.ParseDuration(os.Args[2] + "ms")
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					showhelp()
					os.Exit(1)
				}
			}
		default:
			showhelp()

		}
	}

	// Slow stdin with delay
	slow(delay)

}

func slow(delay time.Duration) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Printf(s)
		fmt.Printf("\n")
		time.Sleep(delay)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
