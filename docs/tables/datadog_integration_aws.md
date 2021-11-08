# Table: datadog_integration_aws

Datadog AWS integration resource. Integrations allows to bring together all of the metrics and logs from your infrastructure and gain insight into the unified system.

## Examples

### Basic info

```sql
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

### List AWS integrations having tags to filter ec2 or lambda resources

```sql
select
  account_id,
  excluded_regions,
  jsonb_pretty(filter_tags) as filter_tags
from
  datadog_integration_aws
where
  filter_tags @> '["env:production"]'::jsonb;
```

### List namespaces enabled for metric collection for specific account

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
  and account_id = '013122550996';
```
