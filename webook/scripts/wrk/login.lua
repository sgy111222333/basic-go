wrk.method="POST"
wrk.headers["Content-Type"] = "application/json"
-- 这个要改为你的注册的数据
wrk.body='{"email":"111111fad6097@qq.com", "password": "Sunguangyu3.14"}'


function response(status, headers, body, request)
    print("Response Status: " .. status)
    --打印响应Headers
    --local headers_str = "{"
    --for k, v in pairs(headers) do
    --    headers_str = headers_str .. k .. ": " .. v .. ", "
    --end
    --headers_str = headers_str:sub(1,-3) .. "}"
    --print("Headers: " .. headers_str)
    print("Response body: " .. body)
end