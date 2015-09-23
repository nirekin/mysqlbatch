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

// A Sum Query allows you to calculate the sum of the desired column.
//
// Will be translated into a "select sum(FIELD) from TABLE WHERE" request.
//
// To create the query you need to specify, the field to sum, "from"
// database table, the complete "where" condition ( including "where" !)
// and the expected result.
type SumQuery struct {
	Field Field `xml:"field"`
	BaseQuery
}

// Returns the SQL correspinding to the query
func (o *SumQuery) getSQL() string {
	return fmt.Sprintf(SUM, o.Field.Content, o.From.Content, o.Where.Content)
}

// Return the expected result for the query
func (o *SumQuery) getExpectedResult() string {
	return o.ExceptedResult.Content

}

// Shows the query content once loaded from the config file
func (query *SumQuery) showLoadedQuery() {
	TraceLoaded.Printf("sum query: \n")
	TraceLoaded.Printf("query field: %s\n", query.Field)
	TraceLoaded.Printf("query from: %s\n", query.From)
	TraceLoaded.Printf("query where: %s\n", query.Where)
	TraceLoaded.Printf("query expected result: %s\n", query.ExceptedResult)
}
