local Hamming = {}

function Hamming.compute(a, b)
    local diff = 0

    if #a ~= #b then
        return -1
    end

    for i = 1, #a do
        if string.sub(a, i, i) ~= string.sub(b, i, i) then
            diff = diff + 1
        end
    end
    return diff
end

return Hamming
