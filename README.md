# Concurrency with Go
Here you can find everyday problems that can be solved better with concurrency.

## Content

### Gofi
Reading files concurrently

### FastFetch
Fetching multiple requests in microservice environment. Common scenario: fetching results (indepedent from each other) from several microservices. Some of them may require a **complex computation** or the **response is slow**. 

A nice alternative way of doing these request is asynchronously fetching each of them in a separate ```go routine```. This way you can save time due to the fact that the time needed for fethching all the requests equals the time of the slowest request (~3s in the example). Doing it syncronously (~9s in the example), the fetching time is the sum of each request timeframes. 

<div width="100%">
<img src="https://github.com/JordanRad/concurrency-practice/blob/main/resources/fast_fetch/async.png" alt="async" width="49%"/>
<img src="https://github.com/JordanRad/concurrency-practice/blob/main/resources/fast_fetch/sync.png" alt="sync" width="49%"/>
</div>


