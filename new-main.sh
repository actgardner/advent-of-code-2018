#!/bin/bash

set -euo pipefail

if [[ -z $1 || -z $2 ]]; then
   echo "Usage: $0 <day> <cookie>"
   exit 1
fi

OUTPUT="2018/$1/main.go"

if [[ -f "$OUTPUT" ]]; then
   echo "$OUTPUT already exists, not overwriting"
   exit 2
fi

mkdir -p 2018/$1

curl -f --cookie "session=$2" https://adventofcode.com/2018/day/$1/input -o "$OUTPUT.input"
cat > "$OUTPUT" <<EOF
package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	delim := "\n"
	parts := strings.Split(input, delim)
	intParts := make([]int64, 0)
	for _, p := range parts {
		intPart, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			fmt.Printf("Failed to parse %q", p)
			return
		}
		intParts = append(intParts, intPart)
	}
	fmt.Printf("%v", intParts)
}

var input = \`$(cat $OUTPUT.input)\`
EOF

echo "New main is in 2018/$1/main.go"

vim "2018/$1/main.go"
