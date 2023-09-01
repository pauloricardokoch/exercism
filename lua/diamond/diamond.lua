local diamond = {}
diamond.aByte = string.byte('A')

diamond.setDiff = function(letter)
    diamond.diff = string.byte(letter) - diamond.aByte
end

diamond.line = function(i)
    local pad, fill, out, letter = '', '', '', ''
    letter = string.char(diamond.aByte + i)
    pad = string.rep(' ', diamond.diff - i)
    fill = string.rep(' ', i + (i - 1))

    if i == 0 then
        return out .. pad .. letter .. pad .. '\n'
    end
    return out .. pad .. letter .. fill .. letter .. pad .. '\n'
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
