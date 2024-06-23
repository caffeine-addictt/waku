import type { RepoInfo } from './types';

/** Extract username/org and repo from url */
export const extractRepoInfo = (url: string): RepoInfo => {
  const originUrl = url.trim();
  const match = /^https:\/\/github\.com\/([^/]+)\/([^/]+?)(?:\.git)?$/.exec(
    originUrl,
  );

  if (!match || !match[1] || !match[2])
    throw new AggregateError(
      `Could not extract user and repo from git remote: ${originUrl}`,
    );

  return {
    url: originUrl,
    org: match[1],
    repo: match[2],
  } satisfies RepoInfo;
};
