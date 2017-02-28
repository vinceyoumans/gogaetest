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

The feed has a struct of Entry which is an array ( or slice) of another struct.
This code is able to pick out the individual parts of Entry

Need to add this to a db now.  

Also, the first row is ignored as I think it is suppose to be a Label row
But I cannot find the JSON part that descibed the rows.
If anyone has a comment on that.










