#!/bin/bash

cmd () {

shdir="$MYAPP_DIR/sh" 

 case "$1" in 
    run)
     source $shdir/run/cmd.sh
     run "$@"
     ;;
    connect)
      source $shdir/connect/cmd.sh
      connect "$@"
      ;;
    deploy)
      source $shdir/deploy/cmd.sh
      deploy "$@"
      ;;
    play)
      source $shdir/play/cmd.sh
      play "$@"
      ;;
     ping)
      ansible all -m ping -i $MYAPP_DIR/deploy/ansible/inventory.yaml
      ;;
     tui)
       cd $MYAPP_DIR
       go run ./cmd/tui/main.go
       ;;
     p)
      cat $shdir/cmd.sh
      ;;
      *)
      echo "command '$1' not found. Expected myapp [run | deploy]"
      ;;
  esac
}

