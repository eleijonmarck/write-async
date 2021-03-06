# write-async

Possiblibility to add a write query for the write master where we would add it to the database later on
inspired by the system design architecture where we get most of the functionality in place for a good applciation to scale:

[write-async-system-overview](https://github.com/donnemartin/system-design-primer#system-design-topics-start-here)

## async creates it later

/transfers {customerId}

/loan/create {
}

/transactions/create {
type,
source,
}

## system-design

- write-entity (ex: transfer) needs to implement a worker
- task-queue
- worker
- write-back is still a mystery? how do you perform write back from write-async

write-entity

> register()
> process()

| ...fields:
| status enum: - added - queued - completed

flow:

1. user |> request-time-consuming-task | queue(job_id, type) |> update(user.task.queued)
2. queue |> worker.pulls(job, type) | process | output |> update(user.task.completed)

### queue

## setup - gcloud pubsub

sets up the local use of cloud-sdk with your user on the gcloud project

```
docker run -ti --name gcloud-config google/cloud-sdk gcloud init
```

### locally

```bash
PROJECT_ID:="local"
TOPIC_ID:="topic_local"
SUB_ID:="sub_local"
```

```
gcloud beta emulators pubsub start --project=${PROJECT_ID}
./gcloud/python-pubsub/python publisher.py ${PROJECT_ID} create ${TOPIC_ID}
./gcloud/python-pubsub/python subscriber.py ${PROJECT_ID} create topic ${SUB_ID}
```
