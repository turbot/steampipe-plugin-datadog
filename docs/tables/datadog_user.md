---
title: "Steampipe Table: datadog_user - Query Datadog Users using SQL"
description: "Allows users to query Users in Datadog, specifically the user's details, including their name, email, status, and role."
---

# Table: datadog_user - Query Datadog Users using SQL

Datadog is a monitoring and analytics tool that helps companies gain visibility into application performance, infrastructure, and logs. The user resource in Datadog represents a user in the Datadog organization. It contains information about the user's details, including their name, email, status, and role.

## Table Usage Guide

The `datadog_user` table provides insights into user details within Datadog. As a system administrator, explore user-specific details through this table, including their name, email, status, and role. Utilize it to uncover information about users, such as their status and role within the organization, and to verify the email addresses associated with each user.

## Examples

### Basic info
Explore the user details in your Datadog account by identifying their email, name and associated roles. This can be useful for auditing user access and understanding the distribution of roles within your team.

```sql+postgres
select
  email,
  name,
  role_ids
from
  datadog_user;
```

```sql+sqlite
select
  email,
  name,
  role_ids
from
  datadog_user;
```

### List active users
Explore which users are currently active in your Datadog account. This can aid in understanding user engagement and activity levels within your system.

```sql+postgres
select
  email,
  id,
  disabled
from
  datadog_user
where
  status = 'Active'
```

```sql+sqlite
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
Explore which Datadog user accounts are service accounts to manage access controls and user privileges effectively. This is useful to ensure security and compliance by identifying potentially unauthorized or redundant service accounts.

```sql+postgres
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

```sql+sqlite
select
  email,
  id,
  created_at,
  disabled,
  status
from
  datadog_user
where
  service_account = 1;
```

### List users created in the last 7 days
Discover the details of recently added users within the past week. This can be useful for monitoring account creation trends or identifying potential security concerns.

```sql+postgres
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

```sql+sqlite
select
  handle,
  id,
  status,
  created_at
from
  datadog_user
where
  created_at > datetime('now', '-7 days');
```