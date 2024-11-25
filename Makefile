buf:
	@buf generate

run user:
	@go build -o tesuto-user/app.exe ./tesuto-user
	@./tesuto-user/app.exe --config tesuto-user/config/config.yaml