UPDATE SCHEMA
{
  CREATE CLASS Person 
  {
    id:     Integer,
    name:   String,
    age:    Integer { Storage: B16, Encoding: Unsigned},
    gender:   String,
    phones:   List { Element: Reference{ Referenced: Phone, inverse: owner}}
  }

  CREATE CLASS Phone 
  {
    phoneNumber:  String , 
    owner:      Reference { Referenced: Person, inverse: phones},
    incoming_calls: List { Element: Reference { Referenced: Call, inverse: callee}},
    outgoing_calls: List { Element: Reference { Referenced: Call, inverse: caller}}
  }

  CREATE CLASS Call
  {
    caller: Reference { Referenced: Phone, inverse: outgoing_calls},
    callee: Reference { Referenced: Phone, inverse: incoming_calls},
    callDetails: List { Element: Reference{ Referenced: CallDetail, inverse: call}}
  }

  CREATE CLASS CallDetail
  {
    callerArea:   Integer { Encoding: Unsigned, Storage: B32},
    calleeArea:   Integer { Encoding: Unsigned, Storage: B32},
    duration:     Integer { Encoding: Unsigned, Storage: B32},
    startTime:    DateTime,
    call:         Reference { Referenced: Call, inverse: callDetails}
  }
};
