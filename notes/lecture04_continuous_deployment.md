# Lecture 4: Continuous Integration (CI), Continuous Delivery (CD), and Continuous Deployment
Week 8, 2023

## Step 1: Complete implementing an API for the simulator in your ITU-MiniTwit.

We continued working on implementing the API for the upcoming simulation. To effectively split our workload, we created GitHub Issues for all of the individual endpoints which are to be implemented. We then decided on who was to implement which endpoints, and we then assigned ourselves to the respective Issues. Some of us worked individually, while others worked in smaller groups. To accurately reflect who did what, we made sure to include all co-authors to all commits we made. 

We made an Issue which represented the "main" Issue for implementing the simulation endpoints. This Issue contains a checklist with a checkbox for each individual endpoint. That way we can check off a single endpoint when we feel like that given endpoint is done. All the work on the simulation endpoints was done on a branch called `feature/simulation`, and all the GitHub Issues reference this branch as the target/development branch. We use this naming scheme to clearly distinguish features which are to be implemented, from bugs which are to be fixed. 

After all the checkboxes in the main Issue were checked off, a pull-request was made, which wanted to merge the `feature/simulation` branch into the `main` branch. This pull-request had all members of our group as reviewers, and it could only be merged when it has at least one of these reviewers accept the changes highlighted in the pull-request. This pull-request also links to all of the individual Issues which relate to the simulation endpoints, such that when merging the pull-request all of these Issues will be closed. When our pull-requests need at least one accepting reviewer (different from the author), it acts as a safe-guard which prevents individuals from pushing directly to the main branch.

After some members of our group accepted the changes in the pull-request, we merged the changes, and our minitwit application is now ready for the simulation.

## Step 2: Creating a CI/CD setup for your ITU-MiniTwit.

A CI/CD setup involves a pipeline which automatically builds, tests and deploys our application, once we make changes to it. Continuous Integration (CI) involves automatic building and testing, once changes are pushed to the code repository. This can be extended to involve Continuous Delivery (CD) which means that after the changes have been built and tested, they are also pushed/delivered to some server (e.g. docker) which runs the application. This way the performance can be monitored and logged before the changes are pushed to the customers/users. This can be extended further with Continuous Deployment (CD), which includes the previous steps, but also automatically deploys the new changes directly to the production (source: [simplilearn.com](https://www.simplilearn.com/tutorials/devops-tutorial/continuous-delivery-and-continuous-deployment)).

For our CI/CD setup, we choose to utilize [GitHub Actions](https://github.com/features/actions). This is because:
 - It is integrated in GitHub, which is the version control system we use for our repository. This way we don't have to involve other platforms, and we can keep everything in one place.
 - Since it is integrated in GitHub, we can see the status of individual commits when we push them. This way we can see whether individual commits build (or contain errors), and we can see whether individual commits pass tests. 
 - It allows for Continuous Delivery and Deployment, which means we can automate the delivery or deployment of our changes.

We choose to go for Continuous Deployment, which means that all of our changes are automatically pushed directly to our virtual machine at DigitalOcean, which runs our minitwit application. We do this such that the users always will have access to the newest changes.

To perform the Continuous Deployment, we set up our GitHub Actions as follows:
```
<insert GitHub Actions script>
```
