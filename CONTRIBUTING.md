# Contribute to DevOps Minitwit

Welcome to the Minitwit DevOps project! 

This guide explains how to contribute to the project. 

To read more about the project, visit the [README](https://github.com/Lindharden/DevOps/blob/main/README.md).

## Repository setup

This repository is a public repository with all members having collaborator status.

## Branching model

For this project we have decided on using feature branches. In essence, each new feature should reside in its own branch prefixed with feature/{name of new feature}.

New features are merged into main branch which is used for creating releases.

Avoid rebasing and squashing the commit history to ensure visibility of the work history.

## Distributed development workflow

Depending on the extend and difficulty of issues, work can be done individually, in pair-programming or in mob-programming. It will be visible by the number of authors for each commit whether it was done individually or not.


## Contribution structure

#### Issues

Create issues before developing features or fixing bugs. Depending on the type of issue, a template should be filled out with the request information.

We differ between two types of issues, namely Feature requests and Bug reports. A seperate template exists for the two type of issues.

Create a branch for each issue following the branching convention mentioned in [Branching model](#branching-model)


#### Commits

For commits done in pair-programming or mob-programming, a co-author tag should be added in each commit listing every contributor for the respective commit. This is done by appending the following message to the commit message: 
`Co-authored-by: AUTHOR-NAME <ANOTHER-NAME@EXAMPLE.COM>"`


#### Pull requests

For pull requests, the following steps should be taken into consideration:

* Ensure that at least one reviewer non-contributing reviewer is assigned as reviewer.
* The pull-request template should be filled out with relevant information (i.e. links to relevant issues and more).
* Requested changes for the pull-request should be suggested using the github `suggest changes` feature or in the form of comments.
* Once suggested changes are resolved, mark the conversation/request as solved.

Once a pull-request is approved and merged, the latest release will automatically be updated to include the code within the pull-request.


## Reviewing/integrating contributions

For contributions not done by mob programming, a random member is picked as reviewer and is responsible for reviewing the code. However, every contributor is encouraged to review each pull-request regardless.
