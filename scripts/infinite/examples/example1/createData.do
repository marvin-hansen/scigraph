//Create all Phone nodes
LET Phone1 = CREATE Phone{phoneNumber:'13690065401'};
LET Phone2 = CREATE Phone{phoneNumber:'13424032094'};
LET Phone3 = CREATE Phone{phoneNumber:'15332258934'};
LET Phone4 = CREATE Phone{phoneNumber:'13151366596'};
LET Phone5 = CREATE Phone{phoneNumber:'17447134342'};
LET Phone6 = CREATE Phone{phoneNumber:'15335505806'};

//Create CallDetail nodes
LET CallDetail1 = CREATE CallDetail{
callerArea:46677,
calleeArea:80415,
duration:1692,
startTime:2019/01/17 03:51:14
};
LET CallDetail2 = CREATE CallDetail{
callerArea:46677,
calleeArea:80415,
duration:2675,
startTime:2019/03/21 20:35:07
};
LET CallDetail3 = CREATE CallDetail{
callerArea:42615,
calleeArea:99924,
duration: 3076,
startTime: 2019/11/13 20:49:03
};
LET CallDetail4 = CREATE CallDetail{
callerArea:42615,
calleeArea:99924,
duration: 2089,
startTime: 2019/09/21 17:26:01
};
LET CallDetail5 = CREATE CallDetail{
callerArea:2720,
calleeArea:8503,
duration: 3493,
startTime: 2019/02/04 3:11:22
};
LET CallDetail6 = CREATE CallDetail{
callerArea:2720,
calleeArea:8503,
duration: 438,
startTime: 2019/05/02 21:34:43
};
LET CallDetail7 = CREATE CallDetail{
callerArea:80415,
calleeArea:42615,
duration: 1726,
startTime: 2019/06/20 21:51:21
};
LET CallDetail8 = CREATE CallDetail{
callerArea:99924,
calleeArea:2720,
duration: 1731,
startTime: 2019/03/06 21:43:08
};
LET CallDetail9 = CREATE CallDetail{
callerArea:99924,
calleeArea:2720,
duration: 2789,
startTime: 2019/08/21 7:20:53
};
LET CallDetail10 = CREATE CallDetail{
callerArea:8503,
calleeArea:46677,
duration: 2518,
startTime: 2019/09/09 21:37:36
};
LET CallDetail11 = CREATE CallDetail{
callerArea:8503,
calleeArea:80415,
duration: 2719,
startTime: 2019/07/11 14:06:29
};
LET CallDetail12 = CREATE CallDetail{
callerArea:99924,
calleeArea:8503,
duration: 2789,
startTime: 2019/08/21 7:20:53
};

//Create Person nodes and connect them to list of Phones they own
CREATE Person{
id:11,
name:'Harry',
age:24,
gender:'male',
phones:LIST { $Phone1, $Phone3}
};

CREATE Person{
id:4,
name:'George',
age:67,
gender:'male',
phones:LIST { $Phone2, $Phone4}
};

CREATE Person{
id:5,
name:'Abigail',
age:25,
gender:'female',
phones:LIST { $Phone5}
};

CREATE Person{
id:3,
name:'Lisa',
age:22,
gender:'female',
phones:LIST { $Phone6}
};

//Create Call nodes and all the calldetails for each Call
CREATE Call {
caller:$Phone1,
callee:$Phone2,
callDetails:LIST { $CallDetail1,$CallDetail2}
};

CREATE Call {
caller:$Phone3,
callee:$Phone4,
callDetails:LIST { $CallDetail3,$CallDetail4}
};

CREATE Call {
caller:$Phone5,
callee:$Phone6,
callDetails:LIST { $CallDetail5,$CallDetail6}
};

CREATE Call {
caller:$Phone2,
callee:$Phone3,
callDetails:LIST { $CallDetail7}
};

CREATE Call {
caller:$Phone4,
callee:$Phone5,
callDetails:LIST { $CallDetail8,$CallDetail9}
};

CREATE Call {
caller:$Phone6,
callee:$Phone1,
callDetails:LIST { $CallDetail10}
};

CREATE Call {
caller:$Phone6,
callee:$Phone2,
callDetails:LIST { $CallDetail11}
};

CREATE Call {
caller:$Phone4,
callee:$Phone6,
callDetails:LIST { $CallDetail12}
};











