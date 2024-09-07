bash := if os() == "windows" { "bash.exe" } else { "/usr/bin/env bash" }

# Run a command in a bash shell.
new path lib='lib':
    #!{{ bash }}
    case "{{lib}}" in
        "lib")
            moon generate lib "{{path}}" -- --path "{{path}}"
            go work use "{{path}}"
            ;;
        *)
            echo "Unknown library type: {{lib}}"
            exit 1
            ;;
    esac