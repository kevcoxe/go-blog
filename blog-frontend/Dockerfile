FROM node:12.6.0 as blog-frontend

WORKDIR /frontend

COPY . /frontend

RUN yarn
RUN yarn build


# Use the Nginx image
FROM nginx

# Remove the default nginx.conf
RUN rm /etc/nginx/conf.d/default.conf

# Replace with our own nginx.conf
COPY nginx.conf /etc/nginx/conf.d/
COPY --from=blog-frontend /frontend/build /var/www/html/build/
