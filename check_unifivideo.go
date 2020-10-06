/*-----------------------------------------------------------------------------
# Name:        CHECK_UNIFIVIDEO
# Purpose:     Nagios/Icinga checker for UniFi Video Controller condition
#
# Author:      Rafal Wilk <rw@pcboot.pl>
#
# Created:     30-09-2020
# Modified:    06-10-2020
# Copyright:   (c) PcBoot 2020
# License:     BSD-new
-----------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alexflint/go-arg"
)

var args struct {
	Address       string        `arg:"-a,--addr,required" help:"UVC address"`
	Port          int           `arg:"-p,--port" default:"7080" help:"UVC port"`
	SSL           bool          `arg:"-s,--ssl" default:"false" help:"use HTTPS instead of HTTP"`
	APIKey        string        `arg:"-k,--key,required" help:"UVC API Key"`
	WarningDelay  time.Duration `arg:"-w,--warning" default:"10m" help:"Warning status if last recording is older than this value"`
	CriticalDelay time.Duration `arg:"-c,--critical" default:"30m" help:"Critical status if last recording is older than this value"`
}

func main() {
	if err := arg.Parse(&args); err != nil {
		fmt.Println("CHECK_UNIFIVIDEO for UniFi Video Controller")
		fmt.Println("All rights reserved. (c) PcBoot 2020")
		fmt.Println()
		arg.MustParse(&args)
	}

	arCam := APIResponseCam{}
	if err := arCam.Get(); err != nil {
		handleErr(err)
	}

	var (
		countTotal    int
		countWarning  int
		countCritical int
		countOk       int
	)

	for _, d := range arCam.Data {
		if d.Managed {
			countTotal++
			var recAgo time.Duration
			recAgo = time.Now().Sub(convertToTime(d.LastRecordingStartTime)).Round(time.Second)
			if d.ShouldRecord() {
				if recAgo > args.CriticalDelay {
					fmt.Printf("%s (%s) - %s Last Rec: %s ago \t ShouldRecord?: %v - CRITICAL\n", d.Name, d.InternalHost, d.State, recAgo, d.ShouldRecord())
					countCritical++
					continue
				}
				if recAgo > args.WarningDelay {
					countWarning++
					fmt.Printf("%s (%s) - %s Last Rec: %s ago \t ShouldRecord?: %v - WARNING\n", d.Name, d.InternalHost, d.State, recAgo, d.ShouldRecord())
					continue
				}
			}
			countOk++

		}
	}

	fmt.Printf("\nTotal cameras: %d\tOK: %d\tCritical: %d\tWarning: %d", countTotal, countOk, countCritical, countWarning)

	if countCritical > 0 {
		fmt.Println("\t(Status CRITICAL)")
		os.Exit(2)
	}

	if countWarning > 0 {
		fmt.Println("\t(Status WARNING)")
		os.Exit(1)
	}

	fmt.Println("\t(Status OK)")

}

func convertToTime(i int64) time.Time {
	return time.Unix(0, i*int64(time.Millisecond))
}

func handleErr(err error) {
	panic(err)
}

func getURL(response APIResponse) string {
	proto := "http"
	if args.SSL {
		proto = "https"
	}

	return fmt.Sprintf("%s://%s:%d/api/2.0/%s?apiKey=%s", proto, args.Address, args.Port, response.APIResource(), args.APIKey)
}
