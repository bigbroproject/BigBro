<template>
    <div class="c-wrapper">
        <Header></Header>
        <div class="c-body">
            <main class="c-main">
                <!-- Main content here -->

                <router-view></router-view>

            </main>
        </div>
        <footer class="c-footer">
            <!-- Footer content here -->
        </footer>
    </div>
</template>

<script>
import {ServiceData} from "./models/ServiceData"
import Header from "./components/Header/Header"
import io from 'socket.io-client';
export default {
    components: {Header},
    data() {
        return {
            message: null,
            text: null,
            socket:io()
        }
    },
    created() {
        this.socket.on('connect', () => {
            console.log("Connected!"); // true
        });

        this.socket.on('serviceChange', (data) => {

            this.$store.commit({
                type: 'updateService',
                sData: ServiceData.fromJson(data)
            })
        });


        this.socket.on('disconnect', () => {
            console.log("Disconnected!"); // false
        });
    },
    mounted() {
    },
    methods: {
    },

}
</script>

<style >
@import url("~@coreui/coreui/dist/css/coreui.min.css");
@import url("~@coreui/icons/css/all.min.css");
@import url("~sweetalert2/dist/sweetalert2.min.css");


</style>