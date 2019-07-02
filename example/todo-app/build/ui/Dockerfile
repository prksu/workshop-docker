FROM node:12-alpine AS build
WORKDIR /usr/src/todo-app-ui
ENV PATH /usr/src/todo-app-ui/node_modules/.bin:$PATH
COPY ui/package.json /usr/src/todo-app-ui/package.json
RUN npm install --silent
COPY ui/ /usr/src/todo-app-ui
RUN npm run build

FROM nginx
COPY --from=build /usr/src/todo-app-ui/dist /usr/share/nginx/html
EXPOSE 80 443
CMD ["nginx", "-g", "daemon off;"]