package rotatelogs

import (
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"os"
	"time"
)

// 实现一个日志最小堆，用来删除哪些文件

// Element is an entry in the priority queue
type Element struct {

	// 日志文件的名字
	logFileName string

	// 此日志文件最后被修改的时间
	lastModTime time.Time
}

// Comparator function (sort by element's priority value in descending order)
func byPriority(a, b interface{}) int {
	priorityA := a.(*Element).lastModTime
	priorityB := b.(*Element).lastModTime
	return int(priorityA.UnixMilli() - priorityB.UnixMilli())
}

// TODO 2022-4-15 16:20:58 BUG FIX： 此种方式会导致较老的error.log文件被删除
// 选择要删除哪些日志文件
func selectNeedDeleteLogFile(logFileNameSlice []string, n int) []string {

	if n <= 0 {
		return make([]string, 0)
	}

	// 把每个日志文件的最后修改时间拿到，放入到一个堆中
	queue := pq.NewWith(byPriority)
	for _, logFileName := range logFileNameSlice {
		stat, err := os.Stat(logFileName)
		if err != nil {
			continue
		}
		queue.Enqueue(&Element{
			logFileName: logFileName,
			lastModTime: stat.ModTime(),
		})
	}

	// 然后从堆中取出修改时间距离当前最远的N个文件
	needDeleteLogFileName := make([]string, 0)
	for n > 0 && !queue.Empty() {
		dequeue, ok := queue.Dequeue()
		n--
		if !ok || dequeue == nil {
			break
		}
		element, ok := dequeue.(*Element)
		if !ok || element == nil {
			break
		}
		needDeleteLogFileName = append(needDeleteLogFileName, element.logFileName)
	}
	return needDeleteLogFileName
}
