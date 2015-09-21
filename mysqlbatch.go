// Go MySQL Batch
//
// Copyright 2015 Guillaume Barre. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.package main

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"log"
)

// Message used to store various string content
type Message struct {
	Content string `xml:",chardata"`
}

var (
	TraceLoaded   *log.Logger
	TraceResult   *log.Logger
	TraceActivity *log.Logger

	bLoad     bool
	bActivity bool
	bResult   bool
	bEmail    bool
)

func Init() {

	l := &lumberjack.Logger{
		Filename:   "MySqlBatch.log",
		MaxSize:    250, // mb
		MaxBackups: 5,
		MaxAge:     10, // in days
	}

	TraceLoaded = log.New(l, "Load: ", log.Ldate|log.Ltime|log.Lshortfile)
	TraceResult = log.New(l, "Result: ", log.Ldate|log.Ltime|log.Lshortfile)

	if bActivity {
		TraceActivity = log.New(l, "Activity: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		// If the log of activity is not requested using the flag the we redirect all on ioutil.Discard
		TraceActivity = log.New(ioutil.Discard, "Activity: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

// Main function to launch the script engine
//
// This function must receive as parameter the name of the xml file used to configure the baches
func main() {

	fileFlag := flag.String("file", "batch.xml", "the configuration file")
	flag.BoolVar(&bLoad, "showload", false, "log the loaded content in the trace file")
	flag.BoolVar(&bActivity, "showactivity", false, "log the activity in the trace file")
	flag.BoolVar(&bResult, "showresult", false, "log the execution report in the trace file")
	flag.BoolVar(&bEmail, "email", false, "send the execution report email")
	flag.Parse()

	Init()

	xmlContent, err := ioutil.ReadFile(*fileFlag)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	root := Batches{}

	// Read the xml file
	err = xml.Unmarshal(xmlContent, &root)
	check(err)

	// (showload) show the loaded content
	if bLoad {
		root.showLoadedContent()
	}

	// (debug) run all batches
	root.runBatchOnce()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
