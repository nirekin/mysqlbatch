// Go MySQL Batch
//
// Copyright 2015 Guillaume Barre. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.package main

package main

import ()

// List of email recipients
type Recipients struct {
	RecipientList []*Recipient `xml:"recipient"`
}

// An email recipient
type Recipient struct {
	Address string `xml:"address,attr"`
}

// Shows the recipient list content once loaded from the config file
func (c *Recipients) showLoadedContent() {
	for _, rec := range c.RecipientList {
		TraceLoaded.Printf("recipient: %s\n", rec.Address)
	}
}
