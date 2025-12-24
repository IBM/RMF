import { DATA_SOURCE_TYPE } from "../../constants";
import metrics from '../../metrics/dds/index.json';

const metricMap = prepareMetricMap();

const DATASOURCE_DEFAULT_UID = "${datasource}";
const DATASOURCE_MIXED = "-- Mixed --";
const emptyDashboard = {
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": {
          "type": "grafana",
          "uid": "-- Grafana --"
        },
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "fiscalYearStartMonth": 0,
  "graphTooltip": 0,
  "links": [],
  "liveNow": false,
  "panels": [],
  "refresh": "",
  "schemaVersion": 39,
  "tags": [],
  "templating": {
    "list": []
  },
  "time": {
    "from": "now-5m",
    "to": "now"
  },
  "timeRangeUpdatedDuringEditOrView": false,
  "timepicker": {},
  "timezone": "",
  "title": "Empty Dashboard",
  "version": 1,
  "weekStart": ""
}

function getEmptyDashboard(): any {
    return JSON.parse(JSON.stringify(emptyDashboard));
}

function prepareMetricMap() : Map<string, {id: string; listtype: string}> {
    const metricMap = new Map<string, {id: string; listtype: string}>();
    metrics.metricList = metrics.metricList || [];
    for (let i = 0; i < metrics.metricList.length; i++) {
        let mle = metrics.metricList[i];
        mle.metric = mle.metric || [];
        for (let j = 0; j < mle.metric.length; j++) {
            let m = mle.metric[j];
            metricMap.set(m.id.toUpperCase(), {id: m.id, listtype: m.listtype || ''});
        }
    }
    return metricMap;
}

interface CounterTypeInfo {
    FILTER?: string;
    WSCOPE_TYPE?: string;
}

interface ProxyInfo {
    TYPE?: string;
    NAME?: string;
    HOST?: string;
    TCP_PORT?: number;
    LOCALE?: string;
    TIMEZONE?: string;
    COMTO?: number;
    XML_PORT?: number;
    USE_HTTPS?: string;
}

interface Series {
    CATEGORY?: string;
    CNTTYPE?: string;
    DESC?: string;
    FULLRESNAME?: string;
    PROXYID?: string;
    COLOR?: number;
    COUNTER_TYPE_INFO?: CounterTypeInfo;
    PROXY_INFO?: ProxyInfo;
}

interface Frame {
    X_ORIGIN?: number;
    Y_ORIGIN?: number;
    WIDTH?: number;
    HEIGHT?: number;
}

interface DataView {
    DESC?: string;
    XAXIS?: string;
    CONT_ONFULL?: string;
    BUFSIZE?: number;
    AUTOSCALE?: string;
    FRAME?: Frame;
    SERIES?: Series[];
}

interface VarIString {
    DESCR?: string;
    NAME?: string;
    VALUE?: string;
}

interface VarSpace {
    VAR_ISTRING?: VarIString[];
}

interface PerfDesk {
    DESC?: string;
    STARTUP?: string;
    COLOR?: number;
    DATAVIEW?: DataView[];
    VAR_SPACE?: VarSpace;
}

interface ParsedData {
    PERFDESK?: PerfDesk;
    [key: string]: any;
}

export function parsePmImportFile(content: string | ArrayBuffer | null): ParsedData {
    if (!content) {
        return {};
    }
    content = content.toString();
    const lines = content.split('\n');
    const result: ParsedData = {};
    const stack: any[] = [result];
    const arrayTypes = new Set(['DATAVIEW', 'SERIES', 'VAR_ISTRING']);

    for (let i = 0; i < lines.length; i++) {
        const lineSrc = lines[i];
        const level = lineSrc.search(/\S|$/)/3
        const line = lineSrc.trim();

        if (line.startsWith('[') && line.endsWith(']')) {
            if (stack.length > level + 1) {
                stack.splice(level + 1);
            }
            const sectionName = line.slice(1, -1);
            const newSection: any = {};
            const parent = stack[stack.length - 1];

            if (arrayTypes.has(sectionName)) {
                if (!parent[sectionName]) {
                    parent[sectionName] = [];
                }
                parent[sectionName].push(newSection);
            } else {
                parent[sectionName] = newSection;
            }

            stack.push(newSection);
        } else if (line === '{') {
            continue;
        } else if (line === '}') {
            stack.pop();
        } else if (line.includes('=')) {
            const equalIndex = line.indexOf('=');
            const key = line.slice(0, equalIndex).trim();
            let value: any = line.slice(equalIndex + 1).trim();

            if (value.startsWith('"') && value.endsWith('"')) {
                value = value.slice(1, -1);
            } else if (value === 'Y' || value === 'N') {
                // Keep as string
            } else if (!isNaN(Number(value))) {
                value = Number(value);
            }

            const current = stack[stack.length - 1];
            current[key] = value;
        }
    }
    if (Object.keys(result).length == 0) {
        throw new Error("RMF PM dashboard import file is invalid");
    }
    return result;
}

export function parsePmDatasources(content: string | ArrayBuffer | null): any[] {
    const parsedData = parsePmImportFile(content);
    const datasources: any[] = [];
    parsedData.PERFDESK = parsedData.PERFDESK || {};
    parsedData.PERFDESK.DATAVIEW = parsedData.PERFDESK.DATAVIEW || [];
    for (const dataView of parsedData.PERFDESK.DATAVIEW) {
        for (const series of dataView.SERIES || []) {
            const proxyInfo = series.PROXY_INFO;
            if (proxyInfo) {
                const dsName = proxyInfo.NAME || DATA_SOURCE_TYPE;
                if (!datasources.find(ds => ds.name === dsName)) {
                    datasources.push({
                        name: dsName,
                        type: DATA_SOURCE_TYPE,
                        host: proxyInfo.HOST || '',
                        port: proxyInfo.XML_PORT || 8803,
                        https: proxyInfo.USE_HTTPS === 'YES' ? true : false,
                        locale: proxyInfo.LOCALE || 'en_US',
                        timezone: proxyInfo.TIMEZONE || 'UTC',
                    });
                }
            }
        }
    }
    return datasources;
}

export function parsePmImportFileToDashboard(content: string | ArrayBuffer | null, datasourcesNameUid: Map<string, string>): any {
    const parsedData = parsePmImportFile(content);
    parsedData.PERFDESK = parsedData.PERFDESK || {};
    parsedData.PERFDESK.DESC = parsedData.PERFDESK.DESC || "Default Dashboard Description";

    var needDatasourceVar = false;
    var dashboard = getEmptyDashboard()

    dashboard.title = parsedData.PERFDESK.DESC;
    dashboard.panels = [];
    var deskSize = getDeskSize(parsedData.PERFDESK)
    for (const dataView of parsedData.PERFDESK.DATAVIEW || []) {
        var datasource = getDataSourceFromDataView(dataView, datasourcesNameUid);
        if (datasource.uid === DATASOURCE_DEFAULT_UID) {
            needDatasourceVar = true;
        }
        dashboard.panels.push({
            type: 'barchart',
            title: dataView.DESC || 'Data View',
            datasource: datasource,
            options: {
                orientation: dataView.XAXIS === 'H' ? 'vertical' : 'horizontal',
                xField: getXFieldName(dataView),
                //xTickLabelRotation: 90,
                xTickLabelSpacing: 70,
            },
            transformations: getTransformationsFromDataView(dataView),
            gridPos: getGridPosition(deskSize, dataView),
            targets: dataView.SERIES?.map((series, index) => ({
                selectedVisualisationType: getVisualizationType(series),
                refId: String.fromCharCode(65 + index),
                datasource: getDataSourceFromProxyInfo(series.PROXY_INFO, datasourcesNameUid),
                selectedQuery: getSelectedQueryFromSeries(series),
                selectedResource: {
                    label: getSelectedResourceFromSeries(series),
                    value: getSelectedResourceFromSeries(series)
                },
            })) || []
        });
    }
    if (needDatasourceVar) {
        addDatasourceVar(dashboard);
    }
    
    return dashboard;
}

function addDatasourceVar(dashboard: any) {
    dashboard.templating.list.push({
        type: "datasource",
        name: "datasource",
        query: DATA_SOURCE_TYPE,
        label: "Data source",
        includeAll: false,
        multi: false,
        refresh: 1,
        allowCustomValue: false,
    })
}

function getDataSourceFromDataView(dataView: DataView | undefined, datasourcesNameUid: Map<string, string>): {type: string; uid: string} {
    var uniqueDatasources: Map<string, string>  = new Map<string, string>();
    if (dataView?.SERIES) {
        for (const series of dataView.SERIES) {
            if (series.PROXY_INFO && series.PROXY_INFO.NAME) {
                uniqueDatasources.set(series.PROXY_INFO.NAME, '');
            }
        }
    }
    if (uniqueDatasources.size > 1) {
        return {
            type: 'datasource',
            uid: DATASOURCE_MIXED
        };
    }
    if (dataView?.SERIES && dataView.SERIES.length > 0) {
        return getDataSourceFromProxyInfo(dataView.SERIES[0].PROXY_INFO, datasourcesNameUid);
    }
    return {
        type: DATA_SOURCE_TYPE,
        uid: DATASOURCE_DEFAULT_UID
    };
}

function getDataSourceFromProxyInfo(proxyInfo: ProxyInfo | undefined, datasourcesNameUid: Map<string, string>): {type: string; uid: string} {
    var dsName = DATASOURCE_DEFAULT_UID;
    if (proxyInfo && proxyInfo.NAME) {
        dsName = datasourcesNameUid.get(proxyInfo.NAME) || dsName;
    }
    return {
        type: DATA_SOURCE_TYPE,
        uid: dsName
    };
}

function getXFieldName(dataView: DataView): string {
    if (dataView?.SERIES && dataView.SERIES.length > 0) {
        return getXFieldNameFromSeries(dataView.SERIES[0]);
    }
    return 'time';
}

function getXFieldNameFromSeries(series: Series): string {
    var metricid = getMetricIdFromSeries(series);
    if (!singleMetric(metricid)) {
        return 'partition';
    }
    return 'time';
}

function singleMetric(metricid: string): boolean {
    var metricDef = findmetricById(metricid);
    if (metricDef && metricDef.listtype && metricDef.listtype.trim() === '') {
        return true;
    }
    return false;
}

function getVisualizationType(series: Series): string {
    var metricid = getMetricIdFromSeries(series);
    if (singleMetric(metricid)) {
        return 'TimeSeries';
    }
    return 'bargauge';
}

function findmetricById(metricid: string): {id: string; listtype: string} | null {
    return metricMap.get(metricid) || null;
}

function getSelectedQueryFromSeries(series: Series): string {
    var resource = parseSeriesResname(series.FULLRESNAME || '');
    var ulq = resource.ulq;
    var type = resource.type;
    var name = resource.name;
    var description = parseSeriesDescription(series.DESC || '');
    var query = `${type}.${description.metricDescription} `;
    
    if (ulq || name ||
        (series.COUNTER_TYPE_INFO && (series.COUNTER_TYPE_INFO.FILTER || series.COUNTER_TYPE_INFO.WSCOPE_TYPE))
    ) {
        var params: string[] = [];
        if (ulq) {
           params.push(`ulq=${ulq}`);
        }
        if (name) {
            params.push(`name=${name}`);
        }
        if (series.COUNTER_TYPE_INFO && series.COUNTER_TYPE_INFO.FILTER) {
            params.push(`filter=${prepareFilterValue(series.COUNTER_TYPE_INFO.FILTER, false)}`);
        }
        if (series.COUNTER_TYPE_INFO && series.COUNTER_TYPE_INFO.WSCOPE_TYPE) {
            params.push(`workscope=,,${series.COUNTER_TYPE_INFO.WSCOPE_TYPE}`);
        }
        query += "{" + params.join(",") + "}";
    }
    return query;
}

function getMetricIdFromSeries(series: Series): string {
    var query = series.CNTTYPE || '';
    query = query.replace("RMF#", "8D");
    return query.toUpperCase();
}

function getSelectedResourceFromSeries(series: Series): string {
    var query = "id=" + getMetricIdFromSeries(series);
    query += "&resource=" + (series.FULLRESNAME || '');
    if (series.COUNTER_TYPE_INFO) {
        if (series.COUNTER_TYPE_INFO.FILTER) {
            query += `&filter=${prepareFilterValue(series.COUNTER_TYPE_INFO.FILTER, true)}`;
        }
        if (series.COUNTER_TYPE_INFO.WSCOPE_TYPE) {
            query += `&workscope=,,${series.COUNTER_TYPE_INFO.WSCOPE_TYPE}`;
        }
    }
    return query;
}

function prepareFilterValue(value: string, encode: boolean): string {
    let valueTrimmed = value.trim();
    let valueParts = valueTrimmed.split(/\s*;\s*/);
    let valueProcessedParts = valueParts.map(part => part.trim()).filter(part => part.length > 0);
    let valueProcessed;
    if (encode) {
        valueProcessed = valueProcessedParts.join('%3B');
    } else {
        valueProcessed = valueProcessedParts.join(';');
    }
    let valueFinal = valueProcessed;
    return valueFinal;
}

function parseSeriesResname(fullresname: string): { ulq: string; name: string; type: string } {
    const segments = fullresname.split(',').map(s => s.trim());
    
    return {
        ulq: segments[0] || '',
        name: segments[1] || '',
        type: segments[2] || '',
    };
}

function parseSeriesDescription(description: string): { ulq: string; name: string; type: string; metricDescription: string } {
    const parts = description.split(' - ');
    const metricDescription = parts.length > 1 ? parts[1].trim() : '';
    
    const leftPart = parts[0] || '';
    const segments = leftPart.split(',').map(s => s.trim());
    
    return {
        ulq: segments[0] || '',
        name: segments[1] || '',
        type: segments[2] || '',
        metricDescription: metricDescription
    };
}

function getTransformationsFromDataView(dataView: DataView): any[] {
    const transformations: any[] = []; 
    if (dataView.XAXIS) {
        if (dataView.XAXIS === 'H') {
            
        } else if (dataView.XAXIS === 'V') {
            transformations.push({
                id: "joinByField",
                options: {
                    byField: getGroupByField(dataView),
                    mode: "outer"
                }
            });
        }
    }
    return transformations;
}

function getGroupByField(dataView: DataView): string {
    if (dataView?.SERIES && dataView.SERIES.length > 0) {
        var metricid = getMetricIdFromSeries(dataView.SERIES[0]);
        if (!singleMetric(metricid)) {
            return 'partition';
        }
    }
    return 'time';
}

function getDeskSize(perfDesk: PerfDesk): {width: number, height: number} {
    var maxX = 0;
    var maxY = 0;
    for (const dataView of perfDesk.DATAVIEW || []) {
        var frame = dataView.FRAME;
        if (frame) {
            var xEnd = (frame.X_ORIGIN || 0) + (frame.WIDTH || 0);
            var yEnd = (frame.Y_ORIGIN || 0) + (frame.HEIGHT || 0);
            if (xEnd > maxX) {
                maxX = xEnd;
            }
            if (yEnd > maxY) {
                maxY = yEnd;
            }
        }
    }
    return {width: maxX, height: maxY};
}

function getGridPosition(deskSize: {width: number, height: number}, dataView: DataView): any {
    const gridWidth = 24;
    const gridHeight = 24;
    var scaleX = gridWidth / deskSize.width;
    var scaleY = gridHeight / deskSize.height;
    var dwx = dataView.FRAME?.X_ORIGIN ? Math.abs(dataView.FRAME.X_ORIGIN) : 0;
    var dwy = dataView.FRAME?.Y_ORIGIN ? Math.abs(dataView.FRAME.Y_ORIGIN) : 0;
    var x = Math.round(dwx * scaleX);
    var y = Math.round(dwy * scaleY);
    var w = dataView.FRAME?.WIDTH ? Math.floor((dataView.FRAME.WIDTH) * scaleX) : 8;
    var h = dataView.FRAME?.HEIGHT ? Math.floor((dataView.FRAME.HEIGHT) * scaleY) : 8;
    return {
        x: x,
        y: y,
        h: h,
        w: w,
      }
}