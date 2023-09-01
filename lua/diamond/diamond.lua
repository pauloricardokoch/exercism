local diamond = {}
diamond.aByte = string.byte('A')

diamond.setDiff = function(letter)
    diamond.diff = string.byte(letter) - diamond.aByte
end

diamond.line = function(i)
    local pad, fill, out, temp = '', '', '', ''

    temp = string.char(diamond.aByte + i)
    pad = string.rep(' ', diamond.diff - i)
    fill = string.rep(' ', i + (i - 1))

    if i == 0 then
        out = out .. pad .. temp .. pad
    else
        out = out .. pad .. temp .. fill .. temp .. pad
    end
    return out .. '\n'
end

diamond.draw = function(letter)
    diamond.setDiff(letter)

    local out = ''
    for i = 0, diamond.diff do
        out = out .. diamond.line(i)
    end
    for i = diamond.diff - 1, 0, -1 do
        out = out .. diamond.line(i)
    end
    return out
end

return diamond.draw
