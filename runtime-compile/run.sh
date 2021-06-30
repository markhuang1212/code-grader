#!/bin/bash


tee > code.txt

cat ${TEST_CASE_DIR}/prepend.txt > main.cpp
cat code.txt >> main.cpp
cat ${TEST_CASE_DIR}/append.txt >> main.cpp

${CXX} ${CXXFLAGS} -o /tmp/a.out main.cpp

if [[ $? -ne 0 ]]
then
    exit 1
fi

cat /tmp/a.out