package until

import (
	"time"
)

// LocalWindow represents a window that ignores sync behavior entirely
// and only stores counters in memory.
type LocalWindow struct {
	// The start boundary (timestamp in nanoseconds) of the window.
	// [start, start + size)
	start int64

	// The total count of events happened in the window.
	count int64
}

func NewLocalWindow() (*LocalWindow, StopFunc) {
	return &LocalWindow{}, func() {}
}

func (w *LocalWindow) Start() time.Time {
	return time.Unix(0, w.start)
}

func (w *LocalWindow) Count() int64 {
	return w.count
}

func (w *LocalWindow) AddCount(n int64) {
	w.count += n
}

func (w *LocalWindow) Reset(s time.Time, c int64) {
	w.start = s.UnixNano()
	w.count = c
}

func (w *LocalWindow) Sync(now time.Time) {}

type (
	SyncRequest struct {
		Key     string
		Start   int64
		Count   int64
		Changes int64
	}

	SyncResponse struct {
		// Whether the synchronization succeeds.
		OK    bool
		Start int64
		// The changes accumulated by the local limiter.
		Changes int64
		// The total changes accumulated by all the other limiters.
		OtherChanges int64
	}

	MakeFunc   func() SyncRequest
	HandleFunc func(SyncResponse)
)

type Synchronizer interface {
	// 开始启动同步goroutine
	Start()

	// 停止停止同步goroutine
	Stop()

	// 同步发送一个同步请求
	Sync(time.Time, MakeFunc, HandleFunc)
}

// 同步窗口代表一个窗口,将同步异步计数器数据中央数据存储。
// 注意,最好的窗口之间的协调和同步,同步不是自动的,而是由调用同步。
type SyncWindow struct {
	LocalWindow
	changes int64

	key    string
	syncer Synchronizer
}

// 新的同步窗口
func NewSyncWindow(key string, syncer Synchronizer) (*SyncWindow, StopFunc) {
	w := &SyncWindow{
		key:    key,
		syncer: syncer,
	}

	w.syncer.Start()
	return w, w.syncer.Stop
}

func (w *SyncWindow) AddCount(n int64) {
	w.changes += n
	w.LocalWindow.AddCount(n)
}

func (w *SyncWindow) Reset(s time.Time, c int64) {
	// 明显的改变累积在旧的窗口。
	//剩余请注意,为了简单起见,我们不同步变化,中央数据存储在重置,
	//从而让周期同步充分的窗口的计数的准确性。
	w.changes = 0

	w.LocalWindow.Reset(s, c)
}

func (w *SyncWindow) makeSyncRequest() SyncRequest {
	return SyncRequest{
		Key:     w.key,
		Start:   w.LocalWindow.start,
		Count:   w.LocalWindow.count,
		Changes: w.changes,
	}
}

func (w *SyncWindow) handleSyncResponse(resp SyncResponse) {
	if resp.OK && resp.Start == w.LocalWindow.start {
		// 更新的状态窗口,只有当它没有被重置在最新的同步。考虑更改积累的其他限制。
		w.LocalWindow.count += resp.OtherChanges

		// 减去的数量已经从现有的同步变化
		w.changes -= resp.Changes
	}
}

func (w *SyncWindow) Sync(now time.Time) {
	w.syncer.Sync(now, w.makeSyncRequest, w.handleSyncResponse)
}
