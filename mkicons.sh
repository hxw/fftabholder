#!/bin/sh
ERROR() {
  printf 'error: '
  printf "$@"
  printf '\n'
  exit 1
}

[ -z "$1" ] && ERROR 'missing arguments'

for file in $*
do
  [ -f "${file}" ] || ERROR 'missing file: %s' "${file}"
done

printf '// Copyright (c) 2015-%s, Christopher Hall\n' $(date '+%Y')
printf '// see: LICENSE\n'
printf 'package main\n\n'

printf 'var icons = map[string][]byte{\n'

for file in $*
do
  bn=$(basename "${file}" | tr '[:upper:]' '[:lower:]')
  printf '\t"%s": []byte{\n' "${bn%%.*}"
  xxd -i -c 16 < "${file}" | sed -E 's/^/		/;s/,?$/,/'
  printf '\t},\n'
done
  printf '}\n'
