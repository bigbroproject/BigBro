import {createRouter, createWebHashHistory} from 'vue-router'
import Dashboard from '../views/Dashboard.vue';

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: '/',
            name: 'dashboard',
            component: Dashboard
        },
        { path: "/:pathMatch(.*)*", redirect: '/' }
    ]
});

export default router;