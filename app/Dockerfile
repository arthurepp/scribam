FROM python:3-alpine
ENV PYTHONUNBUFFERED 1
RUN apk update && apk add --no-cache postgresql-dev gcc python3-dev musl-dev
RUN mkdir /code
WORKDIR /code
ADD requirements.txt /code/
RUN pip install --no-cache-dir -r requirements.txt
ADD . /code/
CMD python app.py