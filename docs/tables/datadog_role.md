# Table: datadog_role

Roles categorize users and define what account permissions those users have, such as what data they can read or what account assets they can modify.

## Examples

### Basic info

```sql
select
  name,
  id,
  created_at
from
  datadog_role;
```

### List users attached to Datadog admin role

```sql
select
  name,
  id,
  jsonb_pretty(users) as users
from
  datadog_role
where
  name = 'Datadog Admin Role';
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
