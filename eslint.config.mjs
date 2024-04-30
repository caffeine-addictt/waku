// @ts-check

import pkg from '@eslint/js';
const { configs } = pkg;
import { configs as _configs } from 'typescript-eslint';
import eslintPluginPrettierRecommended from 'eslint-plugin-prettier/recommended';

export default [
  {
    ignores: ['dist/*', 'babel.config.cjs'],
  },
  configs.recommended,
  ..._configs.recommended,
  eslintPluginPrettierRecommended,
];
