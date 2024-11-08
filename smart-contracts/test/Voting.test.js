const Voting = artifacts.require("Voting");
const crypto = require('crypto');

contract("Voting", (accounts)=>{
    before(async ()=>{
        //for testing deployed contracts
        // instance = await Voting.deployed();

        instance = await Voting.new();
    }) 

    it('ensure that add candidate creates a candidate', async ()=>{

        //add a candidate 
        //the receipt is returns when a function emits an event
        let receipt = await instance.addCandidate('Candidate 1', { from: accounts[0] });

        try {
            let receipt = await instance.addCandidate('Candidate 1', { from: accounts[1] });
            assert.fail("Expected error not received");
        } catch (error) {
            assert.equal(error.reason, 'Not authorized');
            
        }
        
        let counts = await instance.getCount();

        console.log(counts[0].toNumber());
        
        console.log(receipt);
        

        //get the candidate
        let candidate1 = await instance.candidates(1);

        //confirm that candidate info is correct 
        assert.equal(candidate1.name, 'Candidate 1', "Candidate 1 contains the correct name");
        assert.equal(candidate1.id.toNumber(), 1, "Candidate 1 contains the correct id");
        assert.equal(candidate1.voteCount.toNumber(), 0, "Candidate 1 contains the correct vote count");

        //get the emitted event
        let event = receipt.logs[0];

        console.log(event.args);
        

        assert.equal(event.event, "CandidateAdded", "CandidateAdded event was emitted");
        assert.equal(event.args.candidateId, 1, "Candidates count should be 1");
        assert.equal(event.args.name, 'Candidate 1', "Candidates name should be Candidate 1");


    })


    it('ensure voters can be registered', async () => {
    
        const secret = 'hello';

        //mimicking voter id hashes
        const userHash = crypto.createHash('sha256', secret).digest('hex');
        const userHash2 = crypto.createHash('sha256', secret).update('2').digest('hex');

        let receipt = await instance.registerVoterHash('0x'+userHash, { from: accounts[0] });
        let receipt2 = await instance.registerVoterHash('0x'+userHash2, { from: accounts[0] });

        //trying to register the same voter 
        try {
            let receipt2 = await instance.registerVoterHash('0x'+userHash2, { from: accounts[0] });   
            assert.fail("Expected error not received");
        } catch (error) {
            assert.equal(error.reason, 'Already Registered');
        }


        //Anyone other than the admin should not be allowed to register a voter
        try {
            let receipt = await instance.registerVoterHash('0x'+userHash, { from: accounts[1] });
            assert.fail("Expected error not received");
        } catch (error) {
            assert.equal(error.reason, 'Not authorized');
            
        }

        //get the emitted event
        let event = receipt.logs[0];
        let event2 = receipt2.logs[0];

        
        let voter1 = await instance.isEligible(event.args.voterHash);
        let voter2 = await instance.isEligible(event2.args.voterHash);

        assert.equal(voter1, true, "Voter1 should be registered.");
        assert.equal(voter2, true, "Voter2 should be registered.");


    })

    it('ensure voting process works', async () => {
        const secret = 'hello';

        //mimicking voter id hashes
        const userHash = crypto.createHash('sha256', secret).digest('hex');
        const userHash2 = crypto.createHash('sha256', secret).update('2').digest('hex');

        const userNotRegistered = crypto.createHash('sha256', secret).update('3').digest('hex');

        //trying to vote for candidate that does not exist 
        try {
            let voteReceipt1 = await instance.vote(2,'0x'+userHash);
            assert.fail("Expected error not received");
        } catch (error) {
            assert.equal(error.reason, "Invalid candidate ID");
        }

        //vote candidate 1
        let voteReceipt1 = await instance.vote(1,'0x'+userHash);
        let voteReceipt2 = await instance.vote(1,'0x'+userHash2);


        let voterStatus1 = await instance.checkVoterStatus('0x'+userHash);
        let voterStatus2 = await instance.checkVoterStatus('0x'+userHash2);

        assert.equal(voterStatus1, true, "Voter 1 status should be true.")
        assert.equal(voterStatus2, true, "Voter 2 status should be true.")

        let event = voteReceipt1.logs[0];
        let event2 = voteReceipt2.logs[0];

        assert.equal(event.event, "VoteCasted", "VoteCasted event was emitted");
        assert.equal(event2.event, "VoteCasted", "VoteCasted event was emitted");

        let candidate1 = await instance.candidates(1);
        assert.equal(candidate1.voteCount.toNumber(), 2, "Candidate 1 should have 2 votes");

        //trying to vote without being registered
        try {
            let voteReceipt3 = await instance.vote(1,'0x'+userNotRegistered);
            assert.fail("Expected error not received");
        } catch (error) {        
            assert.equal(error.reason, "Voter Not Registered");
        }

        //trying to vote again
        try {
            let voteReceipt1 = await instance.vote(1,'0x'+userHash);
            assert.fail("Expected error not received");
        } catch (error) {
            assert.equal(error.reason, "Already Voted");
        }
        
    })

})