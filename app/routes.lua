--
-- Created by IntelliJ IDEA.
-- User: petrus
-- Date: 2020/12/26
-- Time: 22:52
-- To change this template use File | Settings | File Templates.
--

local m = require("modules")

routes = {
    {
        index = "/",
        method = "GET",

    },
    {
        user = "/user/{id}",
        method = "GET"
    }
}

m.startServer('9090')
