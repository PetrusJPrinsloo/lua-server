--
-- Created by IntelliJ IDEA.
-- User: petrus
-- Date: 2021/01/01
-- Time: 09:10
-- To change this template use File | Settings | File Templates.
--
local m = require("services")
mysql = require('mysql')

function get()
--    c = mysql.new()
--    ok, err = c:connect({ host = '127.0.0.1', port = 3306, database = 'test', user = 'lua', password = '@Test1234' })
--    if ok then
--        res, err = c:query('SELECT * FROM user')
--        for key, value in pairs(res) do
--            print(value.name)
--            print(value.last_name)
--            print(value.email)
--        end
--    end
--
--    if not ok then
--        dump(err)
--    end

    for key, value in pairs(POST_DATA) do
        print(key)
        print(value)
    end

    local b = "<h1>Hello There</h1>" ..
            "<p>here's a paragraph</p>" ..
            "<button>Button</button><br>" ..
            "<a href='/'>Some Link</a>"
    return b
end

response.body = get()