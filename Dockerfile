#base image
#TODO: add more documentation on how to create a project specific image
FROM alpine
COPY ./release/alpine-linux/go-mssql-runner .
CMD ["./go-mssql-runner", "version"] 
