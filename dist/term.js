"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.execute = void 0;
const child_process_1 = require("child_process");
const execute = (command, args, spawnOptions) => new Promise((resolve, reject) => {
    const cmd = (0, child_process_1.spawn)(command, args, spawnOptions);
    cmd.stdout.on('data', resolve);
    cmd.stderr.on('data', reject);
});
exports.execute = execute;
