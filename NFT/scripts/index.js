const { ethers } = require("hardhat")

const boxAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3"

async function main(){
    const Box = await ethers.getContractFactory("Box")
    const box = await Box.attach("0x5FbDB2315678afecb367f032d93F642f64180aa3")
    await box.store(14)
    console.log((await box.get()).toString())
}

main().then(()=>process.exit(0))
.catch(error => {
    console.error(error)
    process.exit(1)
})