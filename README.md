# go-backend-comm-patterns
 What this is: Backend communication design patterns in golang

 Description: I have a keen interest in backend development, and I constantly look for ways to improve my skills in backend programming. Learning from Hussein Nasser's Fundamentals of Backend Engineering course on Udemy. Most of the his concepts are covered in Node. Rewriting those concepts in Go. This repo is dedicated to testing out various concepts that pertain to backend development. Mainly for my own reference and acts as a repo to test out various ideas I gain from my studies.

 - Concepts currently iterated on:
    - Short Polling
      - [ ] Client sends a request
      - [ ] Server responds immediately with a handle
      - [ ] Server connotes to process the request
      - [ ] Client uses the handle to check for status
    - Server Sent Events
      - [ ] Response has a start and end
      - [ ] Client sends a request
      - [ ] Server sends logical events as part of response that the client can understand
      - [ ] Server never writes the end/final response
      - [ ] Client parses the data
   -Pub/Sub
      - [ ] Topic: An intermediary channel that maintains a list of subscribers to relay messages to that are received from publishers
      - [ ] Message: Serialized messages sent to a topic by a publisher which has no knowledge of the subscribers
      - [ ] Publisher: The application that publishes a message to a topic
      - [ ] Subscriber: An application that registers itself with the desired topic in order to receive the appropriate messages
