<!--
SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

# Contributing Guidelines

Thank you for your interest in contributing to the Effect Index TripReporter project!

Read our [Code of Conduct](https://github.com/effectindex/.github/blob/master/CODE_OF_CONDUCT.md) to keep our community approachable and respectable.

In this guide you will get an overview of the contribution workflow from opening an issue, creating a PR, reviewing, and merging the PR.

## New Contributor Guide

To get an overview of building and running the project, read the [README](./README.md).
If you would like to ask questions or require help, join our Discord at [effectindex.com/discord](https://effectindex.com/discord), or email <git@frogg.ie>.

## Getting Started

### Issues

If you have a bug or feature request, search through the existing issues to see if a bug report or feature request has already been filed.
If not, [create a new issue](https://github.com/effectindex/tripreporter/issues/new/choose), and select the appropriate template.

### Commit Style

When making a commit, it is preferred that you split your commit by "section" of the project. For example, if you are making a pull request that adds a new frontend feature, you would make multiple commits, splitting them, with a prefix for each relevant section. This might look like 3 separate commits, titled:
- `api: (feature name) Implement [feature] / relevant info`
- `models: (feature name) Add [model] / relevant info`
- `ui: (feature name) Implement [feature] / relevant info`

If you're making a change to existing features, or fixing a bug, it might look like the following:
- `api: (feature name) Fix bug / additional info`
- `ui: (feature name) Fix bug / additional info`

If unsure, take a look at the existing git logs, as well as the git history for the relevant folders / sections of code that you're changing.

### Pull Request

Once your changes are ready for review, make a pull request, with a prefix in the name that describes what the pull request is for. It might look like the following:
- `fix: (account) Wrong endpoint used for updating account email`
- `impr: (ui) Nicer mobile layout on report pages`

Don't forget to link the relevant issue in the description, with something like `Fixes #12 (adds authentication)`, or `Closes #10 (adds analytics)`.

When ready, make the pull request and one of our maintainers will review it as soon as possible. See [above](#new-contributor-guide) for contacts for help or questions about this process, or if you have questions while submitting your pull request.
