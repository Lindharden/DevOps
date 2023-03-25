# Lecture 8: Logging, and Log Analysis. Service-level agreements.
[Week 8](https://github.com/itu-devops/lecture_notes/blob/master/sessions/session_08/README_TASKS.md)

## Add logging to your system
We decided to add EFK logging to our application. To do this we added Elasticsearch, Filebeat and Kibana to our Docker container:
```yaml
elasticsearch:
  image: "docker.elastic.co/elasticsearch/elasticsearch:7.2.0"
  environment:
    - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
    - "discovery.type=single-node"
    - ELASTIC_PASSWORD=${ELASTIC_PASSWORD}
    - xpack.security.enabled=true
  volumes:
    - efk_elasticsearch_data:/usr/share/elasticsearch/data
  networks:
    - efk
  mem_limit: 1g

kibana:
  image: "docker.elastic.co/kibana/kibana:7.2.0"
  depends_on:
    - elasticsearch
  environment:
    - ELASTICSEARCH_HOSTS=["http://elasticsearch:9200"]
    - ELASTICSEARCH_USERNAME=${ELASTIC_USERNAME}
    - ELASTICSEARCH_PASSWORD=${ELASTIC_PASSWORD}
  networks:
    - efk
  ports:
    - "5601:5601"

filebeat:
  image: "docker.elastic.co/beats/filebeat:7.2.0"
  user: root
  command:
    - "-e"
    - "--strict.perms=false"
  volumes:
    - ./logging/filebeat.yml:/usr/share/filebeat/filebeat.yml:ro
    - /var/lib/docker:/var/lib/docker:ro
    - /var/run/docker.sock:/var/run/docker.sock
  networks:
    - efk
  environment:
    - DOCKER_USERNAME=${DOCKER_USERNAME}
    - ELASTIC_USERNAME=${ELASTIC_USERNAME}
    - ELASTIC_PASSWORD=${ELASTIC_PASSWORD}

volumes:
...
  efk_elasticsearch_data:
    driver: local
...

networks:
...
  efk:
    name: efk
...
```

We configured Filebeat and Elasticsearch like such:
```yml
filebeat.inputs:
  - type: container
    paths:
      - "/var/lib/docker/containers/*/*.log"

processors:
  - add_docker_metadata:
      host: "unix:///var/run/docker.sock"

  - decode_json_fields:
      fields: ["message"]
      target: "json"
      overwrite_keys: true

output.elasticsearch:
  hosts: ["elasticsearch:9200"]
  username: ${ELASTIC_USERNAME}
  password: ${ELASTIC_PASSWORD}
  indices:
    - index: "filebeat-elastic-%{[agent.version]}-%{+yyyy.MM.dd}"
      when.or:
        - equals:
            container.image.name: docker.elastic.co/beats/filebeat:7.2.0
        - equals:
            container.image.name: docker.elastic.co/kibana/kibana:7.2.0
    - index: "filebeat-minitwit-%{[agent.version]}-%{+yyyy.MM.dd}"
      when.equals:
        container.image.name: ${DOCKER_USERNAME}/minitwitimage

logging.json: true
logging.metrics.enabled: false
```

We then installed the `Zap` package for Go, which is a tool for logging. We setup the logging like such:
```Go
var logger *zap.SugaredLogger

func SetupLogger() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	_logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	sugar := _logger.Sugar()
	logger = sugar
	sugar.Info("Sugared logger initialized.")
}

func GetLogger() *zap.SugaredLogger {
	return logger
}
```

We call the `SetupLogging()` function from the `minitwit.go` file. We then manually logged exceptions and errors in the application, where we felt necessary and useful.

From now on, everything that happens in our application will be logged and sent to Elasticsearch using Filebeat. From here the logs are sent to our Kibana dashboard where we can view the logs.

## Test that your logging works

### Introduced bug
When setting up the user session as a user performs an login, we return an error when trying to unmarshal the user object. This happens as we have flipped the `==` operator in the `setUserSession()` function. The bug can be seen here:
```Go
if err == nil {
		return errors.New("Could not marshal json")
	}
```

Our CI/CD setup didn't let us deploy this bug to the production code, as our tests catch the bug when the pull request is created (as seen [here](https://github.com/Lindharden/DevOps/actions/runs/4519248472/jobs/7959629022)). Even when forcefully merging the PR, the code would not be sent to production as our CD script would fail and not deploy the changes. This shows that our CI/CD setup works as intended, however, it's a problem when we intentionally try to implement a bug. Therefore the other team must find the bug in a locally running instance of Kibana.

### Find bug in logs
Logging in to the minitwit application gave an error. Looking in Kibana we see the following log:
```
Mar 25, 2023 @ 14:09:38.054	{"level":"error","timestamp":"2023-03-25T13:09:38Z","caller":"controllers/loginController.go:130","msg":"Could not session","user":"zoinks","stacktrace": ...}
```

The log provided an stacktrace which shows us exactly where the error came from. This means we could easily isolate the problematic component.
