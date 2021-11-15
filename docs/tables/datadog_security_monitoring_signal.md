# Table: datadog_security_monitoring_signal

If Datadog detects a threat based on any security monitoring rules, it creates a security signal.

Signals are tied to entities such as the actor conducting the attack (e.g., username, IP) or their targets (e.g., hostname, application) so you can easily correlate signals across your infrastructure, applications, and security products to retrace an entire attack.

**Important notes:**

By default this table will list all the signals generated in last 24 hours. To filter results, we recommend specifying the `timestamp` and/or `filter_query` columns in a `where` clause to explore signals based on requirements.

## Examples

### Basic info

```sql
select
  id,
  title,
  timestamp,
  message,
  jsonb_pretty(attributes) as attributes
from
  datadog_security_monitoring_signal;
```

### List signals created in the last 5 days

```sql
select
  id,
  title,
  timestamp,
  message,
  jsonb_pretty(attributes) as attributes
from
  datadog_security_monitoring_signal
where
  timestamp >= (current_date - interval '5' day);
```

### List high status signals

```sql
select
  id,
  title,
  timestamp,
  jsonb_pretty(attributes) as attributes
from
  datadog_security_monitoring_signal
where
  filter_query = 'status:(critical OR high OR medium)';
```

### List AWS S3 signals created in the last 7 days

```sql
select
  id,
  title,
  timestamp,
  jsonb_pretty(attributes) as attributes
from
  datadog_security_monitoring_signal
where
  filter_query = 'scope:s3' and
  timestamp >= (current_date - interval '7' day);
```
