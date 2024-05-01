# Code Style

The following is a general guide on how to style your work so that the project
remains consistent throughout all files. Please read this document in it's entirety
and refer to it throughout the development of your contribution.

1. [General Guidelines](#general-guidelines)
2. [Markdown Guidelines](#markdown-guidelines)

## General Guidelines

Listed is a example class used demonstrate general rules you should follow
throughout the development of your contribution.

We use [Prettier](https://prettier.io/) and [ESLint](https://eslint.org/)
to ensure that code is consistent and follows our [code style](./CODESTYLE.md).
Please ensure that your code passes linting before merging a Pull Request.

- Docstrings are to follow [JSDoc syntax](https://jsdoc.app).
- Private attributes are to be prefixed with an underscore.

```ts
/** Get a greeting string */
const myFunction = (): string => {
  return 'hi';
};
```

## Markdown Guidelines

Currently, documentation for this project resides in markdown files.

- Use of HTML is permitted
- Use of HTML comments is appreciated
- Exceedingly long lines are to be broken
- [reference style links][reference-style-links] are not required by are appreciated

```markdown
<!--example markdown document-->

# Section

Lorem ipsum dolor sit amet, consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore
magna aliqua. Ut enim ad minim veniam, quis nostrud
exercitation ullamco laboris nisi ut aliquip ex ea
commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu
fugiat nulla pariatur. Excepteur sint occaecat cupidatat
non proident, sunt in culpa qui officia deserunt mollit
anim id est laborum. found [Lorem Ipsum Generator]

# Section 2

<ul>
  <li> Apple
  <li> Orange
  <li> Pineapple
</ul>

[Lorem Ipsum Generator]: https://loremipsum.io/generator/
```

[reference-style-links]: https://www.markdownguide.org/basic-syntax/#reference-style-links
