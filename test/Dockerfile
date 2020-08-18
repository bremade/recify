FROM python:3

RUN pip install requests pymongo pytest pytest-depends

COPY ./test /api_tests
WORKDIR /api_tests
CMD [ "pytest" ]
