STR=$(echo $2 | tr "[:upper:]" "[:lower:]" | tr -cd "[^a-z0-9]")
if [[ $1 == "encode" ]]; then
    STR=$(echo $STR | tr "a-z" "zyxwvutsrqponmlkjihgfedcba")
    ARR=()
    CURR=""
    
    for (( i=0; i<${#STR}; ++i ))
    do
        CURR+=${STR:$i:1}
        if [[ $[i % 5] -eq 4 ]]; then
            ARR+=($CURR)
            CURR=""
        fi
    done

    ARR+=($CURR)
    echo "${ARR[*]}"
    exit 0;
fi

echo $STR | tr "zyxwvutsrqponmlkjihgfedcba" "a-z"