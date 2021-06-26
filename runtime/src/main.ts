import process from 'process'
import fs from 'fs'
import child_process from 'child_process'
import path from 'path'

const test_case_path = process.env.TEST_CASE_PATH

const ErrorCode = {
    CompilationError: 120,
    RuntimeError: 121,
    InternalError: 122
}

if (test_case_path === undefined) {
    console.error('No TEST_CASE_PATH')
    process.exit(122)
}

const copyResult = child_process.spawnSync("sh", ["-c", `cp ${path.join(test_case_path, '*')} /tmp/testcase`], {
    encoding: 'utf-8'
})
if (copyResult.status != 0) {
    console.log(JSON.stringify(copyResult, undefined, 2))
    process.exit(ErrorCode.InternalError)
}

const compilationResult = child_process.spawnSync("sh", ["-c", "cat usercode.cpp | ./preprocess.sh | make"], {
    cwd: '/tmp/testcase',
    encoding: 'utf-8'
})

if (compilationResult.status != 0) {
    console.log(JSON.stringify(compilationResult, undefined, 2))
    process.exit(ErrorCode.CompilationError)
}

const executionResult = child_process.spawnSync("sh", ["-c", "./a.out < input.txt | diff output.txt -"], {
    encoding: 'utf-8',
    cwd: '/tmp/testcase'
})

if (executionResult.status != 0) {
    console.log(JSON.stringify(executionResult, undefined, 2))
    process.exit(ErrorCode.RuntimeError)
}

console.log('Success!')
process.exit(0)