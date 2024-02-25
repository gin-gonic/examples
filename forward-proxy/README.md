# A  proxy integrate both forward and reverse

Run the server and make following reqeust to test forward function. Remember to set `forward=ok` in the header so the middleware can tell which one is for forward and which one is for reverse.  The demo can be adapted to puer forward proxy.

```python
import requests
def test_forward():
  res=requests.get("http://www.baidu.com",headers={"forward":"ok"},proxies={"http":"http://127.0.0.1:8888"})
  print (res.text)
test_forward()

```
