# Lecture 8: Logging, and Log Analysis. Service-level agreements.
[Week 8](https://github.com/itu-devops/lecture_notes/blob/master/sessions/session_08/README_TASKS.md)

## Add logging to your system
We decided to add EFLK logging to our application. To do this we added Elasticsearch, Logstash, Filebeat and Kibana to our Docker container:
```yaml
minitwit:
...
  volumes:
    - filebeat-data:/logs/
...

elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:7.6.2
    volumes:
        - ./logging/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml:ro
    healthcheck:
        test: ["CMD", "curl", "-s", "-f", "http://localhost:9200/_cat/health"]
        interval: 3s
        timeout: 3s
        retries: 10
    ports:
        - 9200:9200

logstash:
    image: docker.elastic.co/logstash/logstash:7.6.2
    volumes:
        - ./logging/logstash.conf:/usr/share/logstash/pipeline/logstash.conf:ro
    depends_on:
        elasticsearch:
            condition: service_healthy

filebeat:
    image: docker.elastic.co/beats/filebeat:7.2.0
    entrypoint: "filebeat modules enable logstash && filebeat setup && filebeat -e -strict.perms=false"
    user: root
    depends_on:
        elasticsearch:
            condition: service_healthy
    volumes:
        - ./logging/filebeat.yml:/usr/share/filebeat/filebeat.yml:ro
        - ./logging/filebeat_setup.sh:/filebeat_setup.sh
        - filebeat-data:/logs/

kibana:
    image: docker.elastic.co/kibana/kibana:7.6.2
    depends_on:
        elasticsearch:
            condition: service_healthy
    healthcheck:
        test: ["CMD", "curl", "-s", "-f", "http://localhost:5601/api/status"]
        interval: 3s
        timeout: 3s
        retries: 50
    ports:
        - 5601:5601

volumes:
...
  filebeat-data:
    driver: local
...
```

We adjusted the application to write to the `/logs` path, and set Filebeat to read from this path.

We created configuration files for Elasticsearch and Filebeat like such:
```yaml
# Elasticsearch
discovery.type: single-node
network.host: 0.0.0.0
cluster.routing.allocation.disk.threshold_enabled: false

# Filebeat
filebeat.inputs:
- type: log
  paths:
  - /logs/*.log

output.elasticsearch:
  hosts: ["elasticsearch:9200"]
  username: "elastic"
  password: ""

setup.kibana:
  host: "kibana:5601"
```

And for logstash like such:
```conf
input {  
  beats {
    port => 5044
  }
}

filter {
  json {
    source => "message"
  }
}

output {
  elasticsearch {
    hosts => [ "elasticsearch" ]
  }
}
```

We then installed the `Zap` package for Go, which is a tool for logging. We then setup Gin to log every action using Zap like such:
```Go
const logPath = "/logs/minitwit.log"

var logger *zap.Logger

func SetupLogging(r *gin.Engine) {
	os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", logPath}
	logger, _ = c.Build()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
}
```

We call the `SetupLogging()` function from the `minitwit.go` file.

From now on, everything that happens in our application will be logged and sent to Elasticsearch using Filebeat. From here the logs are sent to our Kibana dashboard where we can view the logs.

## Test that your logging works

### Introduced bug

### Find bug in logs