#!/usr/bin/env bash
shopt -s nullglob

for FILE in *.go; do
	DAY="$(echo "$FILE" | cut -d'_' -f1)"
	OUTPUT="$(go run "$FILE" <./input/"$DAY" | diff -u - ./output/"$DAY")"

	if [[ "$?" -eq 0 ]]; then
		printf "[\033[0;32mOK\033[0m] $FILE\n"
	else
		printf "[\033[0;31mKO\033[0m] $FILE\n\n$OUTPUT\n\n\033[0;33mCommand\033[0m: go run $FILE <./input/$DAY\n\n"
	fi
done
