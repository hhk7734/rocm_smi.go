// WARNING: This file has automatically been generated on Wed, 22 May 2024 01:06:30 KST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package rsmi

/*
#cgo linux LDFLAGS: -Wl,--export-dynamic -Wl,--unresolved-symbols=ignore-in-object-files
#include "rocm_smi/rocm_smi.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// MAX_NUM_FREQUENCIES as defined in rocm_smi/rocm_smi.h:74
	MAX_NUM_FREQUENCIES = 33
	// MAX_FAN_SPEED as defined in rocm_smi/rocm_smi.h:78
	MAX_FAN_SPEED = 255
	// NUM_VOLTAGE_CURVE_POINTS as defined in rocm_smi/rocm_smi.h:81
	NUM_VOLTAGE_CURVE_POINTS = 3
	// MAX_NUM_POWER_PROFILES as defined in rocm_smi/rocm_smi.h:737
	MAX_NUM_POWER_PROFILES = 0x6260e0
	// NUM_HBM_INSTANCES as defined in rocm_smi/rocm_smi.h:940
	NUM_HBM_INSTANCES = 4
	// MAX_NUM_VCNS as defined in rocm_smi/rocm_smi.h:945
	MAX_NUM_VCNS = 4
	// MAX_NUM_JPEG_ENGS as defined in rocm_smi/rocm_smi.h:950
	MAX_NUM_JPEG_ENGS = 32
	// MAX_NUM_CLKS as defined in rocm_smi/rocm_smi.h:955
	MAX_NUM_CLKS = 4
	// MAX_NUM_XGMI_LINKS as defined in rocm_smi/rocm_smi.h:960
	MAX_NUM_XGMI_LINKS = 8
	// MAX_NUM_GFX_CLKS as defined in rocm_smi/rocm_smi.h:965
	MAX_NUM_GFX_CLKS = 8
	// DEFAULT_VARIANT as defined in rocm_smi/rocm_smi.h:1152
	DEFAULT_VARIANT = 0xFFFFFFFFFFFFFFFF
)

// rsmi_status as declared in rocm_smi/rocm_smi.h:137
type rsmi_status int32

// rsmi_status enumeration from rocm_smi/rocm_smi.h:137
const (
	STATUS_SUCCESS             rsmi_status = iota
	STATUS_INVALID_ARGS        rsmi_status = 1
	STATUS_NOT_SUPPORTED       rsmi_status = 2
	STATUS_FILE_ERROR          rsmi_status = 3
	STATUS_PERMISSION          rsmi_status = 4
	STATUS_OUT_OF_RESOURCES    rsmi_status = 5
	STATUS_INTERNAL_EXCEPTION  rsmi_status = 6
	STATUS_INPUT_OUT_OF_BOUNDS rsmi_status = 7
	STATUS_INIT_ERROR          rsmi_status = 8
	INITIALIZATION_ERROR       rsmi_status = 8
	STATUS_NOT_YET_IMPLEMENTED rsmi_status = 9
	STATUS_NOT_FOUND           rsmi_status = 10
	STATUS_INSUFFICIENT_SIZE   rsmi_status = 11
	STATUS_INTERRUPT           rsmi_status = 12
	STATUS_UNEXPECTED_SIZE     rsmi_status = 13
	STATUS_NO_DATA             rsmi_status = 14
	STATUS_UNEXPECTED_DATA     rsmi_status = 15
	STATUS_BUSY                rsmi_status = 16
	STATUS_REFCOUNT_OVERFLOW   rsmi_status = 17
	STATUS_SETTING_UNAVAILABLE rsmi_status = 18
	STATUS_AMDGPU_RESTART_ERR  rsmi_status = 19
)

// rsmi_init_flags as declared in rocm_smi/rocm_smi.h:154
type rsmi_init_flags int32

// rsmi_init_flags enumeration from rocm_smi/rocm_smi.h:154
const (
	INIT_FLAG_ALL_GPUS rsmi_init_flags = 1
)

// rsmi_dev_perf_level as declared in rocm_smi/rocm_smi.h:181
type rsmi_dev_perf_level int32

// rsmi_dev_perf_level enumeration from rocm_smi/rocm_smi.h:181
const (
	DEV_PERF_LEVEL_AUTO            rsmi_dev_perf_level = iota
	DEV_PERF_LEVEL_FIRST           rsmi_dev_perf_level = 0
	DEV_PERF_LEVEL_LOW             rsmi_dev_perf_level = 1
	DEV_PERF_LEVEL_HIGH            rsmi_dev_perf_level = 2
	DEV_PERF_LEVEL_MANUAL          rsmi_dev_perf_level = 3
	DEV_PERF_LEVEL_STABLE_STD      rsmi_dev_perf_level = 4
	DEV_PERF_LEVEL_STABLE_PEAK     rsmi_dev_perf_level = 5
	DEV_PERF_LEVEL_STABLE_MIN_MCLK rsmi_dev_perf_level = 6
	DEV_PERF_LEVEL_STABLE_MIN_SCLK rsmi_dev_perf_level = 7
	DEV_PERF_LEVEL_DETERMINISM     rsmi_dev_perf_level = 8
	DEV_PERF_LEVEL_LAST            rsmi_dev_perf_level = 8
	DEV_PERF_LEVEL_UNKNOWN         rsmi_dev_perf_level = 256
)

// rsmi_sw_component as declared in rocm_smi/rocm_smi.h:198
type rsmi_sw_component int32

// rsmi_sw_component enumeration from rocm_smi/rocm_smi.h:198
const (
	SW_COMP_FIRST  rsmi_sw_component = iota
	SW_COMP_DRIVER rsmi_sw_component = 0
	SW_COMP_LAST   rsmi_sw_component = 0
)

// rsmi_event_group as declared in rocm_smi/rocm_smi.h:219
type rsmi_event_group int32

// rsmi_event_group enumeration from rocm_smi/rocm_smi.h:219
const (
	EVNT_GRP_XGMI          rsmi_event_group = iota
	EVNT_GRP_XGMI_DATA_OUT rsmi_event_group = 10
)

// rsmi_event_type as declared in rocm_smi/rocm_smi.h:288
type rsmi_event_type int32

// rsmi_event_type enumeration from rocm_smi/rocm_smi.h:288
const (
	EVNT_FIRST               rsmi_event_type = iota
	EVNT_XGMI_FIRST          rsmi_event_type = 0
	EVNT_XGMI_0_NOP_TX       rsmi_event_type = 0
	EVNT_XGMI_0_REQUEST_TX   rsmi_event_type = 1
	EVNT_XGMI_0_RESPONSE_TX  rsmi_event_type = 2
	EVNT_XGMI_0_BEATS_TX     rsmi_event_type = 3
	EVNT_XGMI_1_NOP_TX       rsmi_event_type = 4
	EVNT_XGMI_1_REQUEST_TX   rsmi_event_type = 5
	EVNT_XGMI_1_RESPONSE_TX  rsmi_event_type = 6
	EVNT_XGMI_1_BEATS_TX     rsmi_event_type = 7
	EVNT_XGMI_LAST           rsmi_event_type = 7
	EVNT_XGMI_DATA_OUT_FIRST rsmi_event_type = 10
	EVNT_XGMI_DATA_OUT_0     rsmi_event_type = 10
	EVNT_XGMI_DATA_OUT_1     rsmi_event_type = 11
	EVNT_XGMI_DATA_OUT_2     rsmi_event_type = 12
	EVNT_XGMI_DATA_OUT_3     rsmi_event_type = 13
	EVNT_XGMI_DATA_OUT_4     rsmi_event_type = 14
	EVNT_XGMI_DATA_OUT_5     rsmi_event_type = 15
	EVNT_XGMI_DATA_OUT_LAST  rsmi_event_type = 15
	EVNT_LAST                rsmi_event_type = 15
)

// rsmi_counter_command as declared in rocm_smi/rocm_smi.h:297
type rsmi_counter_command int32

// rsmi_counter_command enumeration from rocm_smi/rocm_smi.h:297
const (
	CNTR_CMD_START rsmi_counter_command = iota
	CNTR_CMD_STOP  rsmi_counter_command = 1
)

// rsmi_evt_notification_type as declared in rocm_smi/rocm_smi.h:323
type rsmi_evt_notification_type int32

// rsmi_evt_notification_type enumeration from rocm_smi/rocm_smi.h:323
const (
	EVT_NOTIF_NONE             rsmi_evt_notification_type = iota
	EVT_NOTIF_VMFAULT          rsmi_evt_notification_type = 1
	EVT_NOTIF_FIRST            rsmi_evt_notification_type = 1
	EVT_NOTIF_THERMAL_THROTTLE rsmi_evt_notification_type = 2
	EVT_NOTIF_GPU_PRE_RESET    rsmi_evt_notification_type = 3
	EVT_NOTIF_GPU_POST_RESET   rsmi_evt_notification_type = 4
	EVT_NOTIF_RING_HANG        rsmi_evt_notification_type = 5
	EVT_NOTIF_LAST             rsmi_evt_notification_type = 5
)

// rsmi_clk_type as declared in rocm_smi/rocm_smi.h:359
type rsmi_clk_type int32

// rsmi_clk_type enumeration from rocm_smi/rocm_smi.h:359
const (
	CLK_TYPE_SYS   rsmi_clk_type = iota
	CLK_TYPE_FIRST rsmi_clk_type = 0
	CLK_TYPE_DF    rsmi_clk_type = 1
	CLK_TYPE_DCEF  rsmi_clk_type = 2
	CLK_TYPE_SOC   rsmi_clk_type = 3
	CLK_TYPE_MEM   rsmi_clk_type = 4
	CLK_TYPE_PCIE  rsmi_clk_type = 5
	CLK_TYPE_LAST  rsmi_clk_type = 4
)

// rsmi_compute_partition_type as declared in rocm_smi/rocm_smi.h:380
type rsmi_compute_partition_type int32

// rsmi_compute_partition_type enumeration from rocm_smi/rocm_smi.h:380
const (
	COMPUTE_PARTITION_INVALID rsmi_compute_partition_type = iota
	COMPUTE_PARTITION_CPX     rsmi_compute_partition_type = 1
	COMPUTE_PARTITION_SPX     rsmi_compute_partition_type = 2
	COMPUTE_PARTITION_DPX     rsmi_compute_partition_type = 3
	COMPUTE_PARTITION_TPX     rsmi_compute_partition_type = 4
	COMPUTE_PARTITION_QPX     rsmi_compute_partition_type = 5
)

// rsmi_memory_partition_type as declared in rocm_smi/rocm_smi.h:404
type rsmi_memory_partition_type int32

// rsmi_memory_partition_type enumeration from rocm_smi/rocm_smi.h:404
const (
	MEMORY_PARTITION_UNKNOWN rsmi_memory_partition_type = iota
	MEMORY_PARTITION_NPS1    rsmi_memory_partition_type = 1
	MEMORY_PARTITION_NPS2    rsmi_memory_partition_type = 2
	MEMORY_PARTITION_NPS4    rsmi_memory_partition_type = 3
	MEMORY_PARTITION_NPS8    rsmi_memory_partition_type = 4
)

// rsmi_temperature_metric as declared in rocm_smi/rocm_smi.h:450
type rsmi_temperature_metric int32

// rsmi_temperature_metric enumeration from rocm_smi/rocm_smi.h:450
const (
	TEMP_CURRENT        rsmi_temperature_metric = iota
	TEMP_FIRST          rsmi_temperature_metric = 0
	TEMP_MAX            rsmi_temperature_metric = 1
	TEMP_MIN            rsmi_temperature_metric = 2
	TEMP_MAX_HYST       rsmi_temperature_metric = 3
	TEMP_MIN_HYST       rsmi_temperature_metric = 4
	TEMP_CRITICAL       rsmi_temperature_metric = 5
	TEMP_CRITICAL_HYST  rsmi_temperature_metric = 6
	TEMP_EMERGENCY      rsmi_temperature_metric = 7
	TEMP_EMERGENCY_HYST rsmi_temperature_metric = 8
	TEMP_CRIT_MIN       rsmi_temperature_metric = 9
	TEMP_CRIT_MIN_HYST  rsmi_temperature_metric = 10
	TEMP_OFFSET         rsmi_temperature_metric = 11
	TEMP_LOWEST         rsmi_temperature_metric = 12
	TEMP_HIGHEST        rsmi_temperature_metric = 13
	TEMP_LAST           rsmi_temperature_metric = 13
)

// rsmi_temperature_type as declared in rocm_smi/rocm_smi.h:472
type rsmi_temperature_type int32

// rsmi_temperature_type enumeration from rocm_smi/rocm_smi.h:472
const (
	TEMP_TYPE_FIRST    rsmi_temperature_type = iota
	TEMP_TYPE_EDGE     rsmi_temperature_type = 0
	TEMP_TYPE_JUNCTION rsmi_temperature_type = 1
	TEMP_TYPE_MEMORY   rsmi_temperature_type = 2
	TEMP_TYPE_HBM_0    rsmi_temperature_type = 3
	TEMP_TYPE_HBM_1    rsmi_temperature_type = 4
	TEMP_TYPE_HBM_2    rsmi_temperature_type = 5
	TEMP_TYPE_HBM_3    rsmi_temperature_type = 6
	TEMP_TYPE_LAST     rsmi_temperature_type = 6
)

// rsmi_activity_metric as declared in rocm_smi/rocm_smi.h:484
type rsmi_activity_metric int32

// rsmi_activity_metric enumeration from rocm_smi/rocm_smi.h:484
const (
	ACTIVITY_GFX rsmi_activity_metric = 1
	ACTIVITY_UMC rsmi_activity_metric = 2
	ACTIVITY_MM  rsmi_activity_metric = 4
)

// rsmi_voltage_metric as declared in rocm_smi/rocm_smi.h:505
type rsmi_voltage_metric int32

// rsmi_voltage_metric enumeration from rocm_smi/rocm_smi.h:505
const (
	VOLT_CURRENT  rsmi_voltage_metric = iota
	VOLT_FIRST    rsmi_voltage_metric = 0
	VOLT_MAX      rsmi_voltage_metric = 1
	VOLT_MIN_CRIT rsmi_voltage_metric = 2
	VOLT_MIN      rsmi_voltage_metric = 3
	VOLT_MAX_CRIT rsmi_voltage_metric = 4
	VOLT_AVERAGE  rsmi_voltage_metric = 5
	VOLT_LOWEST   rsmi_voltage_metric = 6
	VOLT_HIGHEST  rsmi_voltage_metric = 7
	VOLT_LAST     rsmi_voltage_metric = 7
)

// rsmi_voltage_type as declared in rocm_smi/rocm_smi.h:518
type rsmi_voltage_type int32

// rsmi_voltage_type enumeration from rocm_smi/rocm_smi.h:518
const (
	VOLT_TYPE_FIRST  rsmi_voltage_type = iota
	VOLT_TYPE_VDDGFX rsmi_voltage_type = 0
	VOLT_TYPE_LAST   rsmi_voltage_type = 0
)

// rsmi_power_profile_preset_masks as declared in rocm_smi/rocm_smi.h:540
type rsmi_power_profile_preset_masks int32

// rsmi_power_profile_preset_masks enumeration from rocm_smi/rocm_smi.h:540
const (
	PWR_PROF_PRST_CUSTOM_MASK       rsmi_power_profile_preset_masks = 1
	PWR_PROF_PRST_VIDEO_MASK        rsmi_power_profile_preset_masks = 2
	PWR_PROF_PRST_POWER_SAVING_MASK rsmi_power_profile_preset_masks = 4
	PWR_PROF_PRST_COMPUTE_MASK      rsmi_power_profile_preset_masks = 8
	PWR_PROF_PRST_VR_MASK           rsmi_power_profile_preset_masks = 16
	PWR_PROF_PRST_3D_FULL_SCR_MASK  rsmi_power_profile_preset_masks = 32
	PWR_PROF_PRST_BOOTUP_DEFAULT    rsmi_power_profile_preset_masks = 64
	PWR_PROF_PRST_LAST              rsmi_power_profile_preset_masks = 64
)

// rsmi_gpu_block as declared in rocm_smi/rocm_smi.h:571
type rsmi_gpu_block int32

// rsmi_gpu_block enumeration from rocm_smi/rocm_smi.h:571
const (
	GPU_BLOCK_INVALID   rsmi_gpu_block = iota
	GPU_BLOCK_FIRST     rsmi_gpu_block = 1
	GPU_BLOCK_UMC       rsmi_gpu_block = 1
	GPU_BLOCK_SDMA      rsmi_gpu_block = 2
	GPU_BLOCK_GFX       rsmi_gpu_block = 4
	GPU_BLOCK_MMHUB     rsmi_gpu_block = 8
	GPU_BLOCK_ATHUB     rsmi_gpu_block = 16
	GPU_BLOCK_PCIE_BIF  rsmi_gpu_block = 32
	GPU_BLOCK_HDP       rsmi_gpu_block = 64
	GPU_BLOCK_XGMI_WAFL rsmi_gpu_block = 128
	GPU_BLOCK_DF        rsmi_gpu_block = 256
	GPU_BLOCK_SMN       rsmi_gpu_block = 512
	GPU_BLOCK_SEM       rsmi_gpu_block = 1024
	GPU_BLOCK_MP0       rsmi_gpu_block = 2048
	GPU_BLOCK_MP1       rsmi_gpu_block = 4096
	GPU_BLOCK_FUSE      rsmi_gpu_block = 8192
	GPU_BLOCK_LAST      rsmi_gpu_block = 8192
)

// rsmi_ras_err_state as declared in rocm_smi/rocm_smi.h:591
type rsmi_ras_err_state int32

// rsmi_ras_err_state enumeration from rocm_smi/rocm_smi.h:591
const (
	RAS_ERR_STATE_NONE     rsmi_ras_err_state = iota
	RAS_ERR_STATE_DISABLED rsmi_ras_err_state = 1
	RAS_ERR_STATE_PARITY   rsmi_ras_err_state = 2
	RAS_ERR_STATE_SING_C   rsmi_ras_err_state = 3
	RAS_ERR_STATE_MULT_UC  rsmi_ras_err_state = 4
	RAS_ERR_STATE_POISON   rsmi_ras_err_state = 5
	RAS_ERR_STATE_ENABLED  rsmi_ras_err_state = 6
	RAS_ERR_STATE_LAST     rsmi_ras_err_state = 6
)

// rsmi_memory_type as declared in rocm_smi/rocm_smi.h:604
type rsmi_memory_type int32

// rsmi_memory_type enumeration from rocm_smi/rocm_smi.h:604
const (
	MEM_TYPE_FIRST    rsmi_memory_type = iota
	MEM_TYPE_VRAM     rsmi_memory_type = 0
	MEM_TYPE_VIS_VRAM rsmi_memory_type = 1
	MEM_TYPE_GTT      rsmi_memory_type = 2
	MEM_TYPE_LAST     rsmi_memory_type = 2
)

// rsmi_freq_ind as declared in rocm_smi/rocm_smi.h:613
type rsmi_freq_ind int32

// rsmi_freq_ind enumeration from rocm_smi/rocm_smi.h:613
const (
	FREQ_IND_MIN rsmi_freq_ind = iota
	FREQ_IND_MAX rsmi_freq_ind = 1
)

// rsmi_fw_block as declared in rocm_smi/rocm_smi.h:651
type rsmi_fw_block int32

// rsmi_fw_block enumeration from rocm_smi/rocm_smi.h:651
const (
	FW_BLOCK_FIRST    rsmi_fw_block = iota
	FW_BLOCK_ASD      rsmi_fw_block = 0
	FW_BLOCK_CE       rsmi_fw_block = 1
	FW_BLOCK_DMCU     rsmi_fw_block = 2
	FW_BLOCK_MC       rsmi_fw_block = 3
	FW_BLOCK_ME       rsmi_fw_block = 4
	FW_BLOCK_MEC      rsmi_fw_block = 5
	FW_BLOCK_MEC2     rsmi_fw_block = 6
	FW_BLOCK_MES      rsmi_fw_block = 7
	FW_BLOCK_MES_KIQ  rsmi_fw_block = 8
	FW_BLOCK_PFP      rsmi_fw_block = 9
	FW_BLOCK_RLC      rsmi_fw_block = 10
	FW_BLOCK_RLC_SRLC rsmi_fw_block = 11
	FW_BLOCK_RLC_SRLG rsmi_fw_block = 12
	FW_BLOCK_RLC_SRLS rsmi_fw_block = 13
	FW_BLOCK_SDMA     rsmi_fw_block = 14
	FW_BLOCK_SDMA2    rsmi_fw_block = 15
	FW_BLOCK_SMC      rsmi_fw_block = 16
	FW_BLOCK_SOS      rsmi_fw_block = 17
	FW_BLOCK_TA_RAS   rsmi_fw_block = 18
	FW_BLOCK_TA_XGMI  rsmi_fw_block = 19
	FW_BLOCK_UVD      rsmi_fw_block = 20
	FW_BLOCK_VCE      rsmi_fw_block = 21
	FW_BLOCK_VCN      rsmi_fw_block = 22
	FW_BLOCK_LAST     rsmi_fw_block = 22
)

// rsmi_xgmi_status as declared in rocm_smi/rocm_smi.h:660
type rsmi_xgmi_status int32

// rsmi_xgmi_status enumeration from rocm_smi/rocm_smi.h:660
const (
	XGMI_STATUS_NO_ERRORS       rsmi_xgmi_status = iota
	XGMI_STATUS_ERROR           rsmi_xgmi_status = 1
	XGMI_STATUS_MULTIPLE_ERRORS rsmi_xgmi_status = 2
)

// rsmi_memory_page_status as declared in rocm_smi/rocm_smi.h:680
type rsmi_memory_page_status int32

// rsmi_memory_page_status enumeration from rocm_smi/rocm_smi.h:680
const (
	MEM_PAGE_STATUS_RESERVED     rsmi_memory_page_status = iota
	MEM_PAGE_STATUS_PENDING      rsmi_memory_page_status = 1
	MEM_PAGE_STATUS_UNRESERVABLE rsmi_memory_page_status = 2
)

// RSMI_IO_LINK as declared in rocm_smi/rocm_smi.h:691
type RSMI_IO_LINK int32

// RSMI_IO_LINK enumeration from rocm_smi/rocm_smi.h:691
const (
	IOLINK_TYPE_UNDEFINED      RSMI_IO_LINK = iota
	IOLINK_TYPE_PCIEXPRESS     RSMI_IO_LINK = 1
	IOLINK_TYPE_XGMI           RSMI_IO_LINK = 2
	IOLINK_TYPE_NUMIOLINKTYPES RSMI_IO_LINK = 3
)

// RSMI_UTILIZATION_COUNTER as declared in rocm_smi/rocm_smi.h:706
type RSMI_UTILIZATION_COUNTER int32

// RSMI_UTILIZATION_COUNTER enumeration from rocm_smi/rocm_smi.h:706
const (
	UTILIZATION_COUNTER_FIRST RSMI_UTILIZATION_COUNTER = iota
	COARSE_GRAIN_GFX_ACTIVITY RSMI_UTILIZATION_COUNTER = 0
	COARSE_GRAIN_MEM_ACTIVITY RSMI_UTILIZATION_COUNTER = 1
	UTILIZATION_COUNTER_LAST  RSMI_UTILIZATION_COUNTER = 1
)

// RSMI_POWER as declared in rocm_smi/rocm_smi.h:715
type RSMI_POWER int32

// RSMI_POWER enumeration from rocm_smi/rocm_smi.h:715
const (
	AVERAGE_POWER RSMI_POWER = iota
	CURRENT_POWER RSMI_POWER = 1
)

const ()

const ()
