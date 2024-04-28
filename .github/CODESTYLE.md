# Code Style

The following is a general guide on how to style your work so that the project
remains consistent throughout all files. Please read this document in it's entirety
and refer to it throughout the development of your contribution.

1. [General Guidelines](#general-guidelines)
2. [Commit Message Guidelines](#commit-message-guidelines)
3. [Markdown Guidelines](#markdown-guidelines)

## General Guidelines

Listed is a example class used demonstrate general rules you should follow throughout the development of your contribution.

- Docstrings are to follow {{DOCSTRING_FORMAT}}
- Private attributes are to be prefixed with an underscore

```
code example
```

## Commit Message Guidelines

When committing, commit messages are prefixed with one of the following depending on the type of change made.

- `feat:` when a new feature is introduced with the changes.
- `fix:` when a bug fix has occurred.
- `chore:` for changes that do not relate to a fix or feature and do not modify _source_ or _tests_. (like updating dependencies)
- `refactor:` for refactoring code that neither fixes a bug nor adds a feature.
- `docs:` when changes are made to documentation.
- `style:` when changes that do not affect the code, but modify formatting.
- `test:` when changes to tests are made.
- `perf:` for changes that improve performance.
- `ci:` for changes that affect CI.
- `build:` for changes that affect the build system or external dependencies.
- `revert:` when reverting changes.

Commit messages are also to begin with an uppercase character. Below list some example commit messages.

```sh
git commit -m "docs: Added README.md"
git commit -m "revert: Removed README.md"
git commit -m "docs: Moved README.md"
```

## Markdown Guidelines

Currently, documentation for this project resides in markdown files.

- Headings are to be separated with 3 lines
- Use of HTML comments is appreciated
- Use of HTML is permitted
- [reference style links](https://www.markdownguide.org/basic-syntax/#reference-style-links) are not required by are appreciated
- Exceedingly long lines are to be broken
- The indents are to be 4 spaces

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
