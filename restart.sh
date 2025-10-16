cd /www/go_project/journey
git pull origin main
cd alitools
go mod tidy
go build -o alitools
#killall core
#kill $(lsof -i :8083 -t)
#nohup ./alitools --env=prod &