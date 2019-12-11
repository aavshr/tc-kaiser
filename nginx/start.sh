#!/bin/sh

/bin/sh -c exec "nginx"
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start nginx: $status"
  exit $status
fi

./server 