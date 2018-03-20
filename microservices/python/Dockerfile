FROM python:3.6.4-alpine3.7


RUN apk add --update bash curl && \
    rm -rf /var/cache/apk/*

RUN pip install  \
    flask

COPY . /src
WORKDIR /src

ENV SERVICE_VERSION Version-Tw2

CMD ["python", "./app.py"]