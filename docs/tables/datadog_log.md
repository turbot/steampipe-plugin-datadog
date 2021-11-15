# Table: datadog_log

This table lists events from a Datadog logs collection.

**Important notes:**

- By default, this table lists events for the last 15 minutes.
- You can specify `timestamp` in a `where` clause in order to query logs for a time period of your choice.
- You can also use `query` in a `where` clause to search for logs matching a specific criterion. Refer [Log Search Syntax](https://docs.datadoghq.com/logs/explorer/search_syntax/)

## Examples

### Basic info

```sql
select
  id,
  timestamp,
  service,
  status,
  jsonb_pretty(attributes) as attributes
from
  datadog_log;
```

## Timestamp Examples

### List events for the last two days

```sql
select
  timestamp,
  service,
  status,
  message,
  jsonb_pretty(attributes) as attributes
from
  datadog_log
where
  timestamp >= (current_date - interval '2' day);
```

### List events in a specific time range

```sql
select
  timestamp,
  service,
  status,
  message,
  jsonb_pretty(attributes) as attributes
from
  datadog_log
where
  timestamp <= (current_date - interval '2' day)
  and timestamp >= (current_date - interval '5' day);
```

## Query Examples

### List all AWS S3 events for the last two days

```sql
select
  timestamp,
  service,
  status,
  message,
  attributes -> 'detail' ->> 'eventName' as event_name,
  attributes -> 'detail' -> 'requestParameters' ->> 'bucketName' as bucket_name
from
  datadog_log
where
  query = '@detail.eventSource:s3.amazonaws.com'
  and timestamp >= (current_date - interval '2' day);
```

### List AWS S3 buckets created or deleted in the last week

```sql
select
  timestamp,
  service,
  status,
  message,
  attributes -> 'detail' ->> 'eventName' as event_name,
  attributes -> 'detail' -> 'requestParameters' ->> 'bucketName' as bucket_name
from
  datadog_log
where
  query = '@detail.eventName:(CreateBucket OR DeleteBucket)'
  and timestamp >= (current_date - interval '7' day);
```
