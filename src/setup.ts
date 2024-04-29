import fs from 'fs';
import path from 'path';
import readline from 'readline';

import type { ProjectInfo } from './types';
import { replaceInFile, withTempDir } from './io-util';
import { handleError, type NodeErrorMaybe } from './error';

// Constants
const templateSyncIgnore = `
.github/ISSUE_TEMPLATE/*
.github/CODEOWNERS
.github/CODESTYLE.md
.github/PULL_REQUEST_TEMPLATE.md
.github/SECURITY.md
CITATION.cff
LICENSE
README.md` as const;

const templateSyncLabel = `
  - name: 'CI: Template Sync'
  color: AEB1C2
  description: Sync with upstream template
` as const;

/**
 * For interacting with stdin/stdout
 */
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

/** Prompt user for input */
const question = (query: string): Promise<string> =>
  new Promise((resolve) => rl.question(query, resolve));

/** Ask for project information */
const fetchInfo = async (
  cleanup: () => void | unknown,
): Promise<ProjectInfo> => {
  const name = await question('Name? (This will go on the LICENSE)\n=> ');
  const email = await question('Email?\n=> ');
  const username = await question(
    'Username? (https://github.com/<username>)\n=> ',
  );
  const repository = await question(
    'Repository? ((https://github.com/$username/<repo>\n=> ',
  );
  const proj_name = await question('Project name?\n=> ');
  const proj_short_desc = await question('Short description?\n=> ');
  const proj_long_desc = await question('Long description?\n=> ');
  const docs_url = await question('Documentation URL?\n=> ');

  console.log('\n\n===== Log =====');
  console.log('Name:', name);
  console.log('Email:', email);
  console.log('Username:', username);
  console.log('Repository:', repository);
  console.log('Project name:', proj_name);
  console.log('Project short description:', proj_short_desc);
  console.log('Project long description:', proj_long_desc);
  console.log('Docs URL:', docs_url);
  console.log('================');

  // Guard clause for confirmation
  if ((await question('Confirm? (y/n)\n=> ')).toLowerCase() !== 'y') {
    console.log('Aborted.');
    cleanup();
    process.exit(1);
  }

  return {
    name: name,
    email: email,
    username: username,
    repository: repository,
    proj_name: proj_name,
    proj_short_desc: proj_short_desc,
    proj_long_desc: proj_long_desc,
    docs_url: docs_url,
  } satisfies ProjectInfo;
};

/**
 * The main logic
 */
const { func: main } = withTempDir(
  'caffeine-addictt-template-',
  async (tempDir: string) => {
    const data = await fetchInfo(() => rl.close());

    console.log('\nWriting files...');

    // Writing general stuff
    const filesToUpdate = fs.readdirSync('./template', {
      recursive: true,
    }) as string[];
    filesToUpdate.forEach(async (relativePath) => {
      const filePath = path.join('./template', relativePath);
      try {
        const fileInfo = fs.statSync(filePath);
        if (fileInfo.isDirectory()) {
          return;
        }

        await replaceInFile(filePath, tempDir, data);
      } catch (error) {
        // it's a bit different here, won't touch this for now
        if (
          (error as NodeErrorMaybe)?.code !== 'ENOENT' &&
          (error as NodeErrorMaybe)?.code !== 'EEXIST'
        ) {
          console.error(error);
          process.exit(1);
        } else {
          console.log(`File ${filePath} not found.`);
        }
      }
    });

    // Write CODEOWNERS
    try {
      fs.appendFileSync('./template/.github/CODEOWNERS', `* @${data.username}`);
    } catch (error) {
      // also different here
      if (
        (error as NodeErrorMaybe)?.code !== 'ENOENT' &&
        (error as NodeErrorMaybe)?.code !== 'EEXIST'
      ) {
        console.error(error);
        process.exit(1);
      } else {
        fs.renameSync('./template/.github/CODEOWNERS', '.github/CODEOWNERS');
      }
    }

    // Optional keep up-to-date
    if (
      (
        await question(
          'Would you like to keep up-to-date with the template? (y/n)\n=> ',
        )
      ).toLowerCase() === 'y'
    ) {
      console.log('Writing ignore file...');
      try {
        fs.appendFileSync('./template/.templatesyncignore', templateSyncIgnore);
        fs.appendFileSync('./template/.github/settings.yml', templateSyncLabel);
        fs.renameSync('./template/.templatesyncignore', '.templatesyncignore');
        console.log(
          'You can view more configuration here: https://github.com/AndreasAugustin/actions-template-sync',
        );
      } catch (error) {
        handleError(error);
      }
    } else {
      console.log('Removing syncing workflow...');
      try {
        fs.unlinkSync('./template/.github/workflows/sync-template.yml');
      } catch (error) {
        handleError(error);
      }
    }

    // Move from template
    console.log('Moving files...');
    try {
      const filesToMove = fs.readdirSync('./template');
      filesToMove.forEach((file) => {
        fs.renameSync(`./template/${file}`, `./${file}`);
      });
      fs.rmSync('./template', { recursive: true });
      fs.rmSync('.github', { recursive: true });
      fs.renameSync('./template/.github', '.github');
    } catch (error) {
      handleError(error);
    }

    // Clean up development stuff
    console.log('Cleaning up...');
    try {
      // Js
      fs.unlinkSync('package.json');
      fs.unlinkSync('package-lock.json');

      // Ts
      fs.unlinkSync('tsconfig.json');
      fs.rmSync('src', { recursive: true });
      fs.rmSync('tests', { recursive: true });

      // Linting
      fs.unlinkSync('.prettierignore');
      fs.unlinkSync('eslint.config.mjs');

      // Git
      fs.unlinkSync('.gitignore');

      // Node
      fs.rmSync('node_modules', { recursive: true });
    } catch (error) {
      handleError(error);
    }

    // Clean up dist
    try {
      fs.unlinkSync(__filename);
      fs.rmSync('dist', { recursive: true });
    } catch (error) {
      handleError(error);
    }

    // Final stdout
    console.log(
      '\nDone!\nIf you encounter any issues, please report it here: https://github.com/caffeine-addictt/template/issues/new?assignees=caffeine-addictt&labels=Type%3A+Bug&projects=&template=1-bug-report.md&title=[Bug]+',
    );
    rl.close();
  },
);

main();
