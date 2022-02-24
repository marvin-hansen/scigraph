//Find all paths from a specific Phone to all persons in the database(upto 6 degree of separation)
MATCH path = (ph1:Phone{phoneNumber='18691699815'})-[*1..6]->(:Person) RETURN path;

//Find all paths between a specific Person and all others over age 30 (upto 8 degree of separation)
MATCH path = (:Person{id=38})-[*1..8]->(p1:Person{age>30}) RETURN p1.name, p1.age, path;

//Find all paths between a specific Phone Number and all Persons within certain date/time ranges traversing a depth of 7.
MATCH path = (:Person{name='Rob'})-->(:Phone)-->(:Call{callDetails ANY(startTime >=2019-08-01 00:00:00 AND startTime <=2019-12-01 00:00:00)})-->(:Phone)-->(:Call{callDetails ANY(startTime >=2019-08-01 00:00:00 AND startTime <=2019-12-01 00:00:00)})-->(:Phone)-->(:Call{callDetails ANY(startTime >=2019-08-01 00:00:00 AND startTime <=2019-12-01 00:00:00)})-->(:Phone)-->(:Call{callDetails ANY(startTime >=2019-08-01 00:00:00 AND startTime <=2019-12-01 00:00:00)})-->(:Phone)-->(:Person) RETURN path;

//Find all outgoing call details from Phones with phone numbers starts with 159.
MATCH path = (p1:Person)-->(:Phone{phoneNumber=~ '^159'})-[:outgoing_calls]->(c:Call)-->(p2:CallDetail) RETURN p1.name AS caller,c.callee.owner.name AS callee,c.caller.phoneNumber, c.callee.phoneNumber, p2.startTime, p2.duration;

//Find the numbers for the 10 longest calls made by a specific Phone
MATCH path = (:Phone{phoneNumber=='13501155728'})-[:outgoing_calls]->(p1:Call)-->(p2:CallDetail) ORDER BY p2.duration DESC TAKE 10 RETURN  p1.callee.phoneNumber,p2.duration;

//Find the top 10 caller area based on the number of calls within certain date/time ranges for a specific phone number
MATCH path = (:Phone{phoneNumber:'13513317053'})-->(:Call)-[:callDetails]->(p2:CallDetail{startTime >= 2019-07-01 00:00:00 AND startTime <= 2019-11-01 00:00:00}) GROUP BY p2.callerArea ORDER BY COUNT() DESC take 4 RETURN p2.callerArea, COUNT() AS NumberOfCalls;

//Find all the call details whose duration exceeded 2200 for a specific Person.
MATCH path = (:Person { id == 300})-->(:Phone)-->(c1:Call { callDetails ANY(duration > 2200)})-->(:Phone)-->(:Person) RETURN c1.callDetails[duration > 2200], path;

//Update the Call and add CallDetails for two persons.
UPDATE Call WHERE caller.owner.name='Abigail' and callee.owner.name='George' ADD {Create CallDetail{callerArea:40866, calleeArea:67878, duration:9000}} TO callDetails;

//Find all paths to calls made by Person with id 288
MATCH path = (:Person{id=288})--()--()--(:CallDetails) RETURN path;

//Delete the Phones owned by Person with id 288, that made calls within a certain date/time ranges.
UPDATE Person WHERE id = 288 REMOVE phones[ANY(outgoing_calls.callDetails,startTime>2019-11-28 00:00:00) && ANY(outgoing_calls.callDetails,startTime<2019-12-28 00:00:00)] FROM phones ;





