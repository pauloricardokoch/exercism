Robot = {}
local assigned_names = {}

local function random_name()
    return string.char(math.random(65, 90),
        math.random(65, 90),
        math.random(48, 57),
        math.random(48, 57),
        math.random(48, 57))
end

local function new_name()
    local new = ''
    while true do
        new = random_name()
        if assigned_names[new] == nil then
            break
        end
    end
    assigned_names[new] = true
    return new
end

function Robot:new()
    self.__index = self
    return setmetatable({ name = new_name() }, self)
end

function Robot:reset()
    self.name = new_name()
end

return Robot
