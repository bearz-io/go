
$delim = [System.IO.Path]::PathSeparator
$env:PROTO_HOME="$HOME/.proto";
$env:PATH="$PROTO_HOME/shims:$PROTO_HOME/bin:$PATH";
export PATH=$HOME/.local/bin:$PATH
export PATH="$PATH:$HOME/bin"
