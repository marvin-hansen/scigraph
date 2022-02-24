//Simple DO queries for Example 1.

//Retrieve information about Person whose name is Harry
SELECT * FROM Person WHERE name == 'Harry';

//Retrieve 2 Person nodes after skipping the first 2 nodes returned
SELECT * FROM Person SKIP 2 TAKE 2;

//Find the name of Person who owns the Phone with a certain phone number
FROM Phone where phoneNumber == '17447134342' return owner.name;

//Find all paths to neighboring nodes of Person with id = 3
MATCH path=(:Person{id == 3})-->() RETURN path;

//Find all paths between Person with id=5 to anyone else upto 4 degrees. Also retrieve name and age of connected Person
MATCH path = (:Person{id=5})-[*1..4]->(p1:Person) RETURN p1.name, p1.age, path;

//Retrieve all phones that have '134' in the phone number
FROM Phone where CONTAINS(phoneNumber,'134') return *;

//Update a Phone with a new phone number
Update Phone where phoneNumber='15335505806' SET phoneNumber to '16085505806';

//Find the count of Phones owned by all Persons and return the name of the owner and the count of phones in descending order
FROM Person ORDER BY COUNT(phones)DESC RETURN COUNT(phones) AS NumberOfPhones, name;

//Find the average number of phones owned by men and women
FROM Person GROUP BY gender RETURN gender, AVG(COUNT(phones));

//Find the total time of outgoing calls that George made
From Person where name == 'George'  return SUM(phones.outgoing_calls.callDetails.duration) As totalMins;

//Retrieve call information between any two people if any of the calls between them exceeded 3000 seconds
From Call where ANY(callDetails, duration > 3000)  return callDetails.duration, callDetails.startTime, caller.owner.name as Caller , callee.owner.name as Callee;

//Remove one phone from a particular Person
Update Person where name = 'Harry' remove phones[phoneNumber='13690065401'] from phones;

//Delete Phone where phoneNumber is 13690065401
Delete Phone where phoneNumber = '13690065401';





