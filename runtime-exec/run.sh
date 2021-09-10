#!/bin/bash

# Exit Code
# 0: Success
# 1: Wrong Answer
# 2: Execution Error
# 3: Internal Error
# 137: OOM

DATA_DIR=/data

if [ -z "${TEST_CASE_DIR}" ]
then
    echo "Missing TEST_CASE_DIR"
    exit 3
fi

${DATA_DIR}/a.out < ${TEST_CASE_DIR}/input.txt > ${DATA_DIR}/output.txt 2>&1
STATUS=$?

if [[ $STATUS -eq 137 ]]
then
    exit 137
fi

if [[ $STATUS -ne 0 ]]
then
    echo "program exited with status $STATUS"
    if [[ -f output.txt ]];
    then
        cat output.txt
    fi
    exit 2
fi

# Ignore Trailing Spaces
# Ignore chaeges in amount of white spaces
# Ignore Black Lines
diff -Z -b -B ${DATA_DIR}/output.txt ${TEST_CASE_DIR}/output.txt
