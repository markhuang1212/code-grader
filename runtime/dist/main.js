"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const process_1 = __importDefault(require("process"));
const child_process_1 = __importDefault(require("child_process"));
const test_case_path = process_1.default.env.TEST_CASE_PATH;
const ErrorCode = {
    CompilationError: 120,
    RuntimeError: 121,
    InternalError: 122
};
if (test_case_path === undefined) {
    console.error('No TEST_CASE_PATH');
    process_1.default.exit(122);
}
const copyResult = child_process_1.default.spawnSync('cp', ['-r', test_case_path + '/*', '/tmp/testcase'], {
    cwd: '/',
    encoding: 'utf-8'
});
if (copyResult.status != 0) {
    console.log(copyResult, undefined, 2);
    process_1.default.exit(ErrorCode.InternalError);
}
const compilationResult = child_process_1.default.spawnSync("sh", ["-c", "cat usercode.cpp | ./preprocess | make"], {
    cwd: '/tmp/testacse',
    encoding: 'utf-8'
});
if (compilationResult.status != 0) {
    console.log(JSON.stringify(compilationResult, undefined, 2));
    process_1.default.exit(ErrorCode.CompilationError);
}
const executionResult = child_process_1.default.spawnSync("sh", ["-c", "./a.out < input.txt | diff output.txt -"], {
    encoding: 'utf-8',
    cwd: '/tmp/testcase'
});
if (executionResult.status != 0) {
    console.log(JSON.stringify(executionResult, undefined, 2));
    process_1.default.exit(ErrorCode.RuntimeError);
}
console.log('Success!');
process_1.default.exit(0);
