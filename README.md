# mackerel-plugin-battery

## Description

Mackerel metrics plugin to get batteries' info.

## Synopsis

```sh
mackerel-plugin-battery
```

## Installation

```sh
sudo mkr plugin install Arthur1/mackerel-plugin-battery
```

## Setting for mackerel-agent

```
[plugin.metrics.battery]
command = ["/opt/mackerel-agent/plugins/bin/mackerel-plugin-battery"]
```
