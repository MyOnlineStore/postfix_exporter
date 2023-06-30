package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	readFile, err := os.Open("mail.log")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var (
		logLine = regexp.MustCompile(` ?(postfix|opendkim)(.*(\w+))?\[\d+\]: ((?:(warning|error|fatal|panic): )?.*)`)
	)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		logMatches := logLine.FindStringSubmatch(line)
		if logMatches == nil {

		} else {
			process := logMatches[1]
			regexed := regexp.MustCompile(`(.*)/(smtpd)`)
			switch process {
			case "postfix":
				// Group patterns to check by Postfix service.
				subprocess := strings.TrimPrefix(logMatches[2], "/")
				switch {
				case subprocess == "cleanup":
					fmt.Println("Cleanup")
				case regexed.MatchString(subprocess):
					fmt.Println("SMTPD BITCHES")
				default:
					fmt.Println(subprocess)
				}
			}
		}

	}

	line := "Jun 30 06:14:18 mail postfix/bla/postscreen[138]: PASS OLD [107.172.154.9]:41388"
	logMatches := logLine.FindStringSubmatch(line)

	process := logMatches[1]
	level := logMatches[5]
	remainder := logMatches[4]
	switch process {
	case "postfix":
		// Group patterns to check by Postfix service.
		subprocess := logMatches[2]
		subprocess = strings.TrimPrefix(subprocess, "/")
		fmt.Println(subprocess)
	}
	fmt.Println(process)
	fmt.Println(level)
	fmt.Println(remainder)
}
