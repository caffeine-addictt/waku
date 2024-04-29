"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.replaceInFile = exports.withTempDir = void 0;
var fs_1 = require("fs");
var path_1 = require("path");
var stream_1 = require("stream");
var readline_1 = require("readline");
var error_1 = require("./error");
var withTempDir = function (prefix, func, autoCleanup) {
    if (autoCleanup === void 0) { autoCleanup = true; }
    var dirPath = fs_1.default.mkdtempSync(prefix);
    var cleanup = function () { return fs_1.default.rmSync(dirPath, { recursive: true, force: true }); };
    return {
        path: dirPath,
        cleanup: cleanup,
        func: function () {
            try {
                var returnVal = func(dirPath);
                if (autoCleanup)
                    cleanup();
                return returnVal;
            }
            catch (e) {
                (0, error_1.handleError)(e);
            }
        },
    };
};
exports.withTempDir = withTempDir;
/** Replace string in file buffer */
var replaceInFile = function (filePath, tempDir, data) {
    return new Promise(function (resolve) {
        // Revert to legacy
        var fileContent = fs_1.default.readFileSync(filePath, 'utf8');
        fileContent = fileContent
            .replace(/{{REPOSITORY}}/g, data.repository)
            .replace(/{{PROJECT_NAME}}/g, data.proj_name)
            .replace(/{{PROJECT_SHORT_DESCRIPTION}}/g, data.proj_short_desc)
            .replace(/{{PROJECT_LONG_DESCRIPTION}}/g, data.proj_long_desc)
            .replace(/{{DOCS_URL}}/g, data.docs_url)
            .replace(/{{EMAIL}}/g, data.email)
            .replace(/{{USERNAME}}/g, data.username)
            .replace(/{{NAME}}/g, data.name);
        resolve(fs_1.default.writeFileSync(filePath, fileContent));
        return;
        // There was an attempt at buffering R/W
        var outputPath = path_1.default.join(tempDir, path_1.default.basename(filePath));
        fs_1.default.writeFileSync(outputPath, '');
        var inStream = fs_1.default.createReadStream(filePath);
        var outStream = new stream_1.default.Writable();
        readline_1.default
            .createInterface({
            input: inStream,
            output: outStream,
            terminal: false,
        })
            .on('line', function (line) {
            fs_1.default.appendFileSync(outputPath, line
                .replace(/{{REPOSITORY}}/g, "".concat(data.username, "/").concat(data.repository))
                .replace(/{{PROJECT_NAME}}/g, data.proj_name)
                .replace(/{{PROJECT_SHORT_DESCRIPTION}}/g, data.proj_short_desc)
                .replace(/{{PROJECT_LONG_DESCRIPTION}}/g, data.proj_long_desc)
                .replace(/{{DOCS_URL}}/g, data.docs_url)
                .replace(/{{EMAIL}}/g, data.email)
                .replace(/{{USERNAME}}/g, data.username)
                .replace(/{{NAME}}/g, data.name) + '\n');
        })
            .on('close', function () {
            // Move from temp back to original
            fs_1.default.renameSync(outputPath, filePath);
            resolve();
        });
    });
};
exports.replaceInFile = replaceInFile;
