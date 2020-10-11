import { createApp } from 'vue'
import App from './App.vue';
import router from './router/router';

import "@coreui/coreui/dist/js/coreui.bundle.min.js"
import "@coreui/coreui/dist/css/coreui.min.css";
import "@coreui/icons/css/all.min.css";

const app = createApp(App);

app.use(router)

app.mount('body')
