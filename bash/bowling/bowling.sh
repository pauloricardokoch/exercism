#!/usr/bin/env bash

SCORE=()
FRAME=()
KNOCKED=0
STRIKE=0
SPARE=0
THROWS=0
CURRENT_FRAME=1

roll () {
    THROWS=$((THROWS+1))
    FRAME+=($1)
    KNOCKED=$((KNOCKED+$1))    
    SCORE[$((CURRENT_FRAME-1))]=${FRAME[@]}

    [[ $1 -eq 10 ]] && { STRIKE=1; end_frame; }
    [[ $KNOCKED -eq 10 ]] && SPARE=1;
    [[ $THROWS -eq 2 ]] && end_frame

    write_score
}

end_frame () {
    CURRENT_FRAME=$((CURRENT_FRAME+1))
    THROWS=0
    FRAME=()
}

write_score () {
    clear

    SH=""
    SB=""
    P="|"
    SPARE="__"

    for ((i=0; i<10; i++)); do
        F=(${SCORE[$i]})

        FIRST=${F[0]}
        [[ ${#FIRST} -eq 0 ]] && FIRST="__"
        [[ ${#FIRST} -eq 1 ]] && FIRST="0${FIRST}"

        LAST=${F[1]}
        [[ ${#LAST} -eq 0 ]] && LAST="__"
        [[ ${#LAST} -eq 1 ]] && LAST="0${LAST}"

        [[ $i -gt 0 ]] && P=""

        SH+="${P}_____$(($i+1))_____|"
        [[ $i -lt 9 ]] && SB+="${P}__${FIRST}___${LAST}__|" || SB+="${P}__${FIRST}_${LAST}_${SPARE}__|"
    done

    echo $SH
    echo $SB
    echo
}

clear
write_score

for x in $(echo $* | grep -o .); do
    read -p "Press enter key to roll a: ${x}"
    roll $x
done

end_frame