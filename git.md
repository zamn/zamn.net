Git Overview
===========
##### Adam Hamot - January 2014
**Note:** This document uses markdown for formatting which you can learn about here: [StackOverflow Markdown Editing Help](http://stackoverflow.com/editing-help).

Git can become very complex depending on what you want to do but you primarily only care about a couple of commands to do daily tasks. The rest you can just look up on stackoverflow/google. 

Git works through the following process:

1. Stage the files
2. Commit the current staging area
3. Push the commit

Those 3 steps will be your workflow while using git. 

Every time you create a new repository on GitLab you will have a link to clone it on the page itself. These usually look something like: git@host:project.git. In order to get the project onto your computer you simply type:

**git clone git@host:project.git**

This will create a local directory containing all of the files in this repository. 

By default you will have a master branch and an origin remote, which will point to the clone link. 

After that you are ready to start working.

Created Repository Commands
===========

These commands are for when a repository is already made and you just cloned it from a remote source. To initialize your own repository see below.

**git status**

* Shows the files currently in the staging area. 
* This also shows what files are currently staged and which files are not staged/tracked.

**git add <filename>**

* This adds a file to the staging area to be committed.
* A shortcut you can use is *git add .* in order to add all local changes and *git add . -A* to add all local changes including deleted files (by default git add doesn't include them).

**git diff <filename>**

* Before staging a file you can type this command to see the changes that were actually made.
* This is useful if you need to remember what you did to a file for the commit message.

**git commit [-m Message]**

* This commits what is currently added in the staging area and either opens up a new editor window for you to write your commit message or if you use the -m option you can type it inline with the command.
* Commit messages should be useful, NOT things like: 'work in progress', or 'this doesnt work dont look at this.' *cough* martin

**git push [remote] [branch]**

* This pushes all of your local commits to the remote repository.
* You can view your remotes by typing: **git remote show** which will list all of the remotes you have locally. This will normally show names like 'origin' or 'upstream'. These are simply just labels for addresses that git can do stuff from. There will be more on this command later.
* You can view the current branch you are on by typing: **git branch** and the branch with the asterik next to it is the one you are currently on (Alternatively typing git status shows you at the top).
* For quick reference one of the most common uses is: **git push origin master**

**git pull [remote] [branch]**

* This pulls the commits made to the repository that are on the remote server.
* The same rules apply for remotes and branches as above.
* For quick reference: **git pull origin master**

**git stash save [name]**

* This will stash your current changes, giving you a clean slate and allowing you to pull when you get that pesky error message 'cannot pull remote changes; changes are made locally' (or something close to it).
* There are a ton of options to this command which I will briefly go over:
    * git stash list
        * lists all of your current stashes
    * git stash apply
        * applys the last stash you saved.
* If you want to learn more information about git stash then feel free to google it.

**git branch [name]**

* Rather than try to explain branching and mergingto you, I recommend you read the branching page on the [git-scm branching page](http://git-scm.com/book/en/Git-Branching-Basic-Branching-and-Merging).

**git tag [-a] [name] [-m message]**

* This creates a tag for the current commit you are on. Tags are useful for when you want to record a commit of the latest release of a project (linux uses these religiously). 
* I recommend using the command as it is typed here: **git tag -a v1.0 -m '1.0 Release.'**
    * The -a means that the tag is annotated (giving information about who made the tag, in addition to the commit the tag points to). 
    * The -m, like in git commit is for the message. I recommend using -m here rather than typing out a whole long explanation as most UIs only show the short version of the message.
* Once the tag is made you can push it to the remote server so that it appears there as well. You can do that by typing: **git push [remote] [tag]**
* You can learn more about tags here on the [git scm tagging page](http://git-scm.com/book/en/Git-Basics-Tagging).

Getting Started Creating a Repository
=============

**git init**

* Run this command inside of the directory that you want to create a repository in.
* Initially all files will be untracked so you will have to add all the files you want and commit them for them to be tracked.
* To ignore certain files, you can add a **.gitignore** file, which you can look up more information on its syntax.
* This directory will be offline by default but you can push it to a remote repository by following their directions (they will tell you to use **git remote add origin remote_url** to add the remote and then you can work as normal).

Conclusion
=============

All of the commands above will help you get the basic workflow down. You should be able to do your daily tasks using this. The rest of the commands are important but situational (mostly). I highly recommend reading through the git book on the [http://git-scm.org](http://git-scm.com/book/en/Getting-Started) website if you want to get a full understanding of git.