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

// Executes the batch received as parameter against a given database
// and send the populate the result to the channel
func runUrlOnce(ch chan UrlResult, d *DB, batch *Batch) {
	TraceActivity.Printf("runs :%s on %s\n", batch.Name, d.URL)
	r := UrlResult{}
	r.URL = d.URL

	dbC, _ := sql.Open("mysql", d.URL)
	defer dbC.Close()

	for _, query := range batch.Queries.QueryList {
		query.runOnce(dbC, &r)
	}
	ch <- r
}

// Shows the url to reach the database once loaded from the config file
func (c *DBList) showLoadedContent() {
	for _, url := range c.DBList {
		TraceLoaded.Printf("url: %s\n", url.URL)
	}
}
