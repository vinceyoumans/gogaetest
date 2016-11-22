This is an experiment on how to read a Google SpreadSheet
that has been published to the web.
The concept is to get the json version of the web item.

I have a test google spread sheet at:
https://docs.google.com/spreadsheets/d/e/2PACX-1vT0KE8bTPLuh3NtL73pRRmptZsSIc-w1b_dFgGyDU3aNED4Mg9y-Pi-_1dHJQMjGFbMep18qq6Wf5CH/pubhtml

The ID is: 
2PACX-1vT0KE8bTPLuh3NtL73pRRmptZsSIc-w1b_dFgGyDU3aNED4Mg9y-Pi-_1dHJQMjGFbMep18qq6Wf5CH
and
1brWkVApKRJ3aRL6OfIQ4jlyx-wXmDhfpGjHIK6ACdQE

The json get is at:
https://spreadsheets.google.com/feeds/list/1brWkVApKRJ3aRL6OfIQ4jlyx-wXmDhfpGjHIK6ACdQE/1/public/values?alt=json 

Both of the above are confirmed to work.

If you observe the json schema, there are 
version
encoding
feed

Feed has a slice of type struct.
1st, as I am new to goLang, could someone correct my vocabulary description of the above please.
2nd... I am sure I am doing this wrong.
type Entry struct {
    ID string `json:"id"`
}

please see sampentry.json  or goto 
https://spreadsheets.google.com/feeds/list/1brWkVApKRJ3aRL6OfIQ4jlyx-wXmDhfpGjHIK6ACdQE/1/public/values?alt=json


In this sample code, I can extract everything right up to the part where i have to extract the entry.
I get this error:
2016/11/22 14:44:08 json: cannot unmarshal array into Go value of type main.Entry
exit status 1

I believe it is because I am trying to import a slice into a single json array. 
Again forgive the way I am explaining it.  I think there is a specific term for this.

could someone look at the code, and give me a hint on how I can import a slice of structs into my system please.
