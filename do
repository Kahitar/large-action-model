## nginx: Start nginx
function task_nginx() { 
    cd nginx
    
    # Start nginx in the foreground
    # nginx -c nginx.conf
    # Or start it in the background
    # winpty nginx -c nginx.conf &
    # Or use a batch file
    ./start_nginx.bat
}

## kill-nginx: Kill nginx
function task_kill_nginx() {
    export MSYS_NO_PATHCONV=1
    # tasklist /fi "imagename eq nginx.exe"
    # nginx -s stop 
    # nginx -s quit 
    # export MSYS_NO_PATHCONV=1
    taskkill /f /im nginx.exe
}

## ngrok: Start the ngrok tunnel
function task_ngrok() {
    ngrok http --domain=careful-broadly-kitten.ngrok-free.app 8080
}

## run [project-name]: Start the project using it's do script
function task_run() {
    project=${1:-"hello-world"}
    echo "Starting project '$project'"
    cd $project
    ./do run
}

function task_usage {
    echo "Usage: $0"
    sed -n 's/^##//p' <"$0" | column -t -s ':' | sed -E $'s/^/\t/'
}

cmd=${1:-}
shift || true
resolved_command=$(echo "task_${cmd}" | sed 's/-/_/g')
if [[ "$(LC_ALL=C type -t "${resolved_command}")" == "function" ]]; then
    pushd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null
    ${resolved_command} "$@"
else
    task_usage
fi
