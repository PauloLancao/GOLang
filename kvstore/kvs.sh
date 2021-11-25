#!/bin/bash
# set payload
PAYLOAD='{"fullname":"fullname5","firstname":"firstname5","middlename":"middlename5","lastname":"lastname5","email":"test9@test.com","age":"55","phone":"555555555","addresses":[{"id":1,"addressline1":"line1","addressline2":"line2","postcode":"HG1 1AA","city":"Harrogate","county":"York","country":"UK"}]}'
PAYLOAD_UPDATED="UPDATED_PAYLOAD"

# TCP call nc
for n in 1 2 3 4 5 
do
printf "cmd=create|key=$n|body=$PAYLOAD\ncmd=get|key=$n\ncmd=update|key=$n|body=$PAYLOAD_UPDATED\ncmd=get|key=$n\ncmd=delete|key=$n\n" | nc 127.0.0.1 9001
done