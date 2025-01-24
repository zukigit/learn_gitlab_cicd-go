FROM alpine:3.21

COPY test_cicd /usr/local/bin/
RUN chmod +x /usr/local/bin/test_cicd

CMD [ "/usr/local/bin/test_cicd", "$TEST_LISTEN_PORT" ]
