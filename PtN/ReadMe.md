# PtN

This [codeunit](https://github.com/anionDev/PtN/tree/main/PtN) contains the actual source-code of the proxy-server and for this reason some hints are also here.

## Reference

The reference can be found [here](https://github.com/anionDev/PtN/blob/main/PtN/Other/Reference/ReferenceContent/index.md).

## Usage

For an example-docker-compose-file see [the minimal docker-compose-file-example](./Other/Reference/ReferenceContent/Examples/MinimalDockerComposeFile).

Then your `alertmanager.yml`-file should look like this:

```yaml
global:
  resolve_timeout: 5m

route:
  receiver: 'ntfy-via-ptn'
  group_wait: 10s
  group_interval: 30s
  repeat_interval: 1h

receivers:
  - name: 'ntfy-via-ptn'
    webhook_configs:
      - url: 'http://your-ptn-server-ip-address.example.com:8080/my-alert-topic'
        send_resolved: true
```

PtN will automagically recognize `my-alert-topic` as topic and considers this when calling the actual ntfy-server.

## Disclaimer

I am developer, but I am not a Go-developer.
For this reason the source-code of PtN is vibe-coded.
PtN is written in Go because Go is a good language for small container-images without much overhead.
Anyway:
Use on your on risk, no warranty.
