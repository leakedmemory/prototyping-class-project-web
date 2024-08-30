#!/usr/bin/env bash

program_is_installed() {
    command -v $1 &>/dev/null
}

ask_to_install_program_or_exit() {
    read -p "Go's '$1' is not installed on your machine. Do you want to install it? [Y/n] " choice
    choice=${choice:-Y}
    case "$choice" in
    [Yy]*)
        go install $2
        ;;
    *)
        echo "You chose not to install '$1'. Exiting..."
        exit 1
        ;;
    esac
}

declare -A needed_programs=(
    ["templ"]="github.com/a-h/templ/cmd/templ@latest"
    ["air"]="github.com/air-verse/air@latest"
)

for program in "${!needed_programs[@]}"; do
    if ! program_is_installed $program; then
        ask_to_install_program_or_exit $program $repo
    fi
done

if [ ! -f ".env" ]; then
    echo "'.env' file not found"
    exit 1
else
    source .env
fi

if [[ -z "${PORT}" ]]; then
    echo "'PORT' environment variable is not set in .env"
    exit 1
fi

air
