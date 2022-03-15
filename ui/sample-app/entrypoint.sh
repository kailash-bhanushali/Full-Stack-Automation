#!/usr/bin/env bash
export EXISTING_VARS=$(printenv | awk -F= '{print $1}' | sed 's/^/\$/g' | paste  -sd,);
sleep 1
for file in $JSFOLDER;
do
  cat $file | envsubst $EXISTING_VARS | tee $file
done
sleep 1
nginx -g 'daemon off;'