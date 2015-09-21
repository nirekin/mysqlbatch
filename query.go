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

// Error returned by the executable queries.
// This error is used if an executable query fails.
type QueryError struct {
	prob           string
	expectedValue  int
	executionvalue int
}

type ExecutableQyery interface {
	getSQL() string
	getExpectedResult() int
}

func processQuery(q ExecutableQyery, dbC *sql.DB) (error, *QueryError) {
	if st, err := dbC.Prepare(q.getSQL()); err == nil {
		if rows, err := st.Query(); err == nil {
			var n int
			for rows.Next() {
				_ = rows.Scan(&n)
			}
			e := q.getExpectedResult()
			if n == e {
				return nil, nil
			} else {
				s := fmt.Sprintf("error expected %d got %d", e, n)
				return nil, &QueryError{s, e, n}
			}
		} else {
			return err, nil
		}
	} else {
		return err, nil
	}
	return nil, nil
}
