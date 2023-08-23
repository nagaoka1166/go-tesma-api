# #!/bin/bash
# # wait-for-it.sh
# set -e

# host="$1"
# shift
# port="$1"
# shift
# timeout="${1:-15}"
# shift

# cmd="$@"

# until $(nc -zv $host $port); do
#     >&2 echo "Waiting for $host:$port - sleeping"
#     sleep 1
#     timeout=$((timeout - 1))
#     if [ $timeout -le 0 ]; then
#         >&2 echo "Timeout! $host:$port still not reachable"
#         exit 1
#     fi
# done

# >&2 echo "$host:$port is available - executing command"
# exec $cmd