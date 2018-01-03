#base image
FROM alpine
COPY ./release/alpine-linux/go-mssql-runner .
CMD ["./go-mssql-runner", "version"] 
