package mdl

type RequestResult struct {
	Time    uint64
	Succeed bool
	ErrCode int
}

type StaticMetric struct {
	ByNow       uint64
	ReqCostTime uint64
	SucCount    uint64
	FailCount   uint64
	Concurrency uint64
}
