import axios from 'axios'
import {ServiceData} from "../models/ServiceData"
const BACKEND_SERVER = "http://localhost:8181"

export default class SystemService {

    getServicesData() {
        return axios.get(BACKEND_SERVER + '/api/services').then(res => {
            const servicesData = [];
            for (const key of Object.keys(res.data)) {
                servicesData.push(ServiceData.fromJson(res.data[key]));
            }
            return servicesData;

        });
    }

}