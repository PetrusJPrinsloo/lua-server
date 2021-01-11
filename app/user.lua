--
-- Created by IntelliJ IDEA.
-- User: petrus
-- Date: 2021/01/01
-- Time: 09:10
-- To change this template use File | Settings | File Templates.
--
local m = require("services")
mysql = require('mysql')
json = require('json')

function get()

    if POST_DATA ~= nil then
        for key, value in pairs(POST_DATA) do
            print("Key: " .. key .. ",\tValue: " .. value)
        end

        c = mysql.new()
        ok, err = c:connect({ host = '127.0.0.1', port = 3306, database = 'test', user = 'lua', password = '@Test1234' })

        if ok then
            res, err = c:query('INSERT INTO user(name, last_name, email) VALUES("' ..
                POST_DATA.name .. '", "' ..
                POST_DATA.last_name .. '", "' ..
                POST_DATA.email .. '")')

            if err ~= nil then
                print(err)
            end
        end

        mysql.close()

        print(json.endcode(POST_DATA))

        if not ok then
            dump(err)
        end
    end

    local b = "<h1>Hello There</h1>" ..
            "<p>here's a paragraph</p>" ..
            "<form method='post' action='/user'>" ..
            "<input type='text' name='name'>" ..
            "<input type='text' name='last_name'>" ..
            "<input type='text' name='email'>" ..
            "<input type='submit'>" ..
            "</form>" ..
            "<button>Button</button><br>" ..
            "<a href='/'>Some Link</a>"
    return b
end

response.body = get()