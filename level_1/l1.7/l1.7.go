package main

import (
	"fmt"
	"strconv"
	"sync"
)

func addMapRecordSync(k int, v string, syncMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	syncMap.Store(k, v)
}

func runWithSyncMap() {
	var syncMap sync.Map
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			addMapRecordSync(i, strconv.Itoa(i)+"_value", &syncMap, &wg)
		}(i)
	}

	wg.Wait()

	syncMap.Range(func(k, v any) bool {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
		return true
	})
}

// При запуске с опцией `-race` (`go run -race level_1/l1.7/l1.7.go`) детектор не обнаруживает data race

func addMapRecordRegular(k int, v string, regularMap map[int]string, wg *sync.WaitGroup) {
	defer wg.Done()
	regularMap[k] = v
}

func runWithRegularMap() {
	regularMap := make(map[int]string)
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go addMapRecordRegular(i, strconv.Itoa(i)+"_value", regularMap, &wg)
	}

	wg.Wait()

	for k, v := range regularMap {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
}

// При запуске с опцией `-race` (`go run -race level_1/l1.7/l1.7.go`) детектор обнаруживает data race:
// ==================
// WARNING: DATA RACE
// Write at 0x00c0000ba0f0 by goroutine 15:
// ...
//
// Previous write at 0x00c0000ba0f0 by goroutine 11:
// ...
//
// ...
// ==================

func main() {
	runWithRegularMap()
	runWithSyncMap()
}
