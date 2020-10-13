import axios from 'axios'
import SystemInformation from "../models/SystemInformation";
const BACKEND_SERVER = "http://localhost:8181"

export default class SystemService {

    getSystemInformation() {
        return axios.get(BACKEND_SERVER + '/api/system').then(res => {
            const sysInfo = new SystemInformation(res.data.Host, res.data.Cpu, res.data.Gpu, res.data.Memory, res.data.Network);
            return sysInfo;

        });
    }

}