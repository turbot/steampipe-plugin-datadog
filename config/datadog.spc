connection "datadog" {
  plugin = "datadog"

  # Authentication - API Key
  # Steampipe will resolve the API key in below order:
  #  1. The "api_key" specified here in the config
  #  2. The `DD_CLIENT_API_KEY` environment variable
  # api_key   = "1a2345bc6d78e9d98fa7bcd6e5ef56a7"

  # Authentication - APP Key
  # Steampipe will resolve the APP key in below order:
  #  1. The "app_key" specified here in the config
  #  2. The `DD_CLIENT_APP_KEY` environment variable
  # app_key   = "b1cf234c0ed4c567890b524a3b42f1bd91c111a1"
}
