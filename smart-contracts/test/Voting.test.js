const Voting = artifacts.require("Voting");

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
        

        assert.equal(event.event, "CandidateAdded", "CandidateAdded event was not emitted");
        assert.equal(event.args.candidateId, 1, "Candidates count should be 1");
        assert.equal(event.args.name, 'Candidate 1', "Candidates name should be Candidate 1");


    })

})