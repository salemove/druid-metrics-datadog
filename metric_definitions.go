package main

type Definition struct {
	Type         string
	ConvertRange bool
}

const Timer = "timer"
const Gauge = "gauge"
const Count = "count"

func MetricDefinitions() map[string]Definition {
	var m map[string]Definition
	m = make(map[string]Definition)

	m["query/time"] = Definition{Type: Timer}

	m["query/node/time"] = Definition{Type: Timer}
	m["query/node/ttfb"] = Definition{Type: Timer}
	m["query/intervalChunk/time"] = Definition{Type: Timer}

	m["query/segment/time"] = Definition{Type: Timer}
	m["query/wait/time"] = Definition{Type: Timer}
	m["segment/scan/pending"] = Definition{Type: Gauge}
	m["query/segmentAndCache/time"] = Definition{Type: Timer}
	m["query/cpu/time"] = Definition{Type: Timer}

	m["query/cache/delta/numEntries"] = Definition{Type: Count}
	m["query/cache/delta/sizeBytes"] = Definition{Type: Count}
	m["query/cache/delta/hits"] = Definition{Type: Count}
	m["query/cache/delta/misses"] = Definition{Type: Count}
	m["query/cache/delta/evictions"] = Definition{Type: Count}
	m["query/cache/delta/hitRate"] = Definition{Type: Count, ConvertRange: true}
	m["query/cache/delta/averageBytes"] = Definition{Type: Count}
	m["query/cache/delta/timeouts"] = Definition{Type: Count}
	m["query/cache/delta/errors"] = Definition{Type: Count}

	m["query/cache/total/numEntries"] = Definition{Type: Gauge}
	m["query/cache/total/sizeBytes"] = Definition{Type: Gauge}
	m["query/cache/total/hits"] = Definition{Type: Gauge}
	m["query/cache/total/misses"] = Definition{Type: Gauge}
	m["query/cache/total/evictions"] = Definition{Type: Gauge}
	m["query/cache/total/hitRate"] = Definition{Type: Gauge, ConvertRange: true}
	m["query/cache/total/averageBytes"] = Definition{Type: Gauge}
	m["query/cache/total/timeouts"] = Definition{Type: Gauge}
	m["query/cache/total/errors"] = Definition{Type: Gauge}

	m["ingest/events/thrownAway"] = Definition{Type: Count}
	m["ingest/events/unparseable"] = Definition{Type: Count}
	m["ingest/events/processed"] = Definition{Type: Count}
	m["ingest/rows/output"] = Definition{Type: Count}
	m["ingest/persist/count"] = Definition{Type: Count}
	m["ingest/persist/time"] = Definition{Type: Timer}
	m["ingest/persist/cpu"] = Definition{Type: Timer}
	m["ingest/persist/backPressure"] = Definition{Type: Gauge}
	m["ingest/persist/failed"] = Definition{Type: Count}
	m["ingest/handoff/failed"] = Definition{Type: Count}
	m["ingest/merge/time"] = Definition{Type: Timer}
	m["ingest/merge/cpu"] = Definition{Type: Timer}

	m["task/run/time"] = Definition{Type: Timer}
	m["segment/added/bytes"] = Definition{Type: Count}
	m["segment/moved/bytes"] = Definition{Type: Count}
	m["segment/nuked/bytes"] = Definition{Type: Count}

	m["segment/assigned/count"] = Definition{Type: Count}
	m["segment/moved/count"] = Definition{Type: Count}
	m["segment/dropped/count"] = Definition{Type: Count}
	m["segment/deleted/count"] = Definition{Type: Count}
	m["segment/unneeded/count"] = Definition{Type: Count}
	m["segment/cost/raw"] = Definition{Type: Count}
	m["segment/cost/normalization"] = Definition{Type: Count}
	m["segment/cost/normalized"] = Definition{Type: Count}
	m["segment/loadQueue/size"] = Definition{Type: Gauge}
	m["segment/loadQueue/failed"] = Definition{Type: Gauge}
	m["segment/loadQueue/count"] = Definition{Type: Gauge}
	m["segment/dropQueue/count"] = Definition{Type: Gauge}
	m["segment/size"] = Definition{Type: Gauge}
	m["segment/overShadowed/count"] = Definition{Type: Gauge}

	m["segment/max"] = Definition{Type: Gauge}
	m["segment/used"] = Definition{Type: Gauge}
	m["segment/usedPercent"] = Definition{Type: Gauge, ConvertRange: true}

	m["jvm/pool/committed"] = Definition{Type: Gauge}
	m["jvm/pool/init"] = Definition{Type: Gauge}
	m["jvm/pool/max"] = Definition{Type: Gauge}
	m["jvm/pool/used"] = Definition{Type: Gauge}
	m["jvm/bufferpool/count"] = Definition{Type: Gauge}
	m["jvm/bufferpool/used"] = Definition{Type: Gauge}
	m["jvm/bufferpool/capacity"] = Definition{Type: Gauge}
	m["jvm/mem/init"] = Definition{Type: Gauge}
	m["jvm/mem/max"] = Definition{Type: Gauge}
	m["jvm/mem/used"] = Definition{Type: Gauge}
	m["jvm/mem/committed"] = Definition{Type: Gauge}
	m["jvm/gc/count"] = Definition{Type: Count}
	m["jvm/gc/time"] = Definition{Type: Timer}

	m["ingest/events/buffered"] = Definition{Type: Gauge}

	m["sys/swap/free"] = Definition{Type: Gauge}
	m["sys/swap/max"] = Definition{Type: Gauge}
	m["sys/swap/pageIn"] = Definition{Type: Gauge}
	m["sys/swap/pageOut"] = Definition{Type: Gauge}
	m["sys/disk/write/count"] = Definition{Type: Count}
	m["sys/disk/read/count"] = Definition{Type: Count}
	m["sys/disk/write/size"] = Definition{Type: Count}
	m["sys/disk/read/size"] = Definition{Type: Count}
	m["sys/net/write/size"] = Definition{Type: Count}
	m["sys/net/read/size"] = Definition{Type: Count}
	m["sys/fs/used"] = Definition{Type: Gauge}
	m["sys/fs/max"] = Definition{Type: Gauge}
	m["sys/mem/used"] = Definition{Type: Gauge}
	m["sys/mem/max"] = Definition{Type: Gauge}
	m["sys/storage/used"] = Definition{Type: Gauge}
	m["sys/cpu"] = Definition{Type: Gauge}

	m["coordinator-segment/count"] = Definition{Type: Gauge}
	m["historical-segment/count"] = Definition{Type: Gauge}

	return m
}
