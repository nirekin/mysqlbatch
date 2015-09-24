# MySqlBatch

MySqlBatch is a tool allowing you to execute predefined lists of simple queries on an also predefined list of databases.
Each list of queries is wrapped into a batch. The execution of each batch can automatically generate a report and send it by mail to some recipients.

MySqlBatch consume an XML file containing the specification of all the batches to run. 
This files must be valid against this [DTD](https://raw.githubusercontent.com/nirekin/mysqlbatch/master/mysqlbatch.dtd)

By default MySqlBatch will look, on its own directory, for a file named "batch.xml" but you can specify another one using `mysqlbatch -file foo.xml`

In order to automatically send the generated emails you must run MySqlBatch with the `-email` option: `mysqlbatch -file foo.xml -email`


## `<batches>`
The root node of the configuration file is `<batches>`

`<batches>` can contain an optional `<author>` and must contain at least one `<batch>`

### `<author>`
    <author name="jack walsh" organization="YourOrganization" phone="987987987" email="jackwalsh@yourorganization.com" />

The author tag is optional, it allows you to give information about the person responsible of the tool execution and its generated emails.

If defined the content of the author tag will be included into the generated emails.

###  `<batch>`
    <batch name="name_batch1">

A batch is basically the unit of work of MySqlBatch. You can use a batch to specify a list of requests to execute on several databases in order to send a report to identified recipients.

The name of the batch is required it's used to build the subject of the generated emails.

### `<smtp_config>`
    <smtp_config server="smtp.mail.yahoo.com" port="25" user="jack.walsh" password="987987987" sender="jack.walsh@yourorganization.com"/>

The smtp_config tag is required, it defines the smtp server/account to send the generated email for a specific batch.

*The* `sender` *attribute of the smtp_config can be ommitted, in this case the `sender` field of the generated email will be filled with the content of the `user` attribute.*

### `<recipients>`
    <recipients>
      <recipient address="jonathanmardukas@yourorganization.com" />
      <recipient address="alonzomosely@yourorganization.com" />
    </recipients>

The recipients tag is required it allows you to define the list of recipients of the generated email for a specific batch.

### `<mysqls>`
    <mysqls>
      <mysql url="username:password@tcp(10.10.10.10:3306)/your_schema" />
    </mysqls>
    
The mysqls tag is required, it allows you to specify the list of all databases target of the execution of a given batch.

To access MySql MySqlBatch use the following url pattern: `username:password@tcp(10.10.10.10:3306)/your_schema`


Where :

  * "username" is the user login to access the database
  * "password" is the password to access the database
  * "10.10.10.10" is the IP address of the MySql server to reach
  * "3306" is the port of the MySql server to reach
  * "your_schema" is the schema target of the batch execution   

*The "username" and the "password" used to access the database are not shown into the generated emails. In order to identify the databases the emails will contain only the IP address, the port and the targeted schema.*


### `<queries>`






