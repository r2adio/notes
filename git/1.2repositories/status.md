# STATUS:

in a git repo a file can be at one of the states:

- **untracked** : not being tracked by git
- **staged** : marked for inclusion, in next commit
- **commited** : saved to repo's history

## Staged:

- uses git add filename.xyz foo/filename.xyz
- `git add .` : to add all the files in that dir and its sub-dir's only.

## Commit:

- commit is a snapshot of repo a at a given point in time
- git saves the entire state of ur repo per commit, not the diff b/w commits
- `git commit -m "commit message"` : commiting the changes made and addin' a commit message
- `git commit -am "commit message"` : add all the files to the staging area and commit those files
