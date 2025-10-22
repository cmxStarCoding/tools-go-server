#
gentool -dsn "root:@tcp(localhost:3380)/tools?charset=utf8mb4&parseTime=True&loc=Local" --outPath "models" -tables "t_export_task"
