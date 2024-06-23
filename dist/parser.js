"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.extractRepoInfo = void 0;
/** Extract username/org and repo from url */
const extractRepoInfo = (url) => {
    const originUrl = url.trim();
    const match = /^https:\/\/github\.com\/([^/]+)\/([^/]+?)(?:\.git)?$/.exec(originUrl);
    if (!match || !match[1] || !match[2])
        throw new AggregateError(`Could not extract user and repo from git remote: ${originUrl}`);
    return {
        url: originUrl,
        org: match[1],
        repo: match[2],
    };
};
exports.extractRepoInfo = extractRepoInfo;
