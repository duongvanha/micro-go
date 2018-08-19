FROM iron/base
EXPOSE 3000

ADD micro-go /
ADD .env /

ENTRYPOINT ["./micro-go"]