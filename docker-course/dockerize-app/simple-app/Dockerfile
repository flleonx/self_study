FROM node:18.14.1-alpine3.17

RUN addgroup app && adduser -S -G app app

USER app

WORKDIR /app

RUN mkdir data

COPY package.json yarn.lock /app/

RUN yarn install

COPY . .

ENV API_URL=http://api.myapp.com

EXPOSE 5173

# Shell form
# CMD yarn dev

# Exec form
CMD ["yarn", "dev"]

# ENTRYPOINT ["yarn", "dev"]
