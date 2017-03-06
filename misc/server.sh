#!/bin/bash
#

ports="7980 7981 7982 7983"
gqscmd="`which gqs`"
grpccmd="`which greeter_server`"


basedir="/tmp/gobeansdb_$USER"
prog="gobeansdb"

function gen_conf()
{
    ./tests/gen_config.py -d $basedir
}

function start()
{
    if [ `ps -ef | grep -c "$gqscmd"` -ge 2 ]; then
        echo "gqs server already starte"
    else
        $gqscmd &
        echo "Starting the gqs server"
    fi

    if [ `ps -ef | grep -c "$grpccmd"` -ge 2 ]; then
        echo "grpc server already starte"
    else
        $grpccmd &
        echo "Starting the grpc server"
    fi
}

function stop()
{
    if [ `ps -ef | grep -c "$gqscmd"` -eq 1 ]; then
        echo $"Stopped the gqs server"
    else
        kill -TERM `ps -ef | grep "$gqscmd" | grep -v grep | awk '{ print $2 }'`
        echo "Stopping the gqs server on port"
    fi

    if [ `ps -ef | grep -c "$grpccmd"` -eq 1 ]; then
        echo $"Stopped the grpc server"
    else
        kill -TERM `ps -ef | grep "$grpccmd" | grep -v grep | awk '{ print $2 }'`
        echo "Stopping the grpc server"
    fi
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        sleep 1
        start
        ;;
    *)
        printf 'Usage: %s {start|stop|restart}\n' "$prog"
        exit 1
        ;;
esac
