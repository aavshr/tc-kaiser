#!/bin/sh

/bin/sh -c exec "nginx"
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start my_first_process: $status"
  exit $status
fi

./server 

# Naive check runs checks once a minute to see if either of the processes exited.
# The container exits with an error if it detects that either of the processes has exited.
# Otherwise it loops forever, waking up every 60 seconds