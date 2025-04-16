# PtN-reference

## General

The PtN-project archives the goal to provide a minimalistic proxy-server to be able to retrieve messages from Prometheus alertmanager and forward them to a ntfy-server.

For usage-hints of the proxy and its configuration see [the ReadMe.md of the codeunit](https://github.com/anionDev/PtN/blob/main/PtN/ReadMe.md).

## Codeunit-overview

![CodeUnits-Overview.svg](./Technical/Diagrams/CodeUnits-Overview.svg)

## Network-overview

PtN acts in the middle between a Prometheus alertmanager and a ntfy-server.
Technically while communicating with the Prometheus alertmanager PtN is the server (therefor it has to open a port) and while communicating with the ntfy-server PtN is the client.
All network-traffic uses the HTTP-protocol.

![Components.svg](./Technical/Diagrams/Components.svg)

![Communication.svg](./Technical/Diagrams/Communication.svg)
