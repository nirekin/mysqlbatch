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
	"log"
	"strings"
)

// The list of query wrappers associated to one batch
type QueryList struct {
	QueryList []*QueryWrapper `xml:"query"`
}

// A query wrapper must contain a real executable query, for example
// a CountQuery or aCountDistinctQuery.
//
// The wrapper can also include to steps which could be executed
// if the executable query fails or not.
//
// To produce a cleaner an more readable report you can give your
// wrapper a name or even a decription.
type QueryWrapper struct {
	Name               string              `xml:"name,attr"`
	Description        *Message            `xml:"description"`
	OnFail             *FailPass           `xml:"fail"`
	OnPass             *FailPass           `xml:"pass"`
	CountQuery         *CountQuery         `xml:"count_query"`
	CountDistinctQuery *CountDistinctQuery `xml:"count_distinct_query"`
	SumQuery           *SumQuery           `xml:"sum_query"`
	AverageQuery       *AverageQuery       `xml:"average_query"`
	MinQuery           *MinQuery           `xml:"min_query"`
	MaxQuery           *MaxQuery           `xml:"max_query"`
}

// Execute the step associated to the success of the execution
// of the executable query
func (w *QueryWrapper) onPass(dbC *sql.DB, result *QueryResult) {
	result.Status = "PASSED"
	if w.OnPass != nil {
		if w.OnPass.Message.Content != "" {
			s := strings.TrimSpace(w.OnPass.Message.Content)
			result.Message = s
		}

		if w.OnPass.Query != nil {
			qr := QueryResult{}
			qr.InnerLevel = result.InnerLevel + 1
			result.InnerQueryResult = &qr
			w.OnPass.launchInnerQuery(dbC, &qr)
		}
	}
}

/// Execute the step associated to the failure of the execution
// of the executable query
func (w *QueryWrapper) onFail(dbC *sql.DB, result *QueryResult) {
	result.Status = "FAILED"
	if w.OnFail != nil {
		if w.OnFail.Message.Content != "" {
			s := strings.TrimSpace(w.OnFail.Message.Content)
			result.Message = s
		}

		if w.OnFail.Query != nil {
			qr := QueryResult{}
			qr.InnerLevel = result.InnerLevel + 1
			result.InnerQueryResult = &qr
			w.OnFail.launchInnerQuery(dbC, &qr)
		}
	}
}

// Executes the wrapper content.
// If required the execution of the query will populate
// the given result.
func (w *QueryWrapper) runOnce(dbC *sql.DB, result *UrlResult) {
	qr := QueryResult{}
	qr.Name = w.Name
	qr.InnerLevel = 1
	if w.Description != nil {
		qr.Description = w.Description.Content
	} else {
		qr.Description = ""
	}
	q := &qr
	result.QueryResultList = append(result.QueryResultList, q)
	launchQuery(w, dbC, q)
}

// Executes the wrapper content.
// If required the execution of the query will populate
// the given result.
func (w *QueryWrapper) runInnerOnce(dbC *sql.DB, result *QueryResult) {
	result.Name = w.Name
	if w.Description != nil {
		result.Description = w.Description.Content
	} else {
		result.Description = ""
	}
	launchQuery(w, dbC, result)
}

// Runs the executable query associated to the wrapper
// If required the execution of the query will populate
// the given result.
func launchQuery(w *QueryWrapper, dbC *sql.DB, q *QueryResult) {
	if w.CountQuery != nil {
		e, queryError := processQuery(w.CountQuery, dbC)
		processPassFail(w, dbC, q, e, queryError)
	} else if w.CountDistinctQuery != nil {
		e, queryError := processQuery(w.CountDistinctQuery, dbC)
		processPassFail(w, dbC, q, e, queryError)
	} else if w.SumQuery != nil {
		e, queryError := processQuery(w.SumQuery, dbC)
		processPassFail(w, dbC, q, e, queryError)
	} else if w.AverageQuery != nil {
		e, queryError := processQuery(w.AverageQuery, dbC)
		processPassFail(w, dbC, q, e, queryError)
	} else if w.MinQuery != nil {
		e, queryError := processQuery(w.MinQuery, dbC)
		processPassFail(w, dbC, q, e, queryError)
	} else if w.MaxQuery != nil {
		e, queryError := processQuery(w.MaxQuery, dbC)
		processPassFail(w, dbC, q, e, queryError)
	}
}

// Process the result on the execution of an executable query.
// This will populate the result received as parameter and if
// required it will also lauch the inner request corresponding
// to the status of the execution
func processPassFail(w *QueryWrapper, dbC *sql.DB, q *QueryResult, e error, qe *QueryError) {
	if e != nil {
		TraceActivity.Printf("Error running a query  : %s \n", e.Error())
		log.Fatal(e)
	} else {
		if qe == nil {
			w.onPass(dbC, q)
		} else {
			q.QueryError = qe
			w.onFail(dbC, q)
		}
	}
}

// Shows the wrapper content once loaded from the config file
func (w *QueryWrapper) showLoadedQuery() {
	if w.Name != "" {
		TraceLoaded.Printf("wrapper name: %s\n", w.Name)
	}
	if w.Description != nil {
		if w.Description.Content != "" {
			TraceLoaded.Printf("wrapper description: %s\n", w.Description.Content)
		}
	}

	if w.CountQuery != nil {
		w.CountQuery.showLoadedQuery()
	} else if w.CountDistinctQuery != nil {
		w.CountDistinctQuery.showLoadedQuery()
	} else if w.SumQuery != nil {
		w.SumQuery.showLoadedQuery()
	} else if w.AverageQuery != nil {
		w.AverageQuery.showLoadedQuery()
	} else if w.MinQuery != nil {
		w.MinQuery.showLoadedQuery()
	} else if w.MaxQuery != nil {
		w.MaxQuery.showLoadedQuery()
	}

	if w.OnPass != nil {
		TraceLoaded.Printf("%s: wrapper pass: \n", w.Name)
		w.OnPass.showLoadedContent()
	}

	if w.OnFail != nil {
		TraceLoaded.Printf("%s: wrapper fail: \n", w.Name)
		w.OnFail.showLoadedContent()
	}
}
