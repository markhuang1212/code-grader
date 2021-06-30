#!/bin/bash

# Exit Code
# 0: Success
# 1: Compilation Error
# 2: Internal Error

tee > code.txt

if [ -z "${TEST_CASE_DIR}" ]
then
    echo "Missing TEST_CASE_DIR"
    exit 2
fi

if [ -z "${CXX}" ]
then
    echo "Missing CXX"
    exit 2
fi

if [ -z "${CXXFLAGS}" ]
then
    echo "Missing CXXFLAGS"
    exit 2
fi

cat ${TEST_CASE_DIR}/prepend.txt > main.cpp
cat code.txt >> main.cpp
cat ${TEST_CASE_DIR}/append.txt >> main.cpp

${CXX} ${CXXFLAGS} -o /tmp/a.out main.cpp

if [[ $? -ne 0 ]]
then
    exit 1
fi

cat /tmp/a.out