package rsmi

import (
	"github.com/hhk7734/amdsmi.go/pkg/dl"
)

// #include "rocm_smi/rocm_smi.h"
import "C"

type RSMI struct {
	dl *dl.DynamicLibrary
}

func New() *RSMI {
	return &RSMI{
		dl: &dl.DynamicLibrary{},
	}
}

func (r *RSMI) Init(flags uint64) error {
	if err := r.dl.Open("librocm_smi64.so", dl.RTLD_LAZY|dl.RTLD_GLOBAL); err != nil {
		return err
	}

	if err := rsmi_init(flags).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RSMI) Shutdown() error {
	if err := rsmi_shut_down().Err(); err != nil {
		return err
	}

	if err := r.dl.Close(); err != nil {
		return err
	}

	return nil
}

func (r *RSMI) NumMonitorDevices() (uint32, error) {
	var numDevices uint32
	if err := rsmi_num_monitor_devices(&numDevices).Err(); err != nil {
		return 0, err
	}

	return numDevices, nil
}

func (r *RSMI) DevPCIID(deviceIndex uint32) (uint64, error) {
	var pciID uint64
	if err := rsmi_dev_pci_id_get(deviceIndex, &pciID).Err(); err != nil {
		return 0, err
	}

	return pciID, nil
}

func (r *RSMI) DevPowerAve(deviceIndex uint32, powerIndex uint32) (float64, error) {
	var power uint64
	if err := rsmi_dev_power_ave_get(deviceIndex, powerIndex, &power).Err(); err != nil {
		return 0, err
	}

	return float64(power) / 1_000_000, nil
}

func (r *RSMI) DevMemoryTotal(deviceIndex uint32, type_ memoryType) (uint64, error) {
	var total uint64
	if err := rsmi_dev_memory_total_get(deviceIndex, type_, &total).Err(); err != nil {
		return 0, err
	}

	return total, nil
}

func (r *RSMI) DevMemoryUsage(deviceIndex uint32, type_ memoryType) (uint64, error) {
	var usage uint64
	if err := rsmi_dev_memory_usage_get(deviceIndex, type_, &usage).Err(); err != nil {
		return 0, err
	}

	return usage, nil
}

func (r *RSMI) DevFanSpeed(deviceIndex uint32, fanIndex uint32) (int64, error) {
	var fanSpeed int64
	if err := rsmi_dev_fan_speed_get(deviceIndex, fanIndex, &fanSpeed).Err(); err != nil {
		return 0, err
	}

	return fanSpeed, nil
}

func (r *RSMI) DevFanSpeedMax(deviceIndex uint32, fanIndex uint32) (uint64, error) {
	var maxSpeed uint64
	if err := rsmi_dev_fan_speed_max_get(deviceIndex, fanIndex, &maxSpeed).Err(); err != nil {
		return 0, err
	}

	return maxSpeed, nil
}

func (r *RSMI) DevTempMetric(deviceIndex uint32, type_ temperatureType, metric temperatureMetric) (float64, error) {
	var temp int64
	if err := rsmi_dev_temp_metric_get(deviceIndex, uint32(type_), temperatureMetric(metric), &temp).Err(); err != nil {
		return 0, err
	}

	return float64(temp) / 1000, nil
}

func (r *RSMI) DevBusyPercent(deviceIndex uint32) (float64, error) {
	var busyPercent uint32
	if err := rsmi_dev_busy_percent_get(deviceIndex, &busyPercent).Err(); err != nil {
		return 0, err
	}

	return float64(busyPercent), nil
}
