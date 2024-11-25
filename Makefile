buf:
	@buf generate
run user:
	@go build -o app.exe ./tesuto-user
	@./app.exe