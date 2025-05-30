# Change Log

## 1.0.10 (2025-05-30)

- Fixed a Windows-specific regression that caused the plugin to serve data points with the incorrect TZ offset.

## 1.0.9 (2025-05-22)

- IBM RMF for z/OS Grafana plugin has been enhanced to optimize the step size when querying and displaying time series data, based on the selected time range and panel size. This enhancement ensures that panels display the required number of data points.
- A bug has been fixed that was causing specific data points to be dropped from time series views, resulting in a flat line in the graphs.
- Fixed internally found defects.

## 1.0.8 (2025-04-04)

- A new option, Compression, has been added to the data source creation process. This option enables the RMF data source to request HTTP compression from the Data Distributed Server (DDS) during data retrieval, which improves the efficiency of data transfer.
- Fixed certain scenarios of missing data within dashboards for time series visualizations.
- Fixed internally found defects.

## 1.0.7 (2024-11-29)

- Fixed internally found defects.
- You can now download IBM RMF for z/OS Grafana plugin documentation as a PDF file.

## 1.0.6 (2024-10-02)

- The IBM RMF for z/OS Grafana plugin now allows customizing banner and caption fields in IBM RMF Report visualizations.
- After you modify any IBM RMF data source and save the changes, you no longer need to re-enter the password.
- Improved performance.
- Fixed internally found defects.

## 1.0.5 (2024-08-02)

- When creating the datasource, you now have the option to specify the desired cache size in megabytes for the datasource.
- Fixed internally found defects.
- The topic of Creating RMF data sources has been updated per the new UI changes.

## 1.0.4 (2024-06-21)

- The user interface for creating the datasource has been enhanced to improve the user experience.
- Fixed internally found defects.
- The following dashboards are updated:
	- Common Storage Activity (Timeline)
	- Common Storage Activity
	- Coupling Facility Overview (Timeline)
	- Coupling Facility Overview
	- Execution Velocity (Timeline)
	- Execution Velocity
	- General Activity (Timeline)
	- General Activity
	- Overall Image Activity (Timeline)
	- Overall Image Activity
	- Performance Index (Timeline)
	- Response Time (Timeline)
	- Response Time
	- XCF Activity (Timeline)
	- XCF Activity  
	
	**Note:** You must re-import the dashboards to utilize the latest enhancements in the dashboards.
- Improved documentation on datasources and upgrading the plugin.

## 1.0.3 (2024-02-29)

- Time-series dashboards plot graphs correctly even when NaN values are in the historical data.
- Fixed specific issues related to the support of Grafana v10.x.x.
- Improved documentation on Grafana through z/OSMF, installation, and troubleshooting.
- Fixed empty settings issue for IBM RMF data source defined via Home / Apps / IBM RMF page.
- Fixed other internally found defects.

## 1.0.2 (2023-12-08)

- Support spaces in Datasource name.
- Correct quotes in documentation for installation actions.
- Update Golang dependencies.

## 1.0.1 (2023-11-28)

Resolved security issues in dependencies.

## 1.0.0 (2023-11-16)

Initial release.
