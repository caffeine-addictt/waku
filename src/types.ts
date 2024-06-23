// Place some shared types here

export interface ProjectInfo {
  name: string;
  email: string;
  username: string;
  repository: string;
  proj_name: string;
  proj_short_desc: string;
  proj_long_desc: string;
  docs_url: string;
  assignees: string;
}

export interface RepoInfo {
  url?: string;
  org?: string;
  repo?: string;
}

export interface UserInfo {
  name?: string;
  email?: string;
}

export type GitInfo = RepoInfo & UserInfo;
