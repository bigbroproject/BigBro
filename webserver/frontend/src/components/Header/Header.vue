<template>

    <div class="p-panel p-component" >
        <div class="p-panel-header">


            <div class="p-grid p-ai-center vertical-container nested-grid">
                <div class="p-col-2">
                    <img alt="logo" src="../../assets/primevue-logo.png" height="40" class="p-mr-2">
                </div>
                <div class="p-col">
                    <div class="p-grid p-ai-center vertical-container">
                        <div class="p-col p-col-fixed" style="width:50px">
                            <i class="pi pi-spin pi-spinner" style="fontSize: 2rem"></i>
                        </div>
                        <div class="p-col">

                            <ProgressBar :value="10" />
                        </div>
                    </div>
                </div>
                <div class="p-col">
                    <div class="p-grid p-ai-center ">
                        <div class="p-col p-col-fixed" style="width:50px">
                            <i class="pi pi-spin pi-spinner" style="fontSize: 2rem"></i>
                        </div>
                        <div class="p-col">

                            <ProgressBar :value="10" />
                        </div>
                    </div>
                </div>
            </div>


        </div>
    </div>

</template>

<script>
import Menubar from 'primevue/menubar';
import ProgressBar from 'primevue/progressbar';

export default {

name: "Header",
    components:{Menubar, ProgressBar}
}
</script>

<style scoped>

</style>