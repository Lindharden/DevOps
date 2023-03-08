# Lecture 6: Monitoring
[Week 6](https://github.com/itu-devops/lecture_notes/blob/master/sessions/session_06/README_TASKS.md)

## Add Monitoring to Your Systems
The monitoring system we want to implement is Prometheus. Reasons for choosing Prometheus:
 - Supports wide range of metrics, which are important for thorough monitoring of our application.
 - Works with Go, which is the language our application is developed with.
 - Prometheus can be integrated into dashboard tools such as **Grafana** or **Uptime Kuma**, which help us clearly and easily visualize the measured metrics.
 - Prometheus can send us alerts if any metrics go over any thresholds that we determine. This can help us catch problems early.

We are going to implement **Application monitoring** and **Infrastructure monitoring**. The metrics we want to measure are:
 - Number of tweet requests per hour - This will help us determining the amount of load our application is under. We will be able to see when something critical is happening to our tweeting system, which is arguably the most important feature in our application.
 - How long it takes to connect to the application - By constantly monitoring (e.g. every 15 minutes) how long it takes to connect to our service, we can determine when our application is unreachable, and roughly for how long it was unreachable. We can also determine whether we have made any changes which makes our application slower to load.
 - Amount of registrations per hour - This will help us determine how many new users we receive. If we suddenly receive a lot of tweet request, logins or some kind of errors, it would be handy to be able to see whether we recently have received a lot of registration requests.
 - Amount of logins per hour - This will help us determine how many logins we receive. If something happens to our service, it would be handy to be able to see whether we have recently received a lot of login requests.

## Security Features
We have added a contributor bot to our GitHub repository named [Dependabot](https://github.com/dependabot). This bot will automatically go through all our dependencies of our application, and determine whether any of them have been outdated, and therefore pose a security risk. If Dependabot finds any outdated packages, it will create a pull-request in our GitHub repository, where it tries to update the given packages. This will help us update vulnerable packages much faster then if we manually would have to go through each one of them to update them.
 