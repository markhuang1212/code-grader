#!/bin/bash

# Exit Code
# 0: Success
# 1: Wrong Answer
# 2: Internal Error

tee > /tmp/a.out
chmod +x /tmp/a.out

if [ -z "${TEST_CASE_DIR}" ]
then
    echo "Missing TEST_CASE_DIR"
    exit 2
fi

/tmp/a.out < ${TEST_CASE_DIR}/input.txt > /tmp/output.txt

if [[ $? -ne 0 ]]
then
    cat /tmp/output.txt
    exit 2
fi

diff /tmp/output.txt ${TEST_CASE_DIR}/output.txt
