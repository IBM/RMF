# Change Log

## 5.5.9 (2025-05-27)

- New Java 17 Runtime
- Fix Secure FTP connection

## 5.5.8 (2024-09-11)

- The Spreadsheet Reporter has been updated to include the latest Java version 17.

## 5.5.7

- Support CPU reports with multi-page IFL information
- Fix WLMGL working set creation
- Reset buttons in CF Trend macro
- General maintenance

## 5.5.6

- Upgrade to Java 17.
- Sort by date/time no longer default in RepCFSys, RepSubChn1 and RepSubChn2 graphs.

## 5.5.5

- General maintenance.
- Support for AI-infused batch initiators in WLMGL report.

## 5.5.4

- New Java 11 Runtime.

## 5.5.3

- Update XML report display processing.
- Browse XML report through local web server.

## 5.5.2

- New Java 11 Runtime.
- z/OS 3.1 support.

## 5.5.1

- Date ordering fix for CF reports.

## 5.5.0

- Upgrade to Java 11.

## 5.4.46

- New Java 8 Runtime.

## 5.4.45

- Keep remote dataset at transfer failure.
- Additional XCF Spreadsheet performance optimizations.
- CF Spreadsheet named range count overflow fix.
- Adapt to the new XCF report format.
- Fix CF report parsing.

## 5.4.44

- WLM Enclave Transaction metrics parsing.

## 5.4.43

- Fixing "Error(s) occurred during conversion." while creating working set.
- XCF Spreadsheet performance optimization.

## 5.4.42

- New Java 8 Runtime.

## 5.4.41

- New Java 8 Runtime.

## 5.4.40

- Workload Activity report working set loading to spreadsheet regression fix.
- Coupling Facility Trend Report fix for large request number, change Long to Double.

## 5.4.39

- Fix for CF and Workload Activity report parser to support V2R5 version.

## 5.4.38

- Use default translation table while downloading report via FTP.
- Update Online Help.
- Update signing certificate.

## 5.4.37

- Fixed Coupling Facility Trend Report, structure was missing in ReptrdStr and RepIntStr tabs.

## 5.4.36

- Add new fields from XCF Signaling Report Path Usage Statistics.

## 5.4.35

- New Java 8 Runtime.

## 5.4.34

- Fixed RepCFSys tab in Spreadsheet Report failing to update the report heading.
- Fixed unexpected text RepIntAct.

## 5.4.33

- New Java 8 Runtime.

## 5.4.32

- New FFX/QSA fields in CRYPTO report spreadsheet.

## 5.4.31

- New Java 8 Runtime.

## 5.4.30

- New Java 8 Runtime.

## 5.4.29

- Postprocessor Report Layout Changes.

## 5.4.28

- Postprocessor Report Layout Changes.

## 5.4.27

- New Java 8 Runtime.

## 5.4.26

- New Java 8 Runtime.

## 5.4.25

- New Java 8 Runtime.

## 5.4.24

- Postprocessor Report Layout Changes.

## 5.4.23

- Postprocessor Report Layout Changes.

## 5.4.22

- New Crypto Hardware Activity Trend Report Spreadsheet.

## 5.4.21

- New Java 8 Runtime.

## 5.4.20

- Integrity fixes.

## 5.4.19

- Support for changed report layout.
- New Java 8 Runtime.

## 5.4.18

- New Java 8 Runtime.

## 5.4.17

- XCF Trend Report Spreadsheet (Rmfr9xcf.xls):  
  Spreadsheet is now able to handle empty reports.

## 5.4.16

- New Java 8 Runtime.

## 5.4.15

- New Java 8 Runtime.

## 5.4.14

- New Java 8 Runtime.

## 5.4.13

- Support for changed report layout.

## 5.4.12

- New Java 8 Runtime.

## 5.4.11

- Support for changed report layout.

## 5.4.10

- Support for changed report layout.
- New Java 8 Runtime.

## 5.4.9

- Support for changed report layout.

## 5.4.8

- New Java 8 Runtime.
- Enhance Excel compatibility.

## 5.4.7

- Partition Data Report: LPARs may be incorrect converted.
- Support of Asynchronous duplexing CF Lock Structures.

## 5.4.6

- Support for Super PAV Alias Tuning.
- New Java 8 Runtime.

## 5.4.5

- Partition Data Report: IFL LPARs may be incorrect converted.

## 5.4.4

- Support of LPAR Group Absolute Capping.
- CF Report:  
  Large number of requests in the CF to CF section maybe incorrect converted.
- Addressed passive FTP connection problem.

## 5.4.3

- New Java 8 Runtime.

## 5.4.2

- Support of WLM Workload Containers for Mobile Pricing.

## 5.4.1

- New Java 8 Runtime.

## 5.4.0

- New Java 8 Runtime.

## 5.3.14

- New Java Runtime.

## 5.3.13

- New Java Runtime.

## 5.3.12

- Support for changed report layout.
- Enhance Excel compatibility.

## 5.3.11

- Support for new RMF Postprocessor Reports in XML format.
- Coupling Facility Trend Report Spreadsheet (Rmfr9cf.xls):
  - RepTrdStr sheet: Incorrect system list may be created (only systems available in the first
    report are honored) (PR220062MAI).
  - RepCFTrd sheet: Add metric: Effective CPs.
- XCF Trend Report Spreadsheet (Rmfr9xcf.xls):  
  PathOverview Sheet may report incorrect target system name.
- Channel Overview Report Spreadsheet (Rmfx9chn.xls):  
  Channel Utilization on LPAR level may incorrect reported.
- Coupling Facility Activity Report may be incorrect converted.
- XCF Activity Report may be incorrect converted.
- Workload Activity Report may be incorrect converted.

## 5.3.10

- New Certificate.

## 5.3.9

- Support for new RMF Postprocessor Reports in XML format.
- Coupling Facility Trend Report: RepTrdStr chart shows incorrect numbers data for changed requests.

## 5.3.8

- System Overview Spreadsheet (Rmfy9ovw.xls): Data analysis section on Summary sheet shows
  incorrect page faults (PR210056MAI).
- LPAR Trend Report Spreadsheet (Rmfr9lp.xls):  
  % of CEC Guaranteed /Max is incorrect calculated for none-CP Partitions (PR210092MAI).
- Workload Activity Trend Report Spreadsheet: Application Utilization for Report Classes may be
  incorrect reported (PR210109MAI).
- Creating corrupted report index file when registering large report working set (PC/T0321MAI).

## 5.3.7

- Support for new RMF Postprocessor Reports in XML format.
- Converted Overview Record is limited to 8200 rows (PR130182MAI).

## 5.3.6

- New Java runtime.

## 5.3.5

- New Certificate.

## 5.3.4

Spreadsheet Reporter:.
- MigrateWSetsToXLS.bat cause  java.lang.ArrayIndexOutOfBoundsException (at
  com.ibm.erb.rmfsr.wset.ErbXLSFileInstance._extractSST (ErbXLSFileInstance.java:646) PR110135MAI.
- Processing Working Sets cause error messages 'Cannot process data...' when system regional
  settings are switched to Turkish (may also apply to other regions) PR110136MAI.
- Coupling Facility Trend Report:Processing data may cause error message
  "Cannot process data..." (PR120042MAI).
- Coupling Facility Trend Report:Modifing the report option on the RepTrdStr sheet may cause
  Visual Basic run-time-error 1004 message (PC/T0318MAI).
- No Response Time Distribution may be displayed on the RepRsp worksheet (PR120058MAI).
- Workload Activity Spreadsheet (Rmfy9wkl.xls): The WeekChart may display incorrect date values (PR120079MAI).
- Incorrect conversion of z/OS Version number of the Workload Activity Spreadsheet (PR120088MAI).

## 5.3.3

Spreadsheet Reporter:
- XML Report Support update.
- Workload Activity Report: Large CPU Service Times may be incorrect converted (PR110118MAI).

## 5.3.2

Spreadsheet Reporter:
- XML Report Support.

## 5.3.1

Spreadsheet Reporter:
- WLGML Report may be incorrect converted (PR100088MAI).
- Rangenames with nummeric sign "#" cannot be handled (PR100116MAI).
- Numbers of Processors in the Partition Data Report may be incorrect converted (PR100135MAI).
- Using CreateRptWset.bat file may cause error message "Internal index files are not accessible."
  (PR100136MAI).
- Paging Activity Report may be incorrect converted (PR100142MAI).
- Invalid XLS file generated based on overview record (PR100143MAI).
- Corrupt RMFOV.XLS file cause SR hangup without any error messages (PR100144MAI).
- The Spreadsheet Reporter cannot delete or rename working sets (PR100149MAI).
- Processing data may cause error message "The Report Working Set may not contain the required..."
  (PR100114MAI).
- Workload Activity Report: QMPL delays of may be incorrect converted (PR110076MAI).

Spreadsheet Macros:
- XCF Trend Report: The System name may appear mutliple times in the system list (PR100110MAI).
- Channel Overview Report: When adding data to the macro, the workload busy percent is incorrect
  calculated (PR100115MAI).
- LParsTrd sheet may report incorrect Physical Total Utilization for the Physical Partition (PR100146MAI).
- RMF System Overview Report Macro: Adding data may cause error message "Run-time error '9" (PR100148MAI).
