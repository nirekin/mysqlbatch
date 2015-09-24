# MySqlBatch

MySqlBatch is a tool allowing you to :

>  * Execute batches of predefined simple SQL queries.
>  * Target predefined list of databases within each bacth.
>  * Send a report by email for each batch.
>  * Each query can be associated to a `pass` or/and `fail`status which can also contain a query and so on...




By default MySqlBatch will look, on its own directory, for a file named "batch.xml" but you can specify another one using `mysqlbatch -file foo.xml`

In order to automatically send the generated emails you must run MySqlBatch with the `-email` option: `mysqlbatch -file foo.xml -email`


MySqlBatch consume an XML file containing the specification of all the batches to run. 
This files must be valid against this [DTD](https://raw.githubusercontent.com/nirekin/mysqlbatch/master/mysqlbatch.dtd)



## `<batches>`
The root node of the configuration file is `<batches>`

`<batches>` can contain an optional `<author>` and must contain at least one `<batch>`



### `<author>`
    <author name="jack walsh" organization="YourOrganization" phone="987987987" email="jackwalsh@yourorganization.com" />

The author tag is optional, it allows you to give information about the person responsible of the tool execution and its generated emails.

If defined the content of the author tag will be included into the generated emails.



###  `<batch>`
    <batch name="name_batch1">

A batch is basically the unit of work of MySqlBatch. You can use a batch to specify a list of queries to execute on several databases in order to send a report to identified recipients.

The name of the batch is required, it's used to build the subject of the generated emails.



### `<smtp_config>`
    <smtp_config server="smtp.mail.yahoo.com" port="25" user="jack.walsh" password="987987987" sender="jack.walsh@yourorganization.com"/>

The smtp_config tag is required, it defines the smtp server/account to send the generated email for a specific batch.

*The* `sender` *attribute of the smtp_config can be omitted, in this case the `sender` field of the generated email will be filled with the content of the `user` attribute.*



### `<recipients>`
    <recipients>
      <recipient address="jonathanmardukas@yourorganization.com" />
      <recipient address="alonzomosely@yourorganization.com" />
    </recipients>

The recipients tag is required, it allows you to define the list of recipients of the generated email for a specific batch.



### `<mysqls>`
    <mysqls>
      <mysql url="username:password@tcp(10.10.10.10:3306)/your_schema" />
      <mysql url="username:password@tcp(10.10.10.11:3306)/your_schema" />
    </mysqls>
    
The mysqls tag is required, it allows you to specify the list of all databases targeted of the execution of a given batch.

To access MySql MySqlBatch use the following url pattern: `username:password@tcp(10.10.10.10:3306)/your_schema`


Where :

  * "username" is the user login to access the database
  * "password" is the password to access the database
  * "10.10.10.10" is the IP address of the MySql server to reach
  * "3306" is the port of the MySql server to reach
  * "your_schema" is the schema target of the batch execution   

*The "username" and the "password" used to access the database are not shown into the generated emails. In order to identify the databases the emails will contain only the IP address, the port and the targeted schema.*



## `<queries>`

The tag `<queries>` lets you specify the list of queries associated to a batch.



### `<query>`

This tag allows to define a query within the list of queries associated to a batch.

    <query name="first query">
    	<count_query>
    		<from>clients</from>
    		<where>where active = "1"</where>
    		<expected_result>666</expected_result>
    	</count_query>
	    <description>Count active clients</description>
    	<pass>
    		<message>no error counting the active clients</message>
    	</pass>
    	<fail>
    		<message>error counting the active clients</message>
    	</fail>
    </query>



A query must have a `name` attribute, it will be included into the generated emails to identify the queries.

    <query name="first query">

In addition a query must specify one type among those ones :



##### `<min_query>`,`<max_query>`,`<average_query>`,`<sum_query>`,`<count_distinct_query>`,`<count_query>`

Currently you can use only those 6 types of queries within a batch:



#####`<description>`

The description of the query is optional, if defined it will be included into the generated emails.



#####`<pass>` and/or `<fail>`

A query can be associated to one or two status, the `pass` status will be processed if the expected result matches the execution result of the query and the `fail` status will be processed if it doesn't.



### Definition of all query types


#####`<field>field_name</field>`
For some queries this tag will be required, it allows you to specify the table field to use to build the query.


#####`<from>table</from>`
This tag allows you to specify the database table where to execute the query.


#####`<where>where ...</where>`
This tag allows you to specify the complete, well formatted where clause of the query.

*The where clause must be complete and must start with the key word "where". MySqlBatch won't make any change on the tag´s content.*


#####`<expected_result>1</expected_result>`
This tag lets you specify the the expected result returned by the query execution. 
Each execution result and expected result will be processed like "string".

By default to decide if a query has `pass` or `fail` MySqlBatch will check the equality (eq) of the two values.

It's possible, using the attribute `operator` of the `<expected_result>` tag, to specify another operator.

  * `<expected_result operator="lt">` the expected result should be lower than the execution result.
  *  `<expected_result operator="gt">` the expected result should be greater than the execution result.
  * `<expected_result operator="le">` the expected result should be lower or equal to the execution result.
  * `<expected_result operator="ge">` the expected result should be greater or equal to the execution result.
  
 
 
## `<pass>` , `<fail>`

The `pass`and `fail` tags can have an optional message, if defined it will be included into the generated emails.

    <pass>
    	<message>message shown into the generated email id the query passes</message>
    </pass>

    <fail>
    	<message>message shown into the generated email if the query fails</message>
    </fail>

In addition `pass`and `fail` tags can have a nested query following exactly the same pattern than the already detailed `<query>` tag. 

    <query name="first query">
    	<count_query>
    		<from>clients</from>
    		<where>where active = "1"</where>
    		<expected_result>666</expected_result>
    	</count_query>
	    <description>Count active clients</description>
    	<pass>
    		<message>no error counting the active clients</message>
    		<query name="second query">
    			<count_query>
    				<from>clients</from>
    				<where>where with_order = "1"</where>
    				<expected_result operator="gt">500</expected_result>
    			</count_query>
    			<description>Check if we have a least 500 active clients with order</description>
    			<pass>
    				<message>ok ...</message>
    				<query name="third query if the second is okay">
					...
    			</pass>
    			<fail>
    				<message>error ...</message>
    				<query name="fourth query if the second is not okay">
					...
    			</fail>
    		</query>
    	</pass>
    	<fail>
    		<message>error counting the active clients</message>
    	</fail>
    </query>
