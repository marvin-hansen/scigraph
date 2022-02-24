USE Date FORMAT "%Y/%m/%e";
//Create Person nodes
LET people = IMPORT CSV "People.csv" USE HEADER;
FOR EACH person IN $people =>
(
	CREATE Person {
			id		: $person.id,
			name		: $person.name,
			age		: $person.age,
			gender		: $person.gender
		}
);
//Create Phone nodes and connect them to Person nodes
FOR EACH phonenum IN IMPORT CSV 'Phones.csv'=>
(
	Let owner = LOOKUP Person where id = CAST($phonenum[0] AS Integer),
	Let thisPhone = CREATE Phone {
		phoneNumber	: $phonenum[1],
		owner		: ($owner) 
		}
);
//Create Calls between Phones and CallDetails for each Call
FOR EACH mycall IN IMPORT CSV "Calls.csv" COMMIT EVERY 2500 =>
(
	Let callerphone = LOOKUP Phone where phoneNumber == $mycall[0],
	Let calleephone = LOOKUP Phone where phoneNumber == $mycall[1] ,
	Let thiscall = LOOKUP Call where caller = ($callerphone) and callee = ($calleephone),
	IF IS_NULL($thiscall) THEN 
	(
			Let thiscall = CREATE Call{caller:($callerphone), callee:($calleephone)}
	),
	CREATE CallDetail{
				callerArea	: $mycall[2],
				calleeArea	: $mycall[3],
				startTime	: $mycall[4],
				duration	: $mycall[5],
				call		: $thiscall	
			}	
);
 