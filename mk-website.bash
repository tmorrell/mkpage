#!/bin/bash

function checkApp() {
    APP_NAME=$(which $1)
    if [ "$APP_NAME" = "" ] && [ ! -f "./bin/$1" ]; then
        echo "Missing $APP_NAME"
        exit 1
    fi
}

function softwareCheck() {
    for APP_NAME in $@; do
        checkApp $APP_NAME
    done
}

function MakePage () {
    nav="$1"
    content="$2"
    html="$3"
    # Always use the latest compiled mkpage
    APP=$(which mkpage)
    if [ -f ./bin/mkpage ]; then
        APP="./bin/mkpage"
    fi

    echo "Rendering $html from $content and $nav"
    $APP \
	"title=text:mkpage: An experimental template and markdown processor" \
        "nav=$nav" \
        "content=$content" \
	    "sitebuilt=text:Updated $(date)" \
        "copyright=copyright.md" \
        page.tmpl > $html
}

echo "Checking necessary software is installed"
softwareCheck mkpage
echo "Generating website index.html with mkpage"
MakePage nav.md README.md index.html
echo "Generating install.html with mkpage"
MakePage nav.md INSTALL.md install.html
echo "Generating go-template-recipes.html with mkpage"
MakePage nav.md go-template-recipes.md go-template-recipes.html

