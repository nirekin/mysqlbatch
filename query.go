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
	"fmt"
)

// SQL Sentences used by the executable queries
const (
	SELECT_COUNT_ALL      = "SELECT COUNT(*) FROM %v %v"
	SELECT_COUNT_DISTINCT = "SELECT COUNT(DISTINCT %v) FROM %v %v"
	SUM                   = "SELECT SUM(%v) FROM %v %v"
	AVG                   = "SELECT AVG(%v) FROM %v %v"
	MAX                   = "SELECT MAX(%v) FROM %v %v"
	MIN                   = "SELECT MIN(%v) FROM %v %v"
)

// Database field used in executable queries
type Field struct {
	Content string `xml:",chardata"`
}

// From clause used into executable queries
type From struct {
	Content string `xml:",chardata"`
}

// Where clause used into executable queries
type Where struct {
	Content string `xml:",chardata"`
}

// Result expected by the executable queries
type ExceptedResult struct {
	Content string `xml:",chardata"`
}

// To create a query you need to specify, the "from" database table, the
// complete "where" condition ( including "where" !) and the expected result.
type BaseQuery struct {
	From           From           `xml:"from"`
	Where          Where          `xml:"where"`
	ExceptedResult ExceptedResult `xml:"expected_result"`
}

// Error returned by the executable queries.
// This error is used if an executable query fails.
type QueryError struct {
	prob           string
	ExpectedValue  string
	ExecutionValue string
}

type ExecutableQyery interface {
	getSQL() string
	getExpectedResult() string
}

// Shows the content of the error
func (r *QueryError) showContent() {
	TraceResult.Printf("query error %s:\n", r.prob)
	TraceResult.Printf("query error expectedValue %s:\n", r.ExpectedValue)
	TraceResult.Printf("query error executionValue %s:\n", r.ExecutionValue)
}

func processQuery(q ExecutableQyery, dbC *sql.DB) (error, *QueryError) {
	s := q.getSQL()
	TraceActivity.Println("execute query : %s", s)
	if st, err := dbC.Prepare(s); err == nil {
		if rows, err := st.Query(); err == nil {
			var n string
			for rows.Next() {
				_ = rows.Scan(&n)
			}
			e := q.getExpectedResult()
			if n == e {
				return nil, nil
			} else {
				s := fmt.Sprintf("error expected %d got %d", e, n)
				r := &QueryError{}
				r.prob = s
				r.ExpectedValue = e
				r.ExecutionValue = n
				return nil, r
			}
		} else {
			return err, nil
		}
	} else {
		return err, nil
	}
	return nil, nil
}
