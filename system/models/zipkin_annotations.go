package models

type ZipkinAnnotations struct {
	TraceIdHigh         int64  `xorm:"not null default 0 comment('If non zero, this means the trace uses 128 bit traceIds instead of 64 bit') unique(trace_id_high) index(trace_id_high_2) index(trace_id_high_3) unique(trace_id_high_4) index(trace_id_high_5) index(trace_id_high_6) BIGINT(20)"`
	TraceId             int64  `xorm:"not null comment('coincides with zipkin_spans.trace_id') index(trace_id) index(trace_id_2) unique(trace_id_high) index(trace_id_high_2) index(trace_id_high_3) unique(trace_id_high_4) index(trace_id_high_5) index(trace_id_high_6) BIGINT(20)"`
	SpanId              int64  `xorm:"not null comment('coincides with zipkin_spans.id') index(trace_id) index(trace_id_2) unique(trace_id_high) index(trace_id_high_2) unique(trace_id_high_4) index(trace_id_high_5) BIGINT(20)"`
	AKey                string `xorm:"not null comment('BinaryAnnotation.key or Annotation.value if type == -1') index index index(trace_id) index(trace_id_2) unique(trace_id_high) unique(trace_id_high_4) VARCHAR(255)"`
	AValue              []byte `xorm:"default NULL comment('BinaryAnnotation.value(), which must be smaller than 64KB') BLOB"`
	AType               int    `xorm:"not null comment('BinaryAnnotation.type() or -1 if Annotation') index index INT(11)"`
	ATimestamp          int64  `xorm:"default NULL comment('Used to implement TTL; Annotation.timestamp or zipkin_spans.timestamp') unique(trace_id_high) unique(trace_id_high_4) BIGINT(20)"`
	EndpointIpv4        int    `xorm:"default NULL comment('Null when Binary/Annotation.endpoint is null') INT(11)"`
	EndpointIpv6        []byte `xorm:"default NULL comment('Null when Binary/Annotation.endpoint is null, or no IPv6 address') BINARY(16)"`
	EndpointPort        int    `xorm:"default NULL comment('Null when Binary/Annotation.endpoint is null') SMALLINT(6)"`
	EndpointServiceName string `xorm:"default 'NULL' comment('Null when Binary/Annotation.endpoint is null') index index VARCHAR(255)"`
}
