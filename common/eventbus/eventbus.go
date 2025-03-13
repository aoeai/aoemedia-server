package eventbus

import "sync"

type EventBus struct {
	subscribers map[string][]func(interface{}) // 事件类型 -> 处理函数列表
	mu          sync.RWMutex                   // 读写锁保证并发安全
}

var (
	instance *EventBus
	once     sync.Once
)

func Inst() *EventBus {
	once.Do(func() {
		instance = &EventBus{
			subscribers: make(map[string][]func(interface{})),
		}
	})
	return instance
}

func (eb *EventBus) Subscribe(eventType string, handler func(interface{})) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.subscribers[eventType] = append(eb.subscribers[eventType], handler)
}

func (eb *EventBus) Publish(eventType string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	if handlers, exists := eb.subscribers[eventType]; exists {
		for _, handler := range handlers {
			go handler(data) // 异步执行避免阻塞
		}
	}
}
