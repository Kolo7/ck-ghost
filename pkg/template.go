package pkg

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

var (
	tmpl = `CREATE MATERIALIZED VIEW metrics.{{.Appid}}_{{.Table}}_count_mv on cluster {{.Cluster}}
TO metrics.prometheus_metrics_local
AS SELECT 
'clickhouse_log_appid_file_count' AS name,
['__name__=clickhouse_log_appid_file_count', 'instance={{.Instance}}', 'job=clickhouse', 'appid={{.Appid}}', concat('file=', _file), concat('table=', '{{.Table}}'),'idc={{.Idc}}'] AS tags, 
toFloat64(count()) AS val,
toStartOfMinute(_ts) AS ts,
toDate(ts) AS date, 
now() AS updated 
FROM {{.Appid}}.{{.Table}}
GROUP BY _file, ts`
)

type Option struct {
	Cluster  string
	Idc      string
	Instance string
	Appid    string
	Table    string
}

// Build return a clickhouse create sql, by option
func Build(opt Option) (string, error) {

	tmp := template.New("sql")
	tmp = template.Must(tmp.Parse(tmpl))

	buf := new(bytes.Buffer)
	err := tmp.Execute(buf, opt)
	if err != nil {
		return "", err
	}
	bufBt, err := ioutil.ReadAll(buf)
	if err != nil {
		return "", err
	}

	return string(bufBt), nil
}
