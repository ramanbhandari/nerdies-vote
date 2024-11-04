## Goals? Questions?

- Anonymity while still maintaining authorization and transparency
  - we want to only let an authorized voter to vote but we dont want to link them to the candidate they voted for - so one way cryptographic algorithm may be helpful? We cannot reverse engineer
  - we want to show total votes for candidates, being transparent, but do we want to reveal as we go or commit and reveal phase so votes cannot be bought?

## Research

https://www.usvotefoundation.org/blockchain-voting-is-not-a-security-strategy

Etherium consensus is PoS

- own 51% of the network and there is a vulnerability but is it practical? ~48 billion dollars with no guarantees to the outcome in case the community does a hard fork or slashing
- slashing: Etheriums slashing mechanisms are strong enough to detect malicious validators and what happens? Massive losses, get to 10% onlt to loose all and get nothing in back, no attack no money no ETH
- how will the market react to the efforts of someone trying to acquire 51% of the stakes? price of ETH would sky-rocket making it even difficult for them to reach the 51% mark, even if they do, whats next? a very high probablity of ETH value crash
- how will even anyone justify (assuming coming from public reserver from other country) spending 48 billion dollars?

So it definately is very impractical.

Whats practical? its more cost effective to manipulate elections intecepting the paper balots while on their way or break in and take away half of them (funny but I really mean it) or by misinformation

## System Architecture

## Microservices

1. Voter Registration
2. Voting
3. Vote Tallying
4. Auditing

## Smart Contracts

- minimize operations to save the gas costs
- Reentrancy Protection
- Access Control
- Data Integrity

### Voting.sol

- Voter Registration
  - only allow authorized users to register voters, no voter can register themselves, it should be done by an administrator only, like me (Raman rocks)
  - potential functions may be registerVoter, checkVoterRegs
- Vote Casting

  - validate that the user is registered to vote
  - allow voter to cast their vote
  - detect if voter is trying to duplicate vote

- Vote Tallying
  - real-time tallying of votes for each candidate, should provide transparent way to check vote counts

## Database

## User Identification

Database we have for users eligible to Vote
[User's Name, User's DOB, Place of Birth, User's SIN, User's Mobile Number, User's Email]

Identity Verification
They have to create an account with:
[User's Name, User's DOB, Place of Birth, User's SIN, User's Mobile Number, User's Email]
They choose to get OTP either on Mobile or Email. Once OTP is validated, we know this user is legit.
Extra question to create a Voter ID Hash -
Set Password, and 3 Security Questions and their answers.

Then Hash - [User's Name, User's DOB, Place of Birth, User's SIN, User's Mobile Number, User's Email, Password, Security Question1, Answer1, Security Question2, Answer2, Security Question3, Answer3] -> Voter ID Hash -> This user is now registered, put it on the block (Voter ID Hash + Nonce -> Hashed) and Voted (Yes/No)

Whenever user is voting, we create a Hash using the Voter ID Hash + Nonce (which is stored on our end) and check if that hash hasn't voted, then update Voted on the block transaction saying this person voted.

Nonce creation in Solidity - keccak256()
