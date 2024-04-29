import fs from 'fs';
import path from 'path';
import readline from 'readline';

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

// Errors
interface NodeErrorMaybe extends Error {
  code?: string;
}
/** Handle errors and conditionally exit program */
const handleError = (error: Error | unknown) => {
  if (
    (error as NodeErrorMaybe)?.code !== 'ENOENT' &&
    (error as NodeErrorMaybe)?.code !== 'EEXIST'
  ) {
    console.error(error);
    process.exit(1);
  }
};

/**
 * For interacting with stdin/stdout
 */
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

/** Prompt user for input */
const question = (query: string): string =>
  await new Promise((resolve) => {
    rl.question(query, resolve);
  });

/** Create a temp directory with automatic cleanup */
type withTempDirFunc<T = unknown> = (
  prefix: string,
  func: (dirPath: string) => T,
) => { cleanup: () => void; func: () => T };
const withTempDir: withTempDirFunc = (prefix, func) => {
  let dirPath: string;

  const cleanup = () => dirPath && fs.rmdirSync(dirPath);

  return {
    cleanup: cleanup,
    func: () => {
      dirPath = fs.mkdtempSync(prefix);
      try {
        const returnVal = func(dirPath);
        cleanup();
        return returnVal;
      } catch (e) {
        handleError(e);
      }
    },
  };
};

/**
 * Make 1
 */

interface InfoType {
  name: string;
  email: string;
  username: string;
  repository: string;
  proj_name: string;
  proj_short_desc: string;
  proj_long_desc: string;
  docs_url: string;
}

/** Ask for project information */
function fetchInfo(cleanup: () => void | unknown): InfoType {
  const name = question('Name? (This will go on the LICENSE)\n=> ');
  const email = question('Email?\n=> ');
  const username = question('Username? (https://github.com/<username>)\n=> ');
  const repository = question(
    'Repository? ((https://github.com/$username/<repo>\n=> ',
  );
  const proj_name = question('Project name?\n=> ');
  const proj_short_desc = question('Short description?\n=> ');
  const proj_long_desc = question('Long description?\n=> ');
  const docs_url = question('Documentation URL?\n=> ');

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
  if (question('Confirm? (y/n)\n=> ').toLowerCase() !== 'y') {
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
  } satisfies InfoType;
}

/**
 * The main logic
 */
(async () => {
  const data = fetchInfo(() => rl.close());

  console.log('\nWriting files...');

  // Remove prettier stuff
  try {
    fs.unlinkSync('package.json');
    fs.unlinkSync('package-lock.json');
    fs.unlinkSync('.prettierignore');
    fs.rmSync('node_modules', { recursive: true });
  } catch (error) {
    handleError(error);
  }

  // Writing general stuff
  const tempDir = fs.mkdtempSync('caffeine-addictt-template-');

  const filesToUpdate = fs.readdirSync('./template', {
    recursive: true,
  }) as string[];
  filesToUpdate.forEach((relativePath) => {
    const filePath = path.join('./template', relativePath);
    try {
      const fileInfo = fs.statSync(filePath);
      if (fileInfo.isDirectory()) {
        return;
      }

      const readBuffer = fs.readAs;

      let fileContent = fs.readFileSync(filePath, 'utf8');
      fileContent = fileContent
        .replace(/{{REPOSITORY}}/g, `${data.username}/${data.repository}`)
        .replace(/{{PROJECT_NAME}}/g, data.proj_name)
        .replace(/{{PROJECT_SHORT_DESCRIPTION}}/g, data.proj_short_desc)
        .replace(/{{PROJECT_LONG_DESCRIPTION}}/g, data.proj_long_desc)
        .replace(/{{DOCS_URL}}/g, data.docs_url)
        .replace(/{{EMAIL}}/g, data.email)
        .replace(/{{USERNAME}}/g, data.username)
        .replace(/{{NAME}}/g, data.name);

      fs.writeFileSync(filePath, fileContent);
    } catch (error) {
      // it's a bit different here, won't touch this for now
      if (error.code !== 'ENOENT' && error.code !== 'EEXIST') {
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
    if (error.code !== 'ENOENT' && error.code !== 'EEXIST') {
      console.error(error);
      process.exit(1);
    } else {
      fs.renameSync('./template/.github/CODEOWNERS', '.github/CODEOWNERS');
    }
  }

  // Optional keep up-to-date
  const up_to_date = question(
    'Would you like to keep up-to-date with the template? (y/n)\n=> ',
  );
  if (up_to_date.toLowerCase() === 'y') {
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

  // Remove setup script
  if (
    question(
      'Would you like to keep this setup script? (y/n)\n=> ',
    ).toLowerCase() !== 'y'
  ) {
    console.log('Removing setup script...');
    try {
      fs.unlinkSync(__filename);
    } catch (error) {
      handleError(error);
    }
  } else {
    console.log('Okay.');
  }

  // Final stdout
  console.log(
    '\nDone!\nIf you encounter any issues, please report it here: https://github.com/caffeine-addictt/template/issues/new?assignees=caffeine-addictt&labels=Type%3A+Bug&projects=&template=1-bug-report.md&title=[Bug]+',
  );
  rl.close();
})();
