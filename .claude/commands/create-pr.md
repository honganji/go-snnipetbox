# Create a Pull Request

Target branch for merge: $ARGUMENTS

This command commits the current changes and creates a pull request targeting the specified branch.

Please follow these steps:

1. **Verify Arguments**

   - Confirm the target branch name specified by the user

## Commit Granularity Guidelines

Appropriate commit granularity significantly impacts development efficiency and code quality. Follow these criteria when creating commits:

### Basic Concept

Create commits at points where you'd think "If something goes wrong, I want to roll back to here!" Think of it like **save points in Dragon Quest**.

### When to Commit (4 Criteria)

1. **Units where the intent or purpose of changes can be understood later**
   - One feature addition, one bug fix, one refactoring, etc.
   - A scope where the reason for change can be clearly explained in the commit message

2. **When the program is in a working state**
   - Commit when tests pass
   - No build errors present
   - Ensure the app won't break on rollback

3. **When one task from the TODO list is completed**
   - When one unit of planned work is finished
   - As a checkpoint before moving to the next task

4. **Points of work you don't want to lose**
   - Content that took time to implement
   - When complex logic is completed
   - When functionality has been verified

### Commits to Avoid

❌ **Too granular commits**
- Committing line by line
- Commits with only typo fixes (combine with other changes)

❌ **Too large commits**
- Commits bundling multiple features
- Commits containing changes with different purposes
- Commits combining several days of work

### Practical Advice

- **When in doubt, go smaller**: Smaller is safer than larger
- **Consider review ease**: One commit, one intent
- **Consider your future self**: Think about whether you'll understand it 3 months later

2. **Create Commits (Additional Features)**

   - First check current change status: `git status`
   - Review the diff: `git diff`
   - Check recent commit history to understand commit message style: `git log --oneline -5`
   - **Follow the commit granularity guidelines above for appropriate commit units**
   - Auto-generate commit messages based on changes
   - Stage appropriate files: `git add .` or specific files
   - Create commit: `git commit -m "[auto-generated commit message]"`

3. **Check Existing PR Template**

   - Follow the template in `.github/PULL_REQUEST_TEMPLATE.md`
   - Consider all differences from the base branch when drafting PR content

4. **Create Pull Request**
   - Push current branch to origin: `git push -u origin [current-branch-name]`
   - Create PR using `gh pr create` command
   - Use the specified branch as the base branch
   - Set title and body based on changes and template
   - Create PR in Draft state

## PR Creation Example

PR URL: https://github.com/example/project/pull/1234

```
## What
As stated in the title.

## How
Added conditional branching to check if the component is mounted for the screen where the bug was occurring.

## Why (optional)
- Document reasons if you chose this implementation among multiple options
- Include the error message that was occurring before the fix
```

## Important Notes

- Always create PR content in English
- If a PR template exists, always use it and fill it out appropriately
- Do not create a PR without user approval
- Report any errors to the user appropriately

## Usage Example

```
/create-pr release/2.5.0
```
