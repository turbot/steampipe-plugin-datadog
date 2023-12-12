---
title: "Steampipe Table: datadog_log_event - Query Datadog Log Events using SQL"
description: "Allows users to query Datadog Log Events, specifically the event details, status, and associated metadata, providing insights into system logs and potential anomalies."
---

# Table: datadog_log_event - Query Datadog Log Events using SQL

Datadog Log Management is a service that centralizes logs from different sources into a unified platform. It provides real-time search and analytics, which allows users to explore, troubleshoot, and monitor their log data with ease. It also integrates with other Datadog services to provide a comprehensive view of system performance.

## Table Usage Guide

The `datadog_log_event` table provides insights into log events within Datadog Log Management. As a system administrator or a DevOps engineer, explore event-specific details through this table, including event status, associated metadata, and other relevant information. Utilize it to uncover information about system logs, such as those with specific events, the status of these events, and the verification of event metadata.

## Examples

### Basic info
Explore the events in your Datadog logs to gain insights into their status and associated attributes. This can be particularly useful for troubleshooting or identifying patterns in your system's behavior.

```sql+postgres
select
  id,
  timestamp,
  service,
  status,
  jsonb_pretty(attributes) as attributes
from
  datadog_log_event;
```

```sql+sqlite
select
  id,
  timestamp,
  service,
  status,
  attributes
from
  datadog_log_event;
```

## Timestamp Examples

### List events for the last two days
Explore recent activities on your services by identifying events that have occurred in the last two days. This query can help you stay updated with the status and messages of your services, allowing you to quickly react to any potential issues.

```sql+postgres
select
  timestamp,
  service,
  status,
  message,
  jsonb_pretty(attributes) as attributes
from
  datadog_log_event
where
  timestamp >= (current_date - interval '2' day);
```

```sql+sqlite
select
  timestamp,
  service,
  status,
  message,
  attributes
from
  datadog_log_event
where
  timestamp >= date('now', '-2 day');
```

### List events in a specific time range
Explore events that occurred within a specific time frame. This can be useful in identifying patterns or anomalies in service status over a given period.

```sql+postgres
select
  timestamp,
  service,
  status,
  message,
  jsonb_pretty(attributes) as attributes
from
  datadog_log_event
where
  timestamp <= (current_date - interval '2' day)
  and timestamp >= (current_date - interval '5' day);
```

```sql+sqlite
select
  timestamp,
  service,
  status,
  message,
  attributes
from
  datadog_log_event
where
  timestamp <= (date('now','-2 day'))
  and timestamp >= (date('now','-5 day'));
```

## Query Examples

### List all AWS S3 events for the last two days
Explore recent activity in your AWS S3 storage by identifying events that occurred in the last two days. This is useful for monitoring usage, tracking changes, and identifying potential security concerns.

```sql+postgres
select
  timestamp,
  service,
  status,
  message,
  attributes -> 'detail' ->> 'eventName' as event_name,
  attributes -> 'detail' -> 'requestParameters' ->> 'bucketName' as bucket_name
from
  datadog_log_event
where
  query = '@detail.eventSource:s3.amazonaws.com'
  and timestamp >= (current_date - interval '2' day);
```

```sql+sqlite
select
  timestamp,
  service,
  status,
  message,
  json_extract(attributes, '$.detail.eventName') as event_name,
  json_extract(attributes, '$.detail.requestParameters.bucketName') as bucket_name
from
  datadog_log_event
where
  query = '@detail.eventSource:s3.amazonaws.com'
  and timestamp >= date('now','-2 day');
```

### List AWS S3 buckets created or deleted in the last week
Discover recent changes in your AWS S3 storage by identifying buckets that were created or deleted in the past week. This can help maintain an overview of your storage usage and track significant modifications.

```sql+postgres
select
  timestamp,
  service,
  status,
  message,
  attributes -> 'detail' ->> 'eventName' as event_name,
  attributes -> 'detail' -> 'requestParameters' ->> 'bucketName' as bucket_name
from
  datadog_log_event
where
  query = '@detail.eventName:(CreateBucket OR DeleteBucket)'
  and timestamp >= (current_date - interval '7' day);
```

```sql+sqlite
select
  timestamp,
  service,
  status,
  message,
  json_extract(attributes, '$.detail.eventName') as event_name,
  json_extract(attributes, '$.detail.requestParameters.bucketName') as bucket_name
from
  datadog_log_event
where
  query = '@detail.eventName:(CreateBucket OR DeleteBucket)'
  and timestamp >= date('now','-7 day');
```