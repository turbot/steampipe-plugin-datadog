![image](https://hub.steampipe.io/images/plugins/turbot/datadog-social-graphic.png)

# Datadog Plugin for Steampipe

Use SQL to query dashboards, users, roles and more from Datadog.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/datadog)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/datadog/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-datadog/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install datadog
```

Run a query:

```sql
select name, type, overall_state from datadog_monitor where overall_state in ('Alert', 'Warn');
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-datadog.git
cd steampipe-plugin-datadog
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/datadog.spc
```

Try it!

```
steampipe query
> .inspect datadog
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-datadog/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-datadog/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Datadog Plugin](https://github.com/turbot/steampipe-plugin-datadog/labels/help%20wanted)
