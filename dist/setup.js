"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const readline_1 = __importDefault(require("readline"));
const io_util_1 = require("./io-util");
const error_1 = require("./error");
// Constants
const templateSyncIgnore = `
.github/ISSUE_TEMPLATE/*
.github/CODEOWNERS
.github/CODESTYLE.md
.github/PULL_REQUEST_TEMPLATE.md
.github/SECURITY.md
CITATION.cff
LICENSE
README.md`;
const templateSyncLabel = `
  - name: 'CI: Template Sync'
  color: AEB1C2
  description: Sync with upstream template
`;
/**
 * For interacting with stdin/stdout
 */
const rl = readline_1.default.createInterface({
    input: process.stdin,
    output: process.stdout,
});
/** Prompt user for input */
const question = (query) => new Promise((resolve) => rl.question(query, resolve));
/** Ask for project information */
const fetchInfo = async (cleanup) => {
    const name = await question('Name? (This will go on the LICENSE)\n=> ');
    const email = await question('Email?\n=> ');
    const username = await question('Username? (https://github.com/<username>)\n=> ');
    const repository = await question('Repository? ((https://github.com/$username/<repo>\n=> ');
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
    };
};
/**
 * The main logic
 */
const { func: main } = (0, io_util_1.withTempDir)('caffeine-addictt-template-', async (tempDir) => {
    const data = await fetchInfo(() => rl.close());
    console.log('\nWriting files...');
    // Writing general stuff
    const filesToUpdate = fs_1.default.readdirSync('./template', {
        recursive: true,
    });
    filesToUpdate.forEach(async (relativePath) => {
        const filePath = path_1.default.join('./template', relativePath);
        try {
            const fileInfo = fs_1.default.statSync(filePath);
            if (fileInfo.isDirectory()) {
                return;
            }
            await (0, io_util_1.replaceInFile)(filePath, tempDir, data);
        }
        catch (error) {
            // it's a bit different here, won't touch this for now
            if (error?.code !== 'ENOENT' &&
                error?.code !== 'EEXIST') {
                console.error(error);
                process.exit(1);
            }
            else {
                console.log(`File ${filePath} not found.`);
            }
        }
    });
    // Write CODEOWNERS
    try {
        fs_1.default.appendFileSync('./template/.github/CODEOWNERS', `* @${data.username}`);
    }
    catch (error) {
        // also different here
        if (error?.code !== 'ENOENT' &&
            error?.code !== 'EEXIST') {
            console.error(error);
            process.exit(1);
        }
        else {
            fs_1.default.renameSync('./template/.github/CODEOWNERS', '.github/CODEOWNERS');
        }
    }
    // Optional keep up-to-date
    if ((await question('Would you like to keep up-to-date with the template? (y/n)\n=> ')).toLowerCase() === 'y') {
        console.log('Writing ignore file...');
        try {
            fs_1.default.appendFileSync('./template/.templatesyncignore', templateSyncIgnore);
            fs_1.default.appendFileSync('./template/.github/settings.yml', templateSyncLabel);
            fs_1.default.renameSync('./template/.templatesyncignore', '.templatesyncignore');
            console.log('You can view more configuration here: https://github.com/AndreasAugustin/actions-template-sync');
        }
        catch (error) {
            (0, error_1.handleError)(error);
        }
    }
    else {
        console.log('Removing syncing workflow...');
        try {
            fs_1.default.unlinkSync('./template/.github/workflows/sync-template.yml');
        }
        catch (error) {
            (0, error_1.handleError)(error);
        }
    }
    // Move from template
    console.log('Moving files...');
    try {
        const filesToMove = fs_1.default.readdirSync('./template');
        filesToMove.forEach((file) => {
            fs_1.default.renameSync(`./template/${file}`, `./${file}`);
        });
        fs_1.default.rmSync('./template', { recursive: true });
        fs_1.default.rmSync('.github', { recursive: true });
        fs_1.default.renameSync('./template/.github', '.github');
    }
    catch (error) {
        (0, error_1.handleError)(error);
    }
    // Clean up development stuff
    console.log('Cleaning up...');
    try {
        // Js
        fs_1.default.unlinkSync('package.json');
        fs_1.default.unlinkSync('package-lock.json');
        // Ts
        fs_1.default.unlinkSync('tsconfig.json');
        fs_1.default.rmSync('src', { recursive: true });
        fs_1.default.rmSync('tests', { recursive: true });
        // Linting
        fs_1.default.unlinkSync('.prettierignore');
        fs_1.default.unlinkSync('eslint.config.mjs');
        // Git
        fs_1.default.unlinkSync('.gitignore');
        // Node
        fs_1.default.rmSync('node_modules', { recursive: true });
    }
    catch (error) {
        (0, error_1.handleError)(error);
    }
    // Clean up dist
    try {
        fs_1.default.unlinkSync(__filename);
        fs_1.default.rmSync('dist', { recursive: true });
    }
    catch (error) {
        (0, error_1.handleError)(error);
    }
    // Final stdout
    console.log('\nDone!\nIf you encounter any issues, please report it here: https://github.com/caffeine-addictt/template/issues/new?assignees=caffeine-addictt&labels=Type%3A+Bug&projects=&template=1-bug-report.md&title=[Bug]+');
    rl.close();
});
main();
