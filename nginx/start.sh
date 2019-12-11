#!/bin/sh

sed -i -e 's/$PORT/'"$PORT"'/g' /etc/nginx/conf.d/default.conf
nginx
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start nginx: $status"
  exit $status
fi

./server 