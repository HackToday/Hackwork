#! /usr/bin/env bash

add_1=4
add_2=5
# it is go for cross-platform
r_v1=`expr $add_1 + $add_2`
echo "$add_1 + $add_2 = $r_v1"

r_v2=$(( $add_1 + $add_2 ))
echo "v2: $r_v2"

r_v3=$[ $add_1 + $add_2 ]
echo "v3: $r_v3"


