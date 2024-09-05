# Change Log

## 2.6.6 (2024-09-06)
- The Performance Monitoring has been updated to include the latest Java version 17.

## 2.6.5 (2024-06-05)
- PM doesn't gather any data fix

## 2.6.4 (2024-03-04)

- Upgrade to Java 17.
- URL encoding fixes.

## 2.6.3 (2023-12-12)

- New Java version.

## 2.6.2 (2023-10-03)

- Open metric help url by default browser.
- Adjust poll time to get interval data once.

## 2.6.1 (2023-07-14)

- New Java version.

## 2.6.0 (2023-03-02)

- Upgrade to Java 11.

## 2.5.16 (2023-02-01)

- New Java version.

## 2.5.15 (2022-10-12)

- Fix metric list loading.
- General maintenance.

## 2.5.14 (2022-04-11)

- New Java version.

## 2.5.13 (2022-02-23)

- Host name validation fixed.
- General maintenance.
- Sort alphabetically menu item to sort tabs and tree.

## 2.5.12 (2022-01-10)

- New Java version.

## 2.5.11 (2021-04-26)

- Fix server certificate check based on truststore.
- Update installer signing certificate.

## 2.5.10 (2021-02-19)

- New Java version.

## 2.5.09 (2021-01-08)

- New Java version.

## 2.5.08 (2020-09-18)

- Fix error at populating perfdesks list after perfdesk save.
- Alphabetical sort of perfdesks in tree.

## 2.5.07 (2020-09-16)

- Integrity fixes, new Java version.

## 2.5.06 (2020-07-31)

- Integrity fixes, new Java version.

## 2.5.05 (2020-03-24)

- Integrity fixes, new Java version.

## 2.5.04 (2019-12-06)

- Integrity fixes, new Java version.

## 2.5.03 (2019-12-02)

- Various bug fixes.

## 2.5.02 (2019-10-18)

- Integrity fixes, new Java version.

## 2.5.01 (2019-09-24)

- Integrity fixes.

## 2.5.00 (2019-06-26)

Version for z/OS V2R4:
- New Monitor III CRYOVW report and metrics.
- New WLM resources and metrics.
- HTTPS support.

## 2.4.91 (10/22/2018)

- Integrity fixes, new Java version.

## 2.4.90 (2018-06-22)

- Corrected details description for SSID.

## 2.4.89 (2018-04-03)

- Integrity fixes, new Java version.

## 2.4.88 (2018-01-22)

- Integrity fixes, new Java version.

## 2.4.87 (2017-06-02)

- Integrity fixes.

## 2.4.86 (2017-03-14)

Version for z/OS V2R3:
- New metrics for various resources.
- Enable resource attributes for PCIE_FUNCTION resources.
- Add attribute "structure encrypted?" to CF_STRUCTURE.

## 2.4.85 (2017-03-07)

- New IBM Java 8 JRE.
- Installable image signed with new certificate.

## 2.4.84 (2016-11-17)

- Integrating 2GB support (OA48913).

## 2.4.83 (2016-04-07)

- Installable image signed with new certificate.

## 2.4.82 (2016-01-25)

- Enable resource attributes introduced with Absolute hardware group capping and OPT parameter ABSMSUCAPPING.
- Fixed error when asking for metric helps.

## 2.4.81 (2015-07-21)

- Bundled with IBM Java 8 JRE.

## 2.4.73 (2015-04-15)

- Fix error when changing the perfdesk color.

## 2.4.72 (2015-03-17)

- Version for z/OS V2R2.
- New resource for PCIE and SCM.
- New metrics for job USAGE and zFS Sysplex monitoring.

## 2.4.71 (2014-10-22)

- Include metrics for MT mode for z/OS V1R13 and z/OS V2R1.

## 2.4.70 (2013-06-10)

- Fix problem with CPC metric "# processors with medium share (IIP) by partition".

## 2.4.69 (2013-03-15)

- Version for z/OS V2R1.

## 2.4.68 (2011-07-05)

- Version for z/OS V1R13.

## 2.4.67 (2011-02-14)

- Bundled with IBM Java 6 JRE.

## 2.4.66 (2010-10-08)

- Removed limitation when exporting data.
- Fixed plot dialog problems when displaying large values.
- The message browser does no longer pop up to the foreground when a new message arrives.

## 2.4.65 (2010-07-09)

- Version for z/OS V1R12.

## 2.4.62 (2010-01-21)

- Omit '[asid]' when starting Analysis Perfdesk for job workscopes (PMR 19130,49R,000).

## 2.4.61 (2009-07-21)

- Version for z/OS V1R11.

## 2.4.60 (2009-06-22)

- New option in View menu to select a decimal separator used for numbers in CSV export.
- Fixed Plot dialog (selection of list-valued series).

## 2.4.59 (2009-05-27)

- Fixed error where wrapping of dataview sample buffers caused chart values to be overwritten (PMR92206).

## 2.4.58 (2009-04-29)

- 'Saving Plot Data': Output format changed to CSV (from WK1) in the CSV files, a ';' is used as a
  column separator and numbers are represented using english/US format, i.e. a '.' is used as the
  decimal symbol.

## 2.4.57 (2009-02-10)

- Version for z/OS V1R10 (PM 2.4.51).
- Bundled with IBM Java 5 JRE (PM 2.4.50).
- IPv6 support.
- Enable resource names with blanks (e.g. enclaves).
- Fixed Firefox invocation (error "Windows cannot find -requestpending").
- Fix 'remove values' from dataview.
- Re-enable DDS level 2116 XML for Linux image.
- Re-enable z/OS 1.2 XML.
- XML parsing: make 'historical' flag dependent on mintime.
- Fixed ArrayIndexOutOfBounds Exception when opening a PerfDesk.

## 2.4.49 (2008-01-29)

- Initial System z10 and HiperDispatch support.

## 2.4.48 (2007-07-16)

- First version for z/OS V1R9.
- Introduced filters for XCF listtypes.
- If a filter string is chosen which contains special characters '(', ')' or '*', these
  characters are masked with '\'. In case you need the '*' to respresent a string with
  any character, please delete the '\'.
- Changed dataview plot/save functionality so that names containing '(' and ')' are allowed.
- Changed legend in dataviews: Time stamp shown now is the starting time of the intervall.
- Changed computation of space reserved for numbers to the left of the y-axis in horizontal
  charts and plots; space was undercharged in some cases.

## 2.4.44 (2007-02-01)

- Adaption to Daylight Saving Time Changes in 2007.  
  In March 2007, the effective dates of Daylight Savings Time (DST) in the U.S. and
  Canada will change. Dates and times could be affected by this change. If not addressed,
  time and date calculations could be incorrect. The RMF PM is affected because of it's
  dependency on the Java JRE. The JRE of RMF PM is replaced with a DST compatible JRE.
- Deactivated Single Logon.
- Updated resources/simulated.xml for z/OS 1.8.
- Fixed problem for historical sampling in DataView Properties menu and Set Sample Time
  dialog ("Sample..." button of PerfDesk), if timezone of the system where RMF PM is running
  is different from the timezone of the (DDS) host.
- DataView Properties Menu always has a Sampling Panel, also for Dataviews not yet started.
- Deleted unused error messages.
- For z/OS metrics, the help file from the DDS is used.
- Fixed PC/T0300MAI: fixed historical sampling problem, that cause RMF PM error message
  "java.text.ParseException:Unparseable date ..".
- Fixed incorrect invocation of Firefox and Netscape browser.

## 2.4.35 (2006-05-29)

- Fixed error that wrong values were returned for time series data when drawing charts in a dataview.

## 2.4.34 (2006-04-19)

- Brackets around ASID in job names changed from () to [] so that 'Plot/Save Series' function will
  work correctly.

## 2.4.33 (2006-02-01)

- Listtypes (i.e. metrics "by ...") which were newly introduced with z/OS 1.7
  are available in the filter list of the 'New Dataview' dialog.

## 2.4.32 (2005-11-10)

- Fixed PC/T0276MAI: RMF PM doesn't support password in mixed case.
- Fixed PC/T0272MAI: RMF PM working incorrectly after logon to DDS on RMF V1R2 and RMF V1R5 level 2123.

## 2.4.30 (2005-07-05)

- Attributes for dynamic resources are shown via popup menu in charts.
- If pm.ini is corrupted or missing, extracting it from gpmres.jar is attempted.
- If registry keys do not exist, default values are set and resources are unpacked.
- Fixed problem in Linux version that non-standard port (i.e. non-8803) was not accepted.
- For unknown listtypes, an empty filter list (i.e. only shwing --N/A-- ) is generated.

## 2.4.25 (2005-04-14)

- Communication with DDS to obtain performance data now done solely with HTTP (XML) protocol.
- Added resource zFS.
- Make jobnames unique by appending the ASID to the jobname.

## 2.4.20 (2004-12-08)

- Introduce export/import functionality for PerfDesks (for more information please.
  Refer to the online help under menu 'Help' --> 'General Help').
- Fixed the following problems:
  - Problem with data paths containing non-ASCII characters (PC/T0240MAI).
  - Legend texts of samples (i.e. resource name) in Linux PerfDesks are now displayed
    in a better readable format (PC/T0245MAI).
  - Installation failed with InstallShield Error Message: Error Code: -5006  (PC/T0237MAI).
  - After starting Application RMF PM, the workstation freezes (PC/T0244MAI).

## 2.4.16 (2004-01-28)

- Fixed problem migrating PerfDesks from release before 2.3.70 (PC/T0219MAI).
- Fixed problem with Linux konqueror.
- Increase size for metric values from 4 to 6 characters.
- Added metric help for Linux metrics.
- Don't display "My Penguin" for saved Linux PerfDesks (PCI/0218MAI).
- Adjusted XML history gathering date (given time stamp is end, not start of interval).
- Removed problem with Linux-only RMF PM clients (no sysplex definition).

## 2.4.8 (2003-11-06)

- Removed restriction on available time zones.
- Improved accessibility.
- More Java WebStart-related enhancements and fixes.
- Fixed bug for case where DDS closed unexpectedly.

## 2.3.70 (2003-08-05)

- Under Windows, determine automatically the installed default browser
  (Mozilla, Netscape, Internet Explorer, etc.) instead of having the user specify which browser to use.
- Abort Linux Overview PerfDesk startup if "My Penguin" Linux image does not respond.
- Various minor fixes.

## 2.3.60 (2003-06-22)

- Settings are now kept on a per-user basis in c:\Documents and Settings\<user>\Application Data\RMF\...
- Compatible with JAVA WebStart version of RMF PM two JAR files only.

## 2.3.53 (2002-12-10)

- New feature: "SaveAllPerfDesks" menu item.
- New feature: change PerfDesk background color, in order to visually distinguish PerfDesks easily.
- Replaced GPM0405I message (DDS data receive error) by three different, more precise error messages.
- Added timeout error message.
- Support for IBM JDK Java Version 1.4.
- Renamed "Storage" into "Memory (Processor Storage)".
- Display Linux image name instead of resource name in order to make it easier to
  identify which Linux image is actually monitored.
- Added Sysplex ID to every error message, as GPM0502I does not contain this vital information.
- No longer expects space characters before certain XML statements.
- Updated internet link to current DB2 PM web page.
- Added error message if user forgot to specify a resource name in the "I/O resource dialog".
- Removed error message (in log file) if given value in list-valued counter not found,
  as this can be normal behavior.
- Fixed that DataView stopped after change of DataView properties and was not restartable.
- Fixed restart problem for multiple series from different Linux hosts where one Linux image re-IPLed.
- PMR 48660,122,000: re-written exception handling in order to prevent JVM crash.
- Fixed timezone problem.
- Fixed minor HTTP logon problem (only occured if DDS accepted any password, but requested one).

## 2.3.33 (2002-06-12)

- Solved Linux Plot ArrayIndexOutOfBounds problem.
- Added BIG_SCREEN configuration parameter in order to have PerfDesk buttons in one row.
- Fixed problem Analyis could not be opened every time (happened for recursive Analysis scenarios only).
- Made application launchable from other Java applications.

## 2.3.29 (2002-01-31)

- If RMF PM 2.2 PerfDesks were converted to version 2.3, RMF PM erroneously displayed
  "-filtered" for some single filters; fixed.
- Forward/ backward button (under PerfDesks) sometimes displayed additional erroneous
  samples in DataView (with empty values); fixed.
- Sometimes, needed to press "Start" after "Sampling..." multiple times in order to
  re-start all DataViews of PerfDesk; fixed.
- Problem in DataView sampling DataView for selection of To/From times; removed.
- Log, if an unknown resource list was requested from the DDS server.

## 2.3.25 (2002-01-09)

- Fixed bug that Analysis was not available for newly created series (introduced with RMF PM Version 2.3).

## 2.3.24 (2001-12-20)

- Fixed problem empty sample properties dialog may appear for closed sampler.

## 2.3.15-2.3.23 (2001-12-13)

- Changed PerfDesk controls, so you can use RMF PM with a screen resolution of 800x600.
- Changed some filter types changed in APAR fix OW51646.
- Various bug fixes.

## 2.3.13-2.3.14 (2001-10-23)

- Various bug fixes.

## 2.3.12 (2001-10-18)

- Added CPC properties.
- Changed default filter to 20 elements.

## 2.3.10-2.3.11 (2001-10-17)

- Various bug fixes.

## 2.3.9 (2001-09-27)

- Fixed W2K-related bug: open/reopen crashed DDS.

## 2.3.6-2.3.8 (2001-09-26)

- Various bug fixes.

## 2.3.5 (2001-09-21)

- Added "<" and ">" buttons for historical data collection.

## 2.3.4 (2001-09-19)

- Added HTTP Basic Authentication.

## 2.3.3 (2001-09-14)

- Sample PerfDesk for Linux.
- Prevent auto-start of z/OS Overview PerfDesk if no z/OS host was defined.

## 2.3.x (2001-09-06)

Includes various fixes and SPE OW49807.
- Caught all exceptions in runnable threads (SwingUtilities).
- Autostart PerfDesks fixed.
- Single logon bug removed.
- Fixed SSpSection initializitaion bug.
- Removed "-work scope" appendix from DV series description.
- Fixed PDControl bug (new historical data collection).
- Bug in changing workScopes in copied DataView (interfered with parent DataView).
- Enhanced stability for startup perfDesks.
- Fixed PR050063MAI, PR050076MAI.
- New ResPropBrowser (XML-based).
- XML port now needed for z/OS performance monitoring minimal DDS functionality level 1670.
- Added trace command line argument.
- Made Linux configuration dynamic.
- Performance of New/Change Sysplex dialog improved.
- Single Logon.
- IBM License Manager Online support added.
- Some more FVT bug fixes.
- Added PerfDesk "Sample" button for historical data coll.
- XML protocol: changed start range into range (HTTP request).
- SPE OW49807 (License Manager Online support).

## 2.2.16 (2001-05-15)

- New z/OS V1R2 product ID.

## 2.2.15 (2001-05-07)

- Needless XML Open removed.

## 2.2.14 (2001-04-11)

- Added "Add DataView" to LINUX_SYSTEM pulldown menu.

## 2.2.13 (2001-04-09)

- Support for multi-headed W2K installations (should also work for Linux X11 systems).

## 2.2.12 (2001-04-06)

- Bug in historical data collection removed.

## 2.2.11 (2001-04-02)

- Frozen DataView bug (setSuspended) removed.

## 2.2.10 (2001-04-02)

- Linux: propagate gatherer interval seconds from client to server.

## 2.2.9 (2001-03-20)

- New RMF Mail address.
- Java 1.3.
- Xerces XML Parser.
- New tutorial in PDF format.

## 2.2.8 (2001-01-22)

- Added "Open" item to PerfDesk folder context menu.
- Explicitly forbidden duplicate PerfDesk names (can cause problems in saving them).
- Avoid the possibility to open the _same_ PerfDesk from the file menu several times.
- In list-valued counters for channels, PM displayed "?".
- If channel name was identical to some numerical value in list before.
- Data display incorrect if "RMF" and "RMFGAT" in listValued counter.
- "save changes..." even after changing only the window parameters (position, size).
- If adding several list-valued counter in one DataView, PM incorrectly displayed some "?" instead
  of the correct values; fixed (see ListSeries.getValueForNameAt()).
- Bug removed, if opening DataView of closed Sysplex.

## 2.2.0 (2000-12-24)

- Linux support.

## 2.1.13 (2000-10-17)

- Bug after saving PerfDesk twice.

## 2.1.12 (2000-09-18)

- Rem. Copy&Paste exception thrown after complex user error.

## 2.1.11 (2000-09-11)

- Tile analysis PerfDesk bug removed.

## 2.1.10

- Sysover.po cannot be overwritten.
- Installation problem fixed, (rmf.ini can be corrupted when doing a first time install).
- Automatic "Tile Dataviews" for Analysis Perfdesks: problem fixed (Null Pointer).

## 2.1.9

- Memory leak removed (date wrap-around buffer).

## 2.1.8

- Some Java multithreading enhancements.

## 2.1.7

- Fix problem saving zoomed DataView (wrong interval saved).

## 2.1.6

- Fix problems with Analysis PerfDesks: message	"GPM0626I The metric "xxx" is not defined ... ".
