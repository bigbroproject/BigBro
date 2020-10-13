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
            socket:io("http://localhost:8181")
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

<style scoped>
/*#app {*/
/*  font-family: 'Avenir', Helvetica, Arial, sans-serif;*/
/*  -webkit-font-smoothing: antialiased;*/
/*  -moz-osx-font-smoothing: grayscale;*/
/*  color: #2c3e50;*/
/*  margin-top: 60px;*/
/*}*/

/*.app-container {*/
/*  text-align: center;*/
/*}*/

/*body #app .p-button {*/
/*  margin-left: .2em;*/
/*}*/

/*form {*/
/*  margin-top: 2em;*/
/*}*/
</style>