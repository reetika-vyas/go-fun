#!/usr/bin/env bash
set -e

ROOT_PATH=$(dirname $(cd .;pwd -P))
echo -en "\033[1;32m Script is Idempotent can be run multiple times
 Please Run With Sudo as it requires creating soflinks in /etc/
 \033[0m \n"
echo -en "\033[1;34m Using Source Path as $ROOT_PATH \033[0m \n"

#Fun App
echo -en "\033[1;32m Fun App \033[0m \n"
FUNAPP_CONFIG=/etc/fun-app
rm -rf ${FUNAPP_CONFIG}; mkdir -p ${FUNAPP_CONFIG}

#sudo ln -s ${ROOT_PATH}/components/fun-app/config.yml ${FUNAPP_CONFIG}/config.yml