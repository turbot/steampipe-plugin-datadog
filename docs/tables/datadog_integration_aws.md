---
title: "Steampipe Table: datadog_integration_aws - Query Datadog AWS Integrations using SQL"
description: "Allows users to query Datadog AWS Integrations, specifically providing details on integrated AWS accounts, role names, and filter tags."
---

# Table: datadog_integration_aws - Query Datadog AWS Integrations using SQL

Datadog AWS Integration is a feature of Datadog that allows you to monitor your AWS infrastructure in real-time. It provides comprehensive, multi-dimensional views of your AWS environment, helping you to understand performance, troubleshoot issues, and optimize costs. The integration enables you to collect and visualize all your AWS metrics, traces, and logs in one platform.

## Table Usage Guide

The `datadog_integration_aws` table provides insights into AWS integrations within the Datadog platform. As a DevOps engineer, explore integration-specific details through this table, including integrated AWS accounts, role names, and filter tags. Utilize it to uncover information about your integrations, such as which AWS accounts are integrated, the roles used for integration, and the tags applied to filter the data.

## Examples

### Basic info
Explore which roles have access to your AWS account and determine if metrics and resource collection are enabled. This can help identify potential security risks and assess the efficiency of data collection practices.

```sql+postgres
select
  account_id,
  role_name,
  access_key_id,
  excluded_regions,
  metrics_collection_enabled,
  resource_collection_enabled,
  account_specific_namespace_rules
from
  datadog_integration_aws;
```

```sql+sqlite
select
  account_id,
  role_name,
  access_key_id,
  excluded_regions,
  metrics_collection_enabled,
  resource_collection_enabled,
  account_specific_namespace_rules
from
  datadog_integration_aws;
```

### List AWS integrations with "env:production" filter tags
Discover the AWS integrations that have been tagged specifically for a production environment. This can be useful for managing and monitoring resources used in your production workflows.

```sql+postgres
select
  account_id,
  excluded_regions,
  jsonb_pretty(filter_tags) as filter_tags
from
  datadog_integration_aws
where
  filter_tags @> '["env:production"]'::jsonb;
```

```sql+sqlite
Error: The corresponding SQLite query is unavailable.
```

### List namespaces with metric collection enabled for a specific AWS account
This example helps to identify the namespaces within a specific AWS account that have metric collection enabled. This is beneficial for monitoring and analyzing the performance of different services within the account.

```sql+postgres
select
  item.namespace
from
  datadog_integration_aws
  join
    lateral jsonb_each_text(account_specific_namespace_rules) item(namespace, enabled)
    on true
where
  item.enabled::boolean
  and account_id = '123456789012';
```

```sql+sqlite
Error: The corresponding SQLite query is unavailable.
```

### List namespaces with metric collection enabled for a specific AWS account
This example helps to identify the namespaces within a specific AWS account that have metric collection enabled. This is beneficial for monitoring and analyzing the performance of different services within the account.

```sql
select
  item.namespace
from
  datadog_integration_aws
  join
    lateral jsonb_each_text(account_specific_namespace_rules) item(namespace, enabled)
    on true
where
  item.enabled::boolean
  and account_id = '123456789012';
```