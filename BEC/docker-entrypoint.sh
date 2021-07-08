#!/usr/bin/env bash
# exit when any command fails
set -e
cd /specs

# set name of bash
export PS1="フロー BEC #"

# look for a .furobecrc and run it
if test -f .furobecrc; then
  # ensure that we can execute
  chmod +x ./.furobecrc
  . ./.furobecrc
fi


if [ "$1" = 'build' ]; then
  furo run build
fi

if [ "$1" = 'publish' ]; then
  furo run publish
fi


if [ "$1" = 'bash' ]; then
  exec "$@"
fi
