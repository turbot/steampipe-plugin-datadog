---
title: "Steampipe Table: datadog_permission - Query Datadog Permissions using SQL"
description: "Allows users to query Permissions in Datadog, specifically the permissions assigned to roles, providing insights into role capabilities and potential security risks."
---

# Table: datadog_permission - Query Datadog Permissions using SQL

Datadog is a monitoring and analytics platform that allows you to see inside any stack, any app, at any scale, anywhere. With Datadog Permissions, you can manage and control what actions a user or a group of users can perform in your organization. Permissions are assigned to roles, which can then be assigned to users, providing granular control over access and actions within Datadog.

## Table Usage Guide

The `datadog_permission` table provides insights into Permissions within Datadog. As a security analyst, explore permission-specific details through this table, including the roles they are assigned to, and their associated metadata. Utilize it to uncover information about permissions, such as those with high-level access, the roles associated with each permission, and the potential security risks.

## Examples

### Basic info
Explore the permissions within your Datadog setup to understand which are restricted and how they are grouped. This can help ensure appropriate access levels and maintain security standards.

```sql+postgres
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

```sql+sqlite
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
Analyze the settings to understand which permissions are restricted in Datadog. This is beneficial in managing user access and ensuring security protocols are adhered to.

```sql+postgres
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

```sql+sqlite
select
  name,
  id,
  restricted,
  group_name
from
  datadog_permission
where
  restricted = 1;
```

### List all the permissions for a specific role
Determine the areas in which a particular role has access by identifying the permissions associated with it. This can be useful for auditing security measures and ensuring appropriate access levels.

```sql+postgres
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

```sql+sqlite
select
  role.name as role_name,
  dd_perms.name as permission_name,
  dd_perms.description as permission_description
from
  datadog_role as role,
  json_each(permissions) as role_perms,
  datadog_permission as dd_perms
where
  role.name = 'Datadog Standard Role'
  and dd_perms.id = json_extract(role_perms.value, '$.id');
```