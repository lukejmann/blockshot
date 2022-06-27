# FROM node:14-alpine as dev
# RUN apk add --no-cache libc6-compat
# RUN apk add --no-cache git
# WORKDIR /root/
# CMD echo "yarn" \
#     && yarn \
#     && echo "yarn start" \
#     && yarn dev

# FROM node:14-alpine as build
# RUN apk add --no-cache libc6-compat
# RUN apk add --no-cache git
# WORKDIR /root
# COPY ./package*.json ./
# RUN yarn
# COPY . . 
# RUN yarn build

# FROM node:14-alpine as prod
# WORKDIR /root
# EXPOSE 3000
# COPY --from=build /root/public ./public
# COPY --from=build /root/.next ./.next
# COPY --from=build /root/node_modules ./node_modules
# COPY --from=build /root/package.json ./package.json
# CMD ["yarn", "start"]



FROM node:lts as dev
WORKDIR /root/
CMD echo "yarn" \
    && yarn \
    && echo "yarn start" \
    && yarn dev

FROM node:lts as dependencies
WORKDIR /root
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

FROM node:lts as build
WORKDIR /root
COPY . .
COPY --from=dependencies /root/node_modules ./node_modules
RUN echo $(ls)
RUN yarn build

FROM node:lts as prod
WORKDIR /root
ENV NODE_ENV production
COPY --from=build /root/. .
COPY --from=build /root/next.config.js ./
COPY --from=build /root/public ./public
COPY --from=build /root/.next ./.next
COPY --from=build /root/node_modules ./node_modules
COPY --from=build /root/package.json ./package.json

EXPOSE 3000
CMD ["./node_modules/.bin/next","start"]