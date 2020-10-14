class ProtocolData {

    constructor(type, server, port, interval, customs, err) {
        this.type = type;
        this.server = server;
        this.port = port >0 ? port : "Default port";
        this.interval = interval;
        this.customs = customs;
        this.errStatus = err;
    }

    getError(){
        return this.errStatus;
    }

    static fromJson(json){
        const protData = new ProtocolData(
            json.Protocol.Type,
            json.Protocol.Server,
            json.Protocol.Port,
            json.Protocol.interval,
            json.Protocol.Customs,
            json.Err
        );

        return protData;

    }
}

class ServiceData {
    constructor(name, protocols) {

        this.name = name;
        this.protocols = protocols;
    }


    static fromJson(json){
        const serviceData = new ServiceData(json.Name, []);

        for (const _protocolData of json.Protocols) {
            const protData = ProtocolData.fromJson(_protocolData);
            serviceData.protocols.push(protData);
        }

        return serviceData;
    }
}

export {
    ProtocolData,
    ServiceData,
}