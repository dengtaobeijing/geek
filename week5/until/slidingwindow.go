package until

import (
	"sync"
	"time"
)

// Window 滑动窗口
type Window interface {
	// Start 开始边界
	Start() time.Time

	// Count 计数器
	Count() int64

	// AddCount 累计技术
	AddCount(n int64)

	// Reset 重置时间
	Reset(s time.Time, c int64)

	// Sync 同步时间保持数据一致
	Sync(now time.Time)
}

// StopFunc 停止
type StopFunc func()

// NewWindow 新的计数器
type NewWindow func() (Window, StopFunc)

type Limiter struct {
	size  time.Duration
	limit int64

	mu sync.Mutex

	curr Window
	prev Window
}

//计数器
func NewLimiter(size time.Duration, limit int64, newWindow NewWindow) (*Limiter, StopFunc) {
	currWin, currStop := newWindow()

	prevWin, _ := NewLocalWindow()

	lim := &Limiter{
		size:  size,
		limit: limit,
		curr:  currWin,
		prev:  prevWin,
	}

	return lim, currStop
}

//时间窗口
func (lim *Limiter) Size() time.Duration {
	return lim.size
}

// 访问限制数量
func (lim *Limiter) Limit() int64 {
	lim.mu.Lock()
	defer lim.mu.Unlock()
	return lim.limit
}

// 设置新的限制
func (lim *Limiter) SetLimit(newLimit int64) {
	lim.mu.Lock()
	defer lim.mu.Unlock()
	lim.limit = newLimit
}

// 运行通过的时间
func (lim *Limiter) Allow() bool {
	return lim.AllowN(time.Now(), 1)
}

// 请求是否放行
func (lim *Limiter) AllowN(now time.Time, n int64) bool {
	lim.mu.Lock()
	defer lim.mu.Unlock()

	lim.advance(now)

	elapsed := now.Sub(lim.curr.Start())
	weight := float64(lim.size-elapsed) / float64(lim.size)
	count := int64(weight*float64(lim.prev.Count())) + lim.curr.Count()

	// Trigger the possible sync behaviour.
	defer lim.curr.Sync(now)

	if count+n > lim.limit {
		return false
	}

	lim.curr.AddCount(n)
	return true
}

// 更新窗口时间
func (lim *Limiter) advance(now time.Time) {
	// Calculate the start boundary of the expected current-window.
	newCurrStart := now.Truncate(lim.size)

	diffSize := newCurrStart.Sub(lim.curr.Start()) / lim.size
	if diffSize >= 1 {
		// The current-window is at least one-window-size behind the expected one.

		newPrevCount := int64(0)
		if diffSize == 1 {
			// The new previous-window will overlap with the old current-window,
			// so it inherits the count.
			//
			// Note that the count here may be not accurate, since it is only a
			// SNAPSHOT of the current-window's count, which in itself tends to
			// be inaccurate due to the asynchronous nature of the sync behaviour.
			newPrevCount = lim.curr.Count()
		}
		lim.prev.Reset(newCurrStart.Add(-lim.size), newPrevCount)

		// The new current-window always has zero count.
		lim.curr.Reset(newCurrStart, 0)
	}
}
