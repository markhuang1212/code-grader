#!/bin/bash

# Exit Code
# 0: Success
# 1: Wrong Answer
# 2: Execution Error
# 3: Internal Error

read LENGTH

dd bs=$LENGTH count=1 of=a.out > /dev/null

chmod +x a.out

if [ -z "${TEST_CASE_DIR}" ]
then
    echo "Missing TEST_CASE_DIR"
    exit 3
fi

a.out < ${TEST_CASE_DIR}/input.txt > output.txt

if [[ $? -ne 0 ]]
then
    cat output.txt
    exit 2
fi

diff output.txt ${TEST_CASE_DIR}/output.txt
