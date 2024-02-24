# tt

#### Start http server

````shell
go run main.go
````


#### run httprunner

````shell
httprunner/hrp tt_test/testcases/requests.yml --gen-html-report
````


#### run locust

````shell
source py_venv/bin/active

pip3 install locust

locust -f locustfile.py -H http://localhost:5000
````