# application
web-site


# description


└> $ curl -X POST -F "file=@./Chemniz.jpg" http://127.0.0.1:8080/api/files
{"id":6,"urls":{"150x150":"http://127.0.0.1/img/37cb6719-786d-11f0-b52f-98fa9bbd8c21/150x150.jpg","400x400":"http://127.0.0.1/img/37cb6719-786d-11f0-b52f-98fa9bbd8c21/400x400.jpg","50x50":"http://127.0.0.1/img/37cb6719-786d-11f0-b52f-98fa9bbd8c21/50x50.jpg","800x800":"http://127.0.0.1/img/37cb6719-786d-11f0-b52f-98fa9bbd8c21/800x800.jpg","original":"http://127.0.0.1/img/37cb6719-786d-11f0-b52f-98fa9bbd8c21/original.jpg"}}

 └> $ curl -X POST -F "file=@./Chemniz.jpg" http://127.0.0.1:8080/api/files
{"id":7,"urls":{"150x150":"http://127.0.0.1/img/f0d80770-786d-11f0-b6f7-98fa9bbd8c21/150x150.jpg","400x400":"http://127.0.0.1/img/f0d80770-786d-11f0-b6f7-98fa9bbd8c21/400x400.jpg","50x50":"http://127.0.0.1/img/f0d80770-786d-11f0-b6f7-98fa9bbd8c21/50x50.jpg","800x800":"http://127.0.0.1/img/f0d80770-786d-11f0-b6f7-98fa9bbd8c21/800x800.jpg","original":"http://127.0.0.1/img/f0d80770-786d-11f0-b6f7-98fa9bbd8c21/original.jpg"}}



 └> $ curl -X POST -F "file=@./Chemniz.jpg" http://127.0.0.1:8080/api/files | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 2620k  100   425  100 2620k   1195  7373k --:--:-- --:--:-- --:--:-- 7382k
{
  "id": 8,
  "urls": {
    "150x150": "http://127.0.0.1/img/272eaabf-786f-11f0-b6f7-98fa9bbd8c21/150x150.jpg",
    "400x400": "http://127.0.0.1/img/272eaabf-786f-11f0-b6f7-98fa9bbd8c21/400x400.jpg",
    "50x50": "http://127.0.0.1/img/272eaabf-786f-11f0-b6f7-98fa9bbd8c21/50x50.jpg",
    "800x800": "http://127.0.0.1/img/272eaabf-786f-11f0-b6f7-98fa9bbd8c21/800x800.jpg",
    "original": "http://127.0.0.1/img/272eaabf-786f-11f0-b6f7-98fa9bbd8c21/original.jpg"
  }
}
