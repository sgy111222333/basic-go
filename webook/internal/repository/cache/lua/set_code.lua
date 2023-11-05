local key = KEYS[1]
local cntKey = key .. ":cnt"
local val = ARGV[1]
local ttl = tonumber(redis.call("ttl", key))
if ttl == -1 then
    -- redis取ttl得到-1表示没有过期时间, 返回-2给golang
    return -2
elseif ttl == -2 or ttl < 540 then
    -- redis取ttl得到-2表示已经过期, ttl小于540表示发送验证码已超过1分钟
    -- 于是设置一个新的验证码, 过期时间为10分钟
    redis.call("set", key, val)
    redis.call("expire", key, 600)
    -- 限制一个key在十分钟之内只能被验证3次
    redis.call("set", cntKey, 3)
    redis.call("expire", key, 600)
    return 0
else
    -- 发送太频繁, 返回-1给golang
    return -1
end