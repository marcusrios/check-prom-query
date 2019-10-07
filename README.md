check-prom-query is a sensu check plugin that queries an Prometheus server and return the output in a [sensu check format](https://docs.sensu.io/sensu-core/0.29/reference/checks/#sensu-check-specification)

### Running

```
check-prom-query -host my-prometheus-server.com -port 30900 -critical 0.5 -warninig 0.3 -query 'container_cpu_usage_seconds_total{pod_name=~"my-pod"}'
```

### Execution options

| Options | Description | Default |
| ------- | ----------- | ------- |
| -host   | Pass the host endpoint | "" |
| -port   | Pass the prometheus port | 30900 |
| -query  | Pass the query in promQL format | "" |
| -critical | Pass the critical threshold that will be evaluated | 0.0 |
| -warning | Pass the warning threshold that will be evaluated | 0.0 |

### Building

Run the [build.sh](build.sh) to generate the binaries
