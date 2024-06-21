FROM node:16-alpine3.17 as build
WORKDIR /usr/src/app
COPY package.json package-lock.json ./
RUN npm install --legacy-peer-deps
COPY . .
RUN npm run build --omit=dev
FROM nginx:1.15.8-alpine
COPY --from=build /usr/src/app/build /usr/share/nginx/html
COPY .nginx/nginx.conf /etc/nginx/nginx.conf