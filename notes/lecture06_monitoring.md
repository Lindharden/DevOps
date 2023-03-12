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

In order to add Prometheus and Grafana to our application, we added the following to our `docker-compose.yml` file:
``` yaml
networks:
      - prometheus-network
      ...

prometheus:
    image: prom/prometheus:latest
    container_name: prometheus
    networks:
      - prometheus-network
    ports:
      - "9090:9090"
    volumes:
      - prometheus-data:/prometheus
      - type: bind
        source: ./prometheus.yml
        target: /prometheus/prometheus.yml
        read_only: true
    restart: unless-stopped
    command:
      - "--config.file=prometheus.yml"

  grafana:
    image: grafana/grafana-oss:latest
    container_name: grafana
    ports:
      - "3000:3000"
    volumes:
      - grafana-data:/var/lib/grafana
      - ./grafana/provisioning/:/etc/grafana/provisioning/
      - ./grafana/dashboards:/var/lib/grafana/dashboards
    restart: unless-stopped

volumes:
  postgres-db:
    driver: local
  prometheus-data:
    driver: local
  grafana-data:
    driver: local

networks:
  database-network:
    name: database-network
  prometheus-network:
    name: prometheus-network
```

We added a JSON file for the Grafana dashboard, which defines how the dashboard will look. We also added a new endpoint named `/metrics` which will be exposing the metrics we defined.

## Security Features
We have added a contributor bot to our GitHub repository named [Dependabot](https://github.com/dependabot). This bot will automatically go through all our dependencies of our application, and determine whether any of them have been outdated, and therefore pose a security risk. If Dependabot finds any outdated packages, it will create a pull-request in our GitHub repository, where it tries to update the given packages. This will help us update vulnerable packages much faster then if we manually would have to go through each one of them to update them.

A security policy has been added to our GitHub repository, stating which versions of the application are supported. In the security policy it is stated that no other than the latest version will receive security updates. This information is important to give to the users, such that nobody will be using an older version of the application and expect it to be fully secure.

Another security feature we have added to our GitHub repository is [CodeQL](https://docs.github.com/en/code-security/code-scanning/automatically-scanning-your-code-for-vulnerabilities-and-errors/about-code-scanning-with-codeql), which is an code analysis engine developed by GitHub, which automatically scans the code to perform security checks. We use this tool as it's fully integrated into GitHub, and will automatically scan the code and post alerts to us if any security problems have been detected. By having this process run automatically, we can be sure it will be done as opposed to us having to manually start security scans, which we easily could forget. Security scans are important to do regularly as to catch any vulnerabilities we could be accidentally implementing into our commits.
 