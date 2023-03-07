# Lecture 6: Monitoring
[Week 6](https://github.com/itu-devops/lecture_notes/blob/master/sessions/session_06/README_TASKS.md)

## Step 1: Add Monitoring to Your Systems
The monitoring system we want to implement is Prometheus. Reasons for choosing Prometheus:
 - 

We are going to implement **Application monitoring** and **Infrastructure monitoring**. The metrics we want to measure are:
 - Number of tweet requests per hour - This will help us determining the amount of load our application is under. We will be able to see when something critical is happening to our tweeting system, which is arguably the most important feature in our application.
 