package eventbus

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventBus_Subscribe(t *testing.T) {
	eb := Inst()
	eventType := "test-event"
	handlerCalled := false

	handler := func(data interface{}) {
		handlerCalled = true
	}

	eb.Subscribe(eventType, handler)

	assert.Equal(t, 1, len(eb.subscribers[eventType]), "期望订阅者数量为1")

	// 发布事件并验证处理器是否被调用
	eb.Publish(eventType, "test-data")
	time.Sleep(100 * time.Millisecond) // 等待异步处理器执行

	assert.True(t, handlerCalled, "处理器未被调用")
}

func TestEventBus_Publish(t *testing.T) {
	tests := []struct {
		name           string
		eventType      string
		subscriberNum  int
		data           interface{}
		expectHandled  bool
		waitForHandler bool
	}{
		{
			name:           "单个订阅者正常接收事件",
			eventType:      "test-event-1",
			subscriberNum:  1,
			data:           "test-data",
			expectHandled:  true,
			waitForHandler: true,
		},
		{
			name:           "多个订阅者正常接收事件",
			eventType:      "test-event-2",
			subscriberNum:  3,
			data:           123,
			expectHandled:  true,
			waitForHandler: true,
		},
		{
			name:           "无订阅者的事件发布",
			eventType:      "test-event-3",
			subscriberNum:  0,
			data:           nil,
			expectHandled:  false,
			waitForHandler: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eb := Inst()
			var wg sync.WaitGroup
			handlersCalled := make([]bool, tt.subscriberNum)

			// 注册订阅者
			for i := 0; i < tt.subscriberNum; i++ {
				i := i // 避免闭包问题
				wg.Add(1)
				handler := func(data interface{}) {
					defer wg.Done()
					handlersCalled[i] = true
					assert.Equal(t, tt.data, data, "处理器收到的数据不匹配")
				}
				eb.Subscribe(tt.eventType, handler)
			}

			// 发布事件
			eb.Publish(tt.eventType, tt.data)

			// 等待处理器执行完成
			if tt.waitForHandler {
				wg.Wait()
			} else {
				time.Sleep(100 * time.Millisecond) // 给一些时间让潜在的处理器执行
			}

			// 验证结果
			for i, called := range handlersCalled {
				assert.Equal(t, tt.expectHandled, called, "处理器 %d 的调用状态不符合预期", i)
			}
		})
	}
}

func TestEventBus_ConcurrentOperations(t *testing.T) {
	eb := Inst()
	eventType := "concurrent-test"
	subscriberCount := 10
	publishCount := 5

	// 用于跟踪每个处理器被调用的次数
	handlerCalls := make([]int, subscriberCount)
	var mu sync.Mutex
	var wgHandlers sync.WaitGroup

	// 并发注册订阅者
	var wgSubscribe sync.WaitGroup
	for i := 0; i < subscriberCount; i++ {
		i := i
		wgSubscribe.Add(1)
		go func() {
			defer wgSubscribe.Done()
			handler := func(data interface{}) {
				defer wgHandlers.Done()
				mu.Lock()
				handlerCalls[i]++
				mu.Unlock()
			}
			eb.Subscribe(eventType, handler)
		}()
	}
	wgSubscribe.Wait()

	// 为每次发布的每个处理器添加等待
	wgHandlers.Add(subscriberCount * publishCount)

	// 并发发布事件
	var wgPublish sync.WaitGroup
	wgPublish.Add(publishCount)
	for i := 0; i < publishCount; i++ {
		go func(val int) {
			defer wgPublish.Done()
			eb.Publish(eventType, val)
		}(i)
	}

	// 等待所有发布完成
	wgPublish.Wait()
	// 等待所有处理器执行完成
	wgHandlers.Wait()

	// 验证每个处理器都被调用了正确的次数
	for i, calls := range handlerCalls {
		assert.Equal(t, publishCount, calls, "处理器 %d 的调用次数不正确", i)
	}
}
