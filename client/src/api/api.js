import axios from "axios";
let config = require("../config/config.js");
import Store from "@/store/store";

function isEmptyObject(obj) {
    for (var key in obj) {
        return false;
    }
    if (null != sessionStorage.getItem("token")) {
        Store.state.user = { "token": sessionStorage.getItem("token") }
        return false
    }
    return true;
}
/**
 * Post请求（发送json）
 * @param {String} path 路径
 * @param {Object} params 参数
 * @return {Promise} 请求Promise
 */
function postJSON(path, params) {
    if (isEmptyObject(Store.state.user)) {
        return axios({
            method: "post",
            url: config.url + path,
            data: params,
            headers: {
                "Content-Type": "application/json"
            }
        });
    } else {
        return axios({
            method: "post",
            url: config.url + path,
            data: params,
            headers: {
                "Content-Type": "application/json",
                "Token": Store.state.user.token
            }
        });
    }
}
export default {
    postJSON,
    isEmptyObject
};
