FROM alpine

WORKDIR /App
ADD . /App/
RUN sed -i -e 's/dev/prop/g' -e 's/127.0.0.1/db/g' conf/app.conf 

CMD [ "./backend-shahe" ]
