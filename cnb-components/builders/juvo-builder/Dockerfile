FROM python:3.11-slim-bookworm as base

# Each stack needs an identifier. This is used by the CNB framework to
# determine compatibility among components, for example, matching the build
# and run images of a stack, or stack eligibility to use with a particular buildpack.
ENV CNB_STACK_ID="com.jurisfutura.cnb.stacks.ubuntu"
LABEL io.buildpacks.stack.id="com.jurisfutura.cnb.stacks.ubuntu"

# The CNB framework executes commands during image building as a specified
# user rather than the Docker default of root. In fact, it is expressly
# forbidden to use any root level user for CNB purposes. This improves the
# security and helps prevent permissions problems within the final application image.
ENV CNB_USER_ID=1000
ENV CNB_GROUP_ID=1000
ENV CNB_PLATFORM_API=0.11
RUN groupadd cnb --gid ${CNB_GROUP_ID} && \
    useradd --uid ${CNB_USER_ID} --gid ${CNB_GROUP_ID} -m -s /bin/bash cnb

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN pip install poetry==1.8.2
RUN poetry config virtualenvs.in-project true

# BUILD IMAGE
FROM base as build

RUN apt-get update && \
    apt-get install -y git && \
    rm -rf /var/lib/apt/lists/*

LABEL io.buildpacks.stack.mixins="[\"build:git\"]"

# we set the active user for the image employing the environment variables
# declared in the base image.
USER ${CNB_USER_ID}:${CNB_GROUP_ID}

# RUN IMAGE
FROM base as run

USER ${CNB_USER_ID}:${CNB_GROUP_ID}

