#!/bin/sh

APPEND=$(cat << EOF
#include <bits/stdc++.h>
using namespace std;
EOF
)

echo "$APPEND" && cat