#example image
FROM alpine
RUN mkdir logs
COPY ./release/alpine-linux/go-mssql-runner .
COPY ./example ./example
ENTRYPOINT [ "./go-mssql-runner" ]
CMD ["version"] 
