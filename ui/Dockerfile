ARG NODE_TAG=12
ARG NGINX_TAG=alpine
ARG APP_HOME=/home/node/app

# build stage
FROM node:${NODE_TAG} as build
ARG NODE_TAG
ARG APP_HOME

WORKDIR ${APP_HOME}
COPY package*.json ${APP_HOME}/
RUN echo -n 'node '; node -v; \
    echo -n 'npm '; npm -v ; \
    npm install --production

COPY src ${APP_HOME}/src/
COPY public ${APP_HOME}/public/
RUN npm run build

# deploy stage
FROM nginx:${NGINX_TAG}
ARG NGINX_TAG
ARG APP_HOME
RUN echo "nginx:${NGINX_TAG}" > /docker-image-tag && cat /docker-image-tag
COPY --from=build ${APP_HOME}/build /usr/share/nginx/html

WORKDIR /usr/share/nginx/html

# フォアグラウンドでnginx実行
CMD ["nginx", "-g", "daemon off;"]