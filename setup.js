const readline = require('readline');
const fs = require('fs');
const { execSync } = require('child_process');

const templateSync = `
.github/ISSUE_TEMPLATE/*
.github/CODEOWNERS
.github/CODESTYLE.md
.github/PULL_REQUEST_TEMPLATE.md
.github/SECURITY.md
CITATION.cff
LICENSE
README.md`;

function handleError(error) {
  if (error.code !== 'ENOENT' && error.code !== 'EEXIST') {
    console.error(error);
    process.exit(1);
  }
}

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

const question = (query) =>
  new Promise((resolve) => {
    rl.question(query, resolve);
  });

(async () => {
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

  const confirm = await question('Confirm? (y/n)\n=> ');

  if (confirm.toLowerCase() !== 'y') {
    console.log('Aborted.');
    rl.close();
    return;
  } else {
    console.log('\nWriting files...');

    // Remove prettier stuff
    try {
      fs.unlinkSync('package.json');
      fs.unlinkSync('package-lock.json');
      fs.rmSync('node_modules', { recursive: true });
    } catch (error) {
      handleError(error);
    }

    // Writing general stuff
    const filesToUpdate = ['LICENSE', 'CITATION.cff'];
    filesToUpdate.forEach((fileName) => {
      try {
        let fileContent = fs.readFileSync(`./template/${fileName}`, 'utf8');
        fileContent = fileContent
          .replace(/{{REPOSITORY}}/g, `${username}/${repository}`)
          .replace(/{{PROJECT_NAME}}/g, proj_name)
          .replace(/{{PROJECT_SHORT_DESCRIPTION}}/g, proj_short_desc)
          .replace(/{{PROJECT_LONG_DESCRIPTION}}/g, proj_long_desc)
          .replace(/{{DOCS_URL}}/g, docs_url)
          .replace(/{{EMAIL}}/g, email)
          .replace(/{{USERNAME}}/g, username)
          .replace(/{{NAME}}/g, name);
        fs.writeFileSync(`./template/${fileName}`, fileContent);
      } catch (error) {
        // it's a bit different here, won't touch this for now
        if (error.code !== 'ENOENT' && error.code !== 'EEXIST') {
          console.error(error);
          process.exit(1);
        } else {
          console.log(`File ${fileName} not found.`);
        }
      }
    });

    // Write CODEOWNERS
    try {
      fs.appendFileSync('./template/.github/CODEOWNERS', `* @${username}`);
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
    const up_to_date = await question(
      'Would you like to keep up-to-date with the template? (y/n)\n=> ',
    );
    if (up_to_date.toLowerCase() === 'y') {
      console.log('Writing ignore file...');
      try {
        fs.appendFileSync('./template/.templatesyncignore', templateSync);
        fs.appendFileSync(
          './template/.github/settings.yml',
          `
        - name: 'CI: Template Sync'
        color: AEB1C2
        description: Sync with upstream template
        `,
        );
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

    const keep_script = await question(
      'Would you like to keep this setup script? (y/n)\n=> ',
    );
    if (keep_script.toLowerCase() !== 'y') {
      console.log('Removing setup script...');
      try {
        fs.unlinkSync(__filename);
      } catch (error) {
        handleError(error);
      }
    } else {
      console.log('Okay.');
    }
    console.log(
      '\nDone!\nIf you encounter any issues, please report it here: https://github.com/caffeine-addictt/template/issues/new?assignees=caffeine-addictt&labels=Type%3A+Bug&projects=&template=1-bug-report.md&title=[Bug]+',
    );
    rl.close();
  }
})();
