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
