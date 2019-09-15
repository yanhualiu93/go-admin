package models

type ZipkinSpans struct {
	TraceIdHigh int64  `xorm:"not null pk default 0 comment('If non zero, this means the trace uses 128 bit traceIds instead of 64 bit') unique(trace_id_high) index(trace_id_high_2) index(trace_id_high_3) unique(trace_id_high_4) index(trace_id_high_5) index(trace_id_high_6) BIGINT(20)"`
	TraceId     int64  `xorm:"not null pk unique(trace_id_high) index(trace_id_high_2) index(trace_id_high_3) unique(trace_id_high_4) index(trace_id_high_5) index(trace_id_high_6) BIGINT(20)"`
	Id          int64  `xorm:"pk unique(trace_id_high) index(trace_id_high_2) unique(trace_id_high_4) index(trace_id_high_5) BIGINT(20)"`
	Name        string `xorm:"not null index index VARCHAR(255)"`
	ParentId    int64  `xorm:"default NULL BIGINT(20)"`
	Debug       int    `xorm:"default NULL BIT(1)"`
	StartTs     int64  `xorm:"default NULL comment('Span.timestamp(): epoch micros used for endTs query and to implement TTL') index index BIGINT(20)"`
	Duration    int64  `xorm:"default NULL comment('Span.duration(): micros used for minDuration and maxDuration query') BIGINT(20)"`
}
