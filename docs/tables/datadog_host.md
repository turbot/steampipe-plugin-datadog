# Table: datadog_dashboard

A host is any piece of infrastructure that the Datadog Agent is installed on such as a bare-metal server or a cloud virtual machine, regardless of provider.

## Examples

### Basic info

```sql
select
  name,
  id,
  up,
  last_reported_time,
  is_muted,
  jsonb_pretty(aliases) as aliases
from
  datadog_host;
```

### Find hosts that don't use systemd and contain the AWS region `ap-southeast-2` in their DNS record

```sql
select
  name,
  jsonb_pretty(apps) as apps
from
  datadog_host
where
  not apps @> '["systemd"]'::jsonb
  and name like '%ap-southeast-2%';
```

### Count instance sizes of each host by their attached AWS tags

```sql
select
  tags,
  count(tags)
from
  (select
    jsonb_array_elements_text(tags_by_source->'Amazon Web Services') as tags
    from
      datadog_host
  ) as foo
where
  tags like '%instance-type%'
group by
  tags;
```


### List hosts that have reported metrics within the last 10 minutes

```sql
select
  name,
  last_reported_time,
  up
from
  datadog_host
where
  last_reported_time > current_timestamp - interval '10 minutes';
```
