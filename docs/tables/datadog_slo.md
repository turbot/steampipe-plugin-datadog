# Table: datadog_slo

<description>

## Examples

### Basic info

```sql
select
  name,
  type,
  thresholds,
  created_at,
  creator_email
from
  datadog_slo;
```

### List metric type SLOs

```sql
select
  name,
  type,
  created_at,
  monitor_ids,
from
  datadog_slo
where
  type in ('metric');
```

### List SLOs that are type monitor and have thresholds set to 2.5 9's over 7 days

```sql
select
  name,
  type,
  thresholds,
  created_at
from
  datadog_slo
where
  type in ('monitor') and
  thresholds @> '[{"target":99.5,"target_display":"99.5","timeframe":"7d"}]'::jsonb;
```
