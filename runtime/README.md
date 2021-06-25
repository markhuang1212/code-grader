# runtime

## Usage

`node /dist/main.js --test-case=/path/to/testcase --src=/path/to/source/code.cpp`

Program exits with 0 on success. On failure, it exits with
* `120` when compilation error
* `121` when runtime error
* `122` when internal error