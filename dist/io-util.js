"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.replaceInFile = exports.withTempDir = void 0;
const fs_1 = __importDefault(require("fs"));
const withTempDir = (prefix, func, autoCleanup = true) => {
    const dirPath = fs_1.default.mkdtempSync(prefix);
    const cleanup = () => fs_1.default.rmSync(dirPath, { recursive: true, force: true });
    return {
        path: dirPath,
        cleanup: cleanup,
        func: () => {
            const returnVal = func(dirPath);
            if (autoCleanup)
                cleanup();
            return returnVal;
        },
    };
};
exports.withTempDir = withTempDir;
/** Replace string in file buffer */
const replaceInFile = (filePath, _tempDir, data) => new Promise((resolve) => {
    // Revert to legacy
    let fileContent = fs_1.default.readFileSync(filePath, 'utf8');
    fileContent = fileContent
        .replace(/{{REPOSITORY}}/g, data.repository)
        .replace(/{{PROJECT_NAME}}/g, data.proj_name)
        .replace(/{{PROJECT_SHORT_DESCRIPTION}}/g, data.proj_short_desc)
        .replace(/{{PROJECT_LONG_DESCRIPTION}}/g, data.proj_long_desc)
        .replace(/{{DOCS_URL}}/g, data.docs_url)
        .replace(/{{EMAIL}}/g, data.email)
        .replace(/{{USERNAME}}/g, data.username)
        .replace(/{{NAME}}/g, data.name)
        .replace(/{{ASSIGNEES}}/g, data.assignees);
    resolve(fs_1.default.writeFileSync(filePath, fileContent));
    return;
    // There was an attempt at buffering R/W
    // const outputPath = path.join(tempDir, path.basename(filePath));
    // fs.writeFileSync(outputPath, '');
    //
    // const inStream = fs.createReadStream(filePath);
    // const outStream = new stream.Writable();
    //
    // readline
    //   .createInterface({
    //     input: inStream,
    //     output: outStream,
    //     terminal: false,
    //   })
    //   .on('line', (line) => {
    //     fs.appendFileSync(
    //       outputPath,
    //       line
    //         .replace(/{{REPOSITORY}}/g, `${data.username}/${data.repository}`)
    //         .replace(/{{PROJECT_NAME}}/g, data.proj_name)
    //         .replace(/{{PROJECT_SHORT_DESCRIPTION}}/g, data.proj_short_desc)
    //         .replace(/{{PROJECT_LONG_DESCRIPTION}}/g, data.proj_long_desc)
    //         .replace(/{{DOCS_URL}}/g, data.docs_url)
    //         .replace(/{{EMAIL}}/g, data.email)
    //         .replace(/{{USERNAME}}/g, data.username)
    //         .replace(/{{NAME}}/g, data.name) + '\n',
    //     );
    //   })
    //   .on('close', () => {
    //     // Move from temp back to original
    //     fs.renameSync(outputPath, filePath);
    //     resolve();
    //   });
});
exports.replaceInFile = replaceInFile;
