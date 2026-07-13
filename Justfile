list:
    just --list

dev:
    wails dev

build:
    wails build --clean

install:
    #!/usr/bin/env fish
    set app "$(cat wails.json | jq -r '.name')"
    set source "build/bin/$app.app"
    set dest "/Applications/$app.app"
    echo "source: $source"
    echo "dest: $dest"
    rm -rf $dest
    cp -r $source $dest
    gum confirm "Open $app?" --default=false ;and open -a $app ;or true
