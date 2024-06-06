import fs from 'fs';
import path from 'path';
import readline from 'readline';

import type { ProjectInfo } from './types';
import { replaceInFile, withTempDir } from './io-util';

/**
 * For interacting with stdin/stdout
 */
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

interface Validator {
  validate: (s: string) => boolean;
  onError: () => void;
}

/** Prompt user for input */
const question = (
  query: string,
  validator: Validator[] = [],
  trimWhitespace: boolean = true,
): Promise<string> =>
  new Promise((resolve) =>
    rl.question(query, (s: string) => {
      if (trimWhitespace) s = s.trim();
      validator.forEach((v) => {
        if (!v.validate(s)) {
          v.onError();
          process.exit(1);
        }
      });
      resolve(s);
    }),
  );

/** Ask for project information */
const fetchInfo = async (
  cleanup: () => void | unknown,
): Promise<ProjectInfo> => {
  const name = await question('Name? (This will go on the LICENSE)\n=> ');
  const email = await question('Email?\n=> ', [
    {
      validate: (s: string) => /.+@.+\..+/.test(s),
      onError: () => console.log('Invalid email!'),
    },
  ]);
  const username = await question(
    'Username? (https://github.com/<username>)\n=> ',
  );
  const repository = await question(
    `Repository? (https://github.com/${username}/<repo>)\n=> `,
  );
  const proj_name = await question('Project name?\n=> ');
  const proj_short_desc = await question('Short description?\n=> ');
  const proj_long_desc = await question('Long description?\n=> ');
  const docs_url = await question('Documentation URL?\n=> ');
  const assignees = Array.from(
    new Set(
      [
        ...(
          await question(
            'Additional issue assignees? (Usernames comma separated)\n=> ',
          )
        ).split(','),

        // Add CODEOWNERS
        username,
      ]
        .map((s: string) => s.trim())
        .filter((s: string) => s.length > 0),
    ),
  ).join(`', '`);

  console.log('\n\n===== Log =====');
  console.log('Name:', name);
  console.log('Email:', email);
  console.log('Username:', username);
  console.log('Repository:', repository);
  console.log('Project name:', proj_name);
  console.log('Project short description:', proj_short_desc);
  console.log('Project long description:', proj_long_desc);
  console.log('Docs URL:', docs_url);
  console.log('Issue assignees:', `['${assignees}']`);
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
    assignees: assignees,
  } satisfies ProjectInfo;
};

/**
 * The main logic
 */
const { func: main } = withTempDir(
  'caffeine-addictt-template-',
  async (tempDir: string) => {
    const data = await fetchInfo(() => rl.close());

    /// ######################################## //
    // Stage 1: Update file content in template/ //
    // ######################################### //
    console.log('\nWriting files...');

    const filesToUpdate = fs.readdirSync('./template', {
      recursive: true,
    }) as string[];

    // Use async
    await Promise.all(
      filesToUpdate.map((filename) =>
        (async () => {
          const filePath = path.join('./template', filename);

          const fileInfo = fs.statSync(filePath);
          if (fileInfo.isDirectory()) {
            return;
          }

          await replaceInFile(filePath, tempDir, data);
        })(),
      ),
    );

    // Write CODEOWNERS
    fs.appendFileSync('./template/.github/CODEOWNERS', `* @${data.username}`);

    // ########################################## //
    // Stage 2: Move files from template/ to root //
    // ########################################## //
    console.log('Moving files...');

    // Delete .github/ directory
    fs.rmSync('./.github', { recursive: true, force: true });

    const filesToMove = fs.readdirSync('./template', {
      recursive: true,
    }) as string[];
    filesToMove.forEach((file) => {
      const filePath = path.join('./template', file);

      const fileInfo = fs.statSync(filePath);
      if (fileInfo.isDirectory()) {
        fs.mkdirSync(`${file}`);
        return;
      }
      fs.renameSync(filePath, `./${file}`);
    });
    fs.rmSync('./template', { recursive: true, force: true });

    // ################# //
    // Stage 3: Clean up //
    // ################# //
    // Only add `force: true` for files or directories that
    // will only exist if some development task was carried out
    // like eslintcache
    console.log('Cleaning up...');

    // Js
    fs.unlinkSync('package.json');
    fs.unlinkSync('package-lock.json');

    // Ts
    fs.unlinkSync('tsconfig.json');
    fs.rmSync('src', { recursive: true });

    // Tests
    fs.unlinkSync('babel.config.cjs');
    fs.rmSync('tests', { recursive: true });

    // Linting
    fs.unlinkSync('.eslintignore');
    fs.unlinkSync('.prettierignore');
    fs.unlinkSync('eslint.config.mjs');
    fs.rmSync('.eslintcache', { force: true });

    // Syncing
    fs.unlinkSync('.templatesyncignore');

    // Git
    fs.unlinkSync('.gitignore');

    // Node
    fs.rmSync('node_modules', { recursive: true, force: true });

    // Clean up dist
    fs.unlinkSync(__filename);
    fs.rmSync('dist', { recursive: true });

    // Generate src and test
    fs.mkdirSync('src');
    fs.mkdirSync('tests');

    // Final stdout
    console.log(
      '\nDone!\nIf you encounter any issues, please report it here: https://github.com/caffeine-addictt/template/issues/new?assignees=caffeine-addictt&labels=Type%3A+Bug&projects=&template=1-bug-report.md&title=[Bug]+',
    );
    rl.close();
  },
);

main();
