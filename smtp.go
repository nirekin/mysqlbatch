// Go MySQL Batch
//
// Copyright 2015 Guillaume Barre. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.package main

package main

import (
	//	"bytes"
	"fmt"
	"log"
	"net/smtp"
)

// SMTP configuration associated to a batch
type SmtpConfig struct {
	Server   string `xml:"server,attr"`
	Port     string `xml:"port,attr"`
	User     string `xml:"user,attr"`
	Password string `xml:"password,attr"`
	Sender   string `xml:"sender,attr"`
}

// Returns the sender of the generated email
func (s *SmtpConfig) getSender() string {
	if s.Sender != "" {
		return s.Sender
	} else {
		return s.User
	}
}

// Sends the content passed as parameter to the recipient list
// of the given batch result
func (s *SmtpConfig) sendEmailHtml(r *BatchResult, content string) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := fmt.Sprintf("Subject: MysqlBatch Results for the batch \"%s\"\n", r.BatchName)

	dest := make([]string, 0)
	for _, d := range r.Recipients.RecipientList {
		dest = append(dest, d.Address)
	}

	auth := smtp.PlainAuth(
		"",
		s.User,
		s.Password,
		s.Server,
	)
	err := smtp.SendMail(
		s.Server+":"+s.Port,
		auth,
		s.getSender(),
		dest,
		[]byte(subject+mime+content),
	)
	if err != nil {
		TraceActivity.Printf("Error running a query  : %s \n", err.Error())
		log.Fatal(err)
	}
}
