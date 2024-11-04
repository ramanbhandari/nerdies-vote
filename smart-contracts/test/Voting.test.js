const Voting = artifacts.require("Voting");

contract("Voting", (accounts)=>{
    before(async ()=>{
        instance = await Voting.deployed();
    }) 

    it('ensure that the starting balance of vending machine is 100', async ()=>{
        let balance = await instance.getVendingMachineBalance();

        assert.equal(balance, 100, "the initial balance should be 100");
    })

})