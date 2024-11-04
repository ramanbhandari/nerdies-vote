// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {
    struct Candidate {
        uint id;
        string name;
        uint voteCount;
    }

    // Mappings
    mapping(uint => Candidate) public candidates;
    mapping(bytes32 => bool) public isEligible;
    mapping(bytes32 => bool) public hasVoted;

    // log key actions
    event VoterRegistered(bytes32 voterHash);
    event VoteCasted(uint candidateId, bytes32 voterHash);

    constructor() {

    }

    modifier onlyAdmin() {

    }

    // register the voter, only admins should be able to access
    function registerVoterHash(bytes32 voterHash) public onlyAdmin {
        require(!isEligible[voterHash], "Already Registered");
        isEligible[voterHash] = true;
        emit VoterRegistered(voterHash);
    }

    // checks if user has voted
    function checkVoterStatus() public view return (bool) {
        return hasVoted[voterHash];
    }

    // to vote - increase Candidates votecount++
    function vote(uint candidateId, bytes32 voterHash) public {
        require(isEligible[voterHash], "Already Registered");
        require(!hasVoted[voterHash], "Already Voted");

        // needs validate candidateId validation
        candidates[candidateId].voteCount++;
        hasVoted[voterHash] = true;
        emit VoteCasted(candidateId, voterHash);
    }

    // to check if the user Voter Hash is valid
    function getVoter(bytes32 voterHash) public view return (bool) {
        return isEligible[voterHash];
    }

    // get the voter counts for all candidates - Public
    function getCount() public view returns (uint[] counts) {

    }

    // to add a candidate - Admin Only access
    function addCandidate(string memory _name) public onlyAdmin {

    }
}