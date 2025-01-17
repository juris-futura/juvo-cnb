FROM python:3.12-alpine3.21 AS base

# Each stack needs an identifier. This is used by the CNB framework to
# determine compatibility among components, for example, matching the build
# and run images of a stack, or stack eligibility to use with a particular buildpack.
ENV CNB_STACK_ID="com.jurisfutura.cnb.stacks.alpine"
LABEL io.buildpacks.stack.id="com.jurisfutura.cnb.stacks.alpine"

# The CNB framework executes commands during image building as a specified
# user rather than the Docker default of root. In fact, it is expressly
# forbidden to use any root level user for CNB purposes. This improves the
# security and helps prevent permissions problems within the final application image.
ENV CNB_USER_ID=1000
ENV CNB_GROUP_ID=1000
ENV CNB_PLATFORM_API=0.11
RUN addgroup -g ${CNB_GROUP_ID} cnb
RUN adduser -u ${CNB_USER_ID} -G cnb cnb -D

# alpine-sdk replaces build-essential

RUN apk add --update --no-cache ca-certificates alpine-sdk python3-dev \
                     ncurses gdbm-dev libc6-compat libzmq \
                     zlib-dev libssl3 openssl libffi-dev openssh bash

RUN pip --no-cache-dir install 'poetry==1.8.4'

# BUILD IMAGE

FROM base AS build

RUN apk add --update git
LABEL io.buildpacks.stack.mixins="[\"build:git\"]"

# we set the active user for the image employing the environment variables
# declared in the base image.
USER ${CNB_USER_ID}:${CNB_GROUP_ID}

# RUN IMAGE
FROM base AS run

USER ${CNB_USER_ID}:${CNB_GROUP_ID}

