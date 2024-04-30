import fs from 'fs';
import path from 'path';

import { ProjectInfo } from '../src/types';
import { replaceInFile, withTempDir } from '../src/io-util';

// Setup
const TEST_FILE_TO_REPLACE = path.join('tests/', 'testing_replace.txt');
const toReplace = `
My name is: {{NAME}}

The repo is {{REPOSITORY}}

{{EMAIL}} {{DOCS_URL}}
{{PROJECT_SHORT_DESCRIPTION}}

{{PROJECT_NAME}}
{{EMAIL}} ssssss

{{PROJECT_LONG_DESCRIPTION}}
`;
const replaced = `
My name is: John

The repo is https://github.com/username/repo

HbVYf@example.com https://example.com/docs
Project with some description

My Project
HbVYf@example.com ssssss

Project with some description that is really long
`;

const tmp = withTempDir(
  'io-util-test',
  () => {
    return;
  },
  false,
);
afterAll(() => {
  tmp.cleanup();
});

describe('should create a temporary directory', () => {
  test('Directory exists', () => {
    expect(fs.existsSync(tmp.path)).toBe(true);
  });
});

describe('replace the correct content', () => {
  beforeEach(() => {
    fs.writeFileSync(TEST_FILE_TO_REPLACE, toReplace);
  });

  afterEach(() => {
    fs.rmSync(TEST_FILE_TO_REPLACE);
  });

  const testData = {
    name: 'John',
    email: 'HbVYf@example.com',
    username: 'johnny',
    repository: 'https://github.com/username/repo',
    proj_name: 'My Project',
    proj_short_desc: 'Project with some description',
    proj_long_desc: 'Project with some description that is really long',
    docs_url: 'https://example.com/docs',
  } satisfies ProjectInfo;

  test('File content is replaced matches "replaced" string', async () => {
    expect(
      await replaceInFile(TEST_FILE_TO_REPLACE, tmp.path, testData),
    ).toBeUndefined();
    expect(fs.readFileSync(TEST_FILE_TO_REPLACE, 'utf8')).toEqual(replaced);
  });
});
