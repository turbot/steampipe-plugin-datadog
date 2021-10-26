# Table: datadog_permission

Permissions provide the base level of access for roles.

## Examples

### Basic info

```sql
select
  name,
  id,
  restricted,
  group_name
from
  datadog_permission
order by
  group_name,
  name;
```

### List restricted permissions

```sql
select
  name,
  id,
  restricted,
  group_name
from
  datadog_permission
where
  restricted;
```

### List all the permissions in a specific role

```sql
select
  role.name as role_name,
  dd_perms.name as permission_name,
  dd_perms.description as permission_description
from
  datadog_role as role,
  jsonb_array_elements(permissions) as role_perms,
  datadog_permission as dd_perms
where
  role.name = 'Datadog Standard Role'
  and dd_perms.id = role_perms ->> 'id';
```
