import axios from "axios";

const instance = axios.create({
    // 这边记得修改你对应的配置文件
    // baseURL: "http://192.168.6.155:8080",
    baseURL: "http://172.20.146.154:8080",
    // baseURL: "http://172.16.210.86:80",
    withCredentials: true
})


instance.interceptors.response.use(function (resp) {
    const newToken = resp.headers["x-jwt-token"]
    const newRefreshToken = resp.headers["x-refresh-token"]
    console.log("resp headers", resp.headers)
    if (newToken) {
        localStorage.setItem("token", newToken)
    }
    if (newRefreshToken) {
        localStorage.setItem("refresh_token", newRefreshToken)
    }
    if (resp.status == 401) {
        window.location.href = "/users/login"
    }
    return resp
}, (err) => {
    console.log(err)
    if (err.response.status == 401) {
        window.location.href = "/users/login"
    }
    return err
})

// 在这里让每一个请求都加上 authorization 的头部
instance.interceptors.request.use((req) => {
    const token = localStorage.getItem("token")
    req.headers.setAuthorization("Bearer " + token, true)
    return req
}, (err) => {
    console.log(err)
})

export default instance