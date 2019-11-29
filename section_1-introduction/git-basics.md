# Section 1 - Introduction

## Git Basics

Source code management nowadays is not even a question. Every
developer regardless of the language has to know at least one
VCS (Version Control System)

So when it comes to code versioning there are multiple choices,
however the most popular & versatile and supported by all platforms
is **Git**

It's a very easy to use tool, but so crucial to any project out there.

So here are the basic Git commands which you should use when
working out in this course, to save all your work:

#### Basic commands
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

#### More commands

```bash
# initializes a project directory with git
# making it a fresh git repository
git init

# configures the name and email of the git user
git config user.name "Name Surname"
git config user.email "hello@example.com"

# 
git remote add origin https://github.com/your-user/you-repo.git

git push -u origin master

git push --set-upstream origin branch-name

git push origin :branch-name

git checkout -b branch-name

git branch -D branch-name

git merge master

git checkout master

git log

git help
```

You can find some very good tutorials on learning
**Git** additionally on:

[Atlassian Tutorials](https://www.atlassian.com/git/tutorials)

[Back](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>