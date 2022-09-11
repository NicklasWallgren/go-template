#!/usr/bin/env bash

function help() {
   echo ""
   echo "Usage: $0 -html -func -diff"
   echo -e "\t-html Show coverage as a html report"
   echo -e "\t-func Show coverage a func report"
   echo -e "\t-diff Generate a diff-coverage report against the origin/master branch"
   exit 1
}

function has_param() {
    local terms="$1"
    shift

    for term in $terms; do
        for arg; do
            if [[ $arg == "$term" ]]; then
                echo "yes"
            fi
        done
    done
}

if [[ -n $(has_param "help" "$@") ]]; then
    help
fi

# Creates the default working directory
mkdir -p build

# Executes the tests and generates the coverage.out report
go test $(go list ./...) -p 1 -coverprofile=build/coverage.out -covermode count -coverpkg ./...

# Removes the ignored paths from the coverage.out file
while read p || [ -n "$p" ]
do
if [[ "$OSTYPE" == "darwin"* ]]; then
  sed -i '' "/${p//\//\\/}/d" ./build/coverage.out
else
  sed -i "/${p//\//\\/}/d" ./build/coverage.out
fi
done < ./.coverageignore

if [[ -n $(has_param "--html" "$@") ]]; then
    go tool cover -html=build/coverage.out
fi

if [[ -n $(has_param "--func" "$@") ]]; then
    go tool cover -func=build/coverage.out
fi

if [[ -n $(has_param "--diff" "$@") ]]; then
    gocover-cobertura < build/coverage.out > build/coverage.xml
    diff-cover build/coverage.xml --compare-branch=origin/main --html-report build/report.html
    open build/report.html
fi