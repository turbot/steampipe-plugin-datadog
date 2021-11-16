connection "datadog" {
  plugin = "datadog"

  # Get your API key from https://app.datadoghq.com/organization-settings/api-keys
  # Steampipe will resolve the API key in below order:
  #   1. The "api_key" specified here in the config
  #   2. The `DD_CLIENT_API_KEY` environment variable
  # api_key = "1a2345bc6d78e9d98fa7bcd6e5ef56a7"

  # Get your application key from https://app.datadoghq.com/organization-settings/application-keys
  # Steampipe will resolve the application key in below order:
  #   1. The "app_key" specified here in the config
  #   2. The `DD_CLIENT_APP_KEY` environment variable
  # app_key = "b1cf234c0ed4c567890b524a3b42f1bd91c111a1"

  # The server Host. By default it is pointed to "datadoghq.com"
  # If working with the EU version of Datadog, use "datadoghq.eu"
  # server = "datadoghq.com"
}
