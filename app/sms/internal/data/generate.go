package data

//go:generate sh -c "find . -mindepth 1 -maxdepth 1 -name '*' -type d -exec go run -mod=mod entgo.io/ent/cmd/ent generate --template={}/template {}/schema --target {}/ent --feature sql/execquery \\;"
