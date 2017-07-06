#!/bin/bash

# workaround for issue https://github.com/mattn/goveralls/issues/20
# https://raw.githubusercontent.com/gopns/gopns/master/test-coverage.sh

echo "mode: set" > acc.out

while read dir;
do
    if ls "${dir}"/*.go &> /dev/null; then
        return_val="$(go test -coverprofile=profile.out "${dir}")"
        echo "${return_val}"
        if [[ ${return_val} != *FAIL* ]]
        then
            if [ -f profile.out ]
            then
                grep -v "mode: set" profile.out >> acc.out
            fi
        else
            exit 1
        fi
    fi
done << EOF
$(find ./* -maxdepth 10 -type d -not -path "./vendor*")
EOF


if [[ -n "$COVERALLS" ]]; then
    echo "aaa"
	"${HOME}/gopath/bin/goveralls" -coverprofile=acc.out -service travis-ci
fi

rm -rf ./profile.out
rm -rf ./acc.out
