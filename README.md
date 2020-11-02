# cloud loggin live tail demo

This sample just generates busy log lines as the sample source for Cloud Logging Live Trailing.

## How to use

```
$ nohup PROJECT_ID=<your project ID> ./cloud-logging_tail &
$ export FILTER="severity=warning AND jsonPayload.fruit=banana"
$ gcloud alpha logging tail "$FILTER"
```