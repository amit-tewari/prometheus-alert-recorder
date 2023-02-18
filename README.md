# prometheus-alerts-recorder
Record all alerts triggered by a prometheus instance

## Why this application?
if you wish to
- understand prometheus to alertmanager interactions
- analyze alerts sent by prometheus instance for debugging or fine-tuning
- generate reports on alerts, and slice and dice

## How to use?
- Export environment variables (optionaly use envrc.sample and direnv)
- Build and push docker image
- Create a helm release (test release defined in Makefile, with right env vars)
- Add a config like following (customize service name in target) 
```
alerting:
  alertmanagers:
    - static_configs:
      - targets: ['prometheus-alerts-recorder-dd-prometheus-alerts-recorder:4000']
```

## How to get data

Log lines like following are emmited on pod's logs, accessible using `kubectl logs -f prometheus-alerts-recorde-pod-name`
```
2023/02/18 11:47:26 Using Log file [ /tmp/prometheus-alerts-recorder.txt ]
{"labels":{"alertname":"HostOutOfMemory","instance":"host.docker.internal:9100","job":"host-metrics","severity":"warning"}}
```

Also, image installs jq and vim, which can be used to view entries from `/tmp/prometheus-alerts-recorder.txt` file (or the one you specified.

If you are sending logs of this pod to ElasticSearch/ GCP logs etc, you can view stats there too.
