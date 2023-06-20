FROM nginx
RUN rm /etc/nginx/conf_nginx /etc/nginx/conf.d/default.conf
COPY conf_nginx/* /etc/nginx/

COPY conf_nginx/site.conf /etc/nginx/conf_nginx/site.conf
