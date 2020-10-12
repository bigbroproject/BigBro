(function (window, document, $, io, Swal) {
    'use strict';


    class Application {
        constructor() {
            this.intervalSystemInfo = null
            this.socket = io("http://localhost:8181");

            this.socket.on('connect', () => {
                console.log(this.socket.connected); // true
            });

            this.socket.on('serviceChange', (data) => {
                //console.log(data); // true
                this.updateServiceData(data)
            });


            this.socket.on('disconnect', () => {
                console.log(this.socket.connected); // false
            });

            $('[data-toggle="tooltip"]').tooltip()
        }

        getServices() {
            const $tableList = $("#services-list-table");
            $.get("/api/services", (response) => {
                $tableList.empty();

                for (const serviceKey of Object.keys(response)) {
                    const service = response[serviceKey];

                    for (const protocolData of service.Protocols) {

                        const $tableRow = this.createServiceRow(service, protocolData);
                        $tableList.append($tableRow);

                    }

                }
            })
        }

        getSystemInformation() {
            $.get("/api/system", (response) => {
                $("#cpu-usage").html(Math.round(response.Cpu.Usage) + "%");
                $("#ram-usage").html(Math.round(response.Memory.UsedPercent) + "%");

                $("#os-info").html(response.Host.Platform[0].toUpperCase() + response.Host.Platform.slice(1) + " (" + response.Host.PlatformVersion + ") " + response.Host.KernelArch);
                $("#cpu-info").html(response.Cpu.ModelName + " (" + response.Cpu.Cores + " cores) " + response.Cpu.Frequency + " Mhz");
                if (response.Gpu.Name !== "") {
                    $("#gpu-info").html(response.Gpu.Name + " (" + response.Gpu.Vendor + ")");
                } else {
                    $("#gpu-info").html("No GPUs installed");
                }
                $("#ram-info").html((Math.round(response.Memory.Total / Math.pow(1024, 2))) + " MB");
                if (response.Network.InternetConnected) {
                    $("#internet-info").html("Connected" + " (" + response.Network.PublicIP + ")");
                } else {
                    $("#internet-info").html("Not Connected");
                }


            })
        }

        startSystemInfoInterval() {
            this.intervalSystemInfo = setInterval(this.getSystemInformation, 10000)
        }

        stopSystemInfoInterval() {
            clearInterval(this.intervalSystemInfo)
        }

        updateServiceData(service) {

            for (const protocolData of service.Protocols) {
                const keyMap = service.Name + protocolData.Protocol.Server + protocolData.Protocol.Type + protocolData.Protocol.Port;

                let errStatus = "";
                if (protocolData.Err === "<<pending>>") {
                    // pending
                    errStatus = "<i class=\"fa fa-circle font-small-3 text-warning mr-50\"></i>Pending"
                } else {
                    errStatus = protocolData.Err ? "<i class=\"fa fa-circle font-small-3 text-danger mr-50\"></i>Offline <i  class=\"fa fa-question-circle font-small-3 text-danger error-tooltip ml-50\" data-error='" + protocolData.Err + "' ></i>" : "<i class=\"fa fa-circle font-small-3 text-success mr-50\"></i>Online"
                }

                if ($("#services-list-table tr[data-service-protocol-server='" + keyMap + "']").length <= 0) {

                    const $tableList = $("#services-list-table");
                    const $tableRow = this.createServiceRow(service, protocolData);
                    $tableList.append($tableRow);

                } else {
                    const $statusTd = $("#services-list-table tr[data-service-protocol-server='" + keyMap + "'] .status")
                    $statusTd.html(errStatus)
                    $statusTd.off("click", () => {
                        Swal.fire({
                            type: 'error',
                            title: 'Error on Checking service',
                            text: JSON.stringify(protocolData.Err),
                            confirmButtonClass: 'btn btn-primary',
                            buttonsStyling: false,
                        })
                    })

                    if (protocolData.Err != null) {
                        console.log("Devo mettere evento")
                        $("#services-list-table tr[data-service-protocol-server='" + keyMap + "'] .error-tooltip").attr("data-content", JSON.stringify(protocolData.Err))
                        $statusTd.on("click", () => {

                            Swal.fire({
                                type: 'error',
                                title: 'Error on Checking service',
                                text: JSON.stringify(protocolData.Err),
                                confirmButtonClass: 'btn btn-primary',
                                buttonsStyling: false,
                            })
                        })
                    }
                }

            }

        }

        createServiceRow(service, protocolData) {

                let errStatus = "";

                if (protocolData.Err === "<<pending>>") {
                    // pending
                    errStatus = "<i class=\"fa fa-circle font-small-3 text-warning mr-50\"></i>Pending"
                } else {
                    errStatus = protocolData.Err ? "<i class=\"fa fa-circle font-small-3 text-danger mr-50\"></i>Offline <i  class=\"fa fa-question-circle font-small-3 text-danger error-tooltip ml-50\" data-error='" + protocolData.Err + "' ></i>" : "<i class=\"fa fa-circle font-small-3 text-success mr-50\"></i>Online"
                }

                const $service = {
                    name: service.Name,
                    status: errStatus,
                    error: protocolData.Err,
                    protocol: protocolData.Protocol
                }

                const $tableRow = $("<tr></tr>")
                const keyMap = $service.name + $service.protocol.Server + $service.protocol.Type + $service.protocol.Port
                $tableRow.attr("data-service-protocol-server", keyMap)
                $tableRow.append($("<td></td>").html($service.name));


                const $statusField = $("<td></td>")
                $statusField.addClass("status")

                if (protocolData.Err != null) {

                    $statusField.attr("data-toggle", "popover")
                    $statusField.attr("data-trigger", "focus")
                    $statusField.attr("data-content", "Boh")

                    $statusField.on("click", () => {

                        Swal.fire({
                            type: 'error',
                            title: 'Error on Checking service',
                            text: JSON.stringify(protocolData.Err),
                            confirmButtonClass: 'btn btn-primary',
                            buttonsStyling: false,
                        })
                    })

                }
                $tableRow.append($statusField.html($service.status));

                $tableRow.append($("<td></td>").addClass("p-1").html($service.protocol.Type.toUpperCase()));
                $tableRow.append($("<td></td>").html($service.protocol.Server));
                $tableRow.append($("<td></td>").html($service.protocol.Port ? $service.protocol.Port : "Default Port"));
                $tableRow.append($("<td></td>"));
                $tableRow.append($("<td></td>"));

                return $tableRow



        }
    }

    /*
    NOTE:
    ------
    PLACE HERE YOUR OWN JAVASCRIPT CODE IF NEEDED
    WE WILL RELEASE FUTURE UPDATES SO IN ORDER TO NOT OVERWRITE YOUR JAVASCRIPT CODE PLEASE CONSIDER WRITING YOUR SCRIPT HERE.  */

    $(window).on("load", () => {
        let app = new Application();
        app.getSystemInformation();
        app.startSystemInfoInterval();
        app.getServices();
        app.getServices();
    })

})(window, document, jQuery, io, Swal);