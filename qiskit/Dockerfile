From ubuntu

RUN apt-get update
RUN apt-get install -y python-is-python3 curl
RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
RUN python get-pip.py
RUN pip install qiskit

RUN mkdir -p /tmp/qiskit
WORKDIR /tmp/qiskit
ADD sample.py .

CMD python sample.py 