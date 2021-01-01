--
-- Created by IntelliJ IDEA.
-- User: petru
-- Date: 2020/12/26
-- Time: 22:52
-- To change this template use File | Settings | File Templates.
--

local m = require("modules")

routes = {
    index = "/",
    user = "/user"
}

m.startServer('8080')
