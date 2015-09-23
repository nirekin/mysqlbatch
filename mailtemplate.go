// Go MySQL Batch
//
// Copyright 2015 Guillaume Barre. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.package main

package main

const emailTemplateHtmlHeader = `
{{if .}}
<table bgcolor="#E6E6E6" border="0">
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
</table>	
{{end}}
`

const emailTemplateHtmlBody = `
<table border="0" cellpadding="3">
	<tr bgcolor="#424242">
		<td><font color="white">BatchName:</font></td>
		<td><font color="white">{{.BatchName}}</font></td>
	</tr>
	{{range .URLResultList}}
		<tr bgcolor="#848484">
			<td>Base de datos:</td>
			<td>{{.GetUrlForMail}}</td>
		</tr>
		{{range $i, $e :=.QueryResultList}}
			<tr {{if $i | odd}} bgcolor="#424111" {{end}}>
				<td>Query</td>
				<td>
			{{range .GetInnerResultStack }}  
				
				<table>
				<tr>
					<td>{{.GetSpacing}}Name:{{.Name}} {{if .HasFailed}}<font color="red">{{end}}{{.Status}}{{if .HasFailed}}</font>{{end}}</td>
				</tr>
				{{if .HasError}}
				<tr>
					<td>{{.GetSpacing}}Expected:{{.QueryError.ExpectedValue}}, Received:<font color="red">{{.QueryError.ExecutionValue}}</font></td>
				</tr>
				{{end}}
				{{if .HasDescription}}
				<tr>
					<td>{{.GetSpacing}}Description:{{.Description}}</td>
				</tr>
				{{end}}
				{{if .HasMessage}}
				<tr>
					<td>{{.GetSpacing}}Message:{{.Message}}</td>
				</tr>
				{{end}}
				</table>
			{{end}}

				</td>
			</tr>


		{{end}}
		<tr>
			<td>&nbsp;</td>
			<td>&nbsp;</td>
		</tr>
	{{end}}
</table>
`
