# RMF

Resource Measurement Facility for z/OS

## Grafana Sample Dashboards

The IBM RMF Monitor III metrics can be exposed to third-party tools directly using the OpenMetrics exposition format.
Some examples of the tools are:
- Prometheus
- VictoriaMetrics
- Telegraf

Grafana sample dashboards for Prometheus and VictoriaMetrics as well as usage information are available in the [grafana/dashboards](grafana/dashboards) directory.  

Refer to [z/OS Resource Measurement Facility User's Guide](https://www.ibm.com/docs/en/zos/latest?topic=rmf-zos-resource-measurement-facility-users-guide) for more information on configuring RMF DDS.

## IBM RMF for z/OS Grafana Plugin

The IBM® RMF for z/OS Grafana plugin provides effortless analysis and visualization of RMF Monitor III metrics and reports within the Grafana platform.

See [IBM RMF for z/OS Grafana Plugin](grafana/rmf-app/README.md) for more details.

## IBM RMF Performance Monitoring

The IBM® RMF Performance Monitoring (RMF PM) is a tool that enables you to manage the performance of your z/OS host right from your workstation. By establishing a TCP/IP connection to one or more z/OS sysplexes, you can access real-time data on the resources of the corresponding sysplex.

RMF PM offers flexibility in creating scenarios to monitor the system's performance. You can sample data from different resources and combine it to comprehensively overview the system's performance. With RMF PM, you can visualize the performance data as bar charts to better understand the system's status.

You can download the installer from the following location: 
https://github.com/IBM/RMF/releases

Refer to the [z/OS Resource Measurement Facility User's Guide](https://www.ibm.com/docs/en/zos/latest?topic=monitoring-performance-overview) for more information.

## IBM RMF Spreadsheet Reporter

The RMF Spreadsheet Reporter is a workstation solution that enables users to visualize RMF (Resource Measurement Facility) Postprocessor data in a graphical format. With the RMF Spreadsheet Reporter, you can convert your RMF data into a spreadsheet format, which makes it easier to analyze and manipulate the data.

You can download the installer from the following location: (https://github.com/IBM/RMF/releases).

Refer to the [z/OS Resource Measurement Facility User's Guide](https://www.ibm.com/docs/en/zos/latest?topic=reporter-concepts-performance-analysis-rmf-spreadsheet) for more information.
