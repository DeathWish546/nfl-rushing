FROM node:14-alpine

WORKDIR /app

COPY package.json .

RUN npm install --quiet

EXPOSE 3000

CMD ["npm","start"]
