#!/bin/bash

rm -rf /opt/cert.tar.gz

cd /root/logs && tar -czvf cert.tar.gz *

ls -al .

mv cert.tar.gz /opt
