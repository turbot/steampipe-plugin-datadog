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

### List dashboards with restricted editing access

```sql
select
  dashboard.id,
  title as dashboard_title,
  dr.users as role_users
from
  datadog_dashboard as dashboard,
  jsonb_array_elements_text(restricted_roles) as role,
  datadog_role as dr
where
  dr.id = role;
```

### List read-only dashboards (only dashboard author and admins can make changes to it)

```sql
select
  id,
  title,
  is_read_only
from
  datadog_dashboard
where
  is_read_only;
```
