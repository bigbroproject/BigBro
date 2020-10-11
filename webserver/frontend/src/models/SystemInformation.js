export default class SystemInformation {
    constructor(hostInfo, cpuInfo, gpuInfo, memoryInfo, networkInfo) {

        this.host = {
            platform:hostInfo.Platform,
            platformVersion:hostInfo.PlatformVersion,
            kernelArch:hostInfo.KernelArch,
            upTime: hostInfo.UpTime
        };
        this.cpu = {
            modelName:cpuInfo.ModelName,
            frequency:cpuInfo.Frequency,
            cores: cpuInfo.Cores,
            usage: Math.round(cpuInfo.Usage)
        };

        this.gpu = {
            name : gpuInfo.Name,
            vendor : gpuInfo.Vendor
        };

        this.memory = {
            total : memoryInfo.Total,
            free: memoryInfo.Free,
            used: memoryInfo.Used,
            usedPercent : Math.round(memoryInfo.UsedPercent)
        };

        this.network = {
            internetConnected : networkInfo.InternetConnected,
            publicIP : networkInfo.PublicIP
        }


    }
}