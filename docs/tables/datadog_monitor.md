# Table: datadog_monitor

A monitor provides alerts and notifications if a specific metric is above or below a certain threshold.

## Examples

### Basic info

```sql
select
  name,
  id,
  creator_email,
  overall_state,
  priority,
  query,
  message
from
  datadog_monitor;
```

### List monitors in "Alert" and "Warn" state

```sql
select
  name,
  type,
  created_at,
  message,
  overall_state
from
  datadog_monitor
where
  overall_state in ('Alert', 'Warn');
```

### Filter monitors based on state and tags

```sql
select
  name,
  type,
  created_at,
  overall_state,
  message,
  tags
from
  datadog_monitor
where
  -- monitors in state Alert or Warn
  overall_state in ('Alert', 'Warn') and
  -- monitors having tag aws
  tags @> '["aws"]'::jsonb;
```