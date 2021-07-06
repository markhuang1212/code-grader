#!/bin/bash

TEST_CASE_DIR=$1

cd ${TEST_CASE_DIR}
cat prepend.txt > /tmp/main.cpp
cat answer.txt >> /tmp/main.cpp
cat append.txt >> /tmp/main.cpp

g++ -std=c++11 -o /tmp/a.out /tmp/main.cpp
/tmp/a.out < input.txt > /tmp/output.txt
diff output.txt /tmp/output.txt

if [ $? -eq 0 ]
then
    echo "SUCCESS"
fi