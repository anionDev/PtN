# PtN

How to use:

For an example-docker-compose-file see [the minimal docker-compose-file-example](./Other/Reference/ReferenceContent/Examples/MinimalDockerComposeFile).

As you can see there NtP supports the following environment-variables:

- `NTFY_SERVER`  (required)
- `NTFY_USER` (default: (empty))
- `NTFY_PASS` (default: (empty))
- `PORT` (default: `8080`)

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

Disclaimer: PtN is vibe-coded. No warranty.
