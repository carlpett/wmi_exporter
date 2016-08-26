// returns data points from Win32_PerfRawData_PerfDisk_LogicalDisk
// https://msdn.microsoft.com/en-us/windows/hardware/aa394307(v=vs.71) - Win32_PerfRawData_PerfDisk_LogicalDisk class

package collectors

import (
	"fmt"
	"strings"

	"github.com/StackExchange/wmi"
	"github.com/prometheus/client_golang/prometheus"
)

type PerfCollector struct{}

/*
// A PerfCollector is a Prometheus collector for WMI Win32_PerfRawData_PerfDisk_LogicalDisk metrics
type PerfCollector struct {
	AvgDiskBytesPerRead          *prometheus.Desc
	AvgDiskBytesPerRead_Base     *prometheus.Desc
	AvgDiskBytesPerTransfer      *prometheus.Desc
	AvgDiskBytesPerTransfer_Base *prometheus.Desc
	AvgDiskBytesPerWrite         *prometheus.Desc
	AvgDiskBytesPerWrite_Base    *prometheus.Desc
	AvgDiskQueueLength           *prometheus.Desc
	AvgDiskReadQueueLength       *prometheus.Desc
	AvgDiskSecPerRead            *prometheus.Desc
	AvgDiskSecPerRead_Base       *prometheus.Desc
	AvgDiskSecPerTransfer        *prometheus.Desc
	AvgDiskSecPerTransfer_Base   *prometheus.Desc
	AvgDiskSecPerWrite           *prometheus.Desc
	AvgDiskSecPerWrite_Base      *prometheus.Desc
	AvgDiskWriteQueueLength      *prometheus.Desc
	CurrentDiskQueueLength       *prometheus.Desc
	DiskBytesPerSec              *prometheus.Desc
	DiskReadBytesPerSec          *prometheus.Desc
	DiskReadsPerSec              *prometheus.Desc
	DiskTransfersPerSec          *prometheus.Desc
	DiskWriteBytesPerSec         *prometheus.Desc
	DiskWritesPerSec             *prometheus.Desc
	FreeMegabytes                *prometheus.Desc
	PercentDiskReadTime          *prometheus.Desc
	PercentDiskReadTime_Base     *prometheus.Desc
	PercentDiskTime              *prometheus.Desc
	PercentDiskTime_Base         *prometheus.Desc
	PercentDiskWriteTime         *prometheus.Desc
	PercentDiskWriteTime_Base    *prometheus.Desc
	PercentFreeSpace             *prometheus.Desc
	PercentFreeSpace_Base        *prometheus.Desc
	PercentIdleTime              *prometheus.Desc
	PercentIdleTime_Base         *prometheus.Desc
	SplitIOPerSec                *prometheus.Desc
}
*/

// NewPerfCollector ...
func NewPerfCollector() Collector {

	return &PerfCollector{}
}

/*

   AvgDiskBytesPerRead: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_bytes_per_read"),
       "AvgDiskBytesPerRead",
       nil,
       nil,
   ),

   AvgDiskBytesPerRead_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_bytes_per_read_base"),
       "AvgDiskBytesPerRead_Base",
       nil,
       nil,
   ),

   AvgDiskBytesPerTransfer: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_bytes_per_transfer"),
       "AvgDiskBytesPerTransfer",
       nil,
       nil,
   ),

   AvgDiskBytesPerTransfer_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_bytes_per_transfer_base"),
       "AvgDiskBytesPerTransfer_Base",
       nil,
       nil,
   ),

   AvgDiskBytesPerWrite: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_bytes_per_write"),
       "AvgDiskBytesPerWrite",
       nil,
       nil,
   ),

   AvgDiskBytesPerWrite_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_bytes_per_write_base"),
       "AvgDiskBytesPerWrite_Base",
       nil,
       nil,
   ),

   AvgDiskQueueLength: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_queue_length"),
       "AvgDiskQueueLength",
       nil,
       nil,
   ),

   AvgDiskReadQueueLength: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_read_queue_length"),
       "AvgDiskReadQueueLength",
       nil,
       nil,
   ),

   AvgDiskSecPerRead: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_sec_per_read"),
       "AvgDiskSecPerRead",
       nil,
       nil,
   ),

   AvgDiskSecPerRead_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_sec_per_read_base"),
       "AvgDiskSecPerRead_Base",
       nil,
       nil,
   ),

   AvgDiskSecPerTransfer: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_sec_per_transfer"),
       "AvgDiskSecPerTransfer",
       nil,
       nil,
   ),

   AvgDiskSecPerTransfer_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_sec_per_transfer_base"),
       "AvgDiskSecPerTransfer_Base",
       nil,
       nil,
   ),

   AvgDiskSecPerWrite: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_sec_per_write"),
       "AvgDiskSecPerWrite",
       nil,
       nil,
   ),

   AvgDiskSecPerWrite_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_sec_per_write_base"),
       "AvgDiskSecPerWrite_Base",
       nil,
       nil,
   ),

   AvgDiskWriteQueueLength: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "avg_disk_write_queue_length"),
       "AvgDiskWriteQueueLength",
       nil,
       nil,
   ),

   CurrentDiskQueueLength: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "current_disk_queue_length"),
       "CurrentDiskQueueLength",
       nil,
       nil,
   ),

   DiskBytesPerSec: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "disk_bytes_per_sec"),
       "DiskBytesPerSec",
       nil,
       nil,
   ),

   DiskReadBytesPerSec: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "disk_read_bytes_per_sec"),
       "DiskReadBytesPerSec",
       nil,
       nil,
   ),

   DiskReadsPerSec: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "disk_reads_per_sec"),
       "DiskReadsPerSec",
       nil,
       nil,
   ),

   DiskTransfersPerSec: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "disk_transfers_per_sec"),
       "DiskTransfersPerSec",
       nil,
       nil,
   ),

   DiskWriteBytesPerSec: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "disk_write_bytes_per_sec"),
       "DiskWriteBytesPerSec",
       nil,
       nil,
   ),

   DiskWritesPerSec: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "disk_writes_per_sec"),
       "DiskWritesPerSec",
       nil,
       nil,
   ),

   FreeMegabytes: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "free_megabytes"), // XXX in bytes
       "FreeMegabytes",
       nil,
       nil,
   ),

   PercentDiskReadTime: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "percent_disk_read_time"),
       "PercentDiskReadTime",
       nil,
       nil,
   ),

   PercentDiskReadTime_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "percent_disk_read_time_base"),
       "PercentDiskReadTime_Base",
       nil,
       nil,
   ),

   PercentDiskTime: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "percent_disk_time"),
       "PercentDiskTime",
       nil,
       nil,
   ),

   PercentDiskTime_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "percent_disk_time_base"),
       "PercentDiskTime_Base",
       nil,
       nil,
   ),

   PercentDiskWriteTime: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "percent_disk_write_time"),
       "PercentDiskWriteTime",
       nil,
       nil,
   ),

   PercentDiskWriteTime_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "percent_disk_write_time_base"),
       "PercentDiskWriteTime_Base",
       nil,
       nil,
   ),

   PercentFreeSpace: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "percent_free_space"),
       "PercentFreeSpace",
       nil,
       nil,
   ),

   PercentFreeSpace_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "percent_free_space_base"),
       "PercentFreeSpace_Base",
       nil,
       nil,
   ),

   PercentIdleTime: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "percent_idle_time"),
       "PercentIdleTime",
       nil,
       nil,
   ),

   PercentIdleTime_Base: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "percent_idle_time_base"),
       "PercentIdleTime_Base",
       nil,
       nil,
   ),

   SplitIOPerSec: prometheus.NewDesc(
       prometheus.BuildFQName(wmiNamespace, "perf", "split_io_per_sec"),
       "SplitIOPerSec",
       nil,
       nil,
   ),
*/

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *PerfCollector) Update(ch chan<- prometheus.Metric) error {
	if err := c.update(ch); err != nil {
		return err
	}
	return nil
}

// Describe sends the descriptors of each metric over to the provided channel.
// The corresponding metric values are sent separately.
/*
func (c *PerfCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- c.AvgDiskBytesPerRead
	ch <- c.AvgDiskBytesPerRead_Base
	ch <- c.AvgDiskBytesPerTransfer
	ch <- c.AvgDiskBytesPerTransfer_Base
	ch <- c.AvgDiskBytesPerWrite
	ch <- c.AvgDiskBytesPerWrite_Base
	ch <- c.AvgDiskQueueLength
	ch <- c.AvgDiskReadQueueLength
	ch <- c.AvgDiskSecPerRead
	ch <- c.AvgDiskSecPerRead_Base
	ch <- c.AvgDiskSecPerTransfer
	ch <- c.AvgDiskSecPerTransfer_Base
	ch <- c.AvgDiskSecPerWrite
	ch <- c.AvgDiskSecPerWrite_Base
	ch <- c.AvgDiskWriteQueueLength
	ch <- c.CurrentDiskQueueLength
	ch <- c.DiskBytesPerSec
	ch <- c.DiskReadBytesPerSec
	ch <- c.DiskReadsPerSec
	ch <- c.DiskTransfersPerSec
	ch <- c.DiskWriteBytesPerSec
	ch <- c.DiskWritesPerSec
	ch <- c.FreeMegabytes
	ch <- c.PercentDiskReadTime
	ch <- c.PercentDiskReadTime_Base
	ch <- c.PercentDiskTime
	ch <- c.PercentDiskTime_Base
	ch <- c.PercentDiskWriteTime
	ch <- c.PercentDiskWriteTime_Base
	ch <- c.PercentFreeSpace
	ch <- c.PercentFreeSpace_Base
	ch <- c.PercentIdleTime
	ch <- c.PercentIdleTime_Base
	ch <- c.SplitIOPerSec
}
*/

type Win32_PerfRawData_PerfDisk_LogicalDisk struct {
	AvgDiskBytesPerRead          uint64
	AvgDiskBytesPerRead_Base     uint32
	AvgDiskBytesPerTransfer      uint64
	AvgDiskBytesPerTransfer_Base uint32
	AvgDiskBytesPerWrite         uint64
	AvgDiskBytesPerWrite_Base    uint32
	AvgDiskQueueLength           uint64
	AvgDiskReadQueueLength       uint64
	AvgDiskSecPerRead            uint32
	AvgDiskSecPerRead_Base       uint32
	AvgDiskSecPerTransfer        uint32
	AvgDiskSecPerTransfer_Base   uint32
	AvgDiskSecPerWrite           uint32
	AvgDiskSecPerWrite_Base      uint32
	AvgDiskWriteQueueLength      uint64
	CurrentDiskQueueLength       uint32
	DiskBytesPerSec              uint64
	DiskReadBytesPerSec          uint64
	DiskReadsPerSec              uint32
	DiskTransfersPerSec          uint32
	DiskWriteBytesPerSec         uint64
	DiskWritesPerSec             uint32
	FreeMegabytes                uint32
	Name                         string
	PercentDiskReadTime          uint64
	PercentDiskReadTime_Base     uint64
	PercentDiskTime              uint64
	PercentDiskTime_Base         uint64
	PercentDiskWriteTime         uint64
	PercentDiskWriteTime_Base    uint64
	PercentFreeSpace             uint32
	PercentFreeSpace_Base        uint32
	PercentIdleTime              uint64
	PercentIdleTime_Base         uint64
	SplitIOPerSec                uint32
}

func (c *PerfCollector) collect(ch chan<- prometheus.Metric) error {
	var dst []Win32_PerfRawData_PerfDisk_LogicalDisk
	q := wmi.CreateQuery(&dst, "")
	if err := wmi.Query(q, &dst); err != nil {
		return err
	}

	for _, perf := range dst {

		// XXX use drive letter from "Name" key in resulting name

		if perf.Name == "_Total" {
			fmt.Println("SKIPPING TOTAL!")
			continue
		}

		shortName := strings.ToLower(string(perf.Name[0]))

		v := 13.22

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(wmiNamespace, "perf"+shortName, "xxx"),
				"xxx help",
				nil, nil,
			),
			prometheus.GaugeValue, v,
		)

		/*
			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskBytesPerRead,
				prometheus.GaugeValue,
				float64(perf.AvgDiskBytesPerRead),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskBytesPerRead_Base,
				prometheus.GaugeValue,
				float64(perf.AvgDiskBytesPerRead_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskBytesPerTransfer,
				prometheus.GaugeValue,
				float64(perf.AvgDiskBytesPerTransfer),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskBytesPerTransfer_Base,
				prometheus.GaugeValue,
				float64(perf.AvgDiskBytesPerTransfer_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskBytesPerWrite,
				prometheus.GaugeValue,
				float64(perf.AvgDiskBytesPerWrite),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskBytesPerWrite_Base,
				prometheus.GaugeValue,
				float64(perf.AvgDiskBytesPerWrite_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskQueueLength,
				prometheus.GaugeValue,
				float64(perf.AvgDiskQueueLength),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskReadQueueLength,
				prometheus.GaugeValue,
				float64(perf.AvgDiskReadQueueLength),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskSecPerRead,
				prometheus.GaugeValue,
				float64(perf.AvgDiskSecPerRead),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskSecPerRead_Base,
				prometheus.GaugeValue,
				float64(perf.AvgDiskSecPerRead_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskSecPerTransfer,
				prometheus.GaugeValue,
				float64(perf.AvgDiskSecPerTransfer),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskSecPerTransfer_Base,
				prometheus.GaugeValue,
				float64(perf.AvgDiskSecPerTransfer_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskSecPerWrite,
				prometheus.GaugeValue,
				float64(perf.AvgDiskSecPerWrite),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskSecPerWrite_Base,
				prometheus.GaugeValue,
				float64(perf.AvgDiskSecPerWrite_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.AvgDiskWriteQueueLength,
				prometheus.GaugeValue,
				float64(perf.AvgDiskWriteQueueLength),
			)

			ch <- prometheus.MustNewConstMetric(
				c.CurrentDiskQueueLength,
				prometheus.GaugeValue,
				float64(perf.CurrentDiskQueueLength),
			)

			ch <- prometheus.MustNewConstMetric(
				c.DiskBytesPerSec,
				prometheus.GaugeValue,
				float64(perf.DiskBytesPerSec),
			)

			ch <- prometheus.MustNewConstMetric(
				c.DiskReadBytesPerSec,
				prometheus.GaugeValue,
				float64(perf.DiskReadBytesPerSec),
			)

			ch <- prometheus.MustNewConstMetric(
				c.DiskReadsPerSec,
				prometheus.GaugeValue,
				float64(perf.DiskReadsPerSec),
			)

			ch <- prometheus.MustNewConstMetric(
				c.DiskTransfersPerSec,
				prometheus.GaugeValue,
				float64(perf.DiskTransfersPerSec),
			)

			ch <- prometheus.MustNewConstMetric(
				c.DiskWriteBytesPerSec,
				prometheus.GaugeValue,
				float64(perf.DiskWriteBytesPerSec),
			)

			ch <- prometheus.MustNewConstMetric(
				c.DiskWritesPerSec,
				prometheus.GaugeValue,
				float64(perf.DiskWritesPerSec),
			)

			ch <- prometheus.MustNewConstMetric(
				c.FreeMegabytes,
				prometheus.GaugeValue,
				float64(perf.FreeMegabytes),
			)

			ch <- prometheus.MustNewConstMetric(
				c.PercentDiskReadTime,
				prometheus.GaugeValue,
				float64(perf.PercentDiskReadTime),
			)

			ch <- prometheus.MustNewConstMetric(
				c.PercentDiskReadTime_Base,
				prometheus.GaugeValue,
				float64(perf.PercentDiskReadTime_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.PercentDiskTime,
				prometheus.GaugeValue,
				float64(perf.PercentDiskTime),
			)

			ch <- prometheus.MustNewConstMetric(
				c.PercentDiskTime_Base,
				prometheus.GaugeValue,
				float64(perf.PercentDiskTime_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.PercentDiskWriteTime,
				prometheus.GaugeValue,
				float64(perf.PercentDiskWriteTime),
			)

			ch <- prometheus.MustNewConstMetric(
				c.PercentDiskWriteTime_Base,
				prometheus.GaugeValue,
				float64(perf.PercentDiskWriteTime_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.PercentFreeSpace,
				prometheus.GaugeValue,
				float64(perf.PercentFreeSpace),
			)

			ch <- prometheus.MustNewConstMetric(
				c.PercentFreeSpace_Base,
				prometheus.GaugeValue,
				float64(perf.PercentFreeSpace_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.PercentIdleTime,
				prometheus.GaugeValue,
				float64(perf.PercentIdleTime),
			)

			ch <- prometheus.MustNewConstMetric(
				c.PercentIdleTime_Base,
				prometheus.GaugeValue,
				float64(perf.PercentIdleTime_Base),
			)

			ch <- prometheus.MustNewConstMetric(
				c.SplitIOPerSec,
				prometheus.GaugeValue,
				float64(perf.SplitIOPerSec),
			)
		*/
	}

	return nil
}
