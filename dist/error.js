"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.handleError = void 0;
/** Handle errors and conditionally exit program */
const handleError = (error) => {
    if (error?.code !== 'ENOENT' &&
        error?.code !== 'EEXIST') {
        console.error(error);
        process.exit(1);
    }
};
exports.handleError = handleError;
