package linux

const (
	//Thermal constants extracted from linux kernel source code
	//user space folder with thermal information
	KernelSysThermalClass string = "/sys/class/thermal/"

	//Types constants

	KernelSysThermalTypeBcm2835 string = "bcm2835_thermal"
	//drivers/thermal/intel/int340x_thermal/int3400_thermal.c
	KernelSysThermalTypeInt3400Temp string = "INT3400 Thermal"
	//drivers/acpi/thermal.c
	KernelSysThermalTypeAcpitz string = "acpitz"
	//from drivers/thermal/qcom/tsens.c
	KernelSysThermalTypeQcom string = "tsens"
	//drivers/thermal/intel/x86_pkg_temp_thermal.c
	KernelSysThermalTypeX86PkgTemp string = "x86_pkg_temp"

	KernelSysThermalTypeTmem  = "TMEM"
	KernelSysThermalTypeTskn  = "TSKN"
	KernelSysThermalTypeNgff  = "NGFF"
	KernelSysThermalTypeB0d4  = "B0D4"
	KernelSysThermalTypeIwifi = "iwlwifi_1"

	PathThermalType = "type"
	PathThermalTemp = "temp"
)

//GetThermals - retrieve all the thermal constants in a slice
func GetThermals() []string {
	return []string{
		KernelSysThermalTypeX86PkgTemp,
		KernelSysThermalTypeBcm2835,
		KernelSysThermalTypeAcpitz,
		KernelSysThermalTypeQcom,
		KernelSysThermalTypeTmem,
		KernelSysThermalTypeTskn,
		KernelSysThermalTypeNgff,
		KernelSysThermalTypeB0d4,
		KernelSysThermalTypeIwifi}
}
