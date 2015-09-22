// Go MySQL Batch
//
// Copyright 2015 Guillaume Barre. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.package main

package main

import (
	"database/sql"
)

// The list of databases associated to a batch
type DBList struct {
	DBList []*DB `xml:"mysql"`
}

// URL used to connect to a database
type DB struct {
	URL string `xml:"url,attr"`
}

// SQL Sentences used by the executable queries
const (
	SELECT_COUNT_ALL      = "SELECT COUNT(*) FROM %v %v"
	SELECT_COUNT_DISTINCT = "SELECT COUNT DISTINCT(%v) FROM %v %v"
	SUM                   = "SELECT SUM(%v) FROM %v %v"
	AVG                   = "SELECT AVG(%v) FROM %v %v"
	MAX                   = "SELECT MAX(%v) FROM %v %v"
	MIN                   = "SELECT MIN(%v) FROM %v %v"
)

// Executes the batch received as parameter against a given database
// and send the populate the result to the channel
func runUrlOnce(ch chan UrlResult, d *DB, batch *Batch) {
	r := UrlResult{}
	r.URL = d.URL

	dbC, _ := sql.Open("mysql", d.URL)
	defer dbC.Close()

	for _, query := range batch.Queries.QueryList {
		query.runOnce(dbC, &r)
	}
	ch <- r
}

func (c *DBList) showLoadedContent() {
	for _, url := range c.DBList {
		TraceLoaded.Printf("url: %s\n", url.URL)
	}
}
