# Contributing
## Found a Bug?
If you find a bug in the source code, you can help us by submitting an issue to our GitHub Repository. Even better, you can submit a Pull Request with a fix.

## Submit an Issue
We welcome your submission of issues.

To submit a bug report, enhancement request, or any feedback, please open a GitHub issue using the appropriate issue template.

## Pull Requests 
1. Fork the project
2. Implement feature/fix bug & add test cases
3. Ensure test cases & static analysis runs succesfully
4. Submit a pull request to `master` branch
 
Please include unit tests where necessary to cover any functionality that is introduced.
More details,please check [development.md](./docs/development.md).
 
## Commit messages
The message has a special format that includes a type, a scope and a subject:

```bash
git commit -m "<type>(<scope>): Subject"
```

The **header** is mandatory and the **scope** of the header is optional.

### Revert
If the commit reverts a previous commit, it should begin with `revert: `, followed by the header of the reverted commit. In the body it should say: `This reverts commit <hash>.`, where the hash is the SHA of the commit being reverted.
 
### Type
Must be one of the following:
 
* **build**: Changes that affect the build system or external dependencies (example scopes: gradle, fastlane, npm)
* **ci**: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
* **docs**: Documentation only changes
* **feat**: A new feature
* **fix**: A bug fix
* **perf**: A code change that improves performance
* **refactor**: A code change that neither fixes a bug nor adds a feature
* **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
* **test**: Adding missing tests or correcting existing tests
 
### Scope
The scope should be the name of the npm package affected (as perceived by person reading changelog generated from commit messages.

The following is the list of supported scopes:
 
* **TODO**

### Subject 
The subject contains succinct description of the change:

* use the imperative, present tense: "change" not "changed" nor "changes"
* don't capitalize first letter
* no dot (.) at the end
