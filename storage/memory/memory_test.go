package memory

import (
	"sync"
	"testing"

	"github.com/appleboy/gorush/core"

	"github.com/stretchr/testify/assert"
)

func TestMemoryEngine(t *testing.T) {
	var val int64

	memory := New()
	err := memory.Init()
	assert.Nil(t, err)

	memory.Add(core.AndroidSuccessKey, 10)
	val = memory.Get(core.AndroidSuccessKey)
	assert.Equal(t, int64(10), val)
	memory.Add(core.AndroidSuccessKey, 10)
	val = memory.Get(core.AndroidSuccessKey)
	assert.Equal(t, int64(20), val)

	memory.Set(core.AndroidSuccessKey, 0)
	val = memory.Get(core.AndroidSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			memory.Add(core.AndroidSuccessKey, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	val = memory.Get(core.AndroidSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, memory.Close())
}
