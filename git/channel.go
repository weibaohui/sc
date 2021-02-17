package git

import (
	"sync"
	"time"

	"go.uber.org/atomic"
)

type CommitChan struct {
	AuthorEmail  chan *Commit
	Done         chan bool
	receiveCount *atomic.Int32
	sumCount     *atomic.Int32
}

var chanOnce sync.Once
var commitChan *CommitChan

func init() {
	chanOnce.Do(func() {

		commitChan = &CommitChan{
			AuthorEmail:  make(chan *Commit, 100),
			Done:         make(chan bool),
			receiveCount: atomic.NewInt32(0),
			sumCount:     atomic.NewInt32(0),
		}
	})
}
func GetChanInstance() *CommitChan {
	return commitChan
}
func (ch *CommitChan) Add(c *Commit) {
	// fmt.Println("channel 来了一个c", c.Author.Email)
	ch.AuthorEmail <- c
	ch.receiveCount.Inc()
}
func (ch *CommitChan) Sum() {
	timer := time.NewTicker(time.Second * 1)
	g := GetInstance()
	for {
		select {
		case c := <-commitChan.AuthorEmail:
			_, exists := g.Summary.authorCountsMap.Load(c.Author.Email)
			if !exists {
				Debugf("统计作者%s\n", c.Author.Email)
				ac := g.repo.SumAuthor(c.Author)
				g.Summary.authorCountsMap.Store(c.Author.Email, ac)
			}
			ch.sumCount.Inc()
		case <-timer.C:
			Debugf("统计收到%d，完成%d\n", ch.receiveCount.Load(), ch.sumCount.Load())
			if ch.sumCount.Load() > 0 && ch.sumCount.Load() == ch.receiveCount.Load() {
				Debugf("全部统计完毕-%d\n", ch.sumCount)
				ch.Done <- true
				return
			}
		}
	}
}
