# Contributing

See our specific contributing docs [here](https://goplugin.github.io/plugin-testing-framework/contributing/), along with the general [Plugin contributing guidelines](https://docs.chain.link/docs/contributing-to-plugin/).

# Creating new PRs

When creating a new PR remember to:

- include the id of JIRA ticket in the format `[PROJECT_ID-ticket-id]` (e.g.`[TT-1234]` for a Test Tooling ticket with ID 1234) in the PR title
- include a short description of the PR

# Drafting new releases

To draft a new release:

- open the repository's releases page in [Github](https://github.com/goplugin/plugin-testing-framework/releases)
- click on "Draft a new release"
- chose correct version (see below) and use it both a tag and release title
- write a short description of the release
- click "Publish release"

## Release versioning

When releasing a new version remember to follow correct semver:

- `major` version when you make incompatible API changes,
- `minor` version when you add functionality in a backwards compatible manner, and
- `patch` version when you make backwards compatible bug fixes.

You can read more about semver [here](https://semver.org/).
