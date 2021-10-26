# Table: datadog_dashboard

A dashboard is Datadog’s tool for visually tracking, analyzing, and displaying key performance metrics.

## Examples

### Basic info

```sql
select
  id,
  author_handle,
  layout_type,
  url,
  is_read_only,
  created_at,
  jsonb_pretty(restricted_roles) as restricted_roles
from
  datadog_dashboard;
```
