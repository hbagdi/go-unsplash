#!/bin/bash

#For testing on local desktop
TOKEN_FILE="auth.env"
if [ -r $TOKEN_FILE ];
then
		source $TOKEN_FILE
fi
if ! [[ -z "$unsplash_appID" ]]
then
		export unsplash_appID
fi
if ! [[ -z "$unsplash_secret" ]]
then
		export unsplash_secret
fi
if ! [[ -z "$unsplash_usertoken" ]]
then
		export unsplash_usertoken
fi
#
rm coverage.txt
go test -coverprofile=coverage.txt
go tool cover -html=coverage.txt
