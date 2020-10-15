package system

import (
	"github.com/bigbroproject/bigbrocore/utilities"
	"github.com/fatih/color"
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"
	"time"
)

type HostInfo struct {
	Platform        string
	PlatformVersion string
	KernelArch      string
	Uptime          uint64
}

type CpuInfo struct {
	ModelName string
	Frequency float64
	Cores     int
	Usage     float64
}

type GpuInfo struct {
	Name   string
	Vendor string
}

type MemoryInfo struct {
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}
type NetworkInfo struct {
	InternetConnected bool
	PublicIP          string
}

type SystemInfo struct {
	Host    HostInfo
	Cpu     CpuInfo
	Gpu     GpuInfo
	Memory  MemoryInfo
	Network NetworkInfo
}

func PrintSystemInfo() {
	sysInfo, err := GetSystemInfo()
	if err != nil {
		log.Printf("[%s] %s \n", utilities.CreateColorString("Error", color.FgHiRed), err)
		return
	}

	log.Printf("[%s] \n", utilities.CreateColorString("=== System Information ===", color.FgBlue))
	log.Printf("[%s] %s (%s) %s \n", utilities.CreateColorString("OS", color.FgYellow), strings.Title(sysInfo.Host.Platform), sysInfo.Host.PlatformVersion, sysInfo.Host.KernelArch)
	log.Printf("[%s] %s (%d cores) %d Mhz \n", utilities.CreateColorString("CPU", color.FgCyan), sysInfo.Cpu.ModelName, sysInfo.Cpu.Cores, int(sysInfo.Cpu.Frequency))
	if sysInfo.Gpu.Name != "" {
		log.Printf("[%s] %s (%s) \n", utilities.CreateColorString("GPU", color.FgMagenta), sysInfo.Gpu.Name, sysInfo.Gpu.Vendor)
	}
	log.Printf("[%s] Installed: %d MB | Used: %d MB | Free: %d MB \n", utilities.CreateColorString("RAM", color.FgRed), int(float64(sysInfo.Memory.Total)/math.Pow(1024, 2)), int(float64(sysInfo.Memory.Used)/math.Pow(1024, 2)), int(float64(sysInfo.Memory.Free)/math.Pow(1024, 2)))

	if sysInfo.Network.InternetConnected {
		log.Printf("[%s] %s (%s) \n", utilities.CreateColorString("Internet", color.FgHiCyan), utilities.CreateColorString("Connected", color.FgHiGreen), sysInfo.Network.PublicIP)
	} else {
		log.Printf("[%s] %s \n", utilities.CreateColorString("Internet", color.FgHiCyan), utilities.CreateColorString("Not Connected", color.FgHiRed))
	}

}

func GetSystemInfo() (SystemInfo, error) {

	c, err := cpu.Info()

	if err != nil {
		log.Printf("[%s] %s \n", utilities.CreateColorString("Error", color.FgHiRed), err)
		return SystemInfo{}, err
	}

	h, err := host.Info()
	if err != nil {
		log.Printf("[%s] %s \n", utilities.CreateColorString("Error", color.FgHiRed), err)
		return SystemInfo{}, err
	}

	p, err := cpu.Percent(time.Millisecond*100, false)
	if err != nil {
		log.Printf("[%s] %s \n", utilities.CreateColorString("Error", color.FgHiRed), err)
		return SystemInfo{}, err
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("[%s] %s \n", utilities.CreateColorString("Error", color.FgHiRed), err)
		return SystemInfo{}, err
	}

	_gpu := true
	gpu, err := ghw.GPU()
	if err != nil {
		log.Printf("[%s] %s \n", utilities.CreateColorString("Error", color.FgHiRed), err)
		_gpu = false
		//return SystemInfo{}, err
	}

	_connected := false
	publicIp := ""

	resp, errHttp := http.Get("https://api.ipify.org")
	if errHttp != nil {
		_connected = false
	} else {
		_connected = true
		body, _ := ioutil.ReadAll(resp.Body)
		publicIp = string(body)
	}

	gpuInf := GpuInfo{
		Name:   "",
		Vendor: "",
	}


	if _gpu && len(gpu.GraphicsCards) > 0 && gpu.GraphicsCards[0].DeviceInfo != nil{
		gpuInf.Name = gpu.GraphicsCards[0].DeviceInfo.Product.Name
		gpuInf.Vendor = gpu.GraphicsCards[0].DeviceInfo.Vendor.Name
	}



	systemInfo := SystemInfo{
		Host: HostInfo{
			Platform:        h.Platform,
			PlatformVersion: h.PlatformVersion,
			KernelArch:      h.KernelArch,
			Uptime:          h.Uptime,
		},
		Cpu: CpuInfo{
			ModelName: c[0].ModelName,
			Frequency: c[0].Mhz,
			Cores:     len(c),
			Usage:     p[0],
		},
		Gpu: gpuInf,
		Memory: MemoryInfo{
			Total:       v.Total,
			Free:        v.Free,
			Used:        v.Total - v.Free,
			UsedPercent: v.UsedPercent,
		},
		Network: NetworkInfo{
			InternetConnected: _connected,
			PublicIP:          publicIp,
		},
	}



	return systemInfo, nil
}
