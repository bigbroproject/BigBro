<template>
    <div class="container-fluid">
        <div class="row">
            <div class="col-md-12">
                <div class="card">
                    <div class="card-header">System Information</div>

                    <div class="card-body">

                        <template v-if="systemInfo">



                            <div class="row">
                                <div class="col-sm-3">
                                    <div class="c-callout c-callout-info"><small class="text-muted">Operative
                                        System</small>
                                        <div class="text-value">
                                            {{ systemInfo.host.platform }} {{ systemInfo.host.platformVersion }}
                                            ({{ systemInfo.host.kernelArch }})
                                        </div>
                                    </div>

                                </div>
                                <div class="col-sm-3">
                                    <div class="c-callout c-callout-danger"><small class="text-muted">CPU</small>
                                        <div class="text-value">{{ systemInfo.cpu.modelName }}
                                            {{ systemInfo.cpu.cores }} cores ({{ systemInfo.cpu.frequency }} Mhz)
                                        </div>
                                    </div>
                                </div>
                                <div class="col-sm-3">
                                    <div class="c-callout c-callout-warning"><small class="text-muted">RAM</small>
                                        <div class="text-value">
                                            Installed {{ bytesToHuman(systemInfo.memory.total, 2) }} <br>
                                            <small>Used {{ bytesToHuman(systemInfo.memory.used, 2) }} |
                                                Free {{ bytesToHuman(systemInfo.memory.free, 2) }}</small>
                                        </div>
                                    </div>
                                </div>

                                <div class="col-sm-3">
                                    <div class="c-callout c-callout-success"><small class="text-muted">GPU</small>
                                        <div class="text-value">
                                            <template v-if="systemInfo.gpu">
                                                {{ systemInfo.gpu.name }} ({{ systemInfo.gpu.vendor }})
                                            </template>
                                            <template v-else="systemInfo.gpu">
                                                No videocard installed
                                            </template>
                                        </div>
                                    </div>
                                </div>

                            </div>
                            <div class="row">

                                <div class="col-sm-6">

                                    <hr class="mt-0">
                                    <div class="progress-group">
                                        <div class="progress-group-header">
                                            <i class="cil-bolt" style="font-size: 1.5em;"></i>
                                            <div class="ml-1">CPU Usage</div>
                                            <div class="mfs-auto font-weight-bold">{{ systemInfo.cpu.usage }}%</div>
                                        </div>
                                        <div class="progress-group-bars">
                                            <div class="progress progress-xs">
                                                <div class="progress-bar bg-gradient-danger" role="progressbar"
                                                     :style="'width: '+systemInfo.cpu.usage+'%'"
                                                     :aria-valuenow="systemInfo.cpu.usage" aria-valuemin="0"
                                                     aria-valuemax="100"></div>
                                            </div>
                                        </div>
                                    </div>

                                </div>
                                <div class="col-sm-6">

                                    <hr class="mt-0">
                                    <div class="progress-group">
                                        <div class="progress-group-header">
                                            <i class="cil-memory" style="font-size: 1.5em;"></i>
                                            <div class="ml-1">Ram Usage</div>
                                            <div class="mfs-auto font-weight-bold">
                                                {{ systemInfo.memory.usedPercent }}%
                                            </div>
                                        </div>
                                        <div class="progress-group-bars">
                                            <div class="progress progress-xs">
                                                <div class="progress-bar bg-gradient-warning" role="progressbar"
                                                     :style="'width: '+systemInfo.memory.usedPercent+'%'"
                                                     :aria-valuenow="systemInfo.memory.usedPercent" aria-valuemin="0"
                                                     aria-valuemax="100"></div>
                                            </div>
                                        </div>
                                    </div>
                                </div>

                            </div>
                        </template>
                        <div class="row" v-if="!systemInfo">
                            <div class="col-12">
                                <h4>Loading System information.....</h4>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

        </div>

    </div>
</template>

<script>
import SystemService from "../services/SystemService";

export default {
    name: "SystemStats",
    data() {
        return {
            systemInfo: null,
            systemInfoInterval: null
        }
    },
    unmounted() {
        if (this.systemInfoInterval) clearInterval(this.systemInfoInterval);
    },
    created() {
        this.systemService = new SystemService();
    },
    components: {},

    mounted() {
        this.getSystemInfo();
        this.systemInfoInterval = setInterval(() => {
            this.getSystemInfo();
        }, 5000);

    },
    methods: {
        bytesToHuman(bytes, decimals = 2) {

            if (bytes === 0) return '0 Bytes';

            const k = 1024;
            const dm = decimals < 0 ? 0 : decimals;
            const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
            const i = Math.floor(Math.log(bytes) / Math.log(k));

            return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
        },
        getSystemInfo() {
            this.systemService.getSystemInformation().then((data) => {
                this.systemInfo = data;
            }).catch((err) => {
                console.error(err)
            })
        }
    }
}
</script>

<style scoped>

</style>