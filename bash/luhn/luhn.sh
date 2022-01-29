#!/usr/bin/env bash
num=$(echo $* | tr -d ' ' | rev 2>/dev/null)

[[ ${#num} -le 1 ]] && { echo false; exit 0; }
[[ ! "$num" =~ ^[0-9]+$ ]] && { echo false; exit 0; }

sum=0
for ((i=0; i<$((${#num})); i++ )); do
    digit=${num:$i:1}
    if [ $((i%2)) -ne 0 ]; then
        digit=$((digit*2))
        if [ "$digit" -gt 9 ]; then
            digit=$((digit-9))
        fi
    fi
    sum=$((sum+digit))
done

[[ $((sum%10)) -eq 0 ]] && echo true || echo false