/*
 Package stats has methods to update allocations with their CPU, Memory and Disk utilization
 via methods UpdateCPU, UpdateMemory and UpdateDisk. The top K allocations by these three
 criteria are returned by the methods TopKByCPU, TopKByMemory and TopKByDisk respectively.

 CPU stats track percentages spent in user, system and iowait. Initial implementation sorts only
 by user time.
*/
package stats
