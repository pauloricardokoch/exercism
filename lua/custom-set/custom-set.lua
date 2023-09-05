local Set = {}

function Set:is_empty()
    return #self.elements == 0
end

function Set:contains(x)
    for _, v in ipairs(self.elements) do
        if v == x then
            return true
        end
    end
    return false
end

function Set:is_subset(s)
    if #self.elements == 0 then
        return true
    end

    if #s.elements == 0 then
        return false
    end

    for _, v in ipairs(self.elements) do
        if not s:contains(v) then
            return false
        end
    end
    return true
end

function Set:is_disjoint(s)
    for _, v in ipairs(s.elements) do
        if self:contains(v) then
            return false
        end
    end
    return true
end

function Set:new(...)
    local arg = { ... }
    self.__index = self
    return setmetatable({ elements = arg }, self)
end

function Set:union(s)
    for _, v in ipairs(s.elements) do
        if not self:contains(v) then
            table.insert(self.elements, v)
        end
    end
    return self
end

function Set:equals(s)
    if #s.elements ~= #self.elements then
        return false
    end

    for _, v in ipairs(s.elements) do
        if not self:contains(v) then
            return false
        end
    end
    return true
end

function Set:difference(s)
    for i = #self.elements, 1, -1 do
        if s:contains(self.elements[i]) then
            table.remove(self.elements, i)
        end
    end
    return self
end

function Set:intersection(s)
    for i = #self.elements, 1, -1 do
        if not s:contains(self.elements[i]) then
            table.remove(self.elements, i)
        end
    end
    return self
end

function Set:add(x)
    if not self:contains(x) then
        table.insert(self.elements, x)
    end
    return self
end

return function(...)
    return Set:new(...)
end
