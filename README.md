# Stop Orders v0.0.0

An implementation showcasing how one can manage orders in memory to execute Stop Orders.

## Usage

*_TODO_*: How to install/use this module/service

## Development

### Getting Started

    # install golang
    brew install golang

    # install the golangci linter 
    # more details: https://golangci-lint.run/
    brew install golangci-lint
    
    # install pre-commit
    pip install pre-commit
    pre-commit install

### Pre-commit

A number of pre-commit hooks are set up to ensure all commits meet basic code quality standards.

If one of the hooks changes a file, you will need to `git add` that file and re-run `git commit` before being able to continue.


### Git Workflow

This repo is configured for trunk-based development. When adding a new fix or feature, create a new branch off of `main`.

Merges into main *must always be rebased and squashed*. This can be done manually or with GitHub's "Squash and Merge" feature.

### Testing

All test files are kept in ./test/ and named *_test.go. Github workflow automatically run the tests when code is pushed and will return a report with results when finished.
