FROM node:18.14.1-alpine3.17

# RUN addgroup app && adduser -S -G app app

# USER app

WORKDIR /app

COPY package.json yarn.lock /app/

RUN yarn install

COPY . .

EXPOSE 59000

CMD ["yarn", "start"]
