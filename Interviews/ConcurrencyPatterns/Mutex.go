package ConcurrencyPatterns

import (
	"fmt"
	"sync"
	"time"
)

type MutexDataMap struct {
	Data map[string]string
	RwM sync.RWMutex
}

func (cm *MutexDataMap) Upsert(key, val string) {
	cm.RwM.Lock()
	defer cm.RwM.Unlock()

	cm.Data[key] = val
}

func (cm *MutexDataMap) Get(key string) string {

	cm.RwM.RLock()
	defer cm.RwM.RUnlock()

	return cm.Data[key]
}

func (cm *MutexDataMap) Delete(key string) {
	cm.RwM.Lock()
	defer cm.RwM.Unlock()

	delete(cm.Data, key)
}



func Mutex() {

	m := &MutexDataMap{
		Data: make(map[string]string),
	}


	go m.Upsert("India", "New Delhi")
	go m.Upsert("Norway", "Oslo")
	go m.Upsert("Pakistan", "Islamabad")
	go m.Upsert("Bangladesh", "Dhaka")
	go m.Upsert("Nepal", "Kathmandu")


	time.Sleep(2 * time.Second)

	go fmt.Println(m.Get("Nepal"), "--nepal")

	go m.Delete("Nepal")

	go fmt.Println(m.Get("Nepal"), "--nepal")

	go m.Delete("Nepal")
	
	
	
	time.Sleep(3 * time.Second)
	
	go fmt.Println(m.Get("Nepal"), "--nepal")
	time.Sleep(2 * time.Second)

}
