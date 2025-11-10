# How to install IBM RMF for z/OS Grafana Plugin

## Installation steps

- Download a release from https://github.com/IBM/RMF/releases (`ibm-rmf-grafana-<VERSION>.zip`)
  or build it from source code.
- Deploy the downloaded file into a `Grafana` installation (version >= 9.5.1).  
- The `ibm-rmf` Grafana plugin and the two related plugins (`ibm-rmf-datasource` and 
  `ibm-rmf-report-panel`) must be run as unsigned plugins.
  For this specify the below in `Grafana.ini` file’s `[plugins]` section:
  ```ini
  allow_loading_unsigned_plugins = ibm-rmf-app,ibm-rmf-datasource,ibm-rmf-report-panel
  ```
