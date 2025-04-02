# Basic Git Commands:

## git branch xyz

- this make a new branch named xyz
- git will be pointing to the same branch as before

## Deleting a branch

- git branch -d xyz

## git checkout xyz

- will checkout from current branch, master, and start pointing to the xyz branch.

## HEAD

- **HEAD** points to where a branch is currently at
- usually HEAD points to the end of the branch.
- can allow HEAD to point one back as well, as at last its a checkpoint.

## Merging the branches

- Fast forward merging:
  - work entirely on other branches and later merge them into master branch.
- Not-Fast forward merging:
  - both the alternate and master branch are working, and later merge other branch into master.
- in order to merge:
  - first be on branch, which u dont want to merge.
  - git merge xyz
