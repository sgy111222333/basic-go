wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"

local random = math.random
local function uuid()
    local template = 'xxxxxxxx-xxxx-1xxx-yxxx-xxxxxxxxxxxx'
    return string.gsub(template, '[xy]', function(c)
        local v = (c == 'x') and random(0, 0xf) or random(8, 0xb)
        return string.format('%x', v)
    end)
end

-- 初始化
function init(args)
    -- 每个线程都有一个 cnt，所以是线程安全的
    cnt = 0
    prefix = uuid()
end

function request()
    body = string.format('{"email":"%s%d@qq.com", "password":"Hello#world123", "confirmPassword": "Hello#world123"}', prefix, cnt)
    cnt = cnt + 1
    print("Request URL: " .. wrk.path) -- 打印请求的url
    print("Request Body: " .. body) -- 打印请求的body
    return wrk.format('POST', wrk.path, wrk.headers, body)
end

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
