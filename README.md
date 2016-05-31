# this the endpoint we need to call

TODO:
- call api from main
- sort by rank/star/popularity(using some sort of a sorter struct) and make optional
- allow get details by repo id
- get the arg from user

Workflow:

- call to  `main show [reponame]` also stroes data to the storage
- when call to detail occurs, api call made from corresponding id that got retrieved from disk
- on every call to `main show [reponame]` the storage is overwitten with new data on the disk
- optional, to increase efficiency of the api calls. We should not make any API call if stored data on the  disk 
is not older than a period(??????) of time
