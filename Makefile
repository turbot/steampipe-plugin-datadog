STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo
install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/datadog@latest/steampipe-plugin-datadog.plugin -tags "${BUILD_TAGS}" *.go
