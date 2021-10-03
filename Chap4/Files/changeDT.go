package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

/*
#	Match date abd time string of the log files of Apache web server
#	Change the date and time of the log files to diffrent formats
#	Will read the logs files of the apache log files
*/
func main(){
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide one text file to process!")
		os.Exit(1)
	}

	filename := args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s", err)
		os.Exit(1)
	}

	defer f.Close()

	notMatch := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}else if err != nil {
			fmt.Printf("error reading the file %s", err)
		}
		r1 := regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)
		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])

			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Println(strings.Replace(line, match[1], newFormat, 1))
			}else{
				notMatch++
			}
			continue
		}
	}
}