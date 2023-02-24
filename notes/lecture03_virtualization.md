# Lecture 3: Provision of local and remote virtual machines
Week 7, 2023

## New github setup

Before we started on any work this week, we decided to make some changes to our Github repository. The first thing we did was that we started utilizing the 'issues' feature in Git. This is very similar to many other Kanban tools like *Jira* or *Trello*, where we can create backlog items known as *issues*, and then assign each other to different tasks. 

The reason we decided on picking Github issues over other Kanban solutions, was because of the tight integration between Git and the Kanban board itself. Github issues allow us to natively link and mention different things inside of our Github repository, such as branches. It's also much simpler to just use one platform for everything, instead of having to switch back and forth between different platforms while working. Meanwhile Github issues support all of the features we need, such as:
 - Creating issues/backlog items.
 - Creating milestones, and linking issues to milestones.
 - Assigning authors to different issues.

 We also changed some rules in our Github repository such that no one can push directly to the main branch. From now on everything has to be implemented on separate branches and then we have to create pull requests, which have to be peer-reviewed, before the content can be merged into the main branch.

## Step 1: Implement an API for the simulator in your ITU-MiniTwit.

To implement the API for the simulator, we went through the simulator file [minitwit_sim_api.py](https://github.com/itu-devops/lecture_notes/blob/master/sessions/session_03/API_Spec/minitwit_sim_api.py), and looked at all the routes we have to implement. We then created the issue [Feature: Implement the endpoints for the simulator](https://github.com/Lindharden/DevOps/issues/13), and some sub-issues which are described inside the main issue.

We then implemented the required controllers and handlers, to support all the new API calls.

## Step 2: Continue refactoring of your ITU-MiniTwit.

We created some issues for the missing features in our application (not related to simulation). 

We then created a Vagrant file for our application and deployed the application here: <http://164.90.223.49:8080/public>
