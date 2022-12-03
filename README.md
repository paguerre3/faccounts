# faccounts [![Build Status](https://travis-ci.com/paguerre3/faccounts.svg?token=vSTu1zSW1ehqZeuodHpi&branch=feature/ci-cd)](https://travis-ci.com/paguerre3/faccounts)

f3 accounts challenge



### run challenge
execute <code>accounts_lib_tests</code> (of client lib) based on docker compose defined dependencies
<pre>
docker-compose up --abort-on-container-exit
</pre>

alternative: rebuild test image everytime
<pre>
docker-compose up --build --abort-on-container-exit
</pre>

clean up resources stored in volumes
<pre>
docker-compose down --volumes
</pre>



### architecture
#### package view
![Screenshot](https://github.com/paguerre3/faccounts/blob/main/assets/pkg-diagram.png?raw=true)

#### sequence view
![Screenshot](https://github.com/paguerre3/faccounts/blob/main/assets/seq-diagram.png?raw=true)
---
**NOTE**

- The client lib implements a retry policy for making calls to the account api.
- Also, status codes unexpected are wrapped as error by the client lib. 

---

#### technology stack
![Screenshot](https://github.com/paguerre3/faccounts/blob/main/assets/stack.png?raw=true)



### troubleshoot commands
list network
<pre>
docker network ls  
</pre>
inspect default network created
<pre>
docker network inspect faccounts_default
</pre>



### comments
nice to have in case of a production code:

- logging capability, e.g. exposed in kibana
- metrics capability, e.g. exposed in new-relic and/or grafana/datadog
- enable sh entry, i.e. for debugging inside container client lib
- refine Dockerfile code, e.g. for <code>COPY</code> in order to reduce the layers for <code>cp</code> of go files into the container and/or using a minimal docker image for go like alpine



### versioning semantic
https://semver.org/



### license
Copyright 2021, paguerre3

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.

