Install nginx
`sudo apt install nginx`

Start nginx on machine boot
`sudo systemctl enable nginx`

Start nginx now
`sudo systemctl start nginx`

Check nginx status (it should be active without errors)
`sudo systemctl status nginx`

Create the following file with the following content
`sudo vim /etc/nginx/sites-enabled/gochess_config`

```
upstream gochess {
        server localhost:8081;
}

server {
        listen 80;
        listen [::]:80;

        server_name gochess;

        location / {
                proxy_pass http://gochess;
                include proxy_params;
        }

        location /ws {
                proxy_pass http://gochess;
                include proxy_params;
                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Connection "upgrade";
        }
}
```
------------------------

In nginx config file
`sudo vim /etc/nginx/nginx.conf`

Comment this line and replace it by adding the following line in the http section

```Bash
#include /etc/nginx/sites-enabled/*;
include /etc/nginx/sites-enabled/gochess_config;
```

Restart nginx so it takes the new config into account
`sudo systemctl restart nginx`

Check nginx status (it should be active without errors)
`sudo systemctl nginx status`
