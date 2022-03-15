FROM node:16.5 AS ui-builder
WORKDIR /app
COPY .npmrc .npmrc
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM nginx:latest
ENV JSFOLDER=/usr/share/nginx/html/config.js
COPY ./entrypoint.sh /usr/bin/entrypoint.sh
RUN chmod +x /usr/bin/entrypoint.sh
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d
WORKDIR /usr/share/nginx/html
COPY --from=ui-builder /app/build .
EXPOSE 3000
EXPOSE 80
ENTRYPOINT [ "entrypoint.sh" ]