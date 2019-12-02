# Section 1 - Introduction

## Git Basics

Source code management nowadays is not even a question. Every
developer regardless of the language has to know at least one
**VCS** (Version Control System)

So when it comes to code versioning there are multiple choices,
however the most popular & versatile and supported by all platforms
is **Git**

Git is a very easy to learn & use tool, but so crucial to any project out there.

So here are the basic Git commands which you should use when
working out in this course, to save all your work:

#### Basic commands:
```bash
# clones the repository from a remote (i.e. GitHub)
git clone https://github.com/your-user/you-repo.git

# displays the current state of your project
# which files were added/modified/deleted
# so you can decide what to commit
git status

# adds all files inside the current directory to be committed
git add .
git add -A

# adds "file1" and "file2" inside the current directory to be committed
git add file1 file2

# creates a commit with the specified message with the current staged changes
git commit -m "some commit message"

# pushes all commits to the remote repository (i.e. GitHub)
git push

# pulls all the changes from the remote repository (i.e. GitHub)
git pull
```

#### More commands:

```bash
# initializes a project directory with git
# making it a fresh git repository
git init

# configures the name and email of the git user
git config user.name "Name Surname"
git config user.email "hello@example.com"

# links the local repository to the remote one
git remote add origin https://github.com/your-user/you-repo.git

# pushes changes to "origin" remote
# origin maps to remote repository
git push -u origin master

# pushes changes from local branch directly to remote
# and creates a branch named branch-name on the remote repository
git push --set-upstream origin branch-name

# deletes "branch-name" from the remote repository
git push origin :branch-name

# checks out from current branch to "master" branch
git checkout master

# checks out from current branch to "branch-name" branch
git checkout branch-name

# checks out from current branch to the changes version at commit hash: "820d071"
git checkout 820d071

# creates "branch-name" branch and checks out that branch
git checkout -b branch-name

# deletes the local branch named branch-name
git branch -D branch-name

# merge changes from master with current branch
git merge master

# shows the entire git history
git log

# for more info about git run
git help
```

You can find some very good tutorials on learning
**Git** additionally on:

This is not a **Git** course or tutorial. There are plenty of
resources on how to learn Git. In this course we'll only touch
the basics so that you can, *add*, *commit*, *push* and *pull* changes
on a remote repository

[Atlassian Tutorials](https://www.atlassian.com/git/tutorials)

[Back](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>
