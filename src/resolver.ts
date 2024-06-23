/**
 * This file will hold the configuration options to make the setup script not so annoying
 */

import { extractRepoInfo } from './parser';
import { execute } from './term';
import type { GitInfo, RepoInfo, UserInfo } from './types';

export const resolveRepoInfo = (): Promise<RepoInfo> =>
  execute('git', ['remote', 'get-url', 'origin'])
    .then((_originUrl) => extractRepoInfo(_originUrl.toString()))
    .catch(() => ({}));

export const resolveUserInfo = async (): Promise<UserInfo> => {
  const promises = await Promise.all([
    execute('git', ['config', '--global', 'user.name'])
      .then((r) => r.toString().trim())
      .catch(() => undefined),
    execute('git', ['config', '--global', 'user.email'])
      .then((r) => r.toString().trim())
      .catch(() => undefined),
  ]);

  return {
    name: promises[0],
    email: promises[1],
  } satisfies UserInfo;
};

const resolveGitInfo = async (): Promise<GitInfo> => {
  const promises = await Promise.all([resolveRepoInfo(), resolveUserInfo()]);
  return {
    ...promises[0],
    ...promises[1],
  } as GitInfo;
};
export default resolveGitInfo;
