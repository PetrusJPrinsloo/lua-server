--
-- Created by IntelliJ IDEA.
-- User: petrus
-- Date: 2021/01/01
-- Time: 09:10
-- To change this template use File | Settings | File Templates.
--
local m = require("services")

function get()
    local b = "<h1>Hello There</h1>" ..
            "<p>here's a paragraph</p>" ..
            "<button>Button</button><br>" ..
            "<a href='/user'>Some Link</a>"
    return b
end

response.body = get()