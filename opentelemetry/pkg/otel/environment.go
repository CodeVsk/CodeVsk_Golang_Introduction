package otel

import (
	"context"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type RecordMetrics struct {
	Name string
}

var (
	errorList = []error{}

	cpuUsagePercentage metric.Float64Gauge
	memoryUsageBytes   metric.Float64Gauge
)

func (o *OtelMeter) UseRecordMetrics() *RecordMetrics {
	var err error

	cpuUsagePercentage, err = Meter.Float64Gauge(
		"cpu_usage_percentage",
		metric.WithDescription("Measures the CPU usage percentage"),
	)
	if err != nil {
		errorList = append(errorList, err)
	}

	memoryUsageBytes, err = Meter.Float64Gauge(
		"api_memory_usage_bytes",
		metric.WithDescription("Measures the memory usage in bytes"),
	)
	if err != nil {
		errorList = append(errorList, err)
	}

	return &RecordMetrics{Name: o.Name}
}

func recordCPUUsage(ctx context.Context, serviceName string) {
	percent, _ := cpu.Percent(1*time.Second, false)

	cpuUsagePercentage.Record(ctx, percent[0], metric.WithAttributes(
		attribute.String("service", serviceName),
	))
}

func recordMemoryUsage(ctx context.Context, serviceName string) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	memoryUsageBytes.Record(ctx, float64(mem.HeapAlloc), metric.WithAttributes(
		attribute.String("service", serviceName),
	))
}

func (o *RecordMetrics) Init() error {
	if len(errorList) > 0 {
		return errorList[0]
	}

	ctx := context.Background()
	go func() {
		for {
			recordCPUUsage(ctx, o.Name)
			recordMemoryUsage(ctx, o.Name)
			time.Sleep(2 * time.Second)
		}
	}()

	return nil
}
