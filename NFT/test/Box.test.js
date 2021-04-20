const {expect} = require('chai')
const { ethers } = require('hardhat')

describe('Box',()=>{
    before(async ()=>{
        this.Box = await ethers.getContractFactory("Box")
    })

    beforeEach(async ()=>{
        this.box = await this.Box.deploy()
        await this.box.deployed()
    })

    it('get and store',async ()=>{
        await this.box.store(42)
        expect((await this.box.get()).toString()).to.equal('42');
    })
})