#!/usr/bin/env bash

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
set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -v -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
