"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const readline_1 = __importDefault(require("readline"));
const io_util_1 = require("./io-util");
/**
 * For interacting with stdin/stdout
 */
const rl = readline_1.default.createInterface({
    input: process.stdin,
    output: process.stdout,
});
/** Prompt user for input */
const question = (query, validator = [], trimWhitespace = true) => new Promise((resolve) => rl.question(query, (s) => {
    if (trimWhitespace)
        s = s.trim();
    validator.forEach((v) => {
        if (!v.validate(s)) {
            v.onError();
            process.exit(1);
        }
    });
    resolve(s);
}));
/** Ask for project information */
const fetchInfo = async (cleanup) => {
    const name = await question('Name? (This will go on the LICENSE)\n=> ');
    const email = await question('Email?\n=> ', [
        {
            validate: (s) => /.+@.+\..+/.test(s),
            onError: () => console.log('Invalid email!'),
        },
    ]);
    const username = await question('Username? (https://github.com/<username>)\n=> ');
    const repository = await question(`Repository? (https://github.com/${username}/<repo>)\n=> `);
    const proj_name = await question('Project name?\n=> ');
    const proj_short_desc = await question('Short description?\n=> ');
    const proj_long_desc = await question('Long description?\n=> ');
    const docs_url = await question('Documentation URL?\n=> ');
    const assignees = Array.from(new Set([
        ...(await question('Additional issue assignees? (Usernames comma separated)\n=> ')).split(','),
        // Add CODEOWNERS
        username,
    ]
        .map((s) => s.trim())
        .filter((s) => s.length > 0))).join(`', '`);
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
    };
};
/**
 * The main logic
 */
const { func: main } = (0, io_util_1.withTempDir)('caffeine-addictt-template-', async (tempDir) => {
    const data = await fetchInfo(() => rl.close());
    /// ######################################## //
    // Stage 1: Update file content in template/ //
    // ######################################### //
    console.log('\nWriting files...');
    const filesToUpdate = fs_1.default.readdirSync('./template', {
        recursive: true,
    });
    // Use async
    await Promise.all(filesToUpdate.map((filename) => (async () => {
        const filePath = path_1.default.join('./template', filename);
        const fileInfo = fs_1.default.statSync(filePath);
        if (fileInfo.isDirectory()) {
            return;
        }
        await (0, io_util_1.replaceInFile)(filePath, tempDir, data);
    })()));
    // Write CODEOWNERS
    fs_1.default.appendFileSync('./template/.github/CODEOWNERS', `* @${data.username}`);
    // ########################################## //
    // Stage 2: Move files from template/ to root //
    // ########################################## //
    console.log('Moving files...');
    // Delete .github/ directory
    fs_1.default.rmSync('./.github', { recursive: true, force: true });
    const filesToMove = fs_1.default.readdirSync('./template', {
        recursive: true,
    });
    filesToMove.forEach((file) => {
        const filePath = path_1.default.join('./template', file);
        const fileInfo = fs_1.default.statSync(filePath);
        if (fileInfo.isDirectory()) {
            fs_1.default.mkdirSync(`${file}`);
            return;
        }
        fs_1.default.renameSync(filePath, `./${file}`);
    });
    fs_1.default.rmSync('./template', { recursive: true, force: true });
    // ################# //
    // Stage 3: Clean up //
    // ################# //
    // Only add `force: true` for files or directories that
    // will only exist if some development task was carried out
    // like eslintcache
    console.log('Cleaning up...');
    // Js
    fs_1.default.unlinkSync('package.json');
    fs_1.default.unlinkSync('package-lock.json');
    // Ts
    fs_1.default.unlinkSync('tsconfig.json');
    fs_1.default.rmSync('src', { recursive: true });
    // Tests
    fs_1.default.unlinkSync('babel.config.cjs');
    fs_1.default.rmSync('tests', { recursive: true });
    // Linting
    fs_1.default.unlinkSync('.eslintignore');
    fs_1.default.unlinkSync('.prettierignore');
    fs_1.default.unlinkSync('eslint.config.mjs');
    fs_1.default.rmSync('.eslintcache', { force: true });
    // Syncing
    fs_1.default.unlinkSync('.templatesyncignore');
    // Git
    fs_1.default.unlinkSync('.gitignore');
    // Node
    fs_1.default.rmSync('node_modules', { recursive: true, force: true });
    // Clean up dist
    fs_1.default.unlinkSync(__filename);
    fs_1.default.rmSync('dist', { recursive: true });
    // Generate src and test
    fs_1.default.mkdirSync('src');
    fs_1.default.mkdirSync('tests');
    // Final stdout
    console.log('\nDone!\nIf you encounter any issues, please report it here: https://github.com/caffeine-addictt/template/issues/new?assignees=caffeine-addictt&labels=Type%3A+Bug&projects=&template=1-bug-report.md&title=[Bug]+');
    rl.close();
});
main();
