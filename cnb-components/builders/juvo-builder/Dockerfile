FROM python:3.13-slim-bookworm AS base

# Each stack needs an identifier. This is used by the CNB framework to
# determine compatibility among components, for example, matching the build
# and run images of a stack, or stack eligibility to use with a particular buildpack.
ENV CNB_STACK_ID="com.jurisfutura.cnb.stacks.debian"
LABEL io.buildpacks.stack.id="com.jurisfutura.cnb.stacks.debian"

# The CNB framework executes commands during image building as a specified
# user rather than the Docker default of root. In fact, it is expressly
# forbidden to use any root level user for CNB purposes. This improves the
# security and helps prevent permissions problems within the final application image.
ENV CNB_USER_ID=1000
ENV CNB_GROUP_ID=1000
ENV CNB_PLATFORM_API=0.11

RUN groupadd cnb --gid ${CNB_GROUP_ID}
RUN useradd --uid ${CNB_USER_ID} --gid ${CNB_GROUP_ID} -m -s /bin/bash cnb

RUN apt-get update && \
    apt-get install -y ca-certificates build-essential python3-dev && \
    apt-get install -y libncursesw5-dev libgdbm-dev libc6-dev && \
    apt-get install -y libzmq3-dev && \
    apt-get install -y zlib1g-dev libssl-dev openssl libffi-dev && \
    rm -rf /var/lib/apt/lists/*

RUN pip --no-cache-dir install "poetry==1.8.4"

# BUILD IMAGE

FROM base AS build

RUN apt-get update && \
    apt-get install -y git && \
    rm -rf /var/lib/apt/lists/*

LABEL io.buildpacks.stack.mixins="[\"build:git\"]"

# we set the active user for the image employing the environment variables
# declared in the base image.
USER ${CNB_USER_ID}:${CNB_GROUP_ID}

# RUN IMAGE
FROM base AS run

USER ${CNB_USER_ID}:${CNB_GROUP_ID}

