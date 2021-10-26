# Table: datadog_logs_metric

Log-based metrics are a cost-efficient way to summarize log data from the entire ingest stream. This means that even if you use exclusion filters to limit what you store for exploration, you can still visualize trends and anomalies over all of your log data at 10s granularity for 15 months.

With log-based metrics, you can generate a count metric of logs that match a query or a distribution metric of a numeric value contained in the logs, such as request duration.

## Examples

### Basic info

```sql
select
  id,
  compute_aggregation_type,
  compute_path,
  filter_query,
  jsonb_pretty(group_by) as group_by
from
  datadog_logs_metric;
```
