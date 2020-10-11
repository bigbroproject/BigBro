<template>
    <div class="col-sm-6 col-md-4 col-lg-3 col-xl-2">
        <div class="card">
            <div class="card-body">
                <div class="text-muted text-right mb-4">
                    <i :class="[textClass, faceClass]" style="font-size: 5em;"></i>
                </div>
                <div class="text-value-lg">{{serviceName}}</div><small class="text-muted text-uppercase font-weight-bold"> {{message}}</small>
                <div class="progress progress-xs mt-3 mb-0">
                    <div class="progress-bar" :class="bgClass" role="progressbar" :style="'width:'+ percentage+'%'" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100"></div>
                </div>
            </div>
            <div class="card-footer px-3 py-2"><a class="btn-block text-muted d-flex justify-content-between align-items-center" href="#"><span class="small font-weight-bold">View details</span>
                <i class="cil-arrow-circle-right" > </i>
            </a></div>
        </div>
    </div>
</template>

<script>

/*
* status : success, warning, danger
*
* */
export default {
name: "ServiceStatusWidget",
    data () {
        return {
            percentage: 0,
            textClass : "text-danger",
            faceClass : "cil-mood-very-bad",
            bgClass : "bg-gradient-danger",
            message: ""
        }
    },
    props:{
        serviceName:String,
        totalOnline:Number,
        totalProtocol:Number
    },
    computed: {

    },
    mounted() {
        this.percentage = (Math.round( (this.totalOnline / this.totalProtocol) * 100));
        this.totalOffline = (this.totalProtocol - this.totalOnline)

        if (this.percentage <= 20) {
            this.textClass ="text-danger";
            this.faceClass ="cil-mood-very-bad";
            this.bgClass ="bg-gradient-danger";
            this.message = !this.totalOnline ? "All services offline" : (this.totalOffline)+"/"+this.totalProtocol+" protocols offline";
        } else if (this.percentage >= 100) {
            this.textClass ="text-success";
            this.faceClass ="cil-mood-good";
            this.bgClass ="bg-gradient-success";
            this.message=this.totalProtocol+" protocols online"
        } else {
            this.textClass ="text-warning";
            this.faceClass ="cil-meh";
            this.bgClass ="bg-gradient-warning";
            this.message = !this.totalOnline ? "All services offline" : (this.totalOffline)+"/"+this.totalProtocol+" protocols offline";
        }
    }
}
</script>

<style scoped>

</style>