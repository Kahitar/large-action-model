## ngrok: Start the ngrok tunnel
function task_ngrok() {
    ngrok http --domain=careful-broadly-kitten.ngrok-free.app 8080
}

## run [project-name] (port=8080): Start the project using it's do script
function task_run() {
    project=${1:-"hello-world"}
    port=${2:-8080}
    echo "Starting project '$project' on port '$port'"
    cd $project
    ./do run $port
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
