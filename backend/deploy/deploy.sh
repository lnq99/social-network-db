# nohup app >/dev/null 2>&1 &

# ab -n 10000 -c 1000 -g out/1.txt http://127.0.0.1/api/v1/cmt/3231

go build -o app ../cmd/main.go

# export GIN_MODE=release

pkill app

./app .0.env

# ./app .0.env >/dev/null 2>&1 &
# ./app .1.env >/dev/null 2>&1 &
# ./app .2.env >/dev/null 2>&1 &