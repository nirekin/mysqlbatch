// Go MySQL Batch
//
// Copyright 2015 Guillaume Barre. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.package main

package main

import (
	"bytes"
	"html/template"
	"strings"
)

// The results for all batches, and if defined the author
// specifications
type GlobalResults struct {
	Author          *Author
	BatchResultList []*BatchResult
}

// The result for one batch
type BatchResult struct {
	BatchName     string
	Smtp          SmtpConfig
	Recipients    Recipients
	URLResultList []*UrlResult
}

// The result for one batch against on database
type UrlResult struct {
	URL             string
	QueryResultList []*QueryResult
}

// The result for one query and eventually its nested query
type QueryResult struct {
	Name             string
	Description      string
	Message          string
	Status           string
	InnerQueryResult *QueryResult
	InnerLevel       int
}

// Shows the result of the execution of all batches
func (r *GlobalResults) showContent() {
	TraceResult.Printf("----------------------------------------------------------------------------- \n")
	TraceResult.Printf("processing results: \n")
	TraceResult.Printf("----------------------------------------------------------------------------- \n")
	TraceResult.Printf("process results %d\n", len(r.BatchResultList))
	for _, re := range r.BatchResultList {
		re.showContent()
	}
}

// Shows the result of the execution of one batch
func (r *BatchResult) showContent() {
	TraceResult.Printf("result for %s:\n", r.BatchName)
	for _, ure := range r.URLResultList {
		ure.showContent()
	}
}

// Shows the result of the execution of one batch against
// one database
func (r *UrlResult) showContent() {
	TraceResult.Printf("result on %s:\n", r.URL)
	for _, qre := range r.QueryResultList {
		qre.showContent()
	}
}

// SHows the result of the execution of one query
// and eventually its nested query
func (r *QueryResult) showContent() {
	TraceResult.Printf("query name %s:\n", r.Name)
	TraceResult.Printf("query description %s:\n", r.Description)
	TraceResult.Printf("query status %s:\n", r.Status)
	TraceResult.Printf("query message %s:\n", r.Message)
	if r.InnerQueryResult != nil {
		r.InnerQueryResult.showContent()
	}
}

// Indicates if the query result has a description
func (r *QueryResult) HasDescription() bool {
	return r.Description != ""
}

// Indicates if the query result has a message
func (r *QueryResult) HasMessage() bool {
	return r.Message != ""
}

// Returns the trimed description
func (r *QueryResult) GetTrimedDescription() string {
	return strings.TrimSpace(r.Description)
}

// Adds the inner result to the slice received as parameter
func (r *QueryResult) addInnerResultStack(result *[]*QueryResult) {
	*result = append(*result, r)
	if r.InnerQueryResult != nil {
		r.InnerQueryResult.addInnerResultStack(result)
	}
}

// Retuns a stack of all query result  nested under this one
func (r *QueryResult) GetInnerResultStack() []*QueryResult {
	result := make([]*QueryResult, 0)
	result = append(result, r)
	if r.InnerQueryResult != nil {
		r.InnerQueryResult.addInnerResultStack(&result)
	}
	return result
}

// Sends email based on the result of all batches
func (r *GlobalResults) sendMail() {
	tHeader := template.Must(template.New("header").Parse(emailTemplateHtmlHeader))
	tBody := template.Must(template.New("body").Parse(emailTemplateHtmlBody))

	var docHeader bytes.Buffer
	tHeader.Execute(&docHeader, r.Author)

	for _, br := range r.BatchResultList {
		var docBody bytes.Buffer

		tBody.Execute(&docBody, br)

		TraceActivity.Printf("----------------------------------------------------------------------------- \n")
		TraceActivity.Printf("GeneratedCOntent : \n")
		TraceActivity.Printf("%s\n", docBody.String())
		TraceActivity.Printf("----------------------------------------------------------------------------- \n")

		var buffer bytes.Buffer
		buffer.WriteString(docHeader.String())
		buffer.WriteString(docBody.String())
		br.Smtp.sendEnailHtml(br, buffer.String())
	}
}

// Return the URL without the user and password
func (u *UrlResult) GetUrlForMail() string {
	idx := strings.Index(u.URL, "@")
	if idx > -1 {
		return "..." + u.URL[idx:len(u.URL)]
	} else {
		return u.URL
	}
}
