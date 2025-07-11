// Command earthsign is a command-line implementation of package earthsign.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dkmccandless/earthsign"
)

// dateTime is the same as time.DateTime
// except that it permits single-digit months and days.
const dateTime = "2006-1-2 15:04:05"

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Println("Usage: earthsign lat lon YYYY-MM-DD HH:MM:SS")
		fmt.Println("Latitude and longitude in degrees. North and east are positive.")
		fmt.Println("Time in UTC.")
	}
	flag.Parse()

	if len(os.Args) != 5 {
		flag.Usage()
		if len(os.Args) >= 2 && os.Args[1] == "help" {
			os.Exit(0)
		}
		os.Exit(2)
	}

	lat, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal(errors.New("invalid latitude"))
	}
	lon, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal(errors.New("invalid longitude"))
	}
	t, err := time.Parse(dateTime, os.Args[3]+" "+os.Args[4])
	if err != nil {
		log.Fatal(err)
	}

	name, err := earthsign.At(lat, lon, t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
