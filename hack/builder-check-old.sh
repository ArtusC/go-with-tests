#!/bin/bash

cmd="go test -tags unit -v"
# cmd="go test -v -count=1 -run"
# cmd_race="go test -tags unit -race -v"
# run=""
for arg in "$@"; do
    # Separate parameter and value using '=' as delimiter
    IFS='=' read -r -a arg_parts <<< "$arg"
    
    # Check if both parameter and value are present
    if [ "${#arg_parts[@]}" -eq 2 ]; then
        parameter="${arg_parts[0]}"
        value="${arg_parts[1]}"
        
        echo "Parameter: $parameter, Value: $value"
        
        # Add your logic here for processing each parameter and its value
        # For example, you can use a case statement to handle different parameters
        case "$parameter" in
            "testfile")
                # Handle the 'testfile' parameter
                # echo "Processing 'testfile' parameter with value '$value'"
                # go test -tags unit -v ./context/v3
                cmd="${cmd} ${testfile}"
                ;;
            "testname")
                # Handle the 'testname' parameter
                # echo "Processing 'testname' parameter with value '$value'"
                # go test -tags unit -v ./context/v3 -run ^TestServer$
                cmd="${cmd} -run ^${testname}$"
                echo ""
                ;;
            "race")
                # Handle the 'race' parameter
                # echo "Processing 'race' parameter with value '$value'"
                cmd="${cmd} -race"
                echo ""
                ;;
            *)
                # Handle unknown parameters
                # echo "Unknown parameter: $parameter"
                # cmd="${cmd}"
                echo ""
                ;;
        esac
    else
        echo "Invalid argument: $arg"
        echo ""
    fi
done

echo "AQUI: ${cmd}"
$cmd

# if [ $# -eq 0 ]; then
#     $cmd ./...
#     exit
# fi

# if [ $# -eq 1 ]; then
# 	filetest=$1
# 	echo "Running test $filetest"
# 	$cmd $filetest
# 	exit
# fi

# filetest=$1
# nametest=$2
# echo "Running nametest $nametest on filetest $filetest"
# $cmd $filetest -testify.m $nametest