#!/bin/sh

PREPEND=$(cat << EOF
#include <bits/stdc++.h>
using namespace std;
EOF
)

echo "$PREPEND" && cat
