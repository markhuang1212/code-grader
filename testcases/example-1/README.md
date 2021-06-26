# Writing Test Cases

When you write a test case, you can consider that the user's code is stored in `usercode.cpp`, and the runtime will evaluate the user's code by the following shell script.

```shell
# STEP1: compilation
cat usercode.cpp | bash preprocess.sh | make
# STEP2: execution
./a.out < input.txt | diff output.txt -
```

A test case should at least contain:
* `preprocess.sh`
* `Makefile`
* `input.txt` and `output.txt`

Read these files in this folder for reference.