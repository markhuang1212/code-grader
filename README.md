# code-grader

This is a backend for grading C++ code.

## Overview

The backend receives requests for grading code, it build and execute the code inside a docker environment, and returns the result. The basic unit of code execution is a `test case`.

A test case contains:

* A preprocessing unit that process the c++ code the user uploads. (e.g. help the user include necessary libraries, provide user with the `main` function, etc)
* A set of inputs and outputs
* A set of compiler options (e.g., optimization level)
* A set of runtime options (e.g., memory limit, runtime limit, etc)

## Author

Huang Meng \
2021