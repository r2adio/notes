# PLUMBING:

## it's just files all the way down:

- .**git** includes all the commits, branches, tags, and other objects
- git is made up of objects that are stored in `.git/objects` dir.
- a commit is a type of object.

<!-- NOTE: Inode Busting
    - when a dir havin' to many files, weird things start to happen
    - git, to prevent inode busting, it taked first 2 letters and add this happens:
    .git/objects/5f
    .git/objects/5f/280094d593663a638cad738061de6f90f6a1e4
    .git/objects/33
    .git/objects/33/50a218de1272fc2fe1fbdeacdc27762e1f2878

    find .git/
-->
