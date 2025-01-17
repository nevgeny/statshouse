// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package api

import (
	"strconv"
	"strings"
	"time"

	"github.com/ClickHouse/ch-go/proto"
	"github.com/vkcom/statshouse-go"

	"github.com/vkcom/statshouse/internal/format"
	"github.com/vkcom/statshouse/internal/vkgo/rpc"
	"github.com/vkcom/statshouse/internal/vkgo/srvfunc"
)

const (
	RoutePrefix             = "/api/"
	EndpointMetric          = "metric"
	EndpointMetricList      = "metrics-list"
	EndpointMetricTagValues = "metric-tag-values"
	EndpointQuery           = "query"
	EndpointRender          = "render"
	EndpointResetFlood      = "reset-flood"
	EndpointLegacyRedirect  = "legacy-redirect"
	EndpointDashboard       = "dashboard"
	EndpointDashboardList   = "dashboards-list"
	EndpointGroup           = "group"
	EndpointGroupList       = "group-list"
	EndpointPrometheus      = "prometheus"

	userTokenName = "user"
)

type endpointStat struct {
	endpoint   string
	method     string
	metric     string
	startTime  time.Time
	tokenName  string
	dataFormat string
}

func (es *endpointStat) serviceTime(code int) {
	es.logEvent(format.BuiltinMetricNameAPIEndpointServiceTime, code)
}

func (es *endpointStat) responseTime(code int) {
	es.logEvent(format.BuiltinMetricNameAPIEndpointResponseTime, code)
}

func (es *endpointStat) logEvent(statName string, code int) {
	v := time.Since(es.startTime).Seconds()
	statshouse.AccessMetricRaw(
		statName,
		statshouse.RawTags{
			Tag1: es.endpoint,
			Tag2: es.metric,
			Tag3: strconv.Itoa(code),
			Tag4: es.tokenName,
			Tag5: es.dataFormat,
			Tag6: es.method,
		},
	).Value(v)
}

func (es *endpointStat) setTokenName(user string) {
	es.tokenName = getStatTokenName(user)
}

func getStatTokenName(user string) string {
	if strings.Contains(user, "@") {
		return userTokenName
	}

	return user
}

func newEndpointStat(endpoint, method string, metricID int32, dataFormat string) *endpointStat {
	return &endpointStat{
		endpoint:   endpoint,
		metric:     strconv.Itoa(int(metricID)), // metric ID key is considered "raw"
		startTime:  time.Now(),
		dataFormat: dataFormat,
		method:     method,
	}
}

type rpcMethodStat struct {
	method    string
	startTime time.Time
}

func (ms *rpcMethodStat) serviceTime(ai accessInfo, err error) {
	var errorCode string
	switch e := err.(type) {
	case rpc.Error:
		errorCode = strconv.FormatInt(int64(e.Code), 10)
	case nil:
		errorCode = "0"
	default:
		errorCode = "-1"
	}
	v := time.Since(ms.startTime).Seconds()
	statshouse.AccessMetricRaw(
		format.BuiltinMetricNameAPIRPCServiceTime,
		statshouse.RawTags{
			Tag1: ms.method,
			Tag2: errorCode,
			Tag3: getStatTokenName(ai.user),
			Tag4: srvfunc.HostnameForStatshouse(),
		},
	).Value(v)
}

func CurrentChunksCount(brs *BigResponseStorage) func(*statshouse.Registry) {
	return func(r *statshouse.Registry) {
		r.AccessMetricRaw(
			format.BuiltinMetricNameAPIBRS,
			statshouse.RawTags{
				Tag1: srvfunc.HostnameForStatshouse(),
			},
		).Value(float64(brs.Count()))
	}
}

func ChSelectMetricDuration(duration time.Duration, metricID int32, table, kind string, isFast, isLight bool, err error) {
	ok := "ok"
	if err != nil {
		ok = "error"
	}
	statshouse.AccessMetricRaw(
		format.BuiltinMetricNameAPISelectDuration,
		statshouse.RawTags{
			Tag1: modeStr(isFast, isLight),
			Tag2: strconv.Itoa(int(metricID)),
			Tag3: table,
			Tag4: kind,
			Tag5: ok,
		},
	).Value(duration.Seconds())
}

func ChSelectProfile(isFast, isLight bool, info proto.Profile, err error) {
	chSelectPushMetric(format.BuiltinMetricNameAPISelectBytes, isFast, isLight, float64(info.Bytes), err)
	chSelectPushMetric(format.BuiltinMetricNameAPISelectRows, isFast, isLight, float64(info.Rows), err)
}

func modeStr(isFast, isLight bool) string {
	mode := "slow"
	if isFast {
		mode = "fast"
	}
	if isLight {
		mode += "light"
	} else {
		mode += "heavy"
	}
	return mode
}

func chSelectPushMetric(metric string, isFast, isLight bool, data float64, err error) {

	m := statshouse.AccessMetricRaw(
		metric,
		statshouse.RawTags{
			Tag1: modeStr(isFast, isLight),
		},
	)
	m.Value(data)
	if err != nil {
		m.StringTop(err.Error())
	}
}

func ChCacheRate(cachedRows, chRows int, metricID int32, table, kind string) {
	statshouse.AccessMetricRaw(
		"ch_video_select_test",
		statshouse.RawTags{
			Tag1: "cache",
			Tag2: strconv.Itoa(int(metricID)),
			Tag3: table,
			Tag4: kind,
		},
	).Value(float64(cachedRows))

	statshouse.AccessMetricRaw(
		"ch_video_select_test",
		statshouse.RawTags{
			Tag1: "clickhouse",
			Tag2: strconv.Itoa(int(metricID)),
			Tag3: table,
			Tag4: kind,
		},
	).Value(float64(chRows))
}
