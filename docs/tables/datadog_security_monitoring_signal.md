---
title: "Steampipe Table: datadog_security_monitoring_signal - Query Datadog Security Monitoring Signals using SQL"
description: "Allows users to query Security Monitoring Signals in Datadog, specifically the detection of threats in real-time, providing insights into security incidents and potential vulnerabilities."
---

# Table: datadog_security_monitoring_signal - Query Datadog Security Monitoring Signals using SQL

Datadog Security Monitoring is a feature within Datadog that allows real-time threat detection across your applications and infrastructure. It provides a centralized way to set up and manage alerts for various security incidents, including potential vulnerabilities, unauthorized access, and more. Datadog Security Monitoring helps you stay informed about the security status of your resources and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `datadog_security_monitoring_signal` table provides insights into Security Monitoring Signals within Datadog. As a Security Analyst, explore signal-specific details through this table, including threat levels, incident times, and associated metadata. Utilize it to uncover information about security incidents, such as those related to potential vulnerabilities, the severity of the incidents, and the verification of incident responses.

## Examples

### Basic info
Explore the various signals from your Datadog security monitoring system. This allows you to gain insights into the events and notifications, helping you better understand your system's security status.

```sql+postgres
select
  id,
  title,
  timestamp,
  message,
  jsonb_pretty(attributes) as attributes
from
  datadog_security_monitoring_signal;
```

```sql+sqlite
select
  id,
  title,
  timestamp,
  message,
  attributes
from
  datadog_security_monitoring_signal;
```

### List signals created in the last 5 days
Discover the segments that have generated signals in the past 5 days. This can help in identifying recent security issues or changes within the system that may require attention.

```sql+postgres
select
  id,
  title,
  timestamp,
  message,
  jsonb_pretty(attributes) as attributes
from
  datadog_security_monitoring_signal
where
  timestamp >= (current_date - interval '5' day);
```

```sql+sqlite
select
  id,
  title,
  timestamp,
  message,
  attributes
from
  datadog_security_monitoring_signal
where
  timestamp >= date('now','-5 day');
```

### List high status signals
Explore high-priority security signals in your Datadog environment to proactively address potential threats and maintain system integrity.

```sql+postgres
select
  id,
  title,
  timestamp,
  jsonb_pretty(attributes) as attributes
from
  datadog_security_monitoring_signal
where
  filter_query = 'status:(critical OR high OR medium)';
```

```sql+sqlite
select
  id,
  title,
  timestamp,
  attributes
from
  datadog_security_monitoring_signal
where
  filter_query = 'status:(critical OR high OR medium)';
```

### List AWS S3 signals created in the last 7 days
Determine the areas in which AWS S3 signals have been created in the past week. This is useful for monitoring recent activity and identifying potential security issues.

```sql+postgres
select
  id,
  title,
  timestamp,
  jsonb_pretty(attributes) as attributes
from
  datadog_security_monitoring_signal
where
  filter_query = 'scope:s3' and
  timestamp >= (current_date - interval '7' day);
```

```sql+sqlite
select
  id,
  title,
  timestamp,
  attributes
from
  datadog_security_monitoring_signal
where
  filter_query = 'scope:s3' and
  timestamp >= date('now','-7 day');
```