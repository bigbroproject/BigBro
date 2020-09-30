(function (window, document, $) {
  'use strict';


  class Application {
    constructor(){
      this.intervalSystemInfo = null
    }

    getServices(){
      const $tableList = $("#services-list-table");
      $.get("/api/services",(response) => {
        $tableList.empty();

        for (const serviceKey of Object.keys(response))  {
          const service = response[serviceKey];

          for(const protocolData of service.Protocols){

            let errStatus = "";

            console.log(service)
            if(protocolData.Err === "<<pending>>"){
              // pending
              errStatus = "<i class=\"fa fa-circle font-small-3 text-warning mr-50\"></i>Pending"
            } else {
              errStatus = protocolData.Err ? "<i class=\"fa fa-circle font-small-3 text-danger mr-50\"></i>Offline" : "<i class=\"fa fa-circle font-small-3 text-success mr-50\"></i>Online"
            }

            const $service = {
              name : service.Name,
              status : errStatus,
              error : protocolData.Err,
              protocol : protocolData.Protocol
            }

            const $tableRow = $("<tr></tr>")
            $tableRow.append($("<td></td>").html($service.name));
            $tableRow.append($("<td></td>").html($service.status));
            $tableRow.append($("<td></td>").addClass("p-1").html($service.protocol.Type.toUpperCase()));
            $tableRow.append($("<td></td>").html($service.protocol.Server));
            $tableRow.append($("<td></td>").html($service.protocol.Port ? $service.protocol.Port : "Default Port"));
            $tableRow.append($("<td></td>"));
            $tableRow.append($("<td></td>"));

            $tableList.append($tableRow);
        //    console.log($service)


          }

        }
      })
    }
    getSystemInformation(){
      $.get("/api/system",(response) => {
        $("#cpu-usage").html( Math.round(response.Cpu.Usage) + "%");
        $("#ram-usage").html( Math.round(response.Memory.UsedPercent) + "%");

        $("#os-info").html(   response.Host.Platform[0].toUpperCase() + response.Host.Platform.slice(1)+" ("+response.Host.PlatformVersion+") "+ response.Host.KernelArch);
        $("#cpu-info").html( response.Cpu.ModelName+" ("+response.Cpu.Cores+" cores) "+response.Cpu.Frequency+" Mhz");
        if(response.Gpu.Name !== ""){
          $("#gpu-info").html( response.Gpu.Name+" ("+response.Gpu.Vendor+")");
        } else {
          $("#gpu-info").html("No GPUs installed");
        }
        $("#ram-info").html( ( Math.round(response.Memory.Total / Math.pow(1024,2))) + " MB");
        if(response.Network.InternetConnected){
          $("#internet-info").html( "Connected" + " ("+response.Network.PublicIP+")");
        } else {
          $("#internet-info").html( "Not Connected");
        }


      })
    }

    startSystemInfoInterval(){
      this.intervalSystemInfo = setInterval(this.getSystemInformation, 10000)
    }

    stoptSystemInfoInterval(){
      clearInterval(this.intervalSystemInfo)
    }
  }

  /*
  NOTE:
  ------
  PLACE HERE YOUR OWN JAVASCRIPT CODE IF NEEDED
  WE WILL RELEASE FUTURE UPDATES SO IN ORDER TO NOT OVERWRITE YOUR JAVASCRIPT CODE PLEASE CONSIDER WRITING YOUR SCRIPT HERE.  */

  $(window).on("load",  () => {
    let app = new Application();
    app.getSystemInformation();
    app.startSystemInfoInterval();
    app.getServices();
    app.getServices();
  })

})(window, document, jQuery);