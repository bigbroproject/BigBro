import { createApp } from 'vue'
import App from './App.vue';
import router from './router/router';
import { createStore } from 'vuex'
import VueScrollTo from 'vue-scrollto'
import InlineSvg from 'vue-inline-svg';
import "@coreui/coreui/dist/js/coreui.bundle.min.js";

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
app.component('inline-svg', InlineSvg);

app.use(VueScrollTo);
app.use(router);
app.use(store);
app.mount('body');
