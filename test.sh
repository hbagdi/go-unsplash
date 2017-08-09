#!/usr/bin/env bash

#Do not run tests on other's PRs; secret key issues
if ! [[ -z $TRAVIS_PULL_REQUEST_SLUG ]]
then
	if [[ $TRAVIS_PULL_REQUEST_SLUG != "hbagdi/go-unsplash" ]]
	then
		exit 0
	fi
fi

#For testing on local desktop
curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter
chmod +x ./cc-test-reporter
./cc-test-reporter before-build

declare -r TOKEN_FILE="auth.env"
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

for d in $(go list ./... | grep -v vendor); do
    go test -v -coverprofile="profile.out" -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
/bin/cp coverage.txt c.out
./cc-test-reporter after-build
