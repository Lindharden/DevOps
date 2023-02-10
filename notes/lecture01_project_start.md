# Lecture 1: Project start, forming groups, SSH, SCP, and Bash
February 03, 2023

## Step 1: Adding Version Control

We added version control by creating an github organization, and creating a new repo. We then moved all the source code and database to that github repo.

## Step 2: Try to develop a high-level understanding of ITU-MiniTwit

Mini-twit runs as a python web-application, run by flask. It utilizes an SQL database.

## Step 3: Migrate ITU-MiniTwit to run on a modern computer running Linux

Minitwit uses the package manager to fetch the packages, which means we have to update our packages locally in order to get the newest versions into the application. We updated Python and Flask. We then created an requirements.txt file with the "modern" versions of the packeges, meaning these can be installed simply by running the requirements file.

Since the current application is written in python2, and we need to convert it to python3, we need to fix the few places where the code doesnt work anymore. Here we can simply use the tool named `2to3` which will convert python2 code to python3.

Using `2to3` with the command: `2to3 minitwit.py --add-suffix='3' -n -w` we get the following changes:
```
RefactoringTool: Skipping optional fixer: buffer
RefactoringTool: Skipping optional fixer: idioms
RefactoringTool: Skipping optional fixer: set_literal
RefactoringTool: Skipping optional fixer: ws_comma
RefactoringTool: Refactored minitwit.py
--- minitwit.py (original)
+++ minitwit.py (refactored)
@@ -8,7 +8,7 @@
     :copyright: (c) 2010 by Armin Ronacher.
     :license: BSD, see LICENSE for more details.
 """
-from __future__ import with_statement
+
 import re
 import time
 import sqlite3
@@ -94,7 +94,7 @@
     redirect to the public timeline.  This timeline shows the user's
     messages as well as all the messages of followed users.
     """
-    print "We got a visitor from: " + str(request.remote_addr)
+    print("We got a visitor from: " + str(request.remote_addr))
     if not g.user:
         return redirect(url_for('public_timeline'))
     offset = request.args.get('offset', type=int)
RefactoringTool: Writing converted minitwit.py to minitwit.py3.
RefactoringTool: Files that were modified:
RefactoringTool: minitwit.py
```

We need to change the import statements, and we need to change the print statement.

Now we compile a new `flag_tool` file.

We get some errors with the `werkzeug` package. Some functions have been moved. To solve this we write `werkzeug.security` in the imports. By updating the `colorama` package to version `0.4.6` the application is now able to run using `python3 minitwit.py`.

We now get an error when initializing the database. To fix this we need to update the `control.sh` file. To find the errors we use the tool `shellcheck` to highlight the errors. The output of using shellcheck on the `control.sh` file is this:
```
In control.sh line 1:
if [ $1 = "init" ]; then
^-- SC2148: Tips depend on target shell and yours is unknown. Add a shebang.
     ^-- SC2086: Double quote to prevent globbing and word splitting.
Did you mean: 
if [ "$1" = "init" ]; then

In control.sh line 9:
elif [ $1 = "start" ]; then
       ^-- SC2086: Double quote to prevent globbing and word splitting.

Did you mean: 
elif [ "$1" = "start" ]; then


In control.sh line 11:
    nohup `which python` minitwit.py > /tmp/out.log 2>&1 &
          ^------------^ SC2046: Quote this to prevent word splitting.
          ^------------^ SC2006: Use $(...) notation instead of legacy backticked `...`.
           ^---^ SC2230: which is non-standard. Use builtin 'command -v' instead.

Did you mean: 
    nohup $(which python) minitwit.py > /tmp/out.log 2>&1 &


In control.sh line 12:
elif [ $1 = "stop" ]; then
       ^-- SC2086: Double quote to prevent globbing and word splitting.

Did you mean: 
elif [ "$1" = "stop" ]; then


In control.sh line 15:
elif [ $1 = "inspectdb" ]; then
       ^-- SC2086: Double quote to prevent globbing and word splitting.

Did you mean: 
elif [ "$1" = "inspectdb" ]; then


In control.sh line 17:
elif [ $1 = "flag" ]; then
       ^-- SC2086: Double quote to prevent globbing and word splitting.
```

We then fixed all these things in the `control.sh` file. More precisely we specified a shebang (to tell its a bash shellscript), and quotation marks around all the arguments (that where pointed out by shellcheck)

We made the control.sh file executable by changing its permissions using chmod.

## Step 4: Share your Work on GitHub

After pushing all our changes to our github repo, we created a new release from the main branch.