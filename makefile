DOCKER_REPO = mclabsbr/go-rental-backend
CURRENT_VERSION := $(shell wget -qO- "https://registry.hub.docker.com/v2/repositories/$(DOCKER_REPO)/tags/" | jq -r '.results | map(.name) | max_by(.)' | sed 's/\([0-9]*\.[0-9]*\.[0-9]*\)/\1/')

ifeq ($(CURRENT_VERSION),null)
	NEW_VERSION = 1.0.0
else
	LAST_DIGIT := $(shell echo $(CURRENT_VERSION) | awk -F. '{print $$3}')
	MIDDLE_DIGIT := $(shell echo $(CURRENT_VERSION) | awk -F. '{print $$2}')
	FIRST_DIGIT := $(shell echo $(CURRENT_VERSION) | awk -F. '{print $$1}')

	ifeq ($(LAST_DIGIT),9)
		ifeq ($(MIDDLE_DIGIT),9)
			NEW_VERSION = $$(($(FIRST_DIGIT) + 1)).0.0
		else
			NEW_VERSION = $(FIRST_DIGIT).$$(($(MIDDLE_DIGIT) + 1)).0
		endif
	else
		NEW_VERSION = $(FIRST_DIGIT).$(MIDDLE_DIGIT).$$(($(LAST_DIGIT) + 1))
	endif
endif

IMAGE_NAME = $(DOCKER_REPO):$(NEW_VERSION)

.PHONY: build push release

build:
	docker build --no-cache -t $(IMAGE_NAME) .

push:
	docker push $(IMAGE_NAME)

release: build push
