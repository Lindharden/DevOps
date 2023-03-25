# Lecture 5: What is DevOps and configuration management
[Week 5](https://github.com/itu-devops/lecture_notes/blob/master/sessions/session_05/README_TASKS.md), 2023

## Step 1: DevOps

The DevOps handbook mentions the *Three Ways*. This is our thoughts on how we adhere to each of the principles:

 - **Flow**: We adhere to the *Flow* principle which talks about implementing a fast *left-to-right* flow, which is the time it takes from when requests are put on the backlog, till they implemented and is in production. Our setup has a fast *Flow* as we have an CI/DC setup which automatically builds, tests and deploys changes from the main branch, to our remote server at DigitalOcean. This means that the time it takes from when changes are committed, till they are in production, is just as long as it takes for our team to approve and merge pull-requests. We also try to split tasks into as small as possible backlog items (GitHub Issues), such that they are easier to delegate to different group members, such that they are passed through the pipeline faster.
 - **Feedback**: We adhere to the *Feedback* principle, which focuses on continuous problem solving when they occur. Our application don't have real users, but we simulate the feedback process by receiving error messages from the user simulation run by the teaching team. When we see that errors occur, we create an Issue/ticket which goes in our backlog. We then try to fix the error as fast as possible.
 - **Continual Learning and Experimentation**: We try to adhere to *Continual Learning and Experimentation* where we encourage risk taking and see mistakes as opportunities to learn. We are in a special situation as we are not implementing a real application, and therefore we can afford to make big mistakes without any bigger impacts. We relieve the pressure of having to manually deploy the changes we make, by having an CI/CD setup which automatically deploys all the changes we make.

## Step 2: Complete and polish your ITU-MiniTwit implementation

For our DB abstraction layer we choose to use [Gorm](https://gorm.io/). We choose to use this library as:

 - It works with Go, which is the language we used to implement the minitwit application.
 - It is the most mature and production ready library for DB abstraction for Go, which also makes it the best choice.
 - Gorm supports creating, updating and deleting elements in a database, which is exactly what we need for our abstraction layer.
 - Gorm allows querying the database which allows us to extract information, without having to leave raw SQL code in our code repository.
