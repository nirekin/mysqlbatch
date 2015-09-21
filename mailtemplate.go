// Go MySQL Batch
//
// Copyright 2015 Guillaume Barre. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.package main

package main

const emailTemplateHtmlHeader = `
<table>
{{if .}}
	<tr>
		<td width="30%">Author:</td>
		<td width="70%">{{.Name}}</td>
	</tr>
	<tr>
		<td>Organization:</td>
		<td>{{.Organization}}</td>
	</tr>
	<tr>
		<td>Phone:</td>
		<td>{{.Phone}}</td>
	</tr>
	<tr>
		<td>Email:</td>
		<td>{{.Email}}</td>
	</tr>	
{{end}}
`
const emailTemplateHtmlBody = `
	<tr>
		<td width="30%">BatchName:</td>
		<td width="70%">{{.BatchName}}</td>
	</tr>
	{{range .URLResultList}}
		<tr>
			<td>Base de datos:</td>
			<td>{{.URL}}</td>
		</tr>
		{{range .QueryResultList}}
			{{range .GetInnerResultStack }}  
				<tr>
					<td>Query name:</td>
					<td>{{.InnerLevel}} - {{.Name}}</td>
				</tr>
				<tr>
					<td>Status:</td>
					<td>{{.Status}} </td>
				</tr>
						<tr>
							<td>Description:</td>
							<td>{{.Description}}</td>
						</tr>
					<tr>
						<td>Message:</td>
						<td>{{.Message}}</td>
					</tr>
			{{end}}
		{{end}}
	{{end}}
</table>
`
