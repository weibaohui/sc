package git

import (
	"sync"

	"go.uber.org/atomic"
)

type CommitChan struct {
	AuthorEmail  chan *Commit
	done         chan bool
	receiveCount *atomic.Int32
	processCount *atomic.Int32
}

var chanOnce sync.Once
var commitChan *CommitChan

func init() {
	chanOnce.Do(func() {

		commitChan = &CommitChan{
			AuthorEmail:  make(chan *Commit, 100),
			done:         make(chan bool),
			receiveCount: atomic.NewInt32(-1), // 避免0值，区分是默认值还是累加值，用于极端情况，两个count都是0值
			processCount: atomic.NewInt32(-1),
		}
	})
}
func GetChanInstance() *CommitChan {
	return commitChan
}

// Add
func (ch *CommitChan) Add(c *Commit) {
	ch.AuthorEmail <- c
	ch.receiveCount.Inc()
}

// Process set processCount.Inc()
func (ch *CommitChan) Process(c *Commit) {
	ch.processCount.Inc()
}

// IsDone returns true only if process>0 and process==receive
func (ch *CommitChan) IsDone() bool {
	return ch.processCount.Load() >= 0 && ch.processCount.Load() == ch.receiveCount.Load()
}

func (ch *CommitChan) Complete() {
	ch.done <- true
}
