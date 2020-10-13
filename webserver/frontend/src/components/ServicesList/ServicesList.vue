<template>
    <div class="container-fluid">
        <div class="row">

            <template v-for="serviceData of $store.state.services" :key="serviceData.name"  v-if="$store.state.services">
                <ServiceStatusWidget :service-data="serviceData" ></ServiceStatusWidget>
            </template>

        </div>
    </div>

    <div class="container-fluid">
        <div class="row">
            <div class="col-md-12">
                <div class="card">
                    <div class="card-header">Services List</div>

                    <div class="table-responsive">
                        <table class="table table-border ">

                            <tbody>

                            <template v-for="serviceData of $store.state.services" :key="serviceData.name"  v-if="$store.state.services">

                                <tr class="table-light" :id="'service-row-'+serviceData.name" >
                                    <th scope="col">{{serviceData.name}}</th>
                                    <th scope="col">Protocol</th>
                                    <th scope="col">Status</th>
                                    <th scope="col">Server</th>
                                    <th scope="col">Port</th>
                                </tr>

                                <tr v-for="protocolData of serviceData.protocols">
                                    <td>

                                    </td>
                                    <td>{{ protocolData.type }}</td>
                                    <td>
                                        <template v-if="!protocolData.errStatus">
                                            <i class="cil-happy text-success" style="font-size: 1.2em;"></i> Online
                                        </template>
                                        <template v-else-if="protocolData.errStatus=='<<pending>>'">
                                            <i class="cil-clock text-warning" style="font-size: 1.2em;"></i> Pending
                                        </template>
                                        <template v-else>
                                            <button v-if="protocolData.errStatus" @click="showError(protocolData.errStatus)" class="btn btn-outline-danger btn-sm" type="button">
                                                <i class="cil-sad" style="font-size: 1.2em;"></i> Offline
                                            </button>

                                        </template>
                                    </td>
                                    <td>{{protocolData.server}}</td>
                                    <td>{{protocolData.port}}</td>
                                </tr>

                            </template>


                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import ServiceStatusWidget from "../ServiceStatusWidget"
import ServicesService from "../../services/ServicesService";
import Swal from 'sweetalert2'


export default {
    data() {
        return {
            expandedRowGroups: null
        }
    },
    components:{
        ServicesService,
        ServiceStatusWidget
    },
    servicesService: null,
    created() {
        this.servicesService = new ServicesService();
    },
    mounted() {
        this.servicesService.getServicesData().then(data => {
            this.$store.state.services = data
        });
    },
    methods: {
        showError(err){
            console.log(err);
            Swal.fire({
                title: 'An error as occurred!',
                html: "<pre class='json'>"+this.syntaxHighlight(JSON.stringify(err,undefined, 2))+"</pre>",
                icon: 'error',
                confirmButtonText: 'Ok'
            })
        },
        syntaxHighlight(json) {
            json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
            return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
                var cls = 'number';
                if (/^"/.test(match)) {
                    if (/:$/.test(match)) {
                        cls = 'key';
                    } else {
                        cls = 'string';
                    }
                } else if (/true|false/.test(match)) {
                    cls = 'boolean';
                } else if (/null/.test(match)) {
                    cls = 'null';
                }
                return '<span class="' + cls + '">' + match + '</span>';
            });
        }

    }
}
</script>

<style>
pre.json {
    padding: 5px; margin: 5px; text-align: left;
}
pre.json .string { color: green; }
pre.json .number { color: darkorange; }
pre.json .boolean { color: blue; }
pre.json .null { color: magenta; }
pre.json .key { color: red; }
</style>