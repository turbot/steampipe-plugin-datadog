# Table: datadog_log

This table lists events from a Datadog logs collection.

**Important notes:**

- By default, this table list out events only for last 15 minutes.
- You **_must_** specify `timestamp` in a `where` clause in order to query logs for time period of your choice.
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

### List events for last two days

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

### List AWS `s3.amazonaws.com` service logs for last two days using query

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
  query = '@detail.eventSource:s3.amazonaws.com'
  and timestamp >= (current_date - interval '2' day);
```

### List all AWS S3 events with bucket name for last two days

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

### List AWS S3 Buckets created or deleted in last two days

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
  and timestamp >= (current_date - interval '2' day);
```

### List events for a specific day (example 4th day from now)

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
  timestamp <= (current_date - interval '4' day)
  and timestamp >= (current_date - interval '5' day);
```
