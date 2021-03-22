package stats

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateAllocs(t *testing.T) {
	//test referthing
	t.Parallel()

	assert := assert.New(t)

	// update allocation stats and reads it back
	//empty map
	as := NewAllocStats()
	//creating a struct
	alloc := Alloc{Id: generateUUID(), Name: "test"}

	// update CPU

	cpuStat := CPUStat{User: 120, System: 10, IOWait: 120}

	as.UpdateCPU(alloc, cpuStat)

	stats := as.Stats(alloc)
	assert.Equal(cpuStat, stats.CpuStat)

	// update CPU again
	cpuStat = CPUStat{User: 98, System: 10, IOWait: 15}
	as.UpdateCPU(alloc, cpuStat)
	stats = as.Stats(alloc)
	assert.Equal(cpuStat, stats.CpuStat)

	// update memory
	as.UpdateMemory(alloc, 120000)
	stats = as.Stats(alloc)
	assert.Equal(uint64(120000), stats.MemUsage)

	// update disk
	as.UpdateDisk(alloc, 50000)
	stats = as.Stats(alloc)
	assert.Equal(uint64(50000), stats.DiskUsage)

}

func TestAllocStats_TopKByDisk(t *testing.T) {
	as := NewAllocStats()
	assert := assert.New(t)

	alloc1 := Alloc{Id: generateUUID(), Name: "test1"}
	alloc2 := Alloc{Id: generateUUID(), Name: "test2"}
	alloc3 := Alloc{Id: generateUUID(), Name: "test3"}

	as.UpdateDisk(alloc1, 50000)
	as.UpdateDisk(alloc2, 500)
	as.UpdateDisk(alloc3, 1000000)

	ret := as.TopKByDisk(3)
	expected := []Alloc{alloc3, alloc1, alloc2}
	assert.Equal(expected, ret)

}

func TestAllocStats_TopKByMemory(t *testing.T) {
	as := NewAllocStats()
	assert := assert.New(t)

	alloc1 := Alloc{Id: generateUUID(), Name: "test1"}
	alloc2 := Alloc{Id: generateUUID(), Name: "test2"}
	alloc3 := Alloc{Id: generateUUID(), Name: "test3"}

	as.UpdateMemory(alloc1, 5000)
	as.UpdateMemory(alloc2, 500)
	as.UpdateMemory(alloc3, 50)

	ret := as.TopKByDisk(3)
	expected := []Alloc{alloc1, alloc2, alloc3}
	assert.Equal(expected, ret)

}

// generateUUID is used to generate a random UUID.
func generateUUID() string {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		panic(fmt.Errorf("failed to read random bytes: %v", err))
	}

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16])
}
