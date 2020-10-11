<template>

        <SystemStats ></SystemStats>
        <ServicesList />

</template>

<script>

import Header from '../components/Header/Header.vue'
import ServicesList from '../components/ServicesList/ServicesList.vue'
import SystemStats from '../components/SystemStats'

export default {
name: "Dashboard",
    components: {
        Header,
        ServicesList,
        SystemStats
    }
}
</script>

<style scoped>

</style>