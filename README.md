## Prerequisites
- go version 1.23.5

## Running the Service
1. Clone the repo
2. Navigate to the project directory.
3. Run the following command: `go run main.go`

## On Scalability:
- "The web service should remain responsive under load and be able to support multiple concurrent requests. "
----
- The net/http package uses goroutines to handle incoming HTTP requests concurrently.
  - So ListenAndServe will create new goroutines to process each request.
- If we wanted this to handle a higher load of requests:
  - We can scale horizontally, by making many ec2 instances.
  - Implement an ALB (Application load balancer) to distribute traffic across the many ec2 instances.
    - If implemented, we wouldn't hit "localhost:5000" but the defined load balancer's endpoint so that traffic is spread evenly.

## Unit Testing:
To run the unit tests, you can run this command:
`go test -v`
