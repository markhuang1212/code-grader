# Writing Test Cases

To add testcases to this folder, Please follow guidelines given in folders `example-x`. Upon commit, test cases will be automatically tested. Please make sure the answer can be run within the given runtime and memory.

## Overview

### example-1

This test case is used internally for unit testing. It expect a c++ program that prints `Hello\n`. It does `include <iostream>` and `using namespace std` so that the user's code doesn't needs to. A correct answer is like:

```cpp
int main() { cout << "Hello << endl; }
```

### simple-1

This test case demonstrates how `code-grader` can be used to grade variable definition. Any variable definition of type `int` that has value `1024` pass. 

```cpp
int x = 1024; // this will pass
int y = 1000+24; // this will also pass
auto z = 0x00000400; // this will also pass
// however
char x = 'a'; // this won't pass
```