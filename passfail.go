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
	"strings"
)

// Step associated to the success or the failure of an
// executable query.
//
// The step can contains a message. If defined the message will be
// automatically included into the generated report.
//
// The wrapper can also include a nested query.
//
// To produce a cleaner an more readable report you can give your
// step a name.
//
type FailPass struct {
	Message Message       `xml:"message"`
	Name    string        `xml:"name,attr"`
	Query   *QueryWrapper `xml:"query"`
}

// Shows the step content once loaded from the config file
func (c *FailPass) showLoadedContent() {

	if c.Name != "" {
		TraceLoaded.Printf("name: %s\n", c.Name)
	}

	if c.Message.Content != "" {
		s := strings.TrimSpace(c.Message.Content)
		TraceLoaded.Println(s)
	}

	if c.Query != nil {
		pq := c.Query
		TraceLoaded.Printf("Query:\n")
		pq.showLoadedQuery()
	}
}

// Executes the inner query associated with this step.
// THe result given in paramter should be the "InnerQueryResult"
// of the parent query
func (c *FailPass) launchInnerQuery(dbC *sql.DB, result *QueryResult) {
	c.Query.runInnerOnce(dbC, result)
}
