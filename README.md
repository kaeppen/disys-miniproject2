# Mini project 2 for Distributed Systems 

# Instructions for use
1. Open a terminal in disys-miniproject2\chittychat_server and type "go run ."
2. Open a terminal in disys-miniproject2\chittychat_client and type "go run ."
3. Repeat step 2 in as many different terminals as you want clients
4. Follow the instructions in the terminal window running the client(s)
5. Inspect logs in disys-miniproject2\Logfile.txt

# Description:

You have to implement Chitty-Chat a distributed system, that is providing a chatting service, and keeps track of logical time using Lamport Timestamps.

We call clients of the Chitty-Chat service Participants. 

System Requirements

R1: Clients in Chitty-Chat can Publish a valid chat message at any time the wish.  A valid message is a string of UTF-8 encoded text with a maximum length of 128 characters.
R2: A client publishes a message by making an RPC call Publish() to Chitty-Chat.
R3: The Chitty-Chat service has to Broadcast every published message, together with the current Lamport timestamp, to all participants in the system, by using the RPC call Broadcast(). It is an implementation decision left to the students, whether a Vector Clock or a Lamport timestamp is sent.
R4: When a client receives a broadcasted message, it has to write the message and the current Lamport timestamp to the log
R5: Chat clients can join any time. 
R6: A "Participant X  joined Chitty-Chat at Lamport time L" message is broadcast when client X joins to all Participant, including the new Participant.
R7: Chat clients can drop out any time. 
R8: A "Participant X left Chitty-Chat at Lamport time L" message is broadcast when Participant X leaves to all remaining Participants.
Technical Requirements:

Use gRPC for all message passing between nodes
Use Golang to implement the service and clients
Every client has to be deployed as a separate processes
Log all service calls (Publish, Broadcast, ...) using the log package
Demonstrate that the system can be started with at least 3 client nodes 
Demonstrate that a client node can join the system
Demonstrate that a client node can leave the system
Optional: All elements of the Chitty-Chat service are deployed as Docker containers
Hand-in requirements:

Hand in a single report in a pdf file
Describe your system architecture - do you have a server-client architecture, peer to peer, or something else?
Describe what  RPC methods are implemented: Publish(), Broadcast(), any other ?
Describe how you have implemented calculation of the Lamport timestamps
Provide a diagram, that traces a sequence of RPC calls together with the Lamport timestamps, that corresponds to a chosen sequence of interactions: Client X joines, Client X Publishes, ..., Client X leaves. Include documentation (system logs) in your appendix.
Provide a link to a Git repo with your source code in the report
Include system logs, that document the requirements are met, in the appendix of your report
Grading notes

Partial implementations may be accepted, if the students can reason what they should have done in the report.
