## Example 2 

In Example 1 we went over basic DO syntax to populate a few nodes and edges into the database and then to query this data. In this second example we will use the DO feature to ingest larger amount of data from csv files into the database. With more data we can run more complex queries and filters to find interesting paths and relationships within the data. We will also be defining indexes to improve query performance.

To learn more about the DO query language, read the documentation topics in [Getting Started With DO](https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics/query/doPART-GetStarted.html).

## Setup
	
1. Install the latest version of InfiniteGraph from [InfiniteGraph Downloads](http://support.objectivity.com/downloads/infinitegraph/).
2. Download the files required to run the example. Unzip into your working directory.
3. Check to see if lockserver is running and if it is not, start it:
	From a command prompt run the following command to check if the lockserver is running:
	*objy checkls
	If lockserver is not running start it from a user account that has admin privileges using the following command:
	*objy startlockserver
		
## Goal

Populate data from csv files into an Infinitegraph database using DO
Define Indexes
Run complex queries and filters using DO

Previously covered in Example1(we will be running these steps again to recreate a fresh database for this example)
Create a new Infinitegraph database 
Define schema using DO

## Run Example 2

The example2 package contains the following files:

runExample.bat - Windows batch script to run example1 - it creates a data folder in your working directory where the database and all related files are saved
runExample.sh - Unix shell script to run example1 -  it creates a data folder in your working directory where the database and all related files are saved
createSchema.do - DO statements to create schema
createData.do - DO statements to create data
runQuery.do - DO query examples
People.csv - csv data file containing simulated data for information 400 Person nodes
Phones.csv - csv data file containing simulated data for information 600 Phone nodes 
Calls.csv - csv data file containing simulated call detail records for 5000 CallDetail nodes 

To run this example open a command prompt from your working directory and execute the script(runExample.bat on Windows, ./runExample.sh on Linux). The script does the following:

	* Creates a data folder in your working directory
	* Creates an InfiniteGraph federated database called example2 using the tool objy createfd(deletes and recreates it if it already exists) 
	* Defines schema using DO statements from the file createSchema.do
	* Defines indexes for Person on the 'id' field and for Phone on the 'phoneNumber' field
	* Creates nodes and edges for data in csv files using DO statements from the file createData.do
	* Runs queries from the file runQuery.do and saves the output to the file queryResults.txt in your current working directory


## Additional tasks 

After running the example you can inspect the output file queryResults.txt to observe the results of the different types of DO query statements provided in the runQuery.do file. You can also run the DO runner and connect to the database to run more queries.

From the command prompt run the following tool to open the DO runner where you can run DO statements to query the data:

	*objy do -bootfile data/example2.boot

Alternatively you can open the Studio server to run and visualize queries.
From the command prompt run the following command to start the Studio server:

	*objy startstudioserver -bootfile data/example2.boot

This will start the Studio server at port 8190. From a browser navigate to http://localhost:8190/index.html and connect to the database. 
(If Studio server is already running, you can connect to the database from the Connect tab. Click '+', select 'Add Existing' and enter the path to the bootfile: <example2_working_directory>/data/example2.boot. Then click 'Add' and 'Connect')
Switch to the DO tab and run queries from the DO query window.

Data queries are visualized in a table format. For example the query below will return all details about the Person 'Harry':
	SELECT * FROM Person WHERE name == 'Harry'; 
 
Match queries  are visualized in a graph format. For example the query below will return all the paths upto 4 degrees between a specific Person and all others that are older than 30:
	MATCH path = (:Person{id=38})-[*1..8]->(p1:Person{age>30}) RETURN path;
 


