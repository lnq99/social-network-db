# nohup app >/dev/null 2>&1 &

go build -o app ../cmd/main.go

export GIN_MODE=release

pkill app

./app .0.env >/dev/null 2>&1 &
./app .1.env >/dev/null 2>&1 &
./app .2.env >/dev/null 2>&1 &