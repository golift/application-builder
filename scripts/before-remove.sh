#!/bin/bash

# This file is used by rpm and deb packages. FPM use.
# Edit this file as needed for your application.

if [ "$1" = "upgrade" ] || [ "$1" = "1" ] ; then
  exit 0
fi

if [ -x "/bin/systemctl" ]; then
  /bin/systemctl stop hello-world
  /bin/systemctl disable hello-world
fi
