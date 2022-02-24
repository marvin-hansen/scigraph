## Example 1: Quick Start Example for Infinitegraph to ingest data and query using DO 

This tutorial will help you understand how to use the InfiniteGraph DO query language to ingest data into an InfiniteGraph Database and run simple queries. 

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

Create a new Infinitegraph database 
Define a simple schema using DO
Ingest data - nodes and edges using DO
Query the data using DO

## Run Example 1

The example1 package contains the following files:

runExample.bat - Windows batch script to run example1 - it creates a data folder in your working directory where the database and all related files are saved
runExample.sh - Unix shell script to run example1 -  it creates a data folder in your working directory where the database and all related files are saved
createSchema.do - DO statements to create schema
createData.do - DO statements to create data
runQuery.do - some simple DO query examples

To run this example open a command prompt from your working directory and execute the script(runExample.bat on Windows, ./runExample.sh on Linux). The script does the following:

	* Creates a data folder in your working directory
	* Creates an InfiniteGraph federated database called example1 using the tool objy createfd(deletes and recreates it if it already exists) 
	* Defines schema using DO statements from the file createSchema.do
	* Creates nodes and edges using DO statements from the file createData.do
	* Runs queries from the file runQuery.do and saves the output to the file queryResults.txt in your current working directory


## Additional tasks 

After running the example you can inspect the output file queryResults.txt to observe the results of the different types of DO query statements provided in the runQuery.do file. You can also run the DO runner and connect to the database to run more queries.

From the command prompt run the following tool to open the DO runner where you can run DO statements to query the data:

	*objy do -bootfile data/example1.boot

Alternatively you can open the Studio server to run and visualize queries.
From the command prompt run the following command to start the Studio server(Needs Java 1.8 or above):

	*objy startstudioserver -bootfile data/example1.boot

This will start the Studio server at port 8190. From a browser navigate to http://localhost:8190/index.html and connect to the database. Switch to the DO tab and run queries from the DO query window.

Data queries are visualized in a table format. For example the query below will return all details about the Person 'Harry':
	SELECT * FROM Person WHERE name == 'Harry'; 
 
Match queries  are visualized in a graph format. For example the query below will return all the paths upto 4 degrees between all Persons in the graph:
	MATCH path = (:Person)-[*1..4]->(:Person) RETURN path;
 
## Move to the next example to learn more advanced topics


