package linux

const (
	//Thermal constants extracted from linux kernel source code
	//user space folder with thermal information
	KERNEL_SYS_THERMAL_CLASS string = "/sys/class/thermal/"

	//Types constants

	KERNEL_SYS_THERMAL_TYPE_FAN       string = "fan"
	KERNEL_SYS_THERMAL_TYPE_PROCESSOR string = "processor"

	//from drivers/thermal/broadcom/bcm2835_thermal.c
	KERNEL_SYS_THERMAL_TYPE_BCM2835 string = "bcm2835_thermal"
	//from drivers/thermal/intel/intel_powerclamp.c
	KERNEL_SYS_THERMAL_TYPE_INTEL_POWERCLAMP string = "intel_powerclamp"
	//from drivers/thermal/intel/x86_pkg_temp_thermal.c
	KERNEL_SYS_THERMAL_TYPE_X86_PKG_TEMP string = "pkg_temp_thermal"
	//from drivers/thermal/qcom/tsens.c
	KERNEL_SYS_THERMAL_TYPE_QCOM string = "tsens"
)
