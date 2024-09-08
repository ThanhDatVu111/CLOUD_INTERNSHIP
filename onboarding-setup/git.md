# Git and GitHub Guide for New Interns
Welcome to the team! This guide will help you get started with Git and GitHub, covering everything from cloning a repository to creating pull requests.

## Table of Contents
1. [Setting Up Git](#1-setting-up-git)
2. [Cloning a Repository](#2-cloning-a-repository)
3. [Basic Git Commands](#3-basic-git-commands)
   1. [Checking Status](#31-checking-status)
   2. [Viewing Commit History](#32-viewing-commit-history)
4. [Branching](#4-branching)
   1. [Creating a Branch](#41-creating-a-branch)
   2. [Switching Branches](#42-switching-branches)
   3. [Publish a Branch](#43-publish-branches)
5. [Making Changes](#5-making-changes)
   1. [Adding Changes](#51-adding-changes)
   2. [Committing Changes](#52-committing-changes)
6. [Syncing with Remote](#6-syncing-with-remote)
   1. [Pulling Changes](#61-pulling-changes)
   2. [Pushing Changes](#62-pushing-changes)
7. [Creating a Pull Request](#7-creating-a-pull-request)
8. [Additional Resources](#8-additional-resources)
9. [Contribute Workflow](#9-contribute-workflow)

## 1. Setting Up Git
Download and install Git from [git-scm.com](https://git-scm.com).
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

### Configure Git:
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

## 2. Cloning a Repository
To clone a repository, use the following command:
git clone https://github.com/your-username/repo-name.git
Replace your-username and repo-name with the actual username and repository name.

## 3. Basic Git Commands

### 3.1 Checking Status
To see the status of your working directory and staging area:
git status

### 3.2 Viewing Commit History
To view the commit history:
git log

To view it in a more compact form:
git log --oneline

## 4. Branching

### 4.1 Creating a Branch
To create a new branch:
git checkout -b <new-branch-name>

### 4.2 Switching Branches
To switch to an existing branch:
git checkout <branch-name>

### 4.3 Publish Branches
To publish a branch:
git push --set-upstream origin <branch-name>.

## 5. Making Changes

### 5.1 Adding Changes
To add changes to the staging area:
git add <file-name>

To add all changes:
git add .

### 5.2 Committing Changes
To commit your changes:
git commit -m "Commit message"

## 6. Syncing with Remote

### 6.1 Pulling Changes
To pull the latest changes from the remote repository:
git pull origin <branch-name>

### 6.2 Pushing Changes
To push your changes to the remote repository:
git push origin <branch-name>

## 7. Creating a Pull Request
Push your branch to the remote repository:
Go to your repository on GitHub.
Click the "Pull requests" tab.
Click "New pull request".
Select your branch as the compare branch and the base branch (e.g., main).
Fill in the details and click "Create pull request".

## 8. Additional Resources
[Git Documentation](https://git-scm.com/doc)
[GitHub Guides](https://docs.github.com/en)

## 9. Contribute Workflow
- Cloning a Repository
Navigate to the desired directory: This ensures you clone the repository to a specific location on your machine, keeping your project organized.
git clone https://github.com/your-username/repo-name.git

- Updating the Local Repository
Switch to the main branch: This is the starting point for most development work. It represents the stable, working version of the project.
git checkout origin main

Pull the latest changes: Ensures your local main branch is up-to-date with the remote main branch, preventing conflicts when you create your new branch.
git pull origin main 

- Creating a New Branch
Create a new branch: This isolates your work from the main branch, allowing you to experiment and make changes without affecting the main codebase. It's like creating a sandbox for your specific feature or bug fix.
git checkout -b new-branch-name

- Making Changes and Committing
Add changes to the staging area: This prepares the changes you've made for the next commit. It's like selecting the items you want to pack for a trip.
git add . 

- Commit the changes: This saves a snapshot of your work at a specific point in time, with a clear message describing the changes. It's like saving a checkpoint in a game.
git commit -m

- Publishing the Branch
Push the new branch: This shares your work with the rest of the team, allowing others to review your changes before merge into remote main and collaborate.
git push --set-upstream origin <branch-name>.

Opening a Pull Request
Go to Github Reposotory to create a pull request: This formally proposes your changes for inclusion in the main branch. It starts a discussion and review process with your team.

### Why These Steps?
Pull from main before creating a new branch: This ensures your new branch starts from the latest codebase, minimizing merge conflicts later. It's like starting with the most recent blueprint before making modifications.

Create a new branch for each new feature or bug fix: This keeps your work isolated and organized. It's easier to manage and review changes when they are in separate branches.

Push the branch before creating a pull request: This makes your changes visible to others and allows them to start reviewing while you continue working.

Additional Considerations:
Branch naming conventions: Use descriptive names for your branches to clearly indicate their purpose.
Commit message guidelines: Write clear and concise commit messages to provide context for your changes.
Code reviews: Encourage thorough code reviews to improve code quality and catch potential issues.
Merge conflicts: Be prepared to resolve merge conflicts if multiple developers are working on the same part of the codebase.
By following these guidelines and understanding the reasons behind each step, you can effectively contribute to your team's projects using Git and GitHub.