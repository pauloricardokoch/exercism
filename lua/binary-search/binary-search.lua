return function(array, target)
    local iStart, iEnd, iMid = 1, #array, 0

    while iStart <= iEnd do
        iMid = math.floor((iStart + iEnd) / 2)
        local val = array[iMid]
        if target == val then
            return iMid
        end

        if val > target then
            iEnd = iMid - 1
        else
            iStart = iMid + 1
        end
    end
    return -1
end
