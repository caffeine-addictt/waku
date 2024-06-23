"use strict";
/**
 * This file will hold the configuration options to make the setup script not so annoying
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.resolveUserInfo = exports.resolveRepoInfo = void 0;
const parser_1 = require("./parser");
const term_1 = require("./term");
const resolveRepoInfo = () => (0, term_1.execute)('git', ['remote', 'get-url', 'origin'])
    .then((_originUrl) => (0, parser_1.extractRepoInfo)(_originUrl.toString()))
    .catch(() => ({}));
exports.resolveRepoInfo = resolveRepoInfo;
const resolveUserInfo = async () => {
    const promises = await Promise.all([
        (0, term_1.execute)('git', ['config', '--global', 'user.name'])
            .then((r) => r.toString().trim())
            .catch(() => undefined),
        (0, term_1.execute)('git', ['config', '--global', 'user.email'])
            .then((r) => r.toString().trim())
            .catch(() => undefined),
    ]);
    return {
        name: promises[0],
        email: promises[1],
    };
};
exports.resolveUserInfo = resolveUserInfo;
const resolveGitInfo = async () => {
    const promises = await Promise.all([(0, exports.resolveRepoInfo)(), (0, exports.resolveUserInfo)()]);
    return {
        ...promises[0],
        ...promises[1],
    };
};
exports.default = resolveGitInfo;
