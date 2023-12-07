---
title: "Steampipe Table: datadog_host - Query Datadog Hosts using SQL"
description: "Allows users to query Datadog Hosts, specifically the detailed information about each host monitored by Datadog."
---

# Table: datadog_host - Query Datadog Hosts using SQL

Datadog is a monitoring and analytics platform for cloud-scale applications, providing full-stack observability through logs, metrics, and traces. It allows you to monitor, troubleshoot, and optimize application performance, as well as enhance cross-team collaboration. Datadog Hosts are the individual servers, containers, or devices that your application runs on, and this service provides detailed information about each host.

## Table Usage Guide

The `datadog_host` table provides insights into each host monitored by Datadog. As a DevOps engineer, explore host-specific details through this table, including host names, metrics, and associated metadata. Utilize it to uncover information about hosts, such as their status, uptime, and the apps running on them.

## Examples

### Basic info
Explore which Datadog hosts are currently active and when they last reported in. This is useful to maintain an up-to-date understanding of your system's operational status and to quickly identify any potentially muted or unresponsive hosts.

```sql+postgres
select
  name,
  id,
  up,
  last_reported_time,
  is_muted,
  jsonb_pretty(aliases) as aliases
from
  datadog_host;
```

```sql+sqlite
select
  name,
  id,
  up,
  last_reported_time,
  is_muted,
  aliases
from
  datadog_host;
```

### Find hosts that don't use systemd and contain the AWS region `ap-southeast-2` in their DNS record
Explore which hosts in your infrastructure are not utilizing the 'systemd' application and are associated with the 'ap-southeast-2' AWS region. This can be beneficial in identifying potential inconsistencies in your system setup and ensuring regional compliance.

```sql+postgres
select
  name,
  jsonb_pretty(apps) as apps
from
  datadog_host
where
  not apps @> '["systemd"]'::jsonb
  and name like '%ap-southeast-2%';
```

```sql+sqlite
select
  name,
  apps
from
  datadog_host
where
  json_typeof(json_extract(apps, '$.systemd')) is null
  and name like '%ap-southeast-2%';
```

### Count instance sizes of each host by their attached AWS tags
Analyze your AWS-hosted instances to understand the distribution of instance sizes across different hosts. This can help optimize resource allocation and improve cost efficiency.

```sql+postgres
select
  tags,
  count(tags)
from
  (select
    jsonb_array_elements_text(tags_by_source->'Amazon Web Services') as tags
    from
      datadog_host
  ) as foo
where
  tags like '%instance-type%'
group by
  tags;
```

```sql+sqlite
select
  tags,
  count(tags)
from
  (select
    json_extract(tags_by_source, '$."Amazon Web Services"') as tags
    from
      datadog_host
  ) as foo
where
  tags like '%instance-type%'
group by
  tags;
```

### List hosts that have reported metrics within the last 10 minutes
Explore which hosts have recently reported metrics to stay updated on their status. This is useful for real-time monitoring and prompt issue detection.

```sql+postgres
select
  name,
  last_reported_time,
  up
from
  datadog_host
where
  last_reported_time > current_timestamp - interval '10 minutes';
```

```sql+sqlite
select
  name,
  last_reported_time,
  up
from
  datadog_host
where
  last_reported_time > datetime('now', '-10 minutes');
```