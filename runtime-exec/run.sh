#!/bin/bash

# Exit Code
# 0: Success
# 1: Wrong Answer
# 2: Execution Error
# 3: Internal Error

DATA_DIR=/data

if [ -z "${TEST_CASE_DIR}" ]
then
    echo "Missing TEST_CASE_DIR"
    exit 3
fi

${DATA_DIR}/a.out < ${TEST_CASE_DIR}/input.txt > ${DATA_DIR}/output.txt 2>&1

if [[ $? -ne 0 ]]
then
    cat output.txt
    exit 2
fi

diff -Z -b -B -w ${DATA_DIR}/output.txt ${TEST_CASE_DIR}/output.txt
