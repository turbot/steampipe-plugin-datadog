# Table: datadog_security_monitoring_rule

Security monitoring rules define conditional logic that is applied to all ingested logs and cloud configurations. When at least one case defined in a rule that is matched over a given period of time, Datadog generates a Security Signal.

## Examples

### Basic info

```sql
select
  id,
  name,
  creation_author_id,
  created_at,
  is_default,
  filters,
  tags
from
  datadog_security_monitoring_rule;
```

### List custom monitoring rules

```sql
select
  id,
  name,
  creation_author_id,
  created_at,
  is_default,
  tags,
  filters
from
  datadog_security_monitoring_rule
where
  not is_default;
```

### Filter monitoring rules by tags

```sql
select
  id,
  name,
  creation_author_id,
  created_at,
  is_default,
  tags,
  filters
from
  datadog_security_monitoring_rule
where
  tags @> '["cloud:aws", "source:s3"]'::jsonb
```
