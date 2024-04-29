"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.handleError = void 0;
/** Handle errors and conditionally exit program */
var handleError = function (error) {
    if ((error === null || error === void 0 ? void 0 : error.code) !== 'ENOENT' &&
        (error === null || error === void 0 ? void 0 : error.code) !== 'EEXIST') {
        console.error(error);
        process.exit(1);
    }
};
exports.handleError = handleError;
