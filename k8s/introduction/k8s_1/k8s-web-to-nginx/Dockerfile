FROM node:alpine

WORKDIR /app

EXPOSE 3000

COPY package.json yarn.lock ./

RUN yarn

COPY . ./

CMD ["yarn", "start"]
