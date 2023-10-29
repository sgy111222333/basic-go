wrk.method = "GET"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["User-Agent"] = "PostmanRuntime/7.32.3"
-- 记得修改这个，你在登录页面登录一下，然后复制一个过来这里
wrk.headers["Authorization"] = "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTg1ODI0NDUsIlVpZCI6MjgyMTUsIlVzZXJBZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMTguMC4wLjAgU2FmYXJpLzUzNy4zNiJ9.WDLmoqFrzUEQZTImi8P0WJh27Eds1F09XoX8HyM3tYOWtJHctPZ5XDI4e20W5Btx0QDmcCp2Ag5lQbhBzobs8g"


function response(status, headers, body, request)
    print("Response Status: " .. status)
    --打印响应Headers
    local headers_str = "{"
    for k, v in pairs(headers) do
        headers_str = headers_str .. k .. ": " .. v .. ", "
    end
    headers_str = headers_str:sub(1, -3) .. "}"
    print("Headers: " .. headers_str)
    print("Response body: " .. body)
end