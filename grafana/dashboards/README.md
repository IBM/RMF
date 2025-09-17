# Grafana Sample Dashboards

The IBM RMF Monitor III metrics can be exposed to third-party tools directly using the OpenMetrics exposition format.
Some examples of the tools are:
- Prometheus
- VictoriaMetrics
- Telegraf


The directory contains Grafana sample dashboards for Prometheus and VictoriaMetrics to visualize the Monitor III data.  
Note: VictoriaMetrics is a drop-in replacement for Prometheus, the same dashboards can be used with both.

[Go to Prometheus / VictoriaMetrics Configuration](#prometheus--victoriametrics-configuration)  
[Go to Sample Dashboards Usage](#sample-dashboards-usage)

## Prometheus / VictoriaMetrics Configuration

The Monitor III data can be exposed through the `/metrics/m3` endpoint.  
**Note**: OA67701 has to be applied for the feature.

Make sure you have the Monitor III metrics being scraped by Prometheus or VictoriaMetrics, and a Grafana data source pointing to the Prometheus / VictoriaMetrics instance is configured. 

A target definition example to scrape all the metrics as defined in GPMOMC (to limit metrics being scraped, use the `groups` parameter of the endpoint or edit GPMOMC DDS configuration):

```yaml
- job_name: "m3@plex00"
  scrape_interval: 100s # Should be equal to the Monitor III mintime
  scrape_timeout: 50s
  metrics_path: /metrics/m3
  scheme: https # Remove or change to "http" if AT/TLS for DDS is not set
  tls_config:
  insecure_skip_verify: false # Change to “true” if self-signed certificates are used
  basic_auth: # Use DDS credentials or remove if DDS authentication is disabled
  username: ‘DDSUSER'
  password: 'XXXXXXX'
  static_configs:
  - targets: [ "ddshostname:8803” ]
```

Refer to [z/OS Resource Measurement Facility User's Guide](https://www.ibm.com/docs/en/zos/latest?topic=rmf-zos-resource-measurement-facility-users-guide) for more information.

## Sample Dashboards Usage

The dashboards can be imported into Grafana in three ways:
- [Importing the dashboards directly into Grafana.](#import-dashboards)
- [Provision the dashboards using Grafana's provisioning system.](#provision-dashboards)
- [Installing IBM RMF for z/OS Grafana Plugin.](#install-the-ibm-rmf-for-zos-grafana-plugin)

### Import Dashboards

You can import dashboards directly into your Grafana instance via the UI:
- Log in to Grafana.
- Navigate to Dashboards → New → Import.
- Upload a <dashboard>.json file from this repository or paste its contents.

Refer to [the Import Dashboard docs](https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/import-dashboards/#import-dashboard) for more information.

### Provision Dashboards

For automation and version control, you can provision dashboards using Grafana's provisioning system:
- Copy the .json files from this repo into your `provisioning/directory` (the exact location depends on your Grafana installation and configuration).
- Define them in your provisioning YAML config, for example:
  ```yaml
  apiVersion: 1
  providers:
    # <string> an unique provider name. Required
    - name: 'a unique provider name'
      # <int> Org id. Default to 1
      orgId: 1
      # <string> name of the dashboard folder.
      folder: ''
      # <string> folder UID. will be automatically generated if not specified
      folderUid: ''
      # <string> provider type. Default to 'file'
      type: file
      # <bool> disable dashboard deletion
      disableDeletion: false
      # <int> how often Grafana will scan for changed dashboards
      updateIntervalSeconds: 10
      # <bool> allow updating provisioned dashboards from the UI
      allowUiUpdates: false
      options:
        # <string, required> path to dashboard files on disk. Required when using the 'file' type
        path: /var/lib/grafana/dashboards
        # <bool> use folder names from filesystem to create folders in Grafana
        foldersFromFilesStructure: true
  ```

Refer to [the Provisioning Dashboards docs](https://grafana.com/docs/grafana/latest/administration/provisioning/#dashboards) for more information.


### Install the IBM RMF for z/OS Grafana Plugin

All the sample dashboards are bundled into [IBM RMF for z/OS Grafana Plugin](grafana/rmf-app/README.md).  
Follow the instructions to install it and [deploy the sample dashboards](https://ibm.github.io/RMF/grafana/rmf-app/prometheus_sample_dashboards.html) from IBM RMF App in Grafana.
