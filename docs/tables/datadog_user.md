# Table: datadog_user

Users in the Datadog.

## Examples

### List all users

```sql
select
  *
from
  datadog_user;
```

### List active users

```sql
select
  email,
  id,
  disabled
from
  datadog_user
where
  status = 'Active'
```

### List service accounts

```sql
select
  email,
  id,
  created_at,
  disabled,
  status
from
  datadog_user
where
  service_account;
```

### Users created in the last 7 days

```sql
select
  handle,
  id,
  status,
  created_at
from
  datadog_user
where
  created_at > current_timestamp - interval '7 days';
```
