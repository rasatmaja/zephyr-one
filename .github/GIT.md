# GIT Management Guildeline 
> more info see this articles:
* [Start it up - Apurva Jain](https://medium.com/swlh/writing-better-commit-messages-9b0b6ff60c67) - Writing Better Commit Messages
---
## Writing Better Commit Messages
![](https://miro.medium.com/max/3232/1*7GAbT16Ajy-Y1n680MVCQw.png)


## Commit Message Convention

```
type(scope): subject 

<body>
<footer>
```

### 1. Type of commit
```
feat:     The new feature being added to a particular application
fix:      A bug fix
style:    Feature and updates related to styling
refactor: Refactoring a specific section of the codebase
test:     Everything related to testing
docs:     Everything related to documentation
chore:    Regular code maintenance
```

### 2. Scope (optional)
Provided in parentheses after the type of commit, describing parts of the code affected by the changes or simply the epic name.
```
feat(claims)
fix(orders)
```

### 3. Subject
Short description of the applied changes.
- Limit the subject line to 50 characters
- Your commit message should not contain any whitespace errors or punctuation marks
- Do not end the subject line with a period
- Use the present tense or imperative mood in the subject line
```
feat(claims): add claims detail page
fix(orders): validation of custom specification
```

### 4. Body (Optional)
Use the body to explain what changes you have made and why you made them.
- Separate the subject from the body with a blank line
- Limit each line to 72 characters
- Do not assume the reviewer understands what the original problem was, ensure you add it
- Do not think your code is self-explanatory

```
refactor!: drop support for Node 6
BREAKING CHANGE: refactor to use JavaScript features not available in Node 6.
```

### 5. Footer (optional)
Use this section to reference issues affected by the code changes or comment to another developers or testers.

```
fix(orders): correct minor typos in code
See the issue for details
on typos fixed.
Reviewed-by: @John Doe
Refs #133
```

<!---
SemVer Tips:
fix: a commit of the type fix patches a bug in your codebase (this correlates with PATCH in semantic versioning).
feat: a commit of the type feat introduces a new feature to the codebase (this correlates with MINOR in semantic versioning).
BREAKING CHANGE: a commit that has a footer BREAKING CHANGE:, or appends a ! after the type/scope, introduces a breaking API change (correlating with MAJOR in semantic versioning). A BREAKING CHANGE can be part of commits of any type.
-->