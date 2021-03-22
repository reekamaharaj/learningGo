package stats

import (
	"sort"
)

type Alloc struct {
	Id   string
	Name string
}

type Stats struct {
	CpuStat   CPUStat
	MemUsage  uint64
	DiskUsage uint64
}

type CPUStat struct {
	User   uint64
	System uint64
	IOWait uint64
}

type AllocStats struct {
	StatsMap map[Alloc]*Stats
}

func NewAllocStats() *AllocStats {
	return &AllocStats{make(map[Alloc]*Stats)}
}

//receiver, as-variable AllocStats-type
func (as *AllocStats) UpdateCPU(alloc Alloc, stat CPUStat) {
	stats, ok := as.StatsMap[alloc]
	if !ok {
		stats = &Stats{CpuStat: stat}
	} else {
		stats.CpuStat = stat
	}
	as.StatsMap[alloc] = stats
}

func (as *AllocStats) UpdateMemory(alloc Alloc, memUsage uint64) {
	stats, ok := as.StatsMap[alloc]
	if !ok {
		stats = &Stats{MemUsage: memUsage}
	} else {
		stats.MemUsage = memUsage
	}
	as.StatsMap[alloc] = stats
}

func (as *AllocStats) UpdateDisk(alloc Alloc, diskUsage uint64) {
	stats, ok := as.StatsMap[alloc]
	if !ok {
		stats = &Stats{DiskUsage: diskUsage}
	} else {
		stats.DiskUsage = diskUsage
	}
	as.StatsMap[alloc] = stats
}

func (as *AllocStats) Stats(alloc Alloc) *Stats {
	return as.StatsMap[alloc]
}

type AllocAndStats struct {
	Alloc Alloc
	Stats Stats
}

func (as *AllocStats) TopKByCPU(k int) []Alloc {
	allocs := make([]Alloc, k)

	var allocStats []AllocAndStats
	for k, v := range as.StatsMap {
		allocStats = append(allocStats, AllocAndStats{k, *v})
	}

	sort.Slice(allocStats, func(i, j int) bool {
		return allocStats[i].Stats.CpuStat.User < allocStats[j].Stats.CpuStat.User
	})

	for i := 0; i < k; i++ {
		allocs[i] = allocStats[i].Alloc
	}

	return allocs

}

func (as *AllocStats) TopKByDisk(k int) []Alloc {
	allocs := make([]Alloc, k)

	var allocStats []AllocAndStats
	for k, v := range as.StatsMap {
		allocStats = append(allocStats, AllocAndStats{k, *v})
	}

	sort.Slice(allocStats, func(i, j int) bool {
		return allocStats[i].Stats.DiskUsage > allocStats[j].Stats.DiskUsage
	})

	for i := 0; i < k; i++ {
		allocs[i] = allocStats[i].Alloc
	}

	return allocs

}

func (as *AllocStats) TopKByMemory(k int) []Alloc {
	allocs := make([]Alloc, k)

	var allocStats []AllocAndStats
	for k, v := range as.StatsMap {
		allocStats = append(allocStats, AllocAndStats{k, *v})
	}

	sort.Slice(allocStats, func(i, j int) bool {
		return allocStats[i].Stats.MemUsage > allocStats[i].Stats.MemUsage
	})

	for i := 0; i < k; i++ {
		allocs[i] = allocStats[i].Alloc
	}

	return allocs

}
