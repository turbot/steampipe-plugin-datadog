---
title: "Steampipe Table: datadog_role - Query Datadog Roles using SQL"
description: "Allows users to query Roles in Datadog, providing insights into role-specific details and associated metadata."
---

# Table: datadog_role - Query Datadog Roles using SQL

Datadog is a monitoring service for cloud-scale applications, providing monitoring of servers, databases, tools, and services, through a SaaS-based data analytics platform. Datadog Roles is a feature that allows you to create, modify, and manage custom roles, enabling you to control what users can see and modify within your organization and in the Datadog application. It allows you to provide granular access controls to different types of users.

## Table Usage Guide

The `datadog_role` table provides insights into roles within Datadog. As a DevOps engineer, explore role-specific details through this table, including permissions and associated metadata. Utilize it to uncover information about roles, such as those with specific permissions, and the verification of role policies.

## Examples

### Basic info
Explore which roles have been created in your Datadog account and when they were established. This could be useful for auditing purposes or to understand the evolution of your account's access control structure.

```sql+postgres
select
  name,
  id,
  created_at
from
  datadog_role;
```

```sql+sqlite
select
  name,
  id,
  created_at
from
  datadog_role;
```

### List users assigned the Datadog Admin role
Determine the areas in which users are assigned the Datadog Admin role. This can be useful for managing access control and ensuring only authorized users have administrative privileges.

```sql+postgres
select
  name,
  id,
  jsonb_pretty(users) as users
from
  datadog_role
where
  name = 'Datadog Admin Role';
```

```sql+sqlite
select
  name,
  id,
  users
from
  datadog_role
where
  name = 'Datadog Admin Role';
```

### List all the permissions for a specific role
Explore the various permissions associated with a specific user role to understand the level of access granted. This is useful in managing user roles and ensuring appropriate access rights are provided.

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