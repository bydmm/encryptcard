FROM alpine:3.5

COPY ./encryptcard_server_linux_x64 encryptcard_server

RUN chmod +x /encryptcard_server

RUN mkdir /blocks

CMD ["./encryptcard_server"]