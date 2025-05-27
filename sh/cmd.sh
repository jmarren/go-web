#!/bin/bash

cmd () {

 case "$1" in 
    run)
     source ./sh/run/cmd.sh
     run "$@"
     ;;
    connect)
      source ./sh/connect/cmd.sh
      connect "$@"
      ;;
    deploy)
      source ./sh/deploy/cmd.sh
      deploy "$@"
      ;;
    play)
      source ./sh/play/cmd.sh
      play "$@"
      ;;
     ping)
      ansible all -m ping -i ./deploy/ansible/inventory.yaml
      ;;
     p)
      cat ./sh/cmd.sh
      ;;
      *)
      echo "command '$1' not found. Expected myapp [run | deploy]"
      ;;
  esac
}

