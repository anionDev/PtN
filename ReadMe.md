# PtN

PtN ("Prometheus to ntfy") is a proxy to receive events from a [prometheus-alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/)-server and forwards them to a [ntfy](https://github.com/binwiederhier/ntfy)-server.

For usage-hints of the proxy and its configuration see [the ReadMe.md of the codeunit](https://github.com/anionDev/PtN/blob/main/PtN/ReadMe.md).
For a general reference of the PtN-project see [here](https://github.com/anionDev/PtN/blob/main/Other/Reference/Reference.md).

## Why?

Why do we need another tool for that usecase?
I wanted to forward alerts from the prometheus-alertmanager to a ntfy-server but unfortunately they have different formats/standards/protocols to send/receive data.
And there are already tools which are adding exactly this compatibility-proxy (for example [this](https://github.com/alexbakker/alertmanager-ntfy) and [that](https://github.com/pinpox/alertmanager-ntfy)) but I could not find a container-image for these tools, at least at this moment where I needed it.
So I decided to relase a small proxy-server which is doing exactly the required data-conversion between alertmanager and ntfy by myself.

## Quick-start

Just place PtN beside your alertmanager. See the following `docker-compose.yml` for an example:

```yaml
services:

  alertmanager:
    image: prom/alertmanager:latest
    container_name: alertmanager
    volumes:
      - ./Volumes/Configuration:/etc/alertmanager
    command:
      - '--config.file=/etc/alertmanager/alertmanager.yml'
    networks:
      - alertmanager_net

  alertmanager_ptn:
    container_name: alertmanager_ptn
    environment:
      - NTFY_SERVER=https://my-ntfy-server.example.com
    image: aniondev/ptn:latest
    networks:
      - alertmanager_net

networks:
  alertmanager_net:
    external: true
```

Then your `alertmanager.yml`-file can look something like this:

```yaml
global:
  resolve_timeout: 5m

route:
  receiver: 'ntfy'
  group_wait: 10s
  group_interval: 30s
  repeat_interval: 1h

receivers:
  - name: 'ntfy'
    webhook_configs:
      - url: 'http://alertmanager_ptn:8080/my-alert-topic'
        send_resolved: true
```

## Changelog

See the [Changelog-folder](./Other/Resources/Changelog).

## Contribue

Contributions are always welcome.

This product has the contribution-requirements defines by [DefaultOpenSourceContributionProcess](https://projects.aniondev.de/PublicProjects/Common/ProjectTemplates/-/blob/main/Conventions/Contributing/DefaultOpenSourceContributionProcess/DefaultOpenSourceContributionProcess.md).

## Repository-structure

This product uses the [CommonProjectStructure](https://projects.aniondev.de/PublicProjects/Common/ProjectTemplates/-/blob/main/Conventions/RepositoryStructure/CommonProjectStructure/CommonProjectStructure.md) as repository-structure.

## Branching-system

This product follows the [GitFlowSimplified](https://projects.aniondev.de/PublicProjects/Common/ProjectTemplates/-/blob/main/Conventions/BranchingSystem/GitFlowSimplified/GitFlowSimplified.md)-branching-system.

## Versioning

This product follows the [SemVerPractise](https://projects.aniondev.de/PublicProjects/Common/ProjectTemplates/-/blob/main/Conventions/Versioning/SemVerPractise/SemVerPractise.md)-versioning-system.

## License

See [License.txt](./License.txt) for license-information.
