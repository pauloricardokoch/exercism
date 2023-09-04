local Set = {}

function Set:is_empty()
    return #self.elements == 0
end

function Set:contains(x)
    for i, v in ipairs(self.elements) do
        if v == x then
            return true
        end
    end
    return false
end

function Set:is_subset(s)
    for i, v in ipairs(s.elements) do
        if self:contains(v) == false then
            return false
        end
    end
    return true
end

function Set:is_disjoint()

end

function Set:new(...)
    local arg = { ... }
    self.__index = self
    return setmetatable({ elements = arg }, self)
end

return function(...)
    return Set:new(...)
end
