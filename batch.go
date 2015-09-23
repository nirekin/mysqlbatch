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
	"time"
)

// List of batches to execute
type Batches struct {
	XMLName   xml.Name `xml:"batches"`
	BatchList []*Batch `xml:"batch"`
	Author    *Author  `xml:"author"`
}

// A batch to excute.
//
// A batch must have an smtp configuration, a list of databases
// and a list of query.
//
// If you give a name to the batch this it will be used into the
// generated report.

type Batch struct {
	Name       string     `xml:"name,attr"`
	Smtp       SmtpConfig `xml:"smtp_config"`
	DBs        DBList     `xml:"mysqls"`
	Queries    QueryList  `xml:"queries"`
	Recipients Recipients `xml:"recipients"`
}

// The author of the batches.
// If defined the author information will appear into the
// generated email
type Author struct {
	Name         string `xml:"name,attr"`
	Organization string `xml:"organization,attr"`
	Phone        string `xml:"phone,attr"`
	Email        string `xml:"email,attr"`
}

// Runs the batch and send to the channel the populated result
func runBatchOnce(ch chan BatchResult, b *Batch) {

	chanUrl := make(chan UrlResult, len(b.DBs.DBList))
	r := BatchResult{}
	r.BatchName = b.Name
	r.Recipients = b.Recipients
	r.Smtp = b.Smtp

	for _, url := range b.DBs.DBList {
		TraceActivity.Printf("Launched Url : %s \n", url.URL)
		go runUrlOnce(chanUrl, url, b)
	}

	for i := 0; i < len(b.DBs.DBList); i++ {
		resp := <-chanUrl
		TraceActivity.Printf("received Url : %s \n", resp.URL)
		r.URLResultList = append(r.URLResultList, &resp)
	}
	ch <- r
}

// Runs all batches
func (c *Batches) runBatchOnce() {
	start := time.Now()
	chanBatch := make(chan BatchResult, len(c.BatchList))
	for _, batch := range c.BatchList {
		TraceActivity.Printf("Launched Batch : %s \n", batch.Name)
		go runBatchOnce(chanBatch, batch)
	}
	globalResults := GlobalResults{}
	globalResults.Author = c.Author

	for i := 0; i < len(c.BatchList); i++ {
		resp := <-chanBatch
		TraceActivity.Printf("Received Batch : %s \n", resp.BatchName)
		globalResults.BatchResultList = append(globalResults.BatchResultList, &resp)
	}

	r := globalResults
	elapsed := time.Since(start)
	TraceActivity.Printf("Execution time : %dms \n", elapsed.Nanoseconds()/int64(time.Millisecond))

	if bResult {
		TraceActivity.Printf("----------------------------------------------------------------------------- \n")
		TraceActivity.Printf("Showing the result content: \n")
		TraceActivity.Printf("----------------------------------------------------------------------------- \n")
		r.showContent()
	}

	if bEmail {
		TraceActivity.Printf("----------------------------------------------------------------------------- \n")
		TraceActivity.Printf("Sending the email: \n")
		TraceActivity.Printf("----------------------------------------------------------------------------- \n")
		r.sendMail()
	}
}

// Shows the batches content once loaded from the config file
func (c *Batches) showLoadedContent() {
	TraceLoaded.Printf("----------------------------------------------------------------------------- \n")
	TraceLoaded.Printf("loaded batches: \n")
	TraceLoaded.Printf("----------------------------------------------------------------------------- \n")
	if c.Author != nil {
		TraceLoaded.Printf("Author: \n")
		TraceLoaded.Printf("name: %s\n", c.Author.Name)
		TraceLoaded.Printf("organisation: %s\n", c.Author.Organization)
		TraceLoaded.Printf("phone: %s\n", c.Author.Phone)
		TraceLoaded.Printf("email: %s\n", c.Author.Email)
	}

	for _, batch := range c.BatchList {
		TraceLoaded.Printf("batch name: %s\n", batch.Name)
		TraceLoaded.Printf("smtp: server '%s', port '%s', user '%s', password '%s', sender '%s'\n", batch.Smtp.Server, batch.Smtp.Port, batch.Smtp.User, batch.Smtp.Password, batch.Smtp.Sender)
		batch.Recipients.showLoadedContent()
		batch.DBs.showLoadedContent()
		for _, query := range batch.Queries.QueryList {
			query.showLoadedQuery()
		}
	}
}
