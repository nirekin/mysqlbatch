// Go MySQL Batch
//
// Copyright 2015 Guillaume Barre. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.package main

package main

import (
	"fmt"
)

// A Count Query allows you to count rows matching the given where condition.
//
// Will be translated into a "select count(*) from TABLE WHERE" request.
//
// To create the query you need to specify, the "from" database table, the
// complete "where" condition ( including "where" !) and the expected result.
type CountQuery struct {
	BaseQuery
}

// Returns the SQL correspinding to the query
func (o *CountQuery) getSQL() string {
	return fmt.Sprintf(SELECT_COUNT_ALL, o.From.Content, o.Where.Content)
}

// Return the expected result for the query
func (o *CountQuery) getExpectedResult() string {
	return o.ExceptedResult.Content
}

// Return the operator to apply on expected result of the query
func (o *CountQuery) getOperator() string {
	if o.ExceptedResult.Operator == "" {
		return OPERATOR_EQ
	} else {
		return o.ExceptedResult.Operator
	}
}

// Shows the query content once loaded from the config file
func (query *CountQuery) showLoadedQuery() {
	TraceLoaded.Printf("count query: \n")
	TraceLoaded.Printf("query from: %s\n", query.From)
	TraceLoaded.Printf("query where: %s\n", query.Where)
	TraceLoaded.Printf("query expected result: %s\n", query.ExceptedResult)
	TraceLoaded.Printf("query expected result operator: %s\n", query.ExceptedResult.Operator)
}
