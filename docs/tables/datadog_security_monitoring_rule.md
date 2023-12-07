---
title: "Steampipe Table: datadog_security_monitoring_rule - Query Datadog Security Monitoring Rules using SQL"
description: "Allows users to query Security Monitoring Rules in Datadog, specifically the details and status of each rule, providing insights into security monitoring settings and potential security threats."
---

# Table: datadog_security_monitoring_rule - Query Datadog Security Monitoring Rules using SQL

Datadog Security Monitoring Rules is a feature within Datadog that allows users to define and manage rules for security threats. It provides a centralized way to set up and manage rules for various types of security threats, including network intrusions, unauthorized access, and more. Datadog Security Monitoring Rules helps you stay informed about the security status of your resources and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `datadog_security_monitoring_rule` table provides insights into Security Monitoring Rules within Datadog. As a security engineer, explore rule-specific details through this table, including rule configurations, conditions, and associated metadata. Utilize it to uncover information about rules, such as those related to specific security threats, the conditions that trigger them, and the actions taken when those conditions are met.

## Examples

### Basic info
Explore which security monitoring rules have been created on your Datadog platform. This allows you to understand who created each rule, when they were created, and any filters or tags applied, helping you manage and organize your security protocols effectively.

```sql+postgres
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

```sql+sqlite
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
Uncover the details of custom security monitoring rules in your system, focusing on those that are not default, to better understand your security landscape and identify potential areas of improvement. This query is particularly beneficial for those seeking to optimize their security settings and ensure that custom rules are properly configured and functioning as expected.

```sql+postgres
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

```sql+sqlite
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
  is_default = 0;
```

### Filter monitoring rules by tags
Explore which monitoring rules have been specifically tagged for AWS cloud and S3 source. This allows you to quickly identify and review the rules applicable to your AWS S3 resources.

```sql+postgres
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

```sql+sqlite
Error: SQLite does not support the contains operator (@>) for JSON objects.
```