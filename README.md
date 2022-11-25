# faccounts
f3 accounts challenge



### run challenge
build test image and execute tests based on docker compose dependencies defined
<pre>
docker-compose up --build --abort-on-container-exit
</pre>

clean up resources stored in volumes
<pre>
docker-compose down --volumes
</pre>



### architecture
#### package view
![Screenshot](https://github.com/paguerre3/tpricing/blob/main/design/pckge-diagram.png?raw=true)

#### sequence view
![Screenshot](https://github.com/paguerre3/tpricing/blob/main/design/seq-diagram.png?raw=true)

#### technology stack
![Screenshot](https://github.com/paguerre3/tpricing/blob/main/design/impl-img.png?raw=true)



### troubleshoot commands
list network
<pre>
docker network ls  
</pre>
inspect default network created
<pre>
docker network inspect faccounts_default
</pre>



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

