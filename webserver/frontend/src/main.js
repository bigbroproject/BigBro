import { createApp } from 'vue'
import App from './App.vue';
import router from './router/router';
import { createStore } from 'vuex'
import VueScrollTo from 'vue-scrollto'



import "@coreui/coreui/dist/js/coreui.bundle.min.js"
import "@coreui/coreui/dist/css/coreui.min.css";
import "@coreui/icons/css/all.min.css";
import 'sweetalert2/dist/sweetalert2.min.css';

const app = createApp(App);
const store = createStore({
    state () {
        return {
            services: []
        }
    },
    mutations: {
        updateService (state, obj) {

            for (const index in state.services) {
                if (state.services[index].name === obj.sData.name){
                    state.services[index] = obj.sData;
                }
            }
        }
    }
})

app.use(VueScrollTo)
app.use(router)
app.use(store)
app.mount('body')
