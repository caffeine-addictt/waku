import fs from 'fs';
import path from 'path';
import stream from 'stream';
import readline from 'readline';

import { handleError } from './error';
import type { ProjectInfo } from './types';

/** Create a temp directory with automatic cleanup */
export type withTempDirFunc<T = unknown> = (
  prefix: string,
  func: (dirPath: string) => T,
  autoCleanup?: boolean,
) => { cleanup: () => void; func: () => T; path: string };
export const withTempDir: withTempDirFunc = (
  prefix,
  func,
  autoCleanup = true,
) => {
  const dirPath = fs.mkdtempSync(prefix);

  const cleanup = () => fs.rmSync(dirPath, { recursive: true, force: true });

  return {
    path: dirPath,
    cleanup: cleanup,
    func: () => {
      try {
        const returnVal = func(dirPath);
        if (autoCleanup) cleanup();
        return returnVal;
      } catch (e) {
        handleError(e);
      }
    },
  };
};

/** Replace string in file buffer */
export const replaceInFile = (
  filePath: string,
  tempDir: string,
  data: ProjectInfo,
): Promise<void> =>
  new Promise((resolve) => {
    // Revert to legacy

    let fileContent = fs.readFileSync(filePath, 'utf8');
    fileContent = fileContent
      .replace(/{{REPOSITORY}}/g, data.repository)
      .replace(/{{PROJECT_NAME}}/g, data.proj_name)
      .replace(/{{PROJECT_SHORT_DESCRIPTION}}/g, data.proj_short_desc)
      .replace(/{{PROJECT_LONG_DESCRIPTION}}/g, data.proj_long_desc)
      .replace(/{{DOCS_URL}}/g, data.docs_url)
      .replace(/{{EMAIL}}/g, data.email)
      .replace(/{{USERNAME}}/g, data.username)
      .replace(/{{NAME}}/g, data.name);

    resolve(fs.writeFileSync(filePath, fileContent));
    return;

    // There was an attempt at buffering R/W
    const outputPath = path.join(tempDir, path.basename(filePath));
    fs.writeFileSync(outputPath, '');

    const inStream = fs.createReadStream(filePath);
    const outStream = new stream.Writable();

    readline
      .createInterface({
        input: inStream,
        output: outStream,
        terminal: false,
      })
      .on('line', (line) => {
        fs.appendFileSync(
          outputPath,
          line
            .replace(/{{REPOSITORY}}/g, `${data.username}/${data.repository}`)
            .replace(/{{PROJECT_NAME}}/g, data.proj_name)
            .replace(/{{PROJECT_SHORT_DESCRIPTION}}/g, data.proj_short_desc)
            .replace(/{{PROJECT_LONG_DESCRIPTION}}/g, data.proj_long_desc)
            .replace(/{{DOCS_URL}}/g, data.docs_url)
            .replace(/{{EMAIL}}/g, data.email)
            .replace(/{{USERNAME}}/g, data.username)
            .replace(/{{NAME}}/g, data.name) + '\n',
        );
      })
      .on('close', () => {
        // Move from temp back to original
        fs.renameSync(outputPath, filePath);
        resolve();
      });
  });
