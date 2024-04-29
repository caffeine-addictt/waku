"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
var fs_1 = require("fs");
var path_1 = require("path");
var readline_1 = require("readline");
var io_util_1 = require("./io-util");
var error_1 = require("./error");
// Constants
var templateSyncIgnore = "\n.github/ISSUE_TEMPLATE/*\n.github/CODEOWNERS\n.github/CODESTYLE.md\n.github/PULL_REQUEST_TEMPLATE.md\n.github/SECURITY.md\nCITATION.cff\nLICENSE\nREADME.md";
var templateSyncLabel = "\n  - name: 'CI: Template Sync'\n  color: AEB1C2\n  description: Sync with upstream template\n";
/**
 * For interacting with stdin/stdout
 */
var rl = readline_1.default.createInterface({
    input: process.stdin,
    output: process.stdout,
});
/** Prompt user for input */
var question = function (query) {
    return new Promise(function (resolve) { return rl.question(query, resolve); });
};
/** Ask for project information */
var fetchInfo = function (cleanup) { return __awaiter(void 0, void 0, void 0, function () {
    var name, email, username, repository, proj_name, proj_short_desc, proj_long_desc, docs_url;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, question('Name? (This will go on the LICENSE)\n=> ')];
            case 1:
                name = _a.sent();
                return [4 /*yield*/, question('Email?\n=> ')];
            case 2:
                email = _a.sent();
                return [4 /*yield*/, question('Username? (https://github.com/<username>)\n=> ')];
            case 3:
                username = _a.sent();
                return [4 /*yield*/, question('Repository? ((https://github.com/$username/<repo>\n=> ')];
            case 4:
                repository = _a.sent();
                return [4 /*yield*/, question('Project name?\n=> ')];
            case 5:
                proj_name = _a.sent();
                return [4 /*yield*/, question('Short description?\n=> ')];
            case 6:
                proj_short_desc = _a.sent();
                return [4 /*yield*/, question('Long description?\n=> ')];
            case 7:
                proj_long_desc = _a.sent();
                return [4 /*yield*/, question('Documentation URL?\n=> ')];
            case 8:
                docs_url = _a.sent();
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
                return [4 /*yield*/, question('Confirm? (y/n)\n=> ')];
            case 9:
                // Guard clause for confirmation
                if ((_a.sent()).toLowerCase() !== 'y') {
                    console.log('Aborted.');
                    cleanup();
                    process.exit(1);
                }
                return [2 /*return*/, {
                        name: name,
                        email: email,
                        username: username,
                        repository: repository,
                        proj_name: proj_name,
                        proj_short_desc: proj_short_desc,
                        proj_long_desc: proj_long_desc,
                        docs_url: docs_url,
                    }];
        }
    });
}); };
/**
 * The main logic
 */
var main = (0, io_util_1.withTempDir)('caffeine-addictt-template-', function (tempDir) { return __awaiter(void 0, void 0, void 0, function () {
    var data, filesToUpdate, filesToMove;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, fetchInfo(function () { return rl.close(); })];
            case 1:
                data = _a.sent();
                console.log('\nWriting files...');
                filesToUpdate = fs_1.default.readdirSync('./template', {
                    recursive: true,
                });
                filesToUpdate.forEach(function (relativePath) { return __awaiter(void 0, void 0, void 0, function () {
                    var filePath, fileInfo, error_2;
                    return __generator(this, function (_a) {
                        switch (_a.label) {
                            case 0:
                                filePath = path_1.default.join('./template', relativePath);
                                _a.label = 1;
                            case 1:
                                _a.trys.push([1, 3, , 4]);
                                fileInfo = fs_1.default.statSync(filePath);
                                if (fileInfo.isDirectory()) {
                                    return [2 /*return*/];
                                }
                                return [4 /*yield*/, (0, io_util_1.replaceInFile)(filePath, tempDir, data)];
                            case 2:
                                _a.sent();
                                return [3 /*break*/, 4];
                            case 3:
                                error_2 = _a.sent();
                                // it's a bit different here, won't touch this for now
                                if ((error_2 === null || error_2 === void 0 ? void 0 : error_2.code) !== 'ENOENT' &&
                                    (error_2 === null || error_2 === void 0 ? void 0 : error_2.code) !== 'EEXIST') {
                                    console.error(error_2);
                                    process.exit(1);
                                }
                                else {
                                    console.log("File ".concat(filePath, " not found."));
                                }
                                return [3 /*break*/, 4];
                            case 4: return [2 /*return*/];
                        }
                    });
                }); });
                // Write CODEOWNERS
                try {
                    fs_1.default.appendFileSync('./template/.github/CODEOWNERS', "* @".concat(data.username));
                }
                catch (error) {
                    // also different here
                    if ((error === null || error === void 0 ? void 0 : error.code) !== 'ENOENT' &&
                        (error === null || error === void 0 ? void 0 : error.code) !== 'EEXIST') {
                        console.error(error);
                        process.exit(1);
                    }
                    else {
                        fs_1.default.renameSync('./template/.github/CODEOWNERS', '.github/CODEOWNERS');
                    }
                }
                return [4 /*yield*/, question('Would you like to keep up-to-date with the template? (y/n)\n=> ')];
            case 2:
                // Optional keep up-to-date
                if ((_a.sent()).toLowerCase() === 'y') {
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
                    filesToMove = fs_1.default.readdirSync('./template');
                    filesToMove.forEach(function (file) {
                        fs_1.default.renameSync("./template/".concat(file), "./".concat(file));
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
                return [2 /*return*/];
        }
    });
}); }).func;
main();
