pragma solidity >0.7.0;

/// @notice Kitti is a simulation of cryptoKitti used by CryptoZombi
/// to gen gene of a kitti.
contract Kitti{
    uint private geneLength = 16;
    uint private geneModulus = (10 ** (geneLength+1) ) - 1;

    function getKitti(uint _kittiId) external view returns(uint){
        return uint(keccak256(abi.encodePacked(_kittiId)))%geneModulus;
    }
}