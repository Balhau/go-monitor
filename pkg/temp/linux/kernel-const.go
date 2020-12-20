package linux

const (
	//Thermal constants extracted from linux kernel source code
	//user space folder with thermal information
	KERNEL_SYS_THERMAL_CLASS string = "/sys/class/thermal/"

	//Types constants

	KERNEL_SYS_THERMAL_TYPE_BCM2835 string = "bcm2835_thermal"
	//drivers/thermal/intel/int340x_thermal/int3400_thermal.c
	KERNEL_SYS_THERMAL_TYPE_X86_PKG_TEMP string = "INT3400 Thermal"
	//drivers/acpi/thermal.c
	KERNEL_SYS_THERMAL_TYPE_ACPITZ string = "acpitz"
	//from drivers/thermal/qcom/tsens.c
	KERNEL_SYS_THERMAL_TYPE_QCOM string = "tsens"

	KERNEL_SYS_THERMAL_TYPE_TMEM  = "TMEM"
	KERNEL_SYS_THERMAL_TYPE_TSKN  = "TSKN"
	KERNEL_SYS_THERMAL_TYPE_NGFF  = "NGFF"
	KERNEL_SYS_THERMAL_TYPE_B0D4  = "B0D4"
	KERNEL_SYS_THERMAL_TYPE_IWIFI = "iwlwifi_1"
)

//GetThermals - retrieve all the thermal constants in a slice
func GetThermals() []string {
	return []string{
		KERNEL_SYS_THERMAL_TYPE_X86_PKG_TEMP,
		KERNEL_SYS_THERMAL_TYPE_BCM2835,
		KERNEL_SYS_THERMAL_TYPE_ACPITZ,
		KERNEL_SYS_THERMAL_TYPE_QCOM,
		KERNEL_SYS_THERMAL_TYPE_TMEM,
		KERNEL_SYS_THERMAL_TYPE_TSKN,
		KERNEL_SYS_THERMAL_TYPE_NGFF,
		KERNEL_SYS_THERMAL_TYPE_B0D4,
		KERNEL_SYS_THERMAL_TYPE_IWIFI}
}
