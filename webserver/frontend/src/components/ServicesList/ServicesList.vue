<template>
    <div class="container-fluid">
        <div class="row">
            <ServiceStatusWidget service-name="Pippo puppo" :total-online="2" :total-protocol="10"></ServiceStatusWidget>
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
                            <tr class="table-light">
                                <th scope="col">Facebook</th>
                                <th scope="col">Protocol</th>
                                <th scope="col">Status</th>
                                <th scope="col">Server</th>
                                <th scope="col">Port</th>
                            </tr>
                            <tr>
                                <td></td>
                                <td>HTTP</td>
                                <td><i class="cil-happy text-success" style="font-size: 1.2em;"></i> Online</td>
                                <td>80</td>
                                <td>ICMP</td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import CustomerService from '../../services/CustomerService';
import ServiceStatusWidget from '../../components/ServiceStatusWidget';
import Card from 'primevue/card';

export default {
    data() {
        return {
            customers: null,
            expandedRowGroups: null
        }
    },
    components:{
        Card,
        ServiceStatusWidget
    },
    customerService: null,
    created() {
        this.customerService = new CustomerService();
    },
    mounted() {
        this.customerService.getCustomersMedium().then(data => this.customers = data);
    },
    methods: {
        onRowGroupExpand(event) {
            this.$toast.add({severity: 'info', summary: 'Row Group Expanded', detail: 'Value: ' + event.data, life: 3000});
        },
        onRowGroupCollapse(event) {
            this.$toast.add({severity: 'success', summary: 'Row Group Collapsed', detail: 'Value: ' + event.data, life: 3000});
        },
        calculateCustomerTotal(name) {
            let total = 0;
            if (this.customers) {
                for (let customer of this.customers) {
                    if (customer.representative.name === name) {
                        total++;
                    }
                }
            }
            return total;
        }
    }
}
</script>

<style scoped>

</style>