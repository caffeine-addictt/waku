# Code Style Guide

The following is a general guide on how to style your work so that the project
remains consistent throughout all files. Please read this document in it's entirety
and refer to it throughout the development of your contribution.

## 1. Naming Conventions

- **Variables:** Use lowercase with underscores (e.g. `my_variable`, `home`).
- **Functions:** Use lowercase with underscores (e.g. `my_function`, `home`).
- **Constants:** Use uppercase with underscores (e.g. `MY_CONSTANT`, `HOME`).
- **Classes:** Use CamelCase (e.g. `MyClass`, `Home`).
- **Files:** Use lowercase with hyphens (e.g. `my-file.go`, `home.go`).

## 2. Comments

Comments should be written to explain **why** something is done,
not **what** is done. Avoid comments for obvious code (e.g. `i++; // Increment i`),
good code is self-documenting.

The use of block comments (`/* */`) is preferred for explaining sections or
complex logic, and single-line comments (`//`) are used for brief explanations.

## 3. Error Handling

- Always handle errors explicitly, and do so gracefully.
- Provide meaningful error messages, and avoid silent failures.
- Log or report errors where necessary for debugging.

## 4. Code Structure

- Avoid deep nesting of code—prefer breaking down large functions into smaller ones.
- Each function/method should do only one thing and do it well.
- Group related functions or variables together logically.

## 5. Function Length

Limit the length of functions or methods. Ideally, keep them under 50 lines.
Split large functions into smaller, more manageable ones when they grow too complex.

## 6. Avoid Hard-coding

- Avoid hard-coding values directly into code.
  Use constants or configuration files where appropriate.
- Ensure your code is flexible and can easily adapt to future changes.

## 7. Avoid Premature Optimization

Focus on clarity first; optimize later if necessary.
If optimization is needed, profile the code first to ensure
you're optimizing the right parts.

## 8. Code Duplication

Avoid code duplication, create reusable functions or components instead.
Break code up into functions to reduce repetition and increase readability.

## 9. Security

- Avoid using unsafe functions or libraries.
- Ensure that your code is secure and protected against common vulnerabilities.

## 10. Performance Considerations

Write code with performance in mind, but do not sacrifice readability for optimization
unless truly necessary. Do profile and test your code before and after optimization.
