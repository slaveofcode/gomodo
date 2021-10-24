#!/bin/bash

# Checking NPM installed
if ! command -v node -v &> /dev/null
then
    echo "Nodejs is not installed"
    exit 1
fi

# Checking Go installed
if ! command -v go version &> /dev/null
then
    echo "Golang is not installed"
    exit 1
fi

help_context="\nUsage:"
help_context="${help_context}\n\t./run.sh <command>\n"
help_context="${help_context}\nAvailable commands:\n"
help_context="${help_context}\n\tweb:scaffold\t\tSetup vite project at \"web\" directory"
help_context="${help_context}\n\tweb:start\t\tStart the web development server"
help_context="${help_context}\n\tweb:build\t\tBuild the web static files to \".build\" directory"
help_context="${help_context}\n\tapi:start\t\tStart Go backend server"
help_context="${help_context}\n"
help_context="${help_context}\n\tbuild <env-file>\tCompile the project into a single file binary"
help_context="${help_context}\n"

CMD_WEB_SETUP='web:scaffold'
CMD_WEB_INSTALL_DEPS='web:install'
CMD_WEB_START_SERVER='web:start'
CMD_API_START_SERVER='api:start'
CMD_WEB_BUILD='web:build'

BASE=`pwd`
WEB_PATH="${BASE}/web"
WEB_PROJECT_DIR="vue"
WEB_PATH_VUE="${WEB_PATH}/${WEB_PROJECT_DIR}"

if [[ $1 = $CMD_WEB_START_SERVER ]]
then
    command cd $WEB_PATH_VUE && npm run dev
elif [[ $1 = $CMD_WEB_INSTALL_DEPS ]]
then
    command cd $WEB_PATH_VUE && npm install
elif [[ $1 = $CMD_WEB_SETUP ]]
then
    command cd $WEB_PATH && npm init vite@latest vue --template vue && cd vue && npm install
elif [[ $1 = $CMD_API_START_SERVER ]]
then
    command export $(grep -v '^#' .env | xargs) && go run .
elif [[ $1 = $CMD_WEB_BUILD ]]
then
    echo mv $WEB_PATH_VUE/dist .webstatic
    command cd $WEB_PATH_VUE && npm run build
    success_build=$?
    if [[ $success_build -eq 0 ]]
    then
        # Continue copy .dist
        command mv $WEB_PATH_VUE/dist .webstatic
    else
        echo "Something wrong at build"
    fi
else
    echo -e $help_context
fi