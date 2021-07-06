#!/bin/bash

# Exit Code
# 0: Success
# 1: Compilation Error
# 2: Internal Error

DATA_DIR=/data

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

cat ${TEST_CASE_DIR}/prepend.txt > ${DATA_DIR}/main.cpp
cat ${DATA_DIR}/code.txt >> ${DATA_DIR}/main.cpp
cat ${TEST_CASE_DIR}/append.txt >> ${DATA_DIR}/main.cpp

${CXX} ${CXXFLAGS} -o ${DATA_DIR}/a.out ${DATA_DIR}/main.cpp

if [[ $? -ne 0 ]]
then
    exit 1
fi
