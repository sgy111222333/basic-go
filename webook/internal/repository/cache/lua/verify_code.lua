local key = KEYS[1]
local cntKey = key .. ":cnt"
-- 用户输入的验证码
local expectedKey = ARGV[1]

local cnt = tonumber(redis.call("get", cntKey))
local code = redis.call("get", key)

if cnt == nil or cnt <= 0 then
    --验证次数耗尽
    return -1
end

if code == expectedKey then
    --如果验证码正确, 就将可验证次数设为0, 并返回正确
    redis.call("set", cntKey, 0)
    return 0
else
    -- 不想等, 输入错误
    redis.call("decr", cntKey)
    return -2
end