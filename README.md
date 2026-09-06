# raft

An implementation of the Raft distributed consensus protocol based on _In Search of an Understandable Consensus Algorithm_ by Diego Ongaro and John Ousterhout, and following the [implementation guide by Eli Bendersky](https://eli.thegreenplace.net/2020/implementing-raft-part-0-introduction/).

## Useful Resources

- [Raft paper](https://raft.github.io/raft.pdf)
- [Raft website](https://raft.github.io/)
- [Raft visualizer](https://raft.github.io/visualizer.html)
- [Raft explained](https://thesecretlivesofdata.com/raft/)
- [Raft explained (video)](https://www.youtube.com/watch?v=vYp4LYbnnW8)
- [Raft implementation guide](https://eli.thegreenplace.net/2020/implementing-raft-part-0-introduction/)

## My Personal Advices

- Read the first 2 pages of the Raft paper to understand the problem and the solution, if you didn't get it well, ask an LLM to have a broader idea about the problem we're solving, then watch the video, then the Raft explained website, then read the rest of the paper.
- Write the code yourself, don't copy-paste from the internet.

---

## The Story

The node (or sometimes it's called a server) can be in one of 3 states: `follower`, `candidate`, or `leader`. All nodes start in the _follower_ state. After some time, if the _follower_ did't hear from a _leader_, it becomes a _candidate_ and start an election to become a _leader_.

The _candidate_ then requests votes from other nodes and nodes reply with their votes to the _candidate_. If the _candidate_ receives votes from a majority of nodes, it becomes the _leader_. The _leader_ then starts sending heartbeats to all other nodes to maintain its leadership. If a _follower_ doesn't receive a heartbeat from the _leader_ for some time, it becomes a _candidate_ and starts a new election and so on. This process is called **leader election**.

The clients send requests to the _leader_ and the _leader_ appends the requests to its log (every node has a log). to get this request committed, the _leader_ sends the request to all nodes and waits for replys from them. When the _leader_ receives replys from a majority of nodes, it commits the request and sends a reply to the client. This process is called **log replication**.

In Raft, there is 2 timeouts: the election timeout and the heartbeat timeout. The election timeout is the time a _follower_ waits for a heartbeat from the _leader_ before becoming a _candidate_ (typically randomized between 150-300ms, it must be randomized). The heartbeat timeout is the time a _leader_ waits before sending a heartbeat to all other nodes.

There is something called `term` in Raft, which is a number that increases when a new election starts. The `term` is here to ensure that an old stale _leader_ doesn't become a _leader_ again. if an old _leader_ receives a request from a _candidate_ with a higher `term`, it steps down and becomes a _follower_ again and forget about all of its uncommited changes. The `term` is also used to ensure that the logs are consistent across all nodes.
