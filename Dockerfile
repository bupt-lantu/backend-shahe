FROM alpine

WORKDIR /App
ADD . /App/
RUN sed -i 's/dev/prop/g' conf/app.conf 

CMD [ "./bupt_tour" ]
